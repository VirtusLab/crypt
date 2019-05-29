// +build integration

package azure

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/VirtusLab/crypt/crypto"
	"github.com/VirtusLab/crypt/test"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptFileWithHeader(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// configuration from config.env
	vaultURL := os.Getenv("VAULT_URL")
	vaultKey := os.Getenv("VAULT_KEY")
	vaultKeyVersion := os.Getenv("VAULT_KEY_VERSION")
	require.NotEmpty(t, vaultURL)
	require.NotEmpty(t, vaultKey)
	require.NotEmpty(t, vaultKeyVersion)

	keyVault, err := New(vaultURL, vaultKey, vaultKeyVersion)
	require.NoError(t, err)
	encrypt := crypto.New(keyVault)
	// don't need key info because it is in header in encrypted file
	keyVault, err = New("", "", "")
	require.NoError(t, err)
	decrypt := crypto.New(keyVault)

	inputFile := "test.txt"
	secret := "top secret token"
	err = ioutil.WriteFile(inputFile, []byte(secret), 0644)
	defer func() { _ = os.Remove(inputFile) }()
	require.NoError(t, err, "Can't write plaintext file")

	actual, err := test.EncryptAndDecryptFile(encrypt, decrypt, inputFile)

	assert.NoError(t, err)
	assert.Equal(t, secret, string(actual))
}

func TestEncryptDecryptWithoutHeader(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// configuration from config.env
	vaultURL := os.Getenv("VAULT_URL")
	vaultKey := os.Getenv("VAULT_KEY")
	vaultKeyVersion := os.Getenv("VAULT_KEY_VERSION")
	require.NotEmpty(t, vaultURL)
	require.NotEmpty(t, vaultKey)
	require.NotEmpty(t, vaultKeyVersion)

	keyVault, err := New(vaultURL, vaultKey, vaultKeyVersion)
	require.NoError(t, err)
	secret := "top secret token"
	encrypted, err := keyVault.encrypt([]byte(secret), false)
	require.NoError(t, err)
	decrypted, err := keyVault.Decrypt(encrypted)
	require.NoError(t, err)

	assert.Equal(t, string(decrypted), secret)
}
