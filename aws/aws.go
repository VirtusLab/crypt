package aws

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/sirupsen/logrus"
)

const (
	KMS    = "kms"
	Region = "region"
)

type AmazonKMS struct{}

func NewAmazonKMS() *AmazonKMS {
	return &AmazonKMS{}
}

func (a *AmazonKMS) Encrypt(plaintext []byte, params map[string]interface{}) ([]byte, error) {
	awsKms := params[KMS].(string)
	if len(awsKms) == 0 {
		logrus.Debugf("Error reading awsKms: %v", awsKms)
		return nil, errors.New("awsKms is empty or missing!")
	}

	region := params[Region].(string)
	if len(region) == 0 {
		logrus.Debugf("Error reading region: %v", region)
		return nil, errors.New("region is empty or missing!")
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

func (a *AmazonKMS) Decrypt(ciphertext []byte, params map[string]interface{}) ([]byte, error) {
	region := params[Region].(string)
	if len(region) == 0 {
		logrus.Debugf("Error reading region: %v", region)
		return nil, errors.New("region is empty or missing!")
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
