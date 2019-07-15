package gpg

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncryptDecryptWithKeys(t *testing.T) {
	publicKey, err := ioutil.TempFile(os.TempDir(), "public-*.gpg")
	require.NoError(t, err)
	publicKey.Write([]byte(PublicKey))
	defer publicKey.Close()
	defer os.Remove(publicKey.Name())

	privateKey, err := ioutil.TempFile(os.TempDir(), "private-*.gpg")
	require.NoError(t, err)
	privateKey.Write([]byte(PrivateKey))
	defer privateKey.Close()
	defer os.Remove(privateKey.Name())

	gnupg, err := New(privateKey.Name(), privateKey.Name(), "")
	require.NoError(t, err)

	plaintext := "TOP SECRET"
	ciphertext, err := gnupg.Encrypt([]byte(plaintext))
	require.NoError(t, err)

	decrypted, err := gnupg.Decrypt(ciphertext)
	require.NoError(t, err)

	require.Equal(t, string(decrypted), plaintext)
}

const PublicKey = `
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
=nwRd
-----END PGP PUBLIC KEY BLOCK-----
`

const PrivateKey = `
-----BEGIN PGP PRIVATE KEY BLOCK-----

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
wv/NNKpsKFl+weV+lsBXgKGFQVue
=C7iz
-----END PGP PRIVATE KEY BLOCK-----
`
