package crypto

import (
	"github.com/VirtusLab/crypt/files"

	"github.com/sirupsen/logrus"
)

// Crypt is an abstraction for encryption and decryption with files support
type Crypt interface {
	Crypter
	EncryptFile(inputPath, outputPath string) error
	DecryptFile(inputPath, outputPath string) error
}

// Crypter is an Encrypter and a Decrypter
type Crypter interface {
	Encrypter
	Decrypter
}

// Encrypter must be able to encrypt plaintext into ciphertext, see also Decrypter
type Encrypter interface {
	Encrypt(plaintext []byte) ([]byte, error)
}

// Decrypter must be able to decrypt ciphertext into plaintext, see also Encrypter
type Decrypter interface {
	Decrypt(ciphertext []byte) ([]byte, error)
}

// Crypt type represents the crypt abstraction for simple encryption and decryption.
// A provider (e.g. AWS KMS) determines the detail of the cryptographic operations.
type crypt struct {
	crypter Crypter
}

// New creates a new Crypt with the given provider
func New(crypter Crypter) Crypt {
	return &crypt{crypter: crypter}
}

// EncryptFile encrypts bytes from a file or stdin using a Crypter provider
// and the ciphertext is saved into a file.
// If inputPath is empty, stdin is used as input
// If outputPath is empty, stdout is used as output
func (c *crypt) EncryptFile(inputPath, outputPath string) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		logrus.Debugf("Can't open plaintext file: %v", err)
		return err
	}
	result, err := c.Encrypt(input)
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
func (c *crypt) DecryptFile(inputPath, outputPath string) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		logrus.Debugf("Can't open encrypted file: %v", err)
		return err
	}
	result, err := c.Decrypt(input)
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

// Decrypt decrypts given bytes using the current provider
func (c *crypt) Decrypt(input []byte) ([]byte, error) {
	return c.crypter.Decrypt(input)
}

// Encrypt encrypts given bytes using the current provider
func (c *crypt) Encrypt(input []byte) ([]byte, error) {
	return c.crypter.Encrypt(input)
}
