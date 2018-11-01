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

### Encrypting with GCP KMS

    NAME:
       crypt encrypt gcp - Encrypts files and/or strings with GCP KMS

    USAGE:
       crypt encrypt gcp [command options] [arguments...]

    OPTIONS:
       --project value   the GCP project id for Cloud KMS
       --location value  the location for project and Cloud KMS
       --keyring value   the key ring name
       --key value       the cryptographic key name

## Development

    mkdir $GOPATH/src/github.com/VirtusLab/
    git clone

    go get -u github.com/golang/dep/cmd/dep

    export PATH=$PATH:$GOPATH/bin
    cd $GOPATH/src/github.com/VirtusLab/crypt
    go build

## The name

We believe in obvious names. It encrypts and decrypts.
