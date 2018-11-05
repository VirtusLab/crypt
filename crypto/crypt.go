package crypto

import (
	"github.com/VirtusLab/crypt/files"
	"github.com/sirupsen/logrus"
)

// Crypt structure holds implementation of particular Key Management Service e.g. AWS KMS
type Crypt struct {
	kms KMS
}

// NewCrypt creates a new crypt with corresponding Key Management Service
func NewCrypt(kms KMS) *Crypt {
	return &Crypt{kms: kms}
}

// EncryptFile reads from the inputPath file or stdin if empty.
// Then encrypts content with corresponding Key Management Service.
// Ciphertext is saved into outputPath file or print on stdout if empty.
func (c *Crypt) EncryptFile(inputPath, outputPath string, params map[string]interface{}) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		logrus.Debugf("Can't open plaintext file: %v", err)
		return err
	}
	result, err := c.kms.Encrypt(input, params)
	if err != nil {
		logrus.Debugf("Encrypting failed: %s", err)
		return err
	}
	err = files.WriteOutput(outputPath, result, 0644) // 0644 - user: read&write, group: read, other: read
	if err != nil {
		logrus.Debugf("Can't save the encrypted file: %v", err)
		return err
	}
	return nil
}

// DecryptFile reads from the inputPath file or stdin if empty.
// Then decrypts content with corresponding Key Management Service.
// Plaintext is saved into outputPath file or print on stdout if empty.
func (c *Crypt) DecryptFile(inputPath, outputPath string, params map[string]interface{}) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		logrus.Debugf("Can't open encrypted file: %v", err)
		return err
	}
	result, err := c.kms.Decrypt(input, params)
	if err != nil {
		logrus.Debugf("Decrypting failed: %s", err)
		return err
	}
	err = files.WriteOutput(outputPath, result, 0644) // 0644 - user: read&write, group: read, other: read
	if err != nil {
		logrus.Debugf("Can't save the decrypted file: %v", err)
		return err
	}
	return nil
}
