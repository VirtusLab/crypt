package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
	"github.com/Sirupsen/logrus"
	"github.com/VirtusLab/crypt/aws"
	"github.com/VirtusLab/crypt/azure"
	"github.com/VirtusLab/crypt/constants"
	"github.com/VirtusLab/crypt/crypto"
	"github.com/VirtusLab/crypt/gcp"
	"github.com/VirtusLab/crypt/version"
)

var (
	app        *cli.App
	inputPath  string
	outputPath string

	// gcp kms
	gcpProject  string
	gcpLocation string
	gcpKeyring  string
	gcpKey      string
)

func main() {
	app = cli.NewApp()
	app.Name = constants.Name
	app.Usage = constants.Description
	app.Author = constants.Author
	app.Version = fmt.Sprintf("%s-%s", version.VERSION, version.GITCOMMIT)
	app.Before = preload

	app.Commands = []cli.Command{
		encryptCommand(),
		decryptCommand(),
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

func encryptCommand() cli.Command {
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
						Name:        "in",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out",
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
						Name:        "in",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(aws.NewAmazonKMS())
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
				Name:  "gcp",
				Usage: "Encrypts files and/or strings with GCP KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
					cli.StringFlag{
						Name:        "project",
						Value:       "",
						Usage:       "the GCP project id for Cloud KMS",
						Destination: &gcpProject,
					},
					cli.StringFlag{
						Name:        "location",
						Value:       "",
						Usage:       "the location for project and Cloud KMS",
						Destination: &gcpLocation,
					},
					cli.StringFlag{
						Name:        "keyring",
						Value:       "",
						Usage:       "the key ring name",
						Destination: &gcpKeyring,
					},
					cli.StringFlag{
						Name:        "key",
						Value:       "",
						Usage:       "the cryptographic key name",
						Destination: &gcpKey,
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(gcp.NewGoogleKMS())
					params := map[string]interface{}{
						gcp.ProjectId: gcpProject,
						gcp.Location:  gcpLocation,
						gcp.KeyRing:   gcpKeyring,
						gcp.Key:       gcpKey,
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

func decryptCommand() cli.Command {
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
						Name:        "in",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out",
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
						Name:        "in",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(aws.NewAmazonKMS())
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
				Name:  "gcp",
				Usage: "Decrypts files and/or strings with GCP KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputPath,
					},
					cli.StringFlag{
						Name:        "out",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputPath,
					},
					cli.StringFlag{
						Name:        "project",
						Value:       "",
						Usage:       "the GCP project id for Cloud KMS",
						Destination: &gcpProject,
					},
					cli.StringFlag{
						Name:        "location",
						Value:       "",
						Usage:       "the location for project and Cloud KMS",
						Destination: &gcpLocation,
					},
					cli.StringFlag{
						Name:        "keyring",
						Value:       "",
						Usage:       "the key ring name",
						Destination: &gcpKeyring,
					},
					cli.StringFlag{
						Name:        "key",
						Value:       "",
						Usage:       "the cryptographic key name",
						Destination: &gcpKey,
					},
				},
				Action: func(c *cli.Context) error {
					crypt := crypto.NewCrypt(gcp.NewGoogleKMS())
					params := map[string]interface{}{
						gcp.ProjectId: gcpProject,
						gcp.Location:  gcpLocation,
						gcp.KeyRing:   gcpKeyring,
						gcp.Key:       gcpKey,
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
