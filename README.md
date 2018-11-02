# crypt

Universal cryptographic tool with AWS KMS, GCP KMS and Azure Key Vault support.

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

### Encrypting/Decrypting with GCP KMS

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
       --in value        the input file to decrypt, stdin if empty
       --out value       the output file, stdout if empty
       --project value   the GCP project id for Cloud KMS
       --location value  the location for project and Cloud KMS
       --keyring value   the key ring name
       --key value       the cryptographic key name

### Useful links

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
