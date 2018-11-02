package aws

import "errors"

type AmazonKMS struct{}

func NewAmazonKMS() *AmazonKMS {
	return &AmazonKMS{}
}

func (a *AmazonKMS) Encrypt(plaintext []byte, params map[string]interface{}) ([]byte, error) {
	return nil, errors.New("NOT_IMPLEMENTED_YET")
}

func (a *AmazonKMS) Decrypt(ciphertext []byte, params map[string]interface{}) ([]byte, error) {
	return nil, errors.New("NOT_IMPLEMENTED_YET")
}