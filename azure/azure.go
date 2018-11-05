package azure

import (
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	log "github.com/sirupsen/logrus"

	"aurora-tool/_vendor-20180329085108/github.com/sirupsen/logrus"
	"context"
	"encoding/base64"
)

const (
	// VaultURL - the vault url (example: https://myvault.vault.azure.net/)
	VaultURL = "vaultURL"
	// Name - the name of the key
	Name = "name"
	// Version - the version of the key (example: 1d40f45795a444099c81ca274b88c141)
	Version = "version"
)

// AzureKMS struct represents Azure Key Vault
type AzureKMS struct{}

// Constructor for AzureKMS
func NewAzureKMS() *AzureKMS {
	return &AzureKMS{}
}

func newKeyVaultClient() (keyvault.BaseClient, error) {
	var err error
	keyvaultClient := keyvault.New()
	keyvaultClient.Authorizer, err = auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.WithError(err).Error("Failed to create Azure Authorizer")
		return keyvaultClient, err
	}
	return keyvaultClient, nil
}

// Encrypt is responsible for encrypting plaintext by Azure Key Vault encryption key and returning ciphertext in bytes.
// All configuration is passed in params with according validation.
// See Crypt.EncryptFile
func (a *AzureKMS) Encrypt(plaintext []byte, params map[string]interface{}) ([]byte, error) {
	err := validateParams(params)
	if err != nil {
		return nil, err
	}
	c, err := newKeyVaultClient()
	if err != nil {
		return nil, err
	}
	data := base64.RawURLEncoding.EncodeToString(plaintext)
	p := keyvault.KeyOperationsParameters{Value: &data, Algorithm: keyvault.RSAOAEP256}

	res, err := c.Encrypt(context.Background(),
		params[VaultURL].(string),
		params[Name].(string),
		params[Version].(string),
		p)

	if err != nil {
		return nil, err
	}

	result, err := base64.RawURLEncoding.DecodeString(*res.Result)
	log.WithFields(log.Fields{
		"key":     params[Name].(string),
		"version": params[Version].(string),
	}).Info("Encryption succeeded")
	return result, nil
}

// Decrypt is responsible for decrypting ciphertext by Azure Key Vault encryption key and returning plaintext in bytes.
// All configuration is passed in params with according validation.
// See Crypt.EncryptFile
func (a *AzureKMS) Decrypt(ciphertext []byte, params map[string]interface{}) ([]byte, error) {
	c, err := newKeyVaultClient()
	if err != nil {
		return nil, err
	}
	data := base64.RawURLEncoding.EncodeToString(ciphertext)
	p := keyvault.KeyOperationsParameters{Value: &data, Algorithm: keyvault.RSAOAEP256}

	res, err := c.Decrypt(context.TODO(),
		params[VaultURL].(string),
		params[Name].(string),
		params[Version].(string),
		p)

	if err != nil {
		return nil, err
	}

	plaintext, err := base64.RawURLEncoding.DecodeString(*res.Result)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"key":     params[Name].(string),
		"version": params[Version].(string),
	}).Info("Decryption succeeded")
	return plaintext, nil
}

func validateParams(params map[string]interface{}) error {
	vaultURL := params[VaultURL].(string)
	if len(vaultURL) == 0 {
		logrus.Debugf("Error reading vaultURL: %v", vaultURL)
		return errors.New("vaultURL is empty or missing!")
	}

	name := params[Name].(string)
	if len(name) == 0 {
		logrus.Debugf("Error reading name: %v", name)
		return errors.New("name is empty or missing!")
	}

	version := params[Version].(string)
	if len(version) == 0 {
		logrus.Debugf("Error reading version: %v", version)
		return errors.New("version is empty or missing!")
	}

	return nil
}
