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

func TestEncryptDecryptFileWithAzureKeyVault(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// configuration from config.env
	vaultURL := os.Getenv("VAULT_URL")
	vaultKey := os.Getenv("VAULT_KEY")
	vaultKeyVersion := os.Getenv("VAULT_KEY_VERSION")
	require.NotEmpty(t, vaultURL)
	require.NotEmpty(t, vaultKey)
	require.NotEmpty(t, vaultKeyVersion)

	crypt := crypto.New(New(vaultURL, vaultKey, vaultKeyVersion))

	inputFile := "test.txt"
	expected := "top secret token"
	err := ioutil.WriteFile(inputFile, []byte(expected), 0644)
	defer os.Remove(inputFile)
	require.NoError(t, err, "Can't write plaintext file")

	actual, err := test.EncryptAndDecryptFile(crypt, inputFile)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}
