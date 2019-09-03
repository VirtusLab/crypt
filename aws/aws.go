package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// DefaultProfile is the default profile to be used when loading configuration
	// from the config files if another profile name is not provided.
	DefaultProfile = session.DefaultSharedConfigProfile
)

var (
	// ErrKmsMissing - this is the custom error, returned when name, alias or arn is missing
	ErrKmsMissing = errors.New("kms is empty or missing")
	// ErrRegionMissing - this is the custom error, returned when the region is missing
	ErrRegionMissing = errors.New("region is empty or missing")
)

// KMS struct represents AWS Key Management Service
type KMS struct {
	region  string
	key     string
	profile string
}

// New creates a AWS KMS provider
func New(key, region, profile string) *KMS {
	return &KMS{
		key:     key,
		region:  region,
		profile: profile,
	}
}

// Encrypt is responsible for encrypting plaintext and returning ciphertext in bytes using AWS KMS.
// See Crypt.Encrypt
func (k *KMS) Encrypt(plaintext []byte) ([]byte, error) {
	if len(k.key) == 0 {
		return nil, errors.Wrapf(ErrKmsMissing, "error reading kms: %v", k.key)
	}

	if len(k.region) == 0 {
		return nil, errors.Wrapf(ErrRegionMissing, "error reading region: %v", k.region)
	}

	if len(k.profile) == 0 {
		logrus.Debug("Using default AWS API credentials profile")
		k.profile = DefaultProfile
	}

	// Environment variables can be also used, see: /vendor/github.com/aws/aws-sdk-go/aws/session/env_config.go
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(k.region)},
		Profile:           k.profile,
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := kms.New(awsSession, aws.NewConfig().WithRegion(k.region))
	input := &kms.EncryptInput{
		Plaintext: plaintext,
		KeyId:     aws.String(k.key),
	}
	output, err := svc.Encrypt(input)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return []byte(output.CiphertextBlob), nil
}

// Decrypt is responsible for decrypting ciphertext and returning plaintext in bytes using AWS KMS.
// See Crypt.Decrypt
func (k *KMS) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(k.region) == 0 {
		return nil, errors.Wrapf(ErrRegionMissing, "error reading region: %v", k.region)
	}

	if len(k.profile) == 0 {
		logrus.Debug("Using default AWS API credentials profile")
		k.profile = DefaultProfile
	}

	// Environment variables can be also used, see: /vendor/github.com/aws/aws-sdk-go/aws/session/env_config.go
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           k.profile,
		Config:            aws.Config{Region: aws.String(k.region)},
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := kms.New(awsSession)
	input := &kms.DecryptInput{
		CiphertextBlob: ciphertext,
	}
	output, err := svc.Decrypt(input)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return output.Plaintext, nil
}
