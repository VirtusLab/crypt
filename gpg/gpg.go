package gpg

import (
	"bytes"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// GPG struct represents GPG (GnuPG) service
type GPG struct {
	PublicKeyPath  string
	PrivateKeyPath string
	Passphrase     string
}

// New creates GPG provider
func New(publicKeyPath, privateKeyPath, passphrase string) (*GPG, error) {
	return &GPG{
		PublicKeyPath:  publicKeyPath,
		PrivateKeyPath: privateKeyPath,
		Passphrase:     passphrase,
	}, nil
}

// Encrypt is responsible for encrypting plaintext and returning ciphertext in bytes using GPG (GnuPG).
// See Crypt.Encrypt
func (p *GPG) Encrypt(plaintext []byte) ([]byte, error) {
	return p.encryptWithKey(plaintext)
}

// Decrypt is responsible for decrypting ciphertext and returning plaintext in bytes using GPG (GnuPG).
// See Crypt.Decrypt
func (p *GPG) Decrypt(ciphertext []byte) ([]byte, error) {
	return p.decryptWithKey(ciphertext)
}

func (p *GPG) encryptWithKey(plaintext []byte) ([]byte, error) {
	entity, err := readEntity(p.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	encryptorWriter, err := openpgp.Encrypt(buf, []*openpgp.Entity{entity}, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	_, err = encryptorWriter.Write(plaintext)
	if err != nil {
		return nil, err
	}
	err = encryptorWriter.Close()
	if err != nil {
		return nil, err
	}

	encrypted, err := ioutil.ReadAll(buf)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func (p *GPG) decryptWithKey(ciphertext []byte) ([]byte, error) {
	privateKeyEntity, err := readEntity(p.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	if privateKeyEntity.PrivateKey.Encrypted {
		passphraseByte := []byte(p.Passphrase)
		err = privateKeyEntity.PrivateKey.Decrypt(passphraseByte)
		if err != nil {
			return nil, err
		}
		for _, subkey := range privateKeyEntity.Subkeys {
			err = subkey.PrivateKey.Decrypt(passphraseByte)
			if err != nil {
				return nil, err
			}
		}
	}

	entityList := openpgp.EntityList{privateKeyEntity}
	md, err := openpgp.ReadMessage(bytes.NewBuffer(ciphertext), entityList, nil, nil)
	if err != nil {
		return nil, err
	}
	decrypted, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func readEntity(file string) (*openpgp.Entity, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	block, err := armor.Decode(f)
	if err != nil {
		return nil, err
	}
	return openpgp.ReadEntity(packet.NewReader(block.Body))
}
