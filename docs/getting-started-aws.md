# Getting Started with AWS 

This document describes step by step guide how to use `crypt` with AWS KMS.


AWS KMS uses client from [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/).
You can either run `aws configure` (if you don't have `~/.aws/credentials` already) 
or set [environment variables](https://docs.aws.amazon.com/sdk-for-go/api/aws/session).
To set AWS profile use `--profile` parameter.

## Encryption

### stdin

Example usage with `stdin`:

```console
echo "top secret" | crypt encrypt aws \
    --out file.enc \
    --region eu-west-1 \
    --kms alias/test
```

### Single file

```console
echo "top secret" > file.txt

crypt encrypt aws \
    --in file.txt \
    --out file.enc \
    --region eu-west-1 \
    --kms alias/test
```

### Files in directory

```console
crypt encrypt aws \
    --indir to-encrypt \
    --outdir encrypted \
    --region eu-west-1 \
    --kms alias/test
```

In directory `encrypted` you should have encrypted files with extension `.crypt`.

## Decryption

### Single file

```console
crypt decrypt aws \
    --in file.enc \
    --out file.dec \
    --region eu-west-1
```

 ### Files in directory
 
 ```console
crypt decrypt aws \
    --indir to-decrypt \
    --outdir decrypted \
    --region eu-west-1
 ```
 
 This command scan all files with extension `.crypt` in folder `to-decrypt` and decrypts it to `decrypted` folder.
 Decrypted files doesnt't have `.crypt` extension.