# Getting Started with Azure 

This document describes step by step guide how to use `crypt` with Azure Key Vault.

## Set up Azure credentials

Run:

```console
az login
```

Specify Azure Subscription:

```console
az account list
az account set --subscription <subscription name or ID>
```

## Create Azure Key Vault

Create a new Resource Group:

```console
az group create -n "example-rg" -l "westeurope"
```

Create Azure Key Vault:

```console
NEW_UUID=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 5 | head -n 1)
KEY_VAULT_NAME="example-kv-${NEW_UUID}"
az keyvault create --name "${KEY_VAULT_NAME}" --resource-group "example-rg" --location "westeurope"
```

**Azure Key Vault name has to be globally unique.**

## Generate Encryption Key

Create a software-protected key:

```console
az keyvault key create --vault-name "${KEY_VAULT_NAME}" --name "example-key2" --protection software
```

**Please enable `encrypt` and `decrypt` Access policies in Azure Key Vault.**

## Encryption

### Single file

Run the following command:

```console
echo "top secret" > file.txt

crypt encrypt azure \
        --in file.txt \
        --out file.enc \
        --vaultURL https://${KEY_VAULT_NAME}.vault.azure.net \
        --name example-key \
        --version b3d715d803bb4e3fb07d12701a101dcd
```

You should have output similar to this:

```console
INFO Encryption succeeded   key=example-key keyVersion=b3d715d803bb4e3fb07d12701a101dcd
```

### Files in directory

```console
crypt encrypt azure \
        --indir to-encrypt \
        --outdir encrypted \
        --vaultURL https://${KEY_VAULT_NAME}.vault.azure.net \
        --name example-key \
        --version b3d715d803bb4e3fb07d12701a101dcd
```

In directory `encrypted` you should have encrypted files with extension `.crypt`.

## Decryption

### Single

Run the following command:
 
```console
crypt decrypt azure \
    --in file.txt.enc \
    --out file2.txt \
    --vaultURL=https://${KEY_VAULT_NAME}.vault.azure.net \
    --name=example-key \
    --version b3d715d803bb4e3fb07d12701a101dcd
```

You should have output similar to this:

```console
INFO Decryption succeeded   key=example-key keyVersion=b3d715d803bb4e3fb07d12701a101dcd
```
 
 ### Files in directory
 
 ```console
 crypt decrypt azure \
         --indir to-decrypt \
         --outdir decrypted
 ```
 
 This command scan all files with extension `.crypt` in folder `to-decrypt` and decrypts it to `decrypted` folder.
 Decrypted files doesnt't have `.crypt` extension.