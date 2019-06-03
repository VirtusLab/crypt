package crypto

import (
	"os"
	"path"
	"path/filepath"

	"github.com/VirtusLab/go-extended/pkg/files"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Crypt is an abstraction for encryption and decryption with files support
type Crypt interface {
	Crypter
	EncryptFile(inputPath, outputPath string) error
	EncryptFiles(inputDir, outputDir, inputExtension, outputExtension string) error
	DecryptFile(inputPath, outputPath string) error
	DecryptFiles(inputDir, outputDir, inputExtension, outputExtension string) error
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

func transformFiles(
	inputDir, outputDir, inputExtension, outputExtension string,
	inputFilterFunc func(file files.FileEntry, inputExtension string) bool,
	transformFunc func(inputPath, outputPath string) error,
	outputFilenameFunc func(file files.FileEntry, inputExtension, outputExtension string) files.FileEntry) error {

	logrus.Infof("Directory mode selected: '%s' -> '%s'", inputDir, outputDir)

	fileEntries, err := files.DirTree(inputDir)
	if err != nil {
		return errors.Wrapf(err, "can't scan the directory tree: '%s'", inputDir)
	}

	for _, file := range fileEntries {
		if inputFilterFunc(file, inputExtension) {
			logrus.Debugf("Skipping '%s'", path.Join(file.Path, file.Name))
			continue
		}
		logrus.Debugf("Processing '%s'", path.Join(file.Path, file.Name))

		target := outputFilenameFunc(file, inputExtension, outputExtension)

		rel, err := filepath.Rel(inputDir, file.Path)
		if err != nil {
			return errors.Wrapf(err, "can't get a relative path for: '%s'", file.Path)
		}

		target.Path = path.Join(outputDir, rel)

		err = os.MkdirAll(target.Path, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "can't create the target directory: '%s'", target.Path)
		}
		logrus.Infof("Target directory was created: '%s'", target.Path)

		err = transformFunc(path.Join(file.Path, file.Name), path.Join(target.Path, target.Name))
		if err != nil {
			return errors.Wrap(err, "can't encrypt/decrypt a file")
		}
	}

	return nil
}

func inputFilterFunc(file files.FileEntry, inputExtension string) bool {
	return len(inputExtension) > 0 && inputExtension != file.Extension
}

// EncryptFiles encrypts files from a directory using EncryptFile function
func (c *crypt) EncryptFiles(inputDir, outputDir, inputExtension, outputExtension string) error {
	// FIXME @tsek this does not support all use cases
	outputFunc := func(file files.FileEntry, inputExtension, outputExtension string) files.FileEntry {
		fileEntry := files.FileEntry{
			Path: file.Path,
			Name: file.Name + outputExtension,
		}
		fileEntry.Extension = filepath.Ext(fileEntry.Name)
		return fileEntry
	}
	return transformFiles(inputDir, outputDir, inputExtension, outputExtension, inputFilterFunc, c.EncryptFile, outputFunc)
}

// DecryptFiles decrypts files from a directory using DecryptFile function
func (c *crypt) DecryptFiles(inputDir, outputDir, inputExtension, outputExtension string) error {
	// FIXME @tsek this does not support all use cases
	outputFunc := func(file files.FileEntry, inputExtension, outputExtension string) files.FileEntry {
		if len(inputExtension) == 0 {
			fileEntry := files.FileEntry{
				Path: file.Path,
				Name: file.Name + outputExtension,
			}
			fileEntry.Extension = filepath.Ext(fileEntry.Name)
			return fileEntry
		}
		return files.TrimExtension(file, []string{inputExtension})
	}
	return transformFiles(inputDir, outputDir, inputExtension, outputExtension, inputFilterFunc, c.DecryptFile, outputFunc)
}

// EncryptFile encrypts bytes from a file or stdin using a Crypter provider
// and the ciphertext is saved into a file.
// If inputPath is empty, stdin is used as input
// If outputPath is empty, stdout is used as output
func (c *crypt) EncryptFile(inputPath, outputPath string) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		return errors.Wrap(err, "can't open plaintext file")
	}
	result, err := c.Encrypt(input)
	if err != nil {
		return errors.Wrapf(err, "encrypting failed, file '%s'", inputPath)
	}
	err = files.WriteOutput(outputPath, result, 0644) // 0644 - user: read&write, group: read, other: read
	if err != nil {
		return errors.Wrapf(err, "can't save the encrypted file '%s'", outputPath)
	}
	return nil
}

// DecryptFile reads from the inputPath file or stdin if empty.
// Then decrypts content with corresponding Key Management Service.
// Plaintext is saved into outputPath file or print on stdout if empty.
func (c *crypt) DecryptFile(inputPath, outputPath string) error {
	input, err := files.ReadInput(inputPath)
	if err != nil {
		return errors.Wrap(err, "can't open encrypted file")
	}
	result, err := c.Decrypt(input)
	if err != nil {
		return errors.Wrapf(err, "decrypting failed, file '%s'", inputPath)
	}
	err = files.WriteOutput(outputPath, result, 0644) // 0644 - user: read&write, group: read, other: read
	if err != nil {
		return errors.Wrapf(err, "can't save the decrypted file '%s'", outputPath)
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
