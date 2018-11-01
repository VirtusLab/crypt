package gcp

import "errors"

const (
	ProjectId = "project"
	Location  = "location"
	KeyRing   = "keyring"
	Key       = "key"
)

type GoogleKMS struct{}

func NewGoogleKMS() *GoogleKMS {
	return &GoogleKMS{}
}

func (g *GoogleKMS) Encrypt(plaintext []byte, params map[string]interface{}) (error, []byte) {
	return errors.New("NOT_IMPLEMENTED_YET"), nil
}

func (g *GoogleKMS) Decrypt(ciphertext []byte, params map[string]interface{}) (error, []byte) {
	return errors.New("NOT_IMPLEMENTED_YET"), nil
}
