// +build integration

package azure

import (
	"github.com/VirtusLab/go-extended/pkg/files"
	"io/ioutil"
	"os"
	"path"
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

func TestCrypt_EncryptDecryptFiles(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// configuration from config.env
	vaultURL := os.Getenv("VAULT_URL")
	vaultKey := os.Getenv("VAULT_KEY")
	vaultKeyVersion := os.Getenv("VAULT_KEY_VERSION")
	require.NotEmpty(t, vaultURL)
	require.NotEmpty(t, vaultKey)
	require.NotEmpty(t, vaultKeyVersion)

	encryptedFileExtension := ".crypt"
	rootFileName := "root.txt"
	subdirectoryFileName := "sub-directory.txt"
	subdirectoryName := "sub-directory"
	inDir := "testdata/encryptDecryptFiles"
	encryptedFilesDir := "encryptedFiles"
	decryptedFilesDir := "decryptedFiles"

	keyVault, err := New(vaultURL, vaultKey, vaultKeyVersion)
	require.NoError(t, err)
	crypt := crypto.New(keyVault)
	defer func() { _ = os.RemoveAll(encryptedFilesDir) }()
	err = crypt.EncryptFiles(inDir, encryptedFilesDir, "", encryptedFileExtension)
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
