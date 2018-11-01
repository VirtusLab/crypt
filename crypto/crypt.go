package crypto

import (
	"github.com/Sirupsen/logrus"
	"github.com/VirtusLab/crypt/files"
)

type Crypt struct {
	kms KMS
}

func NewCrypt(kms KMS) *Crypt {
	return &Crypt{kms: kms}
}

func (c *Crypt) EncryptFile(inputPath, outputPath string, params map[string]interface{}) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		logrus.Debugf("Can't open plaintext file: %v", err)
		return err
	}
	err, result := c.kms.Encrypt(input, params)
	if err != nil {
		logrus.Debugf("Encrypting failed: %s", err)
		return err
	}
	err = files.WriteOutput(outputPath, result, 0644)
	if err != nil {
		logrus.Debugf("Can't save the encrypted file: %v", err)
		return err
	}
	return nil
}

func (c *Crypt) DecryptFile(inputPath, outputPath string, params map[string]interface{}) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		logrus.Debugf("Can't open encrypted file: %v", err)
		return err
	}
	err, result := c.kms.Decrypt(input, params)
	if err != nil {
		logrus.Debugf("Decrypting failed: %s", err)
		return err
	}
	err = files.WriteOutput(outputPath, result, 0644)
	if err != nil {
		logrus.Debugf("Can't save the decrypted file: %v", err)
		return err
	}
	return nil
}
