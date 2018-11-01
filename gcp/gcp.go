package gcp

import "errors"

const (
	ProjectId = "project"
	Location  = "location"
	KeyRing   = "keyring"
	Key       = "key"
)

type GCP struct{}

func NewGCP() *GCP {
	return &GCP{}
}

// https://cloud.google.com/kms/docs/quickstart
// https://cloud.google.com/kms/docs/encrypt-decrypt#kms-howto-encrypt-go
func (g *GCP) Encrypt(inputPath, outputPath string, params map[string]interface{}) error {
	return errors.New("NOT_IMPLEMENTED_YET")
}

func (g *GCP) Decrypt(inputPath, outputPath string, params map[string]interface{}) error {
	return errors.New("NOT_IMPLEMENTED_YET")
}
