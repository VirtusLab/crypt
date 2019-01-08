package crypto

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/VirtusLab/crypt/test/kms/fake"

	"github.com/VirtusLab/go-extended/pkg/test"
	"github.com/stretchr/testify/assert"
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
	test.Run(t, test.Test{
		Name: "encrypt decrypt file",
		Fn: func(tt test.Test) {
			kms := fake.Empty()
			crypt := New(kms)

			inputFile := "test.txt"
			expected := "top secret token"
			err := ioutil.WriteFile(inputFile, []byte(expected), 0644)
			if err != nil {
				t.Fatal("Can't write plaintext file", err)
			}
			defer func() { _ = os.Remove(inputFile) }()

			actual, err := when(crypt, inputFile)

			assert.NoError(t, err, tt.Name)
			assert.Equal(t, expected, string(actual))
		},
	})
}

func TestCrypt_EncryptDecryptFileError(t *testing.T) {
	test.Run(t, test.Test{
		Name: "encrypt decrypt non-existing file",
		Fn: func(tt test.Test) {
			kms := fake.Empty()
			crypt := New(kms)

			inputFile := "test.txt"

			_, err := when(crypt, inputFile)

			assert.Error(t, err, tt.Name)
		},
	})
}
