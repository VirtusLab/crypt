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

const (
	// Project - this is constant used in params
	Project = "project"
	// Location - this is constant used in params
	Location = "location"
	// KeyRing - this is constant used in params
	KeyRing = "keyring"
	// Key - this is constant used in params
	Key = "key"
)

var (
	// ErrProjectMissing - this is the custom error, returned when project is missing
	ErrProjectMissing = errors.New("project is empty or missing")
	// ErrLocationMissing = this is the custom error, returned when the location is missing
	ErrLocationMissing = errors.New("location is empty or missing")
	// ErrKeyRingMissing = this is the custom error, returned when the location is missing
	ErrKeyRingMissing = errors.New("key ring is empty or missing")
	// ErrKeyMissing = this is the custom error, returned when the location is missing
	ErrKeyMissing = errors.New("key is empty or missing")
)

// GoogleKMS struct represents GCP Key Management Service
type GoogleKMS struct{}

// NewGoogleKMS new GCP KMS
func NewGoogleKMS() *GoogleKMS {
	return &GoogleKMS{}
}

// Encrypt is responsible for encrypting plaintext and returning ciphertext in bytes using GCP KMS.
// All configuration is passed in params with according validation.
// See Crypt.EncryptFile
func (g *GoogleKMS) Encrypt(plaintext []byte, params map[string]interface{}) ([]byte, error) {
	err := validateParams(params)
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
		params[Project], params[Location], params[KeyRing], params[Key])

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
// All configuration is passed in params with according validation.
// See Crypt.DecryptFile
func (g *GoogleKMS) Decrypt(ciphertext []byte, params map[string]interface{}) ([]byte, error) {
	err := validateParams(params)
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
		params[Project], params[Location], params[KeyRing], params[Key])

	req := &cloudkms.DecryptRequest{
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Decrypt(parentName, req).Do()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(resp.Plaintext)
}

func validateParams(params map[string]interface{}) error {
	project := params[Project].(string)
	if len(project) == 0 {
		logrus.Debugf("Error reading project: %v", project)
		return ErrProjectMissing
	}
	location := params[Location].(string)
	if len(location) == 0 {
		logrus.Debugf("Error reading location: %v", location)
		return ErrLocationMissing
	}
	keyring := params[KeyRing].(string)
	if len(keyring) == 0 {
		logrus.Debugf("Error reading keyring: %v", keyring)
		return ErrKeyRingMissing
	}
	key := params[Key].(string)
	if len(key) == 0 {
		logrus.Debugf("Error reading key: %v", key)
		return ErrKeyMissing
	}
	return nil
}
