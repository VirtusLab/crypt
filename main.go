package main

import (
	"fmt"
	"os"

	"github.com/VirtusLab/crypt/aws"
	"github.com/VirtusLab/crypt/azure"
	"github.com/VirtusLab/crypt/constants"
	"github.com/VirtusLab/crypt/crypto"
	"github.com/VirtusLab/crypt/gcp"
	"github.com/VirtusLab/crypt/version"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
)

var (
	app        *cli.App
	inputPath  string
	outputPath string

	// GCP Cloud KMS resources belong to a project
	gcpProject string
	// The geographical data center location where requests to Cloud KMS are handled
	gcpLocation string
	// A key ring is a grouping of keys for organizational purposes
	gcpKeyring string
	// A key is a named object representing a cryptographic key used for a specific purpose
	gcpKey string

	// Amazon Resource Name (ARN), alias name or alias ARN for the customer master key
	awsKms string
	//  The geographical data center location where requests to AWS KMS are handled
	awsRegion string
)

func main() {
	app = cli.NewApp()
	app.Name = constants.Name
	app.Usage = constants.Description
	app.Author = constants.Author
	app.Version = fmt.Sprintf("%s-%s", version.VERSION, version.GITCOMMIT)
	app.Before = preload
	app.Commands = []cli.Command{
		encrypt(),
		decrypt(),
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "run in debug mode",
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(cli.ErrWriter, "There is no %q command.\n", command)
		cli.OsExiter(1)
	}
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		if isSubcommand {
			return err
		}

		fmt.Fprintf(cli.ErrWriter, "WRONG: %v\n", err)
		return nil
	}
	cli.OsExiter = func(c int) {
		if c != 0 {
			logrus.Debugf("exiting with %d", c)
		}
		os.Exit(c)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(cli.ErrWriter, "ERROR: %v\n", err)
		cli.OsExiter(1)
	}
}

func preload(c *cli.Context) error {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	logrus.SetLevel(logrus.InfoLevel)

	if c.GlobalBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		logrus.Debug("Debug logging enabled")
	}

	if len(c.Args()) == 0 {
		return nil
	}

	if c.Args()[0] == "help" {
		return nil
	}

	return nil
}

func encrypt() cli.Command {
	return cli.Command{
		Name:    "encrypt",
		Aliases: []string{"enc", "en", "e"},
		Usage:   "Encrypts files and/or strings",
		Subcommands: []cli.Command{
			{
				Name:  "azure",
				Usage: "Encrypts files and/or strings with Azure Key Vault",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(azure.NewAzureKMS())
					params := map[string]interface{}{
						// TODO add necessary params
					}
					err := crypt.EncryptFile(inputPath, outputPath, params)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "aws",
				Usage: "Encrypts files and/or strings with AWS KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
					cli.StringFlag{
						Name:        "region",
						Value:       "",
						Usage:       "the AWS region",
						Destination: &awsRegion, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "key-id, kms, kms-alias",
						Value:       "",
						Usage:       "the Amazon Resource Name (ARN), alias name, or alias ARN for the customer master key",
						Destination: &awsKms, // FIXME #2 make this flag required
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(aws.NewAmazonKMS())
					params := map[string]interface{}{
						aws.KMS:    awsKms,
						aws.Region: awsRegion,
					}
					err := crypt.EncryptFile(inputPath, outputPath, params)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "gcp",
				Usage: "Encrypts files and/or strings with GCP KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
					cli.StringFlag{
						Name:        "project",
						Value:       "",
						Usage:       "the GCP project id for Cloud KMS",
						Destination: &gcpProject, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "location",
						Value:       "",
						Usage:       "the location for project and Cloud KMS",
						Destination: &gcpLocation, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "keyring",
						Value:       "",
						Usage:       "the key ring name",
						Destination: &gcpKeyring, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "key",
						Value:       "",
						Usage:       "the cryptographic key name",
						Destination: &gcpKey, // FIXME #2 make this flag required
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(gcp.NewGoogleKMS())
					params := map[string]interface{}{
						gcp.Project:  gcpProject,
						gcp.Location: gcpLocation,
						gcp.KeyRing:  gcpKeyring,
						gcp.Key:      gcpKey,
					}
					err := crypt.EncryptFile(inputPath, outputPath, params)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
}

func decrypt() cli.Command {
	return cli.Command{
		Name:    "decrypt",
		Aliases: []string{"dec", "de", "d"},
		Usage:   "Decrypts files and/or strings",
		Subcommands: []cli.Command{
			{
				Name:  "azure",
				Usage: "Decrypts files and/or strings with Azure Key Vault",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(azure.NewAzureKMS())
					params := map[string]interface{}{
						// TODO add necessary params
					}
					err := crypt.DecryptFile(inputPath, outputPath, params)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "aws",
				Usage: "Decrypts files and/or strings with AmazonKMS KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
					cli.StringFlag{
						Name:        "region",
						Value:       "",
						Usage:       "the AWS region",
						Destination: &awsRegion, // FIXME #2 make this flag required
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(aws.NewAmazonKMS())
					params := map[string]interface{}{
						aws.KMS:    awsKms,
						aws.Region: awsRegion,
					}
					err := crypt.DecryptFile(inputPath, outputPath, params)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "gcp",
				Usage: "Decrypts files and/or strings with GCP KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
					cli.StringFlag{
						Name:        "project",
						Value:       "",
						Usage:       "the GCP project id for Cloud KMS",
						Destination: &gcpProject, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "location",
						Value:       "",
						Usage:       "the location for project and Cloud KMS",
						Destination: &gcpLocation, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "keyring",
						Value:       "",
						Usage:       "the key ring name",
						Destination: &gcpKeyring, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "key",
						Value:       "",
						Usage:       "the cryptographic key name",
						Destination: &gcpKey, // FIXME #2 make this flag required
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(gcp.NewGoogleKMS())
					params := map[string]interface{}{
						gcp.Project:  gcpProject,
						gcp.Location: gcpLocation,
						gcp.KeyRing:  gcpKeyring,
						gcp.Key:      gcpKey,
					}
					err := crypt.DecryptFile(inputPath, outputPath, params)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
}
