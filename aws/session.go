package aws

import (
	"github.com/VirtusLab/go-extended/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	// DefaultProfile is the default profile to be used when loading configuration
	// from the config files if another profile name is not provided.
	DefaultProfile = session.DefaultSharedConfigProfile
)

// SessionConfig returns AWS API client session and config with given region and profile
func SessionConfig(region, profile string) (*session.Session, *aws.Config, error) {
	// Environment variables can be also used, see: /vendor/github.com/aws/aws-sdk-go/aws/session/env_config.go
	config := aws.NewConfig().
		WithRegion(region).
		WithCredentialsChainVerboseErrors(true)

	awsSession, err := session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		Config:            *config,
		SharedConfigState: session.SharedConfigEnable,
	})

	return awsSession, config, errors.Wrap(err)
}
