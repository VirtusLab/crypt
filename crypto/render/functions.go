package render

import (
	"text/template"

	"github.com/VirtusLab/crypt/aws"
	"github.com/VirtusLab/crypt/azure"
	"github.com/VirtusLab/crypt/gcp"
)

/*
TemplateFunctions provides template functions for render or the standard (text/template) template engine

  - encryptAWS - encrypts the data from inside of the template using AWS KMS, for best results use with gzip and b64enc
  - decryptAWS - decrypts the data from inside of the template using AWS KMS, for best results use with ungzip and b64dec
  - encryptGCP - encrypts the data from inside of the template using GCP KMS, for best results use with gzip and b64enc
  - decryptGCP - decrypts the data from inside of the template using GCP KMS, for best results use with ungzip and b64dec
  - encryptAzure - encrypts the data from inside of the template using Azure Key Vault, for best results use with gzip and b64enc
  - decryptAzure - decrypts the data from inside of the template using Azure Key Vault, for best results use with ungzip and b64dec
*/
func TemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"encryptAWS":   EncryptAWS,
		"decryptAWS":   DecryptAWS,
		"encryptGCP":   EncryptGCP,
		"decryptGCP":   DecryptGCP,
		"encryptAzure": EncryptAzure,
		"decryptAzure": DecryptAzure,
	}
}

// EncryptAWS encrypts plaintext using AWS KMS
func EncryptAWS(awsKms, awsRegion, awsProfile, plaintext string) ([]byte, error) {
	amazon := aws.New(awsKms, awsRegion, awsProfile)
	result, err := amazon.Encrypt([]byte(plaintext))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptAWS decrypts ciphertext using AWS KMS
func DecryptAWS(awsRegion, awsProfile, ciphertext string) (string, error) {
	amazon := aws.New("" /* not needed for decryption */, awsRegion, awsProfile)
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
	azr, err := azure.New(azureVaultURL, azureKey, azureKeyVersion)
	if err != nil {
		return nil, err
	}
	result, err := azr.Encrypt([]byte(plaintext))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DecryptAzure decrypts ciphertext using Azure Key Vault
func DecryptAzure(azureVaultURL, azureKey, azureKeyVersion, ciphertext string) (string, error) {
	azr, err := azure.New(azureVaultURL, azureKey, azureKeyVersion)
	if err != nil {
		return "", err
	}
	result, err := azr.Decrypt([]byte(ciphertext))
	if err != nil {
		return "", err
	}
	return string(result), nil
}
