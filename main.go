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

	"github.com/VirtusLab/go-extended/pkg/files"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
)

var (
	app                 *cli.App
	inputFile           string
	outputFile          string
	inputDir            string
	outputDir           string
	inputFileExtension  string
	outputFileExtension string

	// Azure kms
	azureVaultURL   string
	azureKey        string
	azureKeyVersion string

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
	// The geographical data center location where requests to AWS KMS are handled
	awsRegion string
	// The AWS API credentials profile to use
	awsProfile string
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
		_, _ = fmt.Fprintf(cli.ErrWriter, "There is no %q command.\n", command)
		cli.OsExiter(1)
	}
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		if isSubcommand {
			return err
		}

		_, _ = fmt.Fprintf(cli.ErrWriter, "WRONG: %v\n", err)
		return nil
	}
	cli.OsExiter = func(c int) {
		if c != 0 {
			logrus.Debugf("exiting with %d", c)
		}
		os.Exit(c)
	}

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(cli.ErrWriter, "ERROR: %v\n", err)
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
						Name:        "indir",
						Value:       "",
						Usage:       "the input directory, can't be used with --out",
						Destination: &inputDir,
					},
					cli.StringFlag{
						Name:        "outdir",
						Value:       "",
						Usage:       "the output directory, the same as --outdir if empty, can't be used with --in",
						Destination: &outputDir,
					},
					cli.StringFlag{
						Name:        "in-extension",
						Value:       "",
						Usage:       "the extension of input file, used only with --indir",
						Destination: &inputFileExtension,
					},
					cli.StringFlag{
						Name:        "out-extension",
						Value:       ".crypt",
						Usage:       "the extension of output file, used only with --indir",
						Destination: &outputFileExtension,
					},
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputFile,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputFile,
					},
					cli.StringFlag{
						Name:        "vaultURL",
						Value:       "",
						Usage:       "the Azure KeyVault URL",
						Destination: &azureVaultURL,
					},
					cli.StringFlag{
						Name:        "name",
						Value:       "",
						Usage:       "the Azure KeyVault key name",
						Destination: &azureKey,
					},
					cli.StringFlag{
						Name:        "version",
						Value:       "",
						Usage:       "the Azure KeyVault key version",
						Destination: &azureKeyVersion,
					},
				},
				Action: encryptAzure,
			},
			{
				Name:  "aws",
				Usage: "Encrypts files and/or strings with AWS KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputFile,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputFile,
					},
					cli.StringFlag{
						Name:        "region",
						Value:       "",
						Usage:       "the AWS region",
						Destination: &awsRegion, // FIXME #2 make this flag required
					},
					cli.StringFlag{
						Name:        "profile",
						Value:       aws.DefaultProfile,
						Usage:       "the AWS API credentials profile",
						Destination: &awsProfile,
					},
					cli.StringFlag{
						Name:        "key-id, kms, kms-alias",
						Value:       "",
						Usage:       "the Amazon Resource Name (ARN), alias name, or alias ARN for the customer master key",
						Destination: &awsKms, // FIXME #2 make this flag required
					},
				},
				Action: encryptAws,
			},
			{
				Name:  "gcp",
				Usage: "Encrypts files and/or strings with GCP KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputFile,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputFile,
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
				Action: encryptGcp,
			},
		},
	}
}

func action(c *cli.Context, crypt crypto.Crypt, singleFileFunc, directoryFunc func(crypt crypto.Crypt) error) error {
	if c.NArg() > 0 {
		return errors.Errorf("have not expected any arguments, got %d", c.NArg())
	}

	if len(inputDir) > 0 {
		if len(inputFile) > 0 {
			return errors.New("conflict, --in can't be used with --indir or --outdir")
		}
		if len(outputFile) > 0 {
			return errors.New("conflict, --out can't be used with --indir or --outdir")
		}
		if len(inputFileExtension) == 0 && len(outputFileExtension) == 0 {
			return errors.New("--in-extension and --out-extension can't be empty")
		}
		if len(outputDir) == 0 {
			outputDir = inputDir
		}

		return directoryFunc(crypt)
	}

	if len(inputDir) > 0 {
		return errors.New("conflict, --indir can't be used with --in or --out")
	}
	if len(outputDir) > 0 {
		return errors.New("conflict, --outdir can't be used with --in or --out")
	}
	err := singleFileFunc(crypt)
	if err != nil && err == files.ErrExpectedStdin {
		return errors.New("expected either stdin, --indir or --in parameter, for usage use --help")
	}
	return err
}

func encryptAction(c *cli.Context, crypt crypto.Crypt) error {
	if len(inputDir) > 0 {
		if len(outputFileExtension) == 0 {
			return fmt.Errorf("--out-extension can't be empty")
		}
	}

	return action(c, crypt, encryptSingleFile, encryptDirectory)
}

