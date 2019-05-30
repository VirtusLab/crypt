package crypto

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/VirtusLab/crypt/test/fake"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func when(crypt Crypt, inputPath string) (string, error) {
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
	kms := fake.Empty()
	crypt := New(kms)

	inputFile := "test.txt"
	expected := "top secret token"
	err := ioutil.WriteFile(inputFile, []byte(expected), 0644)
	require.NoError(t, err, "Can't write plaintext file")
	defer func() { _ = os.Remove(inputFile) }()

	actual, err := when(crypt, inputFile)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestCrypt_EncryptDecryptFileError(t *testing.T) {
	kms := fake.Empty()
	crypt := New(kms)

	inputFile := "test.txt"

	_, err := when(crypt, inputFile)

	assert.EqualError(t, err, "open test.txt: no such file or directory")
}
