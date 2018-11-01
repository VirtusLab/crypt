package aws

import "errors"

type AWS struct{}

func NewAWS() *AWS {
	return &AWS{}
}

func (a *AWS) Encrypt(inputPath, outputPath string, params map[string]interface{}) error {
	return errors.New("NOT_IMPLEMENTED_YET")
}

func (a *AWS) Decrypt(inputPath, outputPath string, params map[string]interface{}) error {
	return errors.New("NOT_IMPLEMENTED_YET")
}