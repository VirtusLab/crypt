package test

import (
	"io/ioutil"
	"os"

	"github.com/VirtusLab/crypt/crypto"
)

// EncryptAndDecryptFile encrypts and decrypts file using provider Crypt implementation
func EncryptAndDecryptFile(crypt crypto.Crypt, inputPath string) (string, error) {
	defer os.Remove(inputPath + ".encrypted") // clean up
	defer os.Remove(inputPath + ".decrypted") // clean up

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
