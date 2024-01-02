package azure

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/VirtusLab/crypt/version"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

const (
	providerName                        = "az"
	encryptedFileMetadataSeparator byte = '.'
)

var (
	magicNumber []byte
	// ErrVaultURLMissing - this is the custom error, returned when vault url is missing
	ErrVaultURLMissing = errors.New("key vault URL is empty or missing")
	// ErrKeyMissing = this is the custom error, returned when the KeyVault key is missing
	ErrKeyMissing = errors.New("key vault key is empty or missing")
	// ErrKeyVersionMissing = this is the custom error, returned when the KeyVault key version is missing
	ErrKeyVersionMissing = errors.New("key vault key version is empty or missing")
)

// MetadataHeader holds information about KeyVault key used to encrypt
type MetadataHeader struct {
	Provider                string `json:"provider"`
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
	client     KeyVaultClient
}

// KeyVaultClient interface represents Azure Key Vault client for encrypting and decrypting data
type KeyVaultClient interface {
	Encrypt(ctx context.Context, name string, version string, parameters azkeys.KeyOperationsParameters, options *azkeys.EncryptOptions) (azkeys.EncryptResponse, error)
	Decrypt(ctx context.Context, name string, version string, parameters azkeys.KeyOperationsParameters, options *azkeys.DecryptOptions) (azkeys.DecryptResponse, error)
}

// New creates Azure Key Vault KeyVault using chained token credential
func New(vaultURL, key, keyVersion string) (*KeyVault, error) {
	azCreds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	if err != nil {
		return nil, err // err already wrapped in newKeyVaultClient function
	}
	keyVaultClient, err := azkeys.NewClient(vaultURL, azCreds, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &KeyVault{
		client:     keyVaultClient,
		vaultURL:   vaultURL,
		key:        key,
		keyVersion: keyVersion,
	}, nil
}

// Encrypt encrypts plaintext using Azure Key Vault and returns ciphertext
// See Crypt.Encrypt
func (k *KeyVault) Encrypt(plaintext []byte) ([]byte, error) {
	return k.encrypt(plaintext, true)
}

func (k *KeyVault) encrypt(plaintext []byte, includeHeader bool) ([]byte, error) {
	err := k.validate()
	if err != nil {
		return nil, err // err already wrapped in validate function
	}

	alg := azkeys.JSONWebKeyEncryptionAlgorithmRSAOAEP256
	p := azkeys.KeyOperationsParameters{Value: plaintext, Algorithm: &alg}
	res, err := k.client.Encrypt(context.Background(), k.key, k.keyVersion, p, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error with decoding data")
	}
	if includeHeader {
		metadata := MetadataHeader{
			Provider:                providerName,
			CryptVersion:            version.VERSION,
			AzureKeyVaultURL:        k.vaultURL,
			AzureKeyVaultKeyName:    k.key,
			AzureKeyVaultKeyVersion: k.keyVersion,
		}

		metadataBytes, err := json.Marshal(metadata)
		if err != nil {
			return nil, errors.Wrap(err, "error with marshaling header metadata")
		}
		metadataURLEncoded := make([]byte, base64.RawURLEncoding.EncodedLen(len(metadataBytes)))
		base64.RawURLEncoding.Encode(metadataURLEncoded, metadataBytes)

		logrus.WithFields(logrus.Fields{
			"keyVaultURL": k.vaultURL,
			"key":         k.key,
			"keyVersion":  k.keyVersion,
		}).Info("Encryption succeeded")
		result := append(metadataURLEncoded, encryptedFileMetadataSeparator)
		result = append(result, res.Result...)

		return result, nil
	}

	result, err := base64.RawURLEncoding.DecodeString(string(res.Result[:]))
	if err != nil {
		return nil, errors.Wrap(err, "error with decoding data")
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
	if !bytes.HasPrefix(ciphertext, magicNumber) {
		logrus.Debug("Cipher text doesn't contains metadata header")
		err := k.validate()
		if err != nil {
			return nil, err // err already wrapped in validate function
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
			return nil, errors.Wrap(err, "error with decoding header metadata")
		}
		err = json.Unmarshal(metadataURLDecoded, &metadata)
		if err != nil {
			return nil, errors.Wrap(err, "error with unmarshaling header metadata")
		}
		k.vaultURL = metadata.AzureKeyVaultURL
		k.key = metadata.AzureKeyVaultKeyName
		k.keyVersion = metadata.AzureKeyVaultKeyVersion
	}

	var plaintext []byte

	alg := azkeys.JSONWebKeyEncryptionAlgorithmRSAOAEP256
	p := azkeys.KeyOperationsParameters{Value: []byte(dataToDecrypt), Algorithm: &alg}
	res, err := k.client.Decrypt(context.Background(), k.key, k.keyVersion, p, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error with decoding data")
	}
	plaintext = res.Result

	logrus.WithFields(logrus.Fields{
		"keyVaultURL": k.vaultURL,
		"key":         k.key,
		"keyVersion":  k.keyVersion,
	}).Info("Decryption succeeded")

	return plaintext, nil
}

func (k *KeyVault) validate() error {
	if len(k.vaultURL) == 0 {
		return errors.Wrapf(ErrVaultURLMissing, "error reading vaultURL: %v", k.vaultURL)
	}
	if len(k.key) == 0 {
		return errors.Wrapf(ErrKeyMissing, "error reading key: %v", k.key)
	}
	if len(k.keyVersion) == 0 {
		return errors.Wrapf(ErrKeyVersionMissing, "error reading keyVersion: %v", k.keyVersion)
	}
	return nil
}

func init() {
	fileContentPrefix := []byte(`{"provider":"az","crypt"`)
	magicNumber = make([]byte, base64.RawURLEncoding.EncodedLen(len(fileContentPrefix)))
	base64.RawURLEncoding.Encode(magicNumber, fileContentPrefix)
}
