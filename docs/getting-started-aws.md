# Getting Started with AWS 

This document describes step by step guide how to use `crypt` with AWS KMS.


AWS KMS uses client from [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/).
You can either run `aws configure` (if you don't have `~/.aws/credentials` already) 
or set [environment variables](https://docs.aws.amazon.com/sdk-for-go/api/aws/session).
To set AWS profile use `--profile` parameter.

## Encryption

Example usage with file:

    $ echo "top secret" > file.txt
    $ crypt encrypt aws \
        --in file.txt \
        --out file.enc \
        --region eu-west-1 \
        --kms alias/test
    $ crypt decrypt aws \
        --in file.enc \
        --out file.dec \
        --region eu-west-1

## Decryption

Example usage with `stdin`:

    $ echo "top secret" | crypt encrypt aws \
        --out file.enc \
        --region eu-west-1 \
        --kms alias/test