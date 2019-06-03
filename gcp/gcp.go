package gcp

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/api/cloudkms/v1"
)

var (
	// ErrProjectMissing - this is the custom error, returned when project is missing
	ErrProjectMissing = errors.New("project is empty or missing")
	// ErrLocationMissing = this is the custom error, returned when the location is missing
	ErrLocationMissing = errors.New("location is empty or missing")
	// ErrKeyRingMissing = this is the custom error, returned when the key ring is missing
	ErrKeyRingMissing = errors.New("key ring is empty or missing")
	// ErrKeyMissing = this is the custom error, returned when the key is missing
	ErrKeyMissing = errors.New("key is empty or missing")
)

// KMS provides a way to encrypt and decrypt the data using GCP KMS.
type KMS struct {
	project  string
	location string
	keyring  string
	key      string
}

// New new GCP KMS
func New(project, location, keyring, key string) *KMS {
	return &KMS{
		project:  project,
		location: location,
		keyring:  keyring,
		key:      key,
	}
}

// Encrypt is responsible for encrypting plaintext and returning ciphertext in bytes using GCP KMS.
// See Crypt.Encrypt
func (k *KMS) Encrypt(plaintext []byte) ([]byte, error) {
	err := k.validate()
	if err != nil {
		return nil, err // err already wrapped in validate function
	}

	// See https://cloud.google.com/docs/authentication/.
	// Use GOOGLE_APPLICATION_CREDENTIALS environment variable to specify
	// a service account key file to authenticate to the API.
	ctx := context.Background()
	kmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error creating gcp cloudkms service")
	}

	parentName := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		k.project, k.location, k.keyring, k.key)

	req := &cloudkms.EncryptRequest{
		Plaintext: base64.StdEncoding.EncodeToString(plaintext),
	}
	resp, err := kmsService.Projects.Locations.KeyRings.CryptoKeys.Encrypt(parentName, req).Do()
	if err != nil {
		return nil, errors.Wrap(err, "encryption failed")
	}

	return base64.StdEncoding.DecodeString(resp.Ciphertext)
}

// Decrypt is responsible for decrypting ciphertext and returning plaintext in bytes using GCP KMS.
// See Crypt.DecryptFile
func (k *KMS) Decrypt(ciphertext []byte) ([]byte, error) {
	err := k.validate()
	if err != nil {
		return nil, err // err already wrapped in validate function
	}

	// See https://cloud.google.com/docs/authentication/.
	// Use GOOGLE_APPLICATION_CREDENTIALS environment variable to specify
	// a service account key file to authenticate to the API.
	ctx := context.Background()

	cloudkmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error creating gcp cloudkms service")
	}

	parentName := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		k.project, k.location, k.keyring, k.key)

	req := &cloudkms.DecryptRequest{
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Decrypt(parentName, req).Do()
	if err != nil {
		return nil, errors.Wrap(err, "decryption failed")
	}
	return base64.StdEncoding.DecodeString(resp.Plaintext)
}

func (k *KMS) validate() error {
	if len(k.project) == 0 {
		return errors.Wrapf(ErrProjectMissing, "error reading project: %v", k.project)
	}
	if len(k.location) == 0 {
		return errors.Wrapf(ErrLocationMissing, "error reading location: %v", k.location)
	}
	if len(k.keyring) == 0 {
		return errors.Wrapf(ErrKeyRingMissing, "error reading keyring: %v", k.keyring)
	}
	if len(k.key) == 0 {
		return errors.Wrapf(ErrKeyMissing, "error reading key: %v", k.key)
	}
	return nil
}
