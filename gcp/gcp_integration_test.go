// +build integration

package gcp

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

func TestEncryptDecryptWithGCP(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// configuration from config.env
	projectId := os.Getenv("GCP_PROJECT_ID")
	location := os.Getenv("GCP_LOCATION")
	keyring := os.Getenv("GCP_KEY_RING")
	key := os.Getenv("GCP_KEY")
	require.NotEmpty(t, projectId)
	require.NotEmpty(t, location)
	require.NotEmpty(t, keyring)
	require.NotEmpty(t, key)

	crypt := crypto.New(New(projectId, location, keyring, key))

	inputFile := "test.txt"
	expected := "top secret token"
	err := ioutil.WriteFile(inputFile, []byte(expected), 0644)
	defer os.Remove(inputFile)
	require.NoError(t, err, "Can't write plaintext file")

	actual, err := test.EncryptAndDecryptFile(crypt, crypt, inputFile)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}
