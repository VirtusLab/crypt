# crypt

[![Version](https://img.shields.io/badge/version-v0.0.2-brightgreen.svg)](https://github.com/VirtusLab/crypt/releases/tag/v0.0.2)
[![Travis CI](https://img.shields.io/travis/VirtusLab/crypt.svg)](https://travis-ci.org/VirtusLab/crypt)
[![Github All Releases](https://img.shields.io/github/downloads/VirtusLab/crypt/total.svg)](https://github.com/VirtusLab/crypt/releases)

Universal cryptographic tool with AWS KMS, GCP KMS and Azure Key Vault support.

* [Installation](README.md#installation)
  * [Binaries](README.md#binaries)
  * [Via Go](README.md#via-go)
* [Usage](README.md#usage)
  * [Encryption using AWS KMS](README.md#encryption-using-aws-kms)
    * [Examples](README.md#examples)
    * [Useful links](README.md#useful-links)
  * [Encryption using GCP KMS](README.md#encryption-using-gcp-kms)
    * [Examples](README.md#examples-1)
    * [Useful links](README.md#useful-links-1)
  * [Encryption using Azure Key Vault](README.md#encryption-using-azure-key-vault)
    * [Examples](README.md#examples-1)
    * [Useful links](README.md#useful-links-1)
* [Development](README.md#development)
* [The Name](README.md#the-name)

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
       v0.0.2

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

`crypt` uses client from [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/).
You can either run `aws configure` (if you don't have ~/.aws/credentials already) or set [environment variables](https://docs.aws.amazon.com/sdk-for-go/api/aws/session).

#### Examples

Example usage with file:

    $ echo "top secret" > file.txt
    $ crypt en aws --in file.txt --out file.enc --region eu-west-1 --kms alias/test
    $ crypt de aws --in file.enc --out file.dec --region eu-west-1

Example usage with stdin:

    $ echo "top secret" | crypt en aws --out file.enc --region eu-west-1 --kms alias/test

For more details run `crypt en aws --help` or `crypt de aws --help`

    NAME:
       crypt encrypt aws - Encrypts files and/or strings with AWS KMS

    USAGE:
       crypt encrypt aws [command options] [arguments...]

    OPTIONS:
       --in value, --input value                       the input file to decrypt, stdin if empty
       --out value, --output value                     the output file, stdout if empty
       --region value                                  the AWS region
       --key-id value, --kms value, --kms-alias value  the Amazon Resource Name (ARN), alias name, or alias ARN for the customer master key

#### Useful links

- [AWS Key Management Service (KMS)](https://aws.amazon.com/kms/)
- [AWS KMS Creating Keys](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html)
- [AWS KMS CLI command reference](https://docs.aws.amazon.com/cli/latest/reference/kms/index.html#cli-aws-kms)

### Encryption using with GCP KMS

`crypt` uses [DefaultClient](https://godoc.org/golang.org/x/oauth2/google#DefaultClient) from official [Google Cloud Client Libraries for Go](https://github.com/GoogleCloudPlatform/google-cloud-go).
You can either run `gcloud auth application-default login` or set `GOOGLE_APPLICATION_CREDENTIALS` environment variable which points to the file with valid service account.

#### Examples

Example usage with file:

    $ echo "top secret" > file.txt
    $ crypt en gcp --in file.txt --out file.enc --project lunar-compiler-123456 --location global --keyring test --key quickstart
    $ crypt de gcp --in file.enc --out file.dec --project lunar-compiler-123456 --location global --keyring test --key quickstart

Example usage with stdin:

    $ echo "top secret" | crypt en gcp --out file.enc --project lunar-compiler-123456 --location global --keyring test --key quickstart

For more details run `crypt en gcp --help` or `crypt de gcp --help`

    NAME:
       crypt encrypt gcp - Encrypts files and/or strings with GCP KMS

    USAGE:
       crypt encrypt gcp [command options] [arguments...]

    OPTIONS:
       --in value, --input value    the input file to decrypt, stdin if empty
       --out value, --output value  the output file, stdout if empty
       --project value              the GCP project id for Cloud KMS
       --location value             the location for project and Cloud KMS
       --keyring value              the key ring name
       --key value                  the cryptographic key name


#### Useful links

- [Installing Google Cloud SDK](https://cloud.google.com/sdk/install)
- [gcloud auth application-default login](https://cloud.google.com/sdk/gcloud/reference/auth/application-default/login)
- [Setting Up Authentication for Server to Server Production Applications](https://cloud.google.com/docs/authentication/production)
- [Cloud KMS - Quickstart](https://cloud.google.com/kms/docs/quickstart)
- [Cloud KMS - Encrypting and Decrypting Data](https://cloud.google.com/kms/docs/encrypt-decrypt#kms-howto-encrypt-go)

### Encryption using Azure Key Vault

Not supported yet. Stay tuned.

## Development

    mkdir $GOPATH/src/github.com/VirtusLab/
    git clone

    go get -u github.com/golang/dep/cmd/dep

    export PATH=$PATH:$GOPATH/bin
    cd $GOPATH/src/github.com/VirtusLab/crypt
    make all

## The name

We believe in obvious names. It encrypts and decrypts.