func encryptSingleFile(crypt crypto.Crypt) error {
	return crypt.EncryptFile(inputFile, outputFile)
}

func encryptDirectory(crypt crypto.Crypt) error {
	return crypt.EncryptFiles(inputDir, outputDir, inputFileExtension, outputFileExtension)
}

func encryptAzure(c *cli.Context) error {
	azr, err := azure.New(azureVaultURL, azureKey, azureKeyVersion)
	if err != nil {
		return err
	}
	return encryptAction(c, crypto.New(azr))
}

func encryptAws(_ *cli.Context) error {
	amazon := aws.New(awsKms, awsRegion, awsProfile)
	crypt := crypto.New(amazon)
	err := crypt.EncryptFile(inputFile, outputFile)
	if err != nil {
		return err
	}
	return nil
}

func encryptGcp(_ *cli.Context) error {
	google := gcp.New(gcpProject, gcpLocation, gcpKeyring, gcpKey)
	crypt := crypto.New(google)
	err := crypt.EncryptFile(inputFile, outputFile)
	if err != nil {
		return err
	}
	return nil
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
						Name:        "indir",
						Value:       "",
						Usage:       "the input directory, can't be used with --out",
						Destination: &inputDir,
					},
					cli.StringFlag{
						Name:        "outdir",
						Value:       "",
						Usage:       "the output directory, the same as --outdir if empty, can't be used with --in",
						Destination: &outputDir,
					},
					cli.StringFlag{
						Name:        "in-extension",
						Value:       ".crypt",
						Usage:       "the extension of input file, used only with --indir",
						Destination: &inputFileExtension,
					},
					cli.StringFlag{
						Name:        "out-extension",
						Value:       "",
						Usage:       "the extension of output file, used only with --indir",
						Destination: &outputFileExtension,
					},
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputFile,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputFile,
					},
					cli.StringFlag{
						Name:        "vaultURL",
						Value:       "",
						Usage:       "Azure vault URL",
						Destination: &azureVaultURL,
					},
					cli.StringFlag{
						Name:        "name",
						Value:       "",
						Usage:       "the key name",
						Destination: &azureKey,
					},
					cli.StringFlag{
						Name:        "version",
						Value:       "",
						Usage:       "the key version",
						Destination: &azureKeyVersion,
					},
				},
				Action: decryptAzure,
			},
			{
				Name:  "aws",
				Usage: "Decrypts files and/or strings with AmazonKMS KMS",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "in, input",
						Value:       "",
						Usage:       "the input file to decrypt, stdin if empty",
						Destination: &inputFile,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputFile,
					},
					cli.StringFlag{
						Name:        "region",
						Value:       "",
						Usage:       "(required) the AWS region",
						Destination: &awsRegion,
					},
					cli.StringFlag{
						Name:        "profile",
						Value:       aws.DefaultProfile,
						Usage:       "the AWS API credentials profile",
						Destination: &awsProfile,
					},
				},
				Action: func(c *cli.Context) error {
					if len(awsRegion) == 0 {
						return errors.New("pass the AWS region")
					}
					return decryptAws(c)
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
						Destination: &inputFile,
					},
					cli.StringFlag{
						Name:        "out, output",
						Value:       "",
						Usage:       "the output file, stdout if empty",
						Destination: &outputFile,
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
				Action: decryptGcp,
			},
		},
	}
}

func decryptAction(c *cli.Context, crypt crypto.Crypt) error {
	if len(inputDir) > 0 {
		if len(inputFileExtension) == 0 {
			return fmt.Errorf("--in-extension can't be empty")
		}
	}

	return action(c, crypt, decryptSingleFile, decryptDirectory)
}

func decryptSingleFile(crypt crypto.Crypt) error {
	return crypt.DecryptFile(inputFile, outputFile)
}

func decryptDirectory(crypt crypto.Crypt) error {
	return crypt.DecryptFiles(inputDir, outputDir, inputFileExtension, outputFileExtension)
}

func decryptAzure(c *cli.Context) error {
	azr, err := azure.New(azureVaultURL, azureKey, azureKeyVersion)
	if err != nil {
		return err
	}
	return decryptAction(c, crypto.New(azr))
}

func decryptAws(_ *cli.Context) error {
	amazon := aws.New(awsKms, awsRegion, awsProfile)
	crypt := crypto.New(amazon)
	err := crypt.DecryptFile(inputFile, outputFile)
	if err != nil {
		return err
	}
	return nil
}

func decryptGcp(_ *cli.Context) error {
	googleKms := gcp.New(gcpProject, gcpLocation, gcpKeyring, gcpKey)
	crypt := crypto.New(googleKms)
	err := crypt.DecryptFile(inputFile, outputFile)
	if err != nil {
		return err
	}
	return nil
}
