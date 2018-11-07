package aws

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/sirupsen/logrus"
)

var (
	// ErrKmsMissing - this is the custom error, returned when name, alias or arn is missing
	ErrKmsMissing = errors.New("kms is empty or missing")
	// ErrRegionMissing - this is the custom error, returned when the region is missing
	ErrRegionMissing = errors.New("region is empty or missing")
)

// AmazonKMS struct represents AWS Key Management Service
type AmazonKMS struct {
	key    string
	region string
}

// NewAmazonKMS creates AWS KMS
func NewAmazonKMS(key, region string) *AmazonKMS {
	return &AmazonKMS{
		key:    key,
		region: region,
	}
}

// Encrypt is responsible for encrypting plaintext and returning ciphertext in bytes using AWS KMS.
// See Crypt.EncryptFile
func (a *AmazonKMS) Encrypt(plaintext []byte) ([]byte, error) {
	if len(a.key) == 0 {
		logrus.Debugf("Error reading kms: %v", a.key)
		return nil, ErrKmsMissing
	}

	if len(a.region) == 0 {
		logrus.Debugf("Error reading region: %v", a.region)
		return nil, ErrRegionMissing
	}

	// use AWS_DEFAULT_PROFILE environment variable to set profile
	//
	// If not set and environment variables are not set the "default"
	// ill be used as the profile to load the session config from.
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(a.region)},
	}))
	svc := kms.New(awsSession, aws.NewConfig().WithRegion(a.region))
	input := &kms.EncryptInput{
		Plaintext: plaintext,
		KeyId:     aws.String(a.key),
	}
	output, err := svc.Encrypt(input)
	if err != nil {
		return nil, err
	}
	return []byte(output.CiphertextBlob), nil
}

// Decrypt is responsible for decrypting ciphertext and returning plaintext in bytes using AWS KMS.
// See Crypt.DecryptFile
func (a *AmazonKMS) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(a.region) == 0 {
		logrus.Debugf("Error reading region: %v", a.region)
		return nil, ErrRegionMissing
	}

	// use AWS_DEFAULT_PROFILE environment variable to set profile
	//
	// If not set and environment variables are not set the "default"
	// ill be used as the profile to load the session config from.
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(a.region)},
	}))
	svc := kms.New(awsSession)
	input := &kms.DecryptInput{
		CiphertextBlob: ciphertext,
	}
	output, err := svc.Decrypt(input)
	if err != nil {
		return nil, err
	}
	return output.Plaintext, nil
}
