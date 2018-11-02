package gcp

import (
	"fmt"
	"encoding/base64"
	"context"

	"google.golang.org/api/cloudkms/v1"
	"golang.org/x/oauth2/google"
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

const (
	ProjectId = "project"
	Location  = "location"
	KeyRing   = "keyring"
	Key       = "key"
)

type GoogleKMS struct{}

func NewGoogleKMS() *GoogleKMS {
	return &GoogleKMS{}
}

// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/kms/crypter/crypter.go
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
		params[ProjectId], params[Location], params[KeyRing], params[Key])

	req := &cloudkms.EncryptRequest{
		Plaintext: base64.StdEncoding.EncodeToString(plaintext),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Encrypt(parentName, req).Do()
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(resp.Ciphertext)
}

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
		params[ProjectId], params[Location], params[KeyRing], params[Key])

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
	projectId := params[ProjectId].(string)
	if len(projectId) == 0 {
		logrus.Debugf("Error reading project: %v", projectId)
		return errors.New("project is empty or missing!")
	}

	location := params[Location].(string)
	if len(location) == 0 {
		logrus.Debugf("Error reading location: %v", location)
		return errors.New("location is empty or missing!")
	}

	keyring := params[KeyRing].(string)
	if len(keyring) == 0 {
		logrus.Debugf("Error reading keyring: %v", keyring)
		return errors.New("keyring is empty or missing!")
	}

	key := params[Key].(string)
	if len(key) == 0 {
		logrus.Debugf("Error reading key: %v", key)
		return errors.New("key is empty or missing!")
	}

	return nil
}
