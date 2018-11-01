package azure

import "errors"

type AzureKMS struct{}

func NewAzureKMS() *AzureKMS {
	return &AzureKMS{}
}

func (a *AzureKMS) Encrypt(plaintext []byte, params map[string]interface{}) (error, []byte) {
	return errors.New("NOT_IMPLEMENTED_YET"), nil
}

func (a *AzureKMS) Decrypt(ciphertext []byte, params map[string]interface{}) (error, []byte) {
	return errors.New("NOT_IMPLEMENTED_YET"), nil
}