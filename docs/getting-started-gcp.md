# Getting Started with GCP 

This document describes step by step guide how to use `crypt` with GCP KMS.

GCP KMS uses [DefaultClient](https://godoc.org/golang.org/x/oauth2/google#DefaultClient) from [Google Cloud Client Libraries for Go](https://github.com/GoogleCloudPlatform/google-cloud-go).
You can either run `gcloud auth application-default login` or set `GOOGLE_APPLICATION_CREDENTIALS` environment variable which points to the file with valid service account.

# Encryption

Example usage with file:

    $ echo "top secret" > file.txt
    $ crypt encrypt gcp \
        --in file.txt \
        --out file.enc \
        --project lunar-compiler-123456 \
        --location global \
        --keyring test \
        --key quickstart
    $ crypt decrypt gcp \
        --in file.enc \
        --out file.dec \
        --project lunar-compiler-123456 \
        --location global \
        --keyring test \
        --key quickstart
        
# Decryption

Example usage with `stdin`:

    $ echo "top secret" | crypt encrypt gcp \
        --out file.enc \
        --project lunar-compiler-123456 \
        --location global \
        --keyring test \
        --key quickstart