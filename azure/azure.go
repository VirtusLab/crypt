package azure

import (
	"errors"
	"context"
	"encoding/base64"

	"github.com/sirupsen/logrus"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
)

var (
	// ErrVaultURLMissing - this is the custom error, returned when vault url is missing
	ErrVaultURLMissing = errors.New("vault url is empty or missing")
	// ErrKeyMissing = this is the custom error, returned when the key is missing
	ErrKeyMissing = errors.New("key is empty or missing")
	// ErrKeyVersionMissing = this is the custom error, returned when the key version is missing
	ErrKeyVersionMissing = errors.New("key version is empty or missing")
)

// AzureKMS struct represents Azure Key Vault
type AzureKMS struct {
	vaultURL   string
	key        string
	keyVersion string
}

// NewAzureKMS creates Azure Key Vault KMS
func NewAzureKMS(vaultURL, key, keyVersion string) *AzureKMS {
	return &AzureKMS{
		vaultURL:   vaultURL,
		key:        key,
		keyVersion: keyVersion,
	}
}

func newKeyVaultClient() (keyvault.BaseClient, error) {
	var err error
	vaultClient := keyvault.New()
	vaultClient.Authorizer, err = auth.NewAuthorizerFromEnvironment()
	if err != nil {
		logrus.WithError(err).Error("Failed to create Azure Authorizer")
		return vaultClient, err
	}
	return vaultClient, nil
}

// Encrypt is responsible for encrypting plaintext by Azure Key Vault encryption key and returning ciphertext in bytes.
// See Crypt.EncryptFile
func (a *AzureKMS) Encrypt(plaintext []byte) ([]byte, error) {
	err := a.validateParams()
	if err != nil {
		return nil, err
	}

	client, err := newKeyVaultClient()
	if err != nil {
		return nil, err
	}
	data := base64.RawURLEncoding.EncodeToString(plaintext)
	p := keyvault.KeyOperationsParameters{Value: &data, Algorithm: keyvault.RSAOAEP256}

	ctx := context.Background()
	res, err := client.Encrypt(ctx, a.vaultURL, a.key, a.keyVersion, p)
	if err != nil {
		return nil, err
	}

	result, err := base64.RawURLEncoding.DecodeString(*res.Result)
	logrus.WithFields(logrus.Fields{
		"key":        a.key,
		"keyVersion": a.keyVersion,
	}).Info("Encryption succeeded")
	return result, nil
}

// Decrypt is responsible for decrypting ciphertext by Azure Key Vault encryption key and returning plaintext in bytes.
// See Crypt.EncryptFile
func (a *AzureKMS) Decrypt(ciphertext []byte) ([]byte, error) {
	// FIXME a.validateParams()
	client, err := newKeyVaultClient()
	if err != nil {
		return nil, err
	}
	data := base64.RawURLEncoding.EncodeToString(ciphertext)
	p := keyvault.KeyOperationsParameters{Value: &data, Algorithm: keyvault.RSAOAEP256}

	ctx := context.Background()
	res, err := client.Decrypt(ctx, a.vaultURL, a.key, a.keyVersion, p)
	if err != nil {
		return nil, err
	}

	plaintext, err := base64.RawURLEncoding.DecodeString(*res.Result)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"key":        a.key,
		"keyVersion": a.keyVersion,
	}).Info("Decryption succeeded")

	return plaintext, nil
}

func (a *AzureKMS) validateParams() error {
	if len(a.vaultURL) == 0 {
		logrus.Debugf("Error reading vaultURL: %v", a.vaultURL)
		return ErrVaultURLMissing
	}
	if len(a.key) == 0 {
		logrus.Debugf("Error reading key: %v", a.key)
		return ErrKeyMissing
	}
	if len(a.keyVersion) == 0 {
		logrus.Debugf("Error reading keyVersion: %v", a.keyVersion)
		return ErrKeyVersionMissing
	}
	return nil
}
