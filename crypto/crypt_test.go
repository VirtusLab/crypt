package crypto

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/VirtusLab/crypt/test/fake"

	"github.com/VirtusLab/go-extended/pkg/files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func encryptDecryptSingleFile(crypt Crypt, inputPath string) (string, error) {
	defer func() { _ = os.Remove(inputPath + ".encrypted") }() // clean up
	defer func() { _ = os.Remove(inputPath + ".decrypted") }() // clean up

	err := crypt.EncryptFile(inputPath, inputPath+".encrypted")
	if err != nil {
		return "", err
	}

	err = crypt.DecryptFile(inputPath+".encrypted", inputPath+".decrypted")
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadFile(inputPath + ".decrypted")
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func TestCrypt_EncryptDecryptFile(t *testing.T) {
	crypt := New(fake.Empty())

	inputFile := "test.txt"
	expected := "top secret token"
	err := ioutil.WriteFile(inputFile, []byte(expected), 0644)
	require.NoError(t, err, "Can't write plaintext file")
	defer func() { _ = os.Remove(inputFile) }()

	actual, err := encryptDecryptSingleFile(crypt, inputFile)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestCrypt_EncryptDecryptFileError(t *testing.T) {
	crypt := New(fake.Empty())

	inputFile := "test.txt"

	_, err := encryptDecryptSingleFile(crypt, inputFile)

	assert.EqualError(t, err, "can't open plaintext file: open test.txt: no such file or directory")
}

func TestCrypt_EncryptDecryptFiles(t *testing.T) {
	crypt := New(fake.Empty())
	encryptedFileExtension := ".crypt"
	rootFileName := "root.txt"
	subdirectoryFileName := "sub-directory.txt"
	subdirectoryName := "sub-directory"
	inDir := "testdata/encryptDecryptFiles"
	encryptedFilesDir := "encryptedFiles"
	decryptedFilesDir := "decryptedFiles"

	defer func() { _ = os.RemoveAll(encryptedFilesDir) }()
	err := crypt.EncryptFiles(inDir, encryptedFilesDir, "", encryptedFileExtension)
	require.NoError(t, err)
	assert.FileExists(t, path.Join(encryptedFilesDir, rootFileName+encryptedFileExtension))
	assert.FileExists(t, path.Join(encryptedFilesDir, subdirectoryName, subdirectoryFileName+encryptedFileExtension))

	defer func() { _ = os.RemoveAll(decryptedFilesDir) }()
	err = crypt.DecryptFiles(encryptedFilesDir, decryptedFilesDir, encryptedFileExtension, "")
	require.NoError(t, err)
	assert.FileExists(t, path.Join(decryptedFilesDir, rootFileName))
	assert.FileExists(t, path.Join(decryptedFilesDir, subdirectoryName, subdirectoryFileName))

	rootFile, err := files.ReadInput(path.Join(inDir, rootFileName))
	require.NoError(t, err)
	rootFileAfterDecryption, err := files.ReadInput(path.Join(decryptedFilesDir, rootFileName))
	require.NoError(t, err)
	assert.Equal(t, rootFile, rootFileAfterDecryption)

	subdirectoryFile, err := files.ReadInput(path.Join(inDir, subdirectoryName, subdirectoryFileName))
	require.NoError(t, err)
	subdirectoryFileAfterDecryption, err := files.ReadInput(path.Join(decryptedFilesDir, subdirectoryName, subdirectoryFileName))
	require.NoError(t, err)
	assert.Equal(t, subdirectoryFile, subdirectoryFileAfterDecryption)
}

func TestCrypt_DecryptFiles(t *testing.T) {
	crypt := New(fake.Empty())
	encryptedFileExtension := ".crypt"
	rootFileName := "root.txt"
	skipMeFileName := "skip-me.txt"
	encryptedFilesDir := "testdata/decryptFiles"
	decryptedFilesDir := "decryptedFiles"

	defer func() { _ = os.RemoveAll(decryptedFilesDir) }()
	err := crypt.DecryptFiles(encryptedFilesDir, decryptedFilesDir, encryptedFileExtension, "")
	require.NoError(t, err)
	assert.FileExists(t, path.Join(decryptedFilesDir, rootFileName))
	_, err = os.Lstat(path.Join(decryptedFilesDir, skipMeFileName))
	assert.EqualError(t, err, "lstat decryptedFiles/skip-me.txt: no such file or directory")
}

func Test_outputFilenameForEncryption(t *testing.T) {
	t.Run("inputExtension equals outputExtension", func(t *testing.T) {
		inputExtension := ".crypt"
		outputExtension := inputExtension
		file := files.FileEntry{
			Name:      "example.crypt",
			Extension: ".crypt",
		}

		got := outputFilenameForEncryption(file, inputExtension, outputExtension)

		assert.Equal(t, file.Name, got.Name)
		assert.Equal(t, file.Extension, got.Extension)
	})
	t.Run("inputExtension and outputExtension are empty", func(t *testing.T) {
		inputExtension := ""
		outputExtension := inputExtension
		file := files.FileEntry{
			Name:      "example.txt",
			Extension: ".txt",
		}

		got := outputFilenameForEncryption(file, inputExtension, outputExtension)

		assert.Equal(t, file.Name, got.Name)
		assert.Equal(t, file.Extension, got.Extension)
	})
	t.Run("inputExtension not empty and outputExtension is empty", func(t *testing.T) {
		inputExtension := ".crypt"
		outputExtension := ""
		file := files.FileEntry{
			Name:      "example.crypt",
			Extension: ".crypt",
		}

		got := outputFilenameForEncryption(file, inputExtension, outputExtension)

		assert.Equal(t, file.Name, got.Name)
		assert.Equal(t, file.Extension, got.Extension)
	})
	t.Run("inputExtension is empty and outputExtension is empty", func(t *testing.T) {
		inputExtension := ""
		outputExtension := ""
		file := files.FileEntry{
			Name:      "example.crypt",
			Extension: ".crypt",
		}

		got := outputFilenameForEncryption(file, inputExtension, outputExtension)

		assert.Equal(t, file.Name, got.Name)
		assert.Equal(t, file.Extension, got.Extension)
	})
	t.Run("inputExtension is empty and outputExtension is not empty", func(t *testing.T) {
		inputExtension := ""
		outputExtension := ".crypt"
		file := files.FileEntry{
			Name:      "example.txt",
			Extension: ".txt",
		}

		got := outputFilenameForEncryption(file, inputExtension, outputExtension)

		assert.Equal(t, "example.txt.crypt", got.Name)
		assert.Equal(t, outputExtension, got.Extension)
	})
	t.Run("inputExtension is not empty and outputExtension is not empty", func(t *testing.T) {
		inputExtension := ".secret"
		outputExtension := ".crypt"
		file := files.FileEntry{
			Name:      "example.secret",
			Extension: ".secret",
		}

		got := outputFilenameForEncryption(file, inputExtension, outputExtension)

		assert.Equal(t, "example.secret.crypt", got.Name)
		assert.Equal(t, outputExtension, got.Extension)
	})
}

func Test_outputFilenameForDecryption(t *testing.T) {
	t.Run("inputExtension equals outputExtension", func(t *testing.T) {
		inputExtension := ".crypt"
		outputExtension := inputExtension
		file := files.FileEntry{
			Name:      "example.crypt",
			Extension: ".crypt",
		}

		got := outputFilenameForDecryption(file, inputExtension, outputExtension)

		assert.Equal(t, file.Name, got.Name)
		assert.Equal(t, file.Extension, got.Extension)
	})
	t.Run("inputExtension and outputExtension are empty", func(t *testing.T) {
		inputExtension := ""
		outputExtension := inputExtension
		file := files.FileEntry{
			Name:      "example.txt",
			Extension: ".txt",
		}

		got := outputFilenameForDecryption(file, inputExtension, outputExtension)

		assert.Equal(t, file.Name, got.Name)
		assert.Equal(t, file.Extension, got.Extension)
	})
	t.Run("inputExtension not empty and outputExtension is empty", func(t *testing.T) {
		inputExtension := ".crypt"
		outputExtension := ""
		file := files.FileEntry{
			Name:      "example.crypt",
			Extension: ".crypt",
		}

		got := outputFilenameForDecryption(file, inputExtension, outputExtension)

		assert.Equal(t, "example", got.Name)
		assert.Equal(t, "", got.Extension)
	})
	t.Run("inputExtension is empty and outputExtension is empty", func(t *testing.T) {
		inputExtension := ""
		outputExtension := ""
		file := files.FileEntry{
			Name:      "example.crypt",
			Extension: ".crypt",
		}

		got := outputFilenameForDecryption(file, inputExtension, outputExtension)

		assert.Equal(t, file.Name, got.Name)
		assert.Equal(t, file.Extension, got.Extension)
	})
	t.Run("inputExtension is empty and outputExtension is not empty", func(t *testing.T) {
		inputExtension := ""
		outputExtension := ".crypt"
		file := files.FileEntry{
			Name:      "example.txt",
			Extension: ".txt",
		}

		got := outputFilenameForDecryption(file, inputExtension, outputExtension)

		assert.Equal(t, "example.txt.crypt", got.Name)
		assert.Equal(t, outputExtension, got.Extension)
	})
	t.Run("inputExtension is not empty and outputExtension is not empty", func(t *testing.T) {
		inputExtension := ".secret"
		outputExtension := ".crypt"
		file := files.FileEntry{
			Name:      "example.secret",
			Extension: ".secret",
		}

		got := outputFilenameForDecryption(file, inputExtension, outputExtension)

		assert.Equal(t, "example.secret.crypt", got.Name)
		assert.Equal(t, outputExtension, got.Extension)
	})
}
