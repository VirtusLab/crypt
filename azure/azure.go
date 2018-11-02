package azure

import "errors"

type AzureKMS struct{}

func NewAzureKMS() *AzureKMS {
	return &AzureKMS{}
}

func (a *AzureKMS) Encrypt(plaintext []byte, params map[string]interface{}) ([]byte, error) {
	return nil, errors.New("NOT_IMPLEMENTED_YET")
}

func (a *AzureKMS) Decrypt(ciphertext []byte, params map[string]interface{}) ([]byte, error) {
	return nil, errors.New("NOT_IMPLEMENTED_YET")
}