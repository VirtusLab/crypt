package gcp

import (
	"golang.org/x/oauth2/google"
	"fmt"
	"encoding/base64"
	"context"

	"google.golang.org/api/cloudkms/v1"
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
	// FIXME(bantoniak) preconditions
	projectId := params[ProjectId]
	location := params[Location]
	keyring := params[KeyRing]
	key := params[Key]

	ctx := context.Background()

	// See https://cloud.google.com/docs/authentication/.
	// Use GOOGLE_APPLICATION_CREDENTIALS environment variable to specify
	// a service account key file to authenticate to the API.
	client, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	cloudkmsService, err := cloudkms.New(client)
	if err != nil {
		return nil, err
	}

	parentName := fmt.Sprintf("projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		projectId, location, keyring, key)

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
	// FIXME(bantoniak) preconditions
	projectId := params[ProjectId]
	location := params[Location]
	keyring := params[KeyRing]
	key := params[Key]

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
		projectId, location, keyring, key)

	req := &cloudkms.DecryptRequest{
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
	}
	resp, err := cloudkmsService.Projects.Locations.KeyRings.CryptoKeys.Decrypt(parentName, req).Do()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(resp.Plaintext)
}
