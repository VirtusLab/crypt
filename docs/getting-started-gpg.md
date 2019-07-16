# Getting started with GPG

This document describes step by step guide how to use `crypt` with GPG (GnuPG).

Current implementation supports encryption and decryption using armored keys only.
Support for keyring is planned in the next release.

## Generate keys

Generate a new GPG key pair:

    gpg --gen-key

## Encryption using GPG Public Key

Identify your public key by running:

    gpg --list-keys

Run this command to export your GPG Public Key (armored):
    
    gpg --armor --export --output my_pubkey.gpg $ID
    
Encrypt:

    echo test | crypt encrypt gpg --public-key my-public-key.gpg --out test.enc 
       
## Decryption using GPG Private Key

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
    
