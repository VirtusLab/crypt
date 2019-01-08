package render

import (
	"github.com/VirtusLab/crypt/aws"
	"github.com/VirtusLab/crypt/azure"
	"github.com/VirtusLab/crypt/gcp"
)

// EncryptAWS encrypts plaintext using AWS KMS
func EncryptAWS(awsKms, awsRegion, plaintext string) ([]byte, error) {
	amazon := aws.New(awsKms, awsRegion)
	result, err := amazon.Encrypt([]byte(plaintext))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptAWS decrypts ciphertext using AWS KMS
func DecryptAWS(awsRegion, ciphertext string) (string, error) {
	amazon := aws.New("" /* not needed for decryption */, awsRegion)
	result, err := amazon.Decrypt([]byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// EncryptGCP encrypts plaintext using GCP KMS
func EncryptGCP(gcpProject, gcpLocation, gcpKeyring, gcpKey, plaintext string) ([]byte, error) {
	googleKms := gcp.New(gcpProject, gcpLocation, gcpKeyring, gcpKey)
	result, err := googleKms.Encrypt([]byte(plaintext))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptGCP decrypts ciphertext using GCP KMS
func DecryptGCP(gcpProject, gcpLocation, gcpKeyring, gcpKey, ciphertext string) (string, error) {
	googleKms := gcp.New(gcpProject, gcpLocation, gcpKeyring, gcpKey)
	result, err := googleKms.Decrypt([]byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// EncryptAzure encrypts plaintext using Azure Key Vault
func EncryptAzure(azureVaultURL, azureKey, azureKeyVersion, plaintext string) ([]byte, error) {
	azr := azure.New(azureVaultURL, azureKey, azureKeyVersion)
	result, err := azr.Encrypt([]byte(plaintext))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptAzure decrypts ciphertext using Azure Key Vault
func DecryptAzure(azureVaultURL, azureKey, azureKeyVersion, ciphertext string) (string, error) {
	azr := azure.New(azureVaultURL, azureKey, azureKeyVersion)
	result, err := azr.Decrypt([]byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(result), nil
}
