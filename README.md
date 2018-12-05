# crypt

[![Version](https://img.shields.io/badge/version-v0.0.4-brightgreen.svg)](https://github.com/VirtusLab/crypt/releases/tag/v0.0.4)
[![Travis CI](https://img.shields.io/travis/VirtusLab/crypt.svg)](https://travis-ci.org/VirtusLab/crypt)
[![Github All Releases](https://img.shields.io/github/downloads/VirtusLab/crypt/total.svg)](https://github.com/VirtusLab/crypt/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/VirtusLab/crypt "Go Report Card")](https://goreportcard.com/report/github.com/VirtusLab/crypt)

Universal cryptographic tool with AWS KMS, GCP KMS and Azure Key Vault support.

* [Installation](README.md#installation)
  * [Binaries](README.md#binaries)
  * [Via Go](README.md#via-go)
* [Usage](README.md#usage)
  * [Encryption using AWS KMS](README.md#encryption-using-aws-kms)
  * [Encryption using GCP KMS](README.md#encryption-using-gcp-kms)
  * [Encryption using Azure Key Vault](README.md#encryption-using-azure-key-vault)
* [Development](README.md#development)
* [Contribution](README.md#contribution)


## Installation

#### Binaries

For binaries please visit the [Releases Page](https://github.com/VirtusLab/crypt/releases).

#### Via Go

    $ go get github.com/VirtusLab/crypt

## Usage

    NAME:
       crypt - Universal cryptographic tool with AWS KMS, GCP KMS and Azure Key Vault support

    USAGE:
       crypt [global options] command [command options] [arguments...]

    VERSION:
       v0.0.4

    AUTHOR:
       VirtusLab

    COMMANDS:
         encrypt, enc, en, e  Encrypts files and/or strings
         decrypt, dec, de, d  Decrypts files and/or strings
         help, h              Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --debug, -d    run in debug mode
       --help, -h     show help
       --version, -v  print the version

### Encryption using AWS KMS

AWS KMS uses client from [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/).
You can either run `aws configure` (if you don't have `~/.aws/credentials` already) or set [environment variables](https://docs.aws.amazon.com/sdk-for-go/api/aws/session).

Example usage with file:

    $ echo "top secret" > file.txt
    $ crypt encrypt aws --in file.txt --out file.enc --region eu-west-1 --kms alias/test
    $ crypt decrypt aws --in file.enc --out file.dec --region eu-west-1

Example usage with `stdin`:

    $ echo "top secret" | crypt encrypt aws --out file.enc --region eu-west-1 --kms alias/test

### Encryption using with GCP KMS

GCP KMS uses [DefaultClient](https://godoc.org/golang.org/x/oauth2/google#DefaultClient) from [Google Cloud Client Libraries for Go](https://github.com/GoogleCloudPlatform/google-cloud-go).
You can either run `gcloud auth application-default login` or set `GOOGLE_APPLICATION_CREDENTIALS` environment variable which points to the file with valid service account.

Example usage with file:

    $ echo "top secret" > file.txt
    $ crypt encrypt gcp --in file.txt --out file.enc --project lunar-compiler-123456 --location global --keyring test --key quickstart
    $ crypt decrypt gcp --in file.enc --out file.dec --project lunar-compiler-123456 --location global --keyring test --key quickstart

Example usage with `stdin`:

    $ echo "top secret" | crypt encrypt gcp --out file.enc --project lunar-compiler-123456 --location global --keyring test --key quickstart

### Encryption using Azure Key Vault

Azure Key Vault uses [NewAuthorizerFromEnvironment](https://github.com/Azure/azure-sdk-for-go) from [Microsoft Azure SDK for go](https://github.com/Azure/azure-sdk-for-go).
Run `az login` to get your Azure credentials.

Example usage with file:

    $ echo "top secret" > file.txt
    $ crypt encrypt gcp --in file.txt --out file.enc --vaultURL https://example-vault.vault.azure.net --name global --version 77ea..
    $ crypt decrypt gcp --in file.enc --out file.dec --vaultURL https://example-vault.vault.azure.net --name global --version 77ea..

Example usage with `stdin`:

    $ echo "top secret" | crypt encrypt gcp --out file.enc --project lunar-compiler-123456 --location global --keyring test --key quickstart

## Development

    mkdir $GOPATH/src/github.com/VirtusLab/
    git clone

    go get -u github.com/golang/dep/cmd/dep

    export PATH=$PATH:$GOPATH/bin
    cd $GOPATH/src/github.com/VirtusLab/crypt
    make all

### Testing

    make test

### Integration testing

Update properties in `config.env` and run:

    make integrationtest
    
## Contribution

Feel free to file [issues](https://github.com/VirtusLab/crypt/issues) or [pull requests](https://github.com/VirtusLab/crypt/pulls).    