# crypt

Universal cryptographic tool with AWS KMS, GCP KMS and Azure Key Vault support.

* [Installation](README.md#installation)
  * [Binaries](README.md#binaries)
  * [Via Go](README.md#via-go)
* [Usage](README.md#usage)
  * [Encrypting and Decrypting with AWS KMS](README.md#encrypting-and-decrypting-with-aws-kms)
    * [Useful links](README.md#useful-links)
  * [Encrypting and Decrypting with GCP KMS](README.md#encrypting-and-decrypting-with-gcp-kms)
    * [Useful links](README.md#useful-links-1)
* [Development](README.md#development)
* [The Name](README.md#the-name)

## Installation

#### Binaries

For binaries please visit the [Releases Page](https://github.com/VirtusLab/crypt/releases).

#### Via Go

```console
$ go get github.com/VirtusLab/crypt
```

## Usage

    NAME:
       crypt - Universal cryptographic tool with AWS KMS, GCP KMS and Azure Key Vault support

    USAGE:
       crypt [global options] command [command options] [arguments...]

    VERSION:
       v0.0.1

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

### Encrypting and Decrypting with AWS KMS

`crypt` uses standard AWS client from official [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/).
I can be configured using `~/.aws/credentials` and `~/.aws/config` or using environment variables.

    # Access Key ID
    AWS_ACCESS_KEY_ID=AKID
    AWS_ACCESS_KEY=AKID # only read if AWS_ACCESS_KEY_ID is not set.

    # Secret Access Key
    AWS_SECRET_ACCESS_KEY=SECRET
    AWS_SECRET_KEY=SECRET=SECRET # only read if AWS_SECRET_ACCESS_KEY is not set.

    # Session Token
    AWS_SESSION_TOKEN=TOKEN

    AWS_REGION=us-east-1

    # AWS_DEFAULT_REGION is only read if AWS_SDK_LOAD_CONFIG is also set,
    # and AWS_REGION is not also set.
    AWS_DEFAULT_REGION=us-east-1

Note that assuming AWS IAM role is not within crypt scope - you must do it yourself and provide valid credentials.

For more details take a look at [Package session provides configuration for the SDK's service clients.](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/).

Encryption and decryption example

    $ echo "top secret" > file.txt
    $ crypt en aws --in file.txt --out file.enc --region eu-west-1 --kms alias/test
    $ crypt de aws --in file.enc --out file.dec --region eu-west-1

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

### Encrypting and Decrypting with GCP KMS

Set up Cloud KMS and corresponding service account

    $ gcloud init
    $ gcloud kms keyrings create test --location global
    $ gcloud kms keys create quickstart --location global --keyring test --purpose encryption
    $ gcloud kms keys list --location global --keyring test
    $ export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"

Encryption and decryption example

    $ echo "top secret" > file.txt
    $ crypt en gcp --in file.txt --out file.enc --project lunar-compiler-123456 --location global --keyring test --key quickstart
    $ crypt de gcp --in file.enc --out file.dec --project lunar-compiler-123456 --location global --keyring test --key quickstart

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
- [Setting Up Authentication for Server to Server Production Applications](https://cloud.google.com/docs/authentication/production)
- [Cloud KMS - Quickstart](https://cloud.google.com/kms/docs/quickstart)
- [Cloud KMS - Encrypting and Decrypting Data](https://cloud.google.com/kms/docs/encrypt-decrypt#kms-howto-encrypt-go)

## Development

    mkdir $GOPATH/src/github.com/VirtusLab/
    git clone

    go get -u github.com/golang/dep/cmd/dep

    export PATH=$PATH:$GOPATH/bin
    cd $GOPATH/src/github.com/VirtusLab/crypt
    go build

## The name

We believe in obvious names. It encrypts and decrypts.
