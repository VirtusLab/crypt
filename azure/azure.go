package azure

import "errors"

type Azure struct{}

func NewAzure() *Azure {
	return &Azure{}
}

func (a *Azure) Encrypt(inputPath, outputPath string, params map[string]interface{}) error {
	return errors.New("NOT_IMPLEMENTED_YET")
}

func (a *Azure) Decrypt(inputPath, outputPath string, params map[string]interface{}) error {
	return errors.New("NOT_IMPLEMENTED_YET")
}