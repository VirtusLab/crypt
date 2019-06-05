# Getting Started with GCP 

This document describes step by step guide how to use `crypt` with GCP KMS.

GCP KMS uses [DefaultClient](https://godoc.org/golang.org/x/oauth2/google#DefaultClient) from [Google Cloud Client Libraries for Go](https://github.com/GoogleCloudPlatform/google-cloud-go).
You can either run `gcloud auth application-default login` or set `GOOGLE_APPLICATION_CREDENTIALS` environment variable which points to the file with valid service account.

## Encryption

### stdin

Example usage with `stdin`:

```console
echo "top secret" | crypt encrypt gcp \
    --out file.enc \
    --project lunar-compiler-123456 \
    --location global \
    --keyring test \
    --key quickstart
```

### Single file

```console
echo "top secret" > file.txt

crypt encrypt gcp \
    --in file.txt \
    --out file.enc \
    --project lunar-compiler-123456 \
    --location global \
    --keyring test \
    --key quickstart
```

### Files in directory

```console
crypt encrypt gcp \
    --indir to-encrypt \
    --outdir encrypted \
    --project lunar-compiler-123456 \
    --location global \
    --keyring test \
    --key quickstart
```

In directory `encrypted` you should have encrypted files with extension `.crypt`.

## Decryption

### Single file

```console
crypt decrypt gcp \
    --in file.enc \
    --out file.dec \
    --project lunar-compiler-123456 \
    --location global \
    --keyring test \
    --key quickstart
```

 ### Files in directory
 
 ```console
crypt decrypt gcp \
    --indir to-decrypt \
    --outdir decrypted \
    --project lunar-compiler-123456 \
    --location global \
    --keyring test \
    --key quickstart
 ```
 
 This command scan all files with extension `.crypt` in folder `to-decrypt` and decrypts it to `decrypted` folder.
 Decrypted files doesnt't have `.crypt` extension.