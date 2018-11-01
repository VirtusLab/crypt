package aws

import "errors"

type AmazonKMS struct{}

func NewAmazonKMS() *AmazonKMS {
	return &AmazonKMS{}
}

func (a *AmazonKMS) Encrypt(plaintext []byte, params map[string]interface{}) (error, []byte) {
	return errors.New("NOT_IMPLEMENTED_YET"), nil
}

func (a *AmazonKMS) Decrypt(ciphertext []byte, params map[string]interface{}) (error, []byte) {
	return errors.New("NOT_IMPLEMENTED_YET"), nil
}