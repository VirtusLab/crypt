package aws

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/sirupsen/logrus"
)

const (
	// KMS - this is constant used in params
	KMS = "kms"
	// Region - this is constant used in params
	Region = "region"
)

var (
	// ErrKmsMissing - this is the custom error, returned when name, alias or arn is missing
	ErrKmsMissing = errors.New("kms is empty or missing")
	// ErrRegionMissing - this is the custom error, returned when the region is missing
	ErrRegionMissing = errors.New("region is empty or missing")
)

// AmazonKMS struct represents AWS Key Management Service
type AmazonKMS struct{}

// NewAmazonKMS creates AWS KMS
func NewAmazonKMS() *AmazonKMS {
	return &AmazonKMS{}
}

// Encrypt is responsible for encrypting plaintext and returning ciphertext in bytes using AWS KMS.
// All configuration is passed in params with according validation.
// See Crypt.EncryptFile
func (a *AmazonKMS) Encrypt(plaintext []byte, params map[string]interface{}) ([]byte, error) {
	awsKms := params[KMS].(string)
	if len(awsKms) == 0 {
		logrus.Debugf("Error reading kms: %v", awsKms)
		return nil, ErrKmsMissing
	}

	region := params[Region].(string)
	if len(region) == 0 {
		logrus.Debugf("Error reading region: %v", region)
		return nil, ErrRegionMissing
	}

	// use AWS_DEFAULT_PROFILE environment variable to set profile
	//
	// If not set and environment variables are not set the "default"
	// ill be used as the profile to load the session config from.
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(params[Region].(string))},
	}))
	svc := kms.New(awsSession, aws.NewConfig().WithRegion(params[Region].(string)))
	input := &kms.EncryptInput{
		Plaintext: plaintext,
		KeyId:     aws.String(params[KMS].(string)),
	}
	output, err := svc.Encrypt(input)
	if err != nil {
		return nil, err
	}
	return []byte(output.CiphertextBlob), nil
}

// Decrypt is responsible for decrypting ciphertext and returning plaintext in bytes using AWS KMS.
// All configuration is passed in params with according validation.
// See Crypt.DecryptFile
func (a *AmazonKMS) Decrypt(ciphertext []byte, params map[string]interface{}) ([]byte, error) {
	region := params[Region].(string)
	if len(region) == 0 {
		logrus.Debugf("Error reading region: %v", region)
		return nil, ErrRegionMissing
	}

	// use AWS_DEFAULT_PROFILE environment variable to set profile
	//
	// If not set and environment variables are not set the "default"
	// ill be used as the profile to load the session config from.
	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(params[Region].(string))},
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
