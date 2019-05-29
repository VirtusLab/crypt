package azure

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/VirtusLab/crypt/version"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault/keyvaultapi"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/sirupsen/logrus"
)

const (
	encryptedFileMetadataSeparator byte = '.'
)

var (
	magicNumbers []byte
	// ErrVaultURLMissing - this is the custom error, returned when vault url is missing
	ErrVaultURLMissing = errors.New("key vault URL is empty or missing")
	// ErrKeyMissing = this is the custom error, returned when the KeyVault key is missing
	ErrKeyMissing = errors.New("key vault key is empty or missing")
	// ErrKeyVersionMissing = this is the custom error, returned when the KeyVault key version is missing
	ErrKeyVersionMissing = errors.New("key vault key version is empty or missing")
)

// MetadataHeader holds information about KeyVault key used to encrypt
type MetadataHeader struct {
	CryptVersion            string `json:"crypt"`
	AzureKeyVaultURL        string `json:"kvURL"`
	AzureKeyVaultKeyName    string `json:"kvKey"`
	AzureKeyVaultKeyVersion string `json:"kvKeyVer"`
}

// KeyVault struct represents Azure Key Vault
type KeyVault struct {
	vaultURL   string
	key        string
	keyVersion string
	client     keyvaultapi.BaseClientAPI
}

// New creates Azure Key Vault KeyVault
func New(vaultURL, key, keyVersion string) (*KeyVault, error) {
	client, err := newKeyVaultClient()
	if err != nil {
		return nil, err
	}
	return &KeyVault{
		client:     client,
		vaultURL:   vaultURL,
		key:        key,
		keyVersion: keyVersion,
	}, nil
}

func newKeyVaultClient() (keyvaultapi.BaseClientAPI, error) {
	var err error
	vaultClient := keyvault.New()
	vaultClient.Authorizer, err = auth.NewAuthorizerFromCLI()
	if err != nil {
		logrus.WithError(err).Error("failed to create Azure Authorizer")
		return vaultClient, err
	}
	return vaultClient, nil
}

// Encrypt encrypts plaintext using Azure Key Vault and returns ciphertext
// See Crypt.Encrypt
func (k *KeyVault) Encrypt(plaintext []byte) ([]byte, error) {
	return k.encrypt(plaintext, true)
}

func (k *KeyVault) encrypt(plaintext []byte, includeHeader bool) ([]byte, error) {
	err := k.validate()
	if err != nil {
		return nil, err
	}

	data := base64.RawURLEncoding.EncodeToString(plaintext)
	p := keyvault.KeyOperationsParameters{Value: &data, Algorithm: keyvault.RSAOAEP256}

	res, err := k.client.Encrypt(context.Background(), k.vaultURL, k.key, k.keyVersion, p)
	if err != nil {
		return nil, err
	}

	if includeHeader {
		metadata := MetadataHeader{
			CryptVersion:            version.VERSION,
			AzureKeyVaultURL:        k.vaultURL,
			AzureKeyVaultKeyName:    k.key,
			AzureKeyVaultKeyVersion: k.keyVersion,
		}

		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return nil, err
		}
		metadataURLEncoded := make([]byte, base64.RawURLEncoding.EncodedLen(len(metadataBytes)))
		base64.RawURLEncoding.Encode(metadataURLEncoded, metadataBytes)

		logrus.WithFields(logrus.Fields{
			"keyVaultURL": k.vaultURL,
			"key":         k.key,
			"keyVersion":  k.keyVersion,
		}).Info("Encryption succeeded")
		result := append(metadataURLEncoded, encryptedFileMetadataSeparator)
		result = append(result, []byte(*res.Result)...)
		return result, nil
	}
	result, err := base64.RawURLEncoding.DecodeString(*res.Result)
	if err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"key":        k.key,
		"keyVersion": k.keyVersion,
	}).Info("Encryption succeeded")
	return result, nil
}

// Decrypt is responsible for decrypting ciphertext by Azure Key Vault encryption key and returning plaintext in bytes.
// See Crypt.EncryptFile
func (k *KeyVault) Decrypt(ciphertext []byte) ([]byte, error) {
	var dataToDecrypt string
	if !bytes.HasPrefix(ciphertext, magicNumbers) {
		logrus.Debug("Cipher text doesn't contains metadata header")
		err := k.validate()
		if err != nil {
			return nil, err
		}
		dataToDecrypt = base64.RawURLEncoding.EncodeToString(ciphertext)
	} else {
		logrus.Debug("Cipher text contains metadata header")
		indexOfSeparator := bytes.IndexByte(ciphertext, encryptedFileMetadataSeparator)
		dataToDecrypt = string(ciphertext[indexOfSeparator+1:])
		metadata := MetadataHeader{}
		metadataURLDecoded := make([]byte, base64.RawURLEncoding.DecodedLen(len(ciphertext[:indexOfSeparator])))
		_, err := base64.RawURLEncoding.Decode(metadataURLDecoded, ciphertext[:indexOfSeparator])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(metadataURLDecoded, &metadata)
		if err != nil {
			return nil, err
		}
		k.vaultURL = metadata.AzureKeyVaultURL
		k.key = metadata.AzureKeyVaultKeyName
		k.keyVersion = metadata.AzureKeyVaultKeyVersion
	}

	p := keyvault.KeyOperationsParameters{Value: &dataToDecrypt, Algorithm: keyvault.RSAOAEP256}

	res, err := k.client.Decrypt(context.Background(), k.vaultURL, k.key, k.keyVersion, p)
	if err != nil {
		return nil, err
	}

	plaintext, err := base64.RawURLEncoding.DecodeString(*res.Result)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"keyVaultURL": k.vaultURL,
		"key":         k.key,
		"keyVersion":  k.keyVersion,
	}).Info("Decryption succeeded")

	return plaintext, nil
}

func (k *KeyVault) validate() error {
	if len(k.vaultURL) == 0 {
		logrus.Debugf("Error reading vaultURL: %v", k.vaultURL)
		return ErrVaultURLMissing
	}
	if len(k.key) == 0 {
		logrus.Debugf("Error reading key: %v", k.key)
		return ErrKeyMissing
	}
	if len(k.keyVersion) == 0 {
		logrus.Debugf("Error reading keyVersion: %v", k.keyVersion)
		return ErrKeyVersionMissing
	}
	return nil
}

func init() {
	fileContentPrefix := []byte("{\"crypt\":")
	magicNumbers = make([]byte, base64.RawURLEncoding.EncodedLen(len(fileContentPrefix)))
	base64.RawURLEncoding.Encode(magicNumbers, fileContentPrefix)
}
