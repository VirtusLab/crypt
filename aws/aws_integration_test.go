// +build integration

package aws

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

func TestEncryptDecryptWithAWS(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// configuration from config.env
	key := os.Getenv("AWS_KEY")
	region := os.Getenv("AWS_REGION")
	profile := os.Getenv("AWS_PROFILE")
	require.NotEmpty(t, key)
	require.NotEmpty(t, region)
	require.NotEmpty(t, profile)

	crypt := crypto.New(New(key, region, profile))

	inputFile := "test.txt"
	expected := "top secret token"
	err := ioutil.WriteFile(inputFile, []byte(expected), 0644)
	defer os.Remove(inputFile)
	require.NoError(t, err, "Can't write plaintext file")

	actual, err := test.EncryptAndDecryptFile(crypt, inputFile)

	assert.NoError(t, err)
	assert.Equal(t, expected, string(actual))
}
