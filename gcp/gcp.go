package gcp

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
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

// GoogleKMS provides a way to encrypt and decrypt the data using GCP KMS.
type GoogleKMS struct {
	project  string
	location string
	keyring  string
	key      string
}

// NewGoogleKMS new GCP KMS
func NewGoogleKMS(project, location, keyring, key string) *GoogleKMS {
	return &GoogleKMS{
		project:  project,
		location: location,
		keyring:  keyring,
		key:      key,
	}
}

// Encrypt is responsible for encrypting plaintext and returning ciphertext in bytes using GCP KMS.
// See Crypt.EncryptFile
func (g *GoogleKMS) Encrypt(plaintext []byte) ([]byte, error) {
	err := g.validateParams()
	if err != nil {
		return nil, err
	}

	// See https://cloud.google.com/docs/authentication/.
	// Use GOOGLE_APPLICATION_CREDENTIALS environment variable to specify
	// a service account key file to authenticate to the API.
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	cloudkmsService, err := cloudkms.New(client)
	if err != nil {
		return nil, err
	}

	parentName := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		g.project, g.location, g.keyring, g.key)

	req := &cloudkms.EncryptRequest{
		Plaintext: base64.StdEncoding.EncodeToString(plaintext),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Encrypt(parentName, req).Do()
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(resp.Ciphertext)
}

// Decrypt is responsible for decrypting ciphertext and returning plaintext in bytes using GCP KMS.
// See Crypt.DecryptFile
func (g *GoogleKMS) Decrypt(ciphertext []byte) ([]byte, error) {
	err := g.validateParams()
	if err != nil {
		return nil, err
	}

	// See https://cloud.google.com/docs/authentication/.
	// Use GOOGLE_APPLICATION_CREDENTIALS environment variable to specify
	// a service account key file to authenticate to the API.
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	cloudkmsService, err := cloudkms.New(client)
	if err != nil {
		return nil, err
	}

	parentName := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		g.project, g.location, g.keyring, g.key)

	req := &cloudkms.DecryptRequest{
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Decrypt(parentName, req).Do()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(resp.Plaintext)
}

func (g *GoogleKMS) validateParams() error {
	if len(g.project) == 0 {
		logrus.Debugf("Error reading project: %v", g.project)
		return ErrProjectMissing
	}
	if len(g.location) == 0 {
		logrus.Debugf("Error reading location: %v", g.location)
		return ErrLocationMissing
	}
	if len(g.keyring) == 0 {
		logrus.Debugf("Error reading keyring: %v", g.keyring)
		return ErrKeyRingMissing
	}
	if len(g.key) == 0 {
		logrus.Debugf("Error reading key: %v", g.key)
		return ErrKeyMissing
	}
	return nil
}
