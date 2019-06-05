package azure

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/VirtusLab/go-extended/pkg/cli"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	killIn   = 10 * time.Second
	cryptCmd = "../crypt" // crypt is compiled before run this test in makefile
)

func TestEncryptDecryptFile(t *testing.T) {
	// configuration from config.env
	vaultURL := os.Getenv("VAULT_URL")
	vaultKey := os.Getenv("VAULT_KEY")
	vaultKeyVersion := os.Getenv("VAULT_KEY_VERSION")
	require.NotEmpty(t, vaultURL)
	require.NotEmpty(t, vaultKey)
	require.NotEmpty(t, vaultKeyVersion)

	logger := logrus.New()
	secret := "uber-secret"
	encryptedFileName := "encrypted"
	ctx, cancel := context.WithTimeout(context.TODO(), killIn)
	defer cancel()

	stdout, stderr, err := cli.Sh(ctx, logger, []string{}, &secret,
		cryptCmd, "encrypt", "azure", "--out", encryptedFileName,
		"--vaultURL", vaultURL, "--name", vaultKey, "--version", vaultKeyVersion)
	defer func() { _ = os.Remove(encryptedFileName) }()
	require.NoError(t, err, stdout, stderr)

	stdout, stderr, err = cli.Sh(ctx, logger, []string{}, &secret,
		cryptCmd, "decrypt", "azure", "--in", encryptedFileName)
	require.NoError(t, err, stdout, stderr)
	assert.Equal(t, secret, stdout)
}
