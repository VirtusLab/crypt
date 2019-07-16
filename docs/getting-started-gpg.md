# Getting started with GPG

This document describes step by step guide how to use `crypt` with GPG (GnuPG).

Current implementation supports encryption and decryption using armored keys only.
Support for keyring is planned in the next release.

## Generate keys

Generate a new GPG key pair:

    gpg --gen-key

## Encryption using local GPG Public Key

Identify your public key by running:

    gpg --list-keys

Run this command to export your GPG Public Key (armored):
    
    gpg --armor --export --output my_pubkey.gpg $ID
    
Encrypt:

    echo test | crypt encrypt gpg --public-key my-public-key.gpg --out test.enc 
       
## Encryption using GPG Public Key from the Key Server

    echo test | ./crypt enc gpg --keyserver keyserver.ubuntu.com --key-id 51716619E084DAB9 --out test.enc
       
## Decryption using local GPG Private Key

Identify your public key by running:

    gpg --list-secret-keys

Run this command to export your GPG Private Key (armored):
    
    gpg --export-secret-keys --armor $ID > my-private-key.asc   
    
Decrypt:

    crypt decrypt --private-key my-private-key.asc --in test.enc

Decrypt with passphrase for the private key: 

    crypt decrypt --private-key my-private-key.asc --passphrase <passphrase> --in test.enc

Also, you can set `GPG_PASSPHRASE` environment variable:
    
    export GPG_PASSPHRASE=<passphrase>
    crypt decrypt --private-key my-private-key.asc --in test.enc
    
## Current limitations

This section describes current limitations and planned fixes in the next releases. 

### No support for keyring

In the GnuPG 2.1 keyring format has changed, see [What’s new in GnuPG 2.1](https://www.gnupg.org/faq/whats-new-in-2.1.html) for more details.

Use `--export-secret-keys` to export the secret keys , or `--export` to export your public keys.

Upstream issues:
- https://github.com/helm/helm/issues/2843#issuecomment-424926564
- https://github.com/golang/go/issues/29082