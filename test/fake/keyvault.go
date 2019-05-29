package fake

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

// KeyVaultAPIClient is a fake KeyVault API client provider
type KeyVaultAPIClient struct {
}

// BackupKey unused
func (KeyVaultAPIClient) BackupKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.BackupKeyResult, err error) {
	panic("implement me")
}

// BackupSecret unused
func (KeyVaultAPIClient) BackupSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.BackupSecretResult, err error) {
	panic("implement me")
}

// CreateCertificate unused
func (KeyVaultAPIClient) CreateCertificate(ctx context.Context, vaultBaseURL string, certificateName string, parameters keyvault.CertificateCreateParameters) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// CreateKey unused
func (KeyVaultAPIClient) CreateKey(ctx context.Context, vaultBaseURL string, keyName string, parameters keyvault.KeyCreateParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// Decrypt unused
func (KeyVaultAPIClient) Decrypt(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	result.Result = parameters.Value
	return
}

// DeleteCertificate unused
func (KeyVaultAPIClient) DeleteCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.DeletedCertificateBundle, err error) {
	panic("implement me")
}

// DeleteCertificateContacts unused
func (KeyVaultAPIClient) DeleteCertificateContacts(ctx context.Context, vaultBaseURL string) (result keyvault.Contacts, err error) {
	panic("implement me")
}

// DeleteCertificateIssuer unused
func (KeyVaultAPIClient) DeleteCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// DeleteCertificateOperation unused
func (KeyVaultAPIClient) DeleteCertificateOperation(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// DeleteKey unused
func (KeyVaultAPIClient) DeleteKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.DeletedKeyBundle, err error) {
	panic("implement me")
}

// DeleteSasDefinition unused
func (KeyVaultAPIClient) DeleteSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// DeleteSecret unused
func (KeyVaultAPIClient) DeleteSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.DeletedSecretBundle, err error) {
	panic("implement me")
}

// DeleteStorageAccount unused
func (KeyVaultAPIClient) DeleteStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// Encrypt unused
func (KeyVaultAPIClient) Encrypt(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	result.Result = parameters.Value
	return
}

// GetCertificate unused
func (KeyVaultAPIClient) GetCertificate(ctx context.Context, vaultBaseURL string, certificateName string, certificateVersion string) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// GetCertificateContacts unused
func (KeyVaultAPIClient) GetCertificateContacts(ctx context.Context, vaultBaseURL string) (result keyvault.Contacts, err error) {
	panic("implement me")
}

// GetCertificateIssuer unused
func (KeyVaultAPIClient) GetCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// GetCertificateIssuers unused
func (KeyVaultAPIClient) GetCertificateIssuers(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.CertificateIssuerListResultPage, err error) {
	panic("implement me")
}

// GetCertificateOperation unused
func (KeyVaultAPIClient) GetCertificateOperation(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// GetCertificatePolicy unused
func (KeyVaultAPIClient) GetCertificatePolicy(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificatePolicy, err error) {
	panic("implement me")
}

// GetCertificates unused
func (KeyVaultAPIClient) GetCertificates(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.CertificateListResultPage, err error) {
	panic("implement me")
}

// GetCertificateVersions unused
func (KeyVaultAPIClient) GetCertificateVersions(ctx context.Context, vaultBaseURL string, certificateName string, maxresults *int32) (result keyvault.CertificateListResultPage, err error) {
	panic("implement me")
}

// GetDeletedCertificate unused
func (KeyVaultAPIClient) GetDeletedCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.DeletedCertificateBundle, err error) {
	panic("implement me")
}

// GetDeletedCertificates unused
func (KeyVaultAPIClient) GetDeletedCertificates(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedCertificateListResultPage, err error) {
	panic("implement me")
}

// GetDeletedKey unused
func (KeyVaultAPIClient) GetDeletedKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.DeletedKeyBundle, err error) {
	panic("implement me")
}

// GetDeletedKeys unused
func (KeyVaultAPIClient) GetDeletedKeys(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedKeyListResultPage, err error) {
	panic("implement me")
}

// GetDeletedSecret unused
func (KeyVaultAPIClient) GetDeletedSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.DeletedSecretBundle, err error) {
	panic("implement me")
}

// GetDeletedSecrets unused
func (KeyVaultAPIClient) GetDeletedSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.DeletedSecretListResultPage, err error) {
	panic("implement me")
}

// GetKey unused
func (KeyVaultAPIClient) GetKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// GetKeys unused
func (KeyVaultAPIClient) GetKeys(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.KeyListResultPage, err error) {
	panic("implement me")
}

// GetKeyVersions unused
func (KeyVaultAPIClient) GetKeyVersions(ctx context.Context, vaultBaseURL string, keyName string, maxresults *int32) (result keyvault.KeyListResultPage, err error) {
	panic("implement me")
}

// GetSasDefinition unused
func (KeyVaultAPIClient) GetSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// GetSasDefinitions unused
func (KeyVaultAPIClient) GetSasDefinitions(ctx context.Context, vaultBaseURL string, storageAccountName string, maxresults *int32) (result keyvault.SasDefinitionListResultPage, err error) {
	panic("implement me")
}

// GetSecret unused
func (KeyVaultAPIClient) GetSecret(ctx context.Context, vaultBaseURL string, secretName string, secretVersion string) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// GetSecrets unused
func (KeyVaultAPIClient) GetSecrets(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.SecretListResultPage, err error) {
	panic("implement me")
}

// GetSecretVersions unused
func (KeyVaultAPIClient) GetSecretVersions(ctx context.Context, vaultBaseURL string, secretName string, maxresults *int32) (result keyvault.SecretListResultPage, err error) {
	panic("implement me")
}

// GetStorageAccount unused
func (KeyVaultAPIClient) GetStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// GetStorageAccounts unused
func (KeyVaultAPIClient) GetStorageAccounts(ctx context.Context, vaultBaseURL string, maxresults *int32) (result keyvault.StorageListResultPage, err error) {
	panic("implement me")
}

// ImportCertificate unused
func (KeyVaultAPIClient) ImportCertificate(ctx context.Context, vaultBaseURL string, certificateName string, parameters keyvault.CertificateImportParameters) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// ImportKey unused
func (KeyVaultAPIClient) ImportKey(ctx context.Context, vaultBaseURL string, keyName string, parameters keyvault.KeyImportParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// MergeCertificate unused
func (KeyVaultAPIClient) MergeCertificate(ctx context.Context, vaultBaseURL string, certificateName string, parameters keyvault.CertificateMergeParameters) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// PurgeDeletedCertificate unused
func (KeyVaultAPIClient) PurgeDeletedCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result autorest.Response, err error) {
	panic("implement me")
}

// PurgeDeletedKey unused
func (KeyVaultAPIClient) PurgeDeletedKey(ctx context.Context, vaultBaseURL string, keyName string) (result autorest.Response, err error) {
	panic("implement me")
}

// PurgeDeletedSecret unused
func (KeyVaultAPIClient) PurgeDeletedSecret(ctx context.Context, vaultBaseURL string, secretName string) (result autorest.Response, err error) {
	panic("implement me")
}

// RecoverDeletedCertificate unused
func (KeyVaultAPIClient) RecoverDeletedCertificate(ctx context.Context, vaultBaseURL string, certificateName string) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// RecoverDeletedKey unused
func (KeyVaultAPIClient) RecoverDeletedKey(ctx context.Context, vaultBaseURL string, keyName string) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// RecoverDeletedSecret unused
func (KeyVaultAPIClient) RecoverDeletedSecret(ctx context.Context, vaultBaseURL string, secretName string) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// RegenerateStorageAccountKey unused
func (KeyVaultAPIClient) RegenerateStorageAccountKey(ctx context.Context, vaultBaseURL string, storageAccountName string, parameters keyvault.StorageAccountRegenerteKeyParameters) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// RestoreKey unused
func (KeyVaultAPIClient) RestoreKey(ctx context.Context, vaultBaseURL string, parameters keyvault.KeyRestoreParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// RestoreSecret unused
func (KeyVaultAPIClient) RestoreSecret(ctx context.Context, vaultBaseURL string, parameters keyvault.SecretRestoreParameters) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// SetCertificateContacts unused
func (KeyVaultAPIClient) SetCertificateContacts(ctx context.Context, vaultBaseURL string, contacts keyvault.Contacts) (result keyvault.Contacts, err error) {
	panic("implement me")
}

// SetCertificateIssuer unused
func (KeyVaultAPIClient) SetCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string, parameter keyvault.CertificateIssuerSetParameters) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// SetSasDefinition unused
func (KeyVaultAPIClient) SetSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string, parameters keyvault.SasDefinitionCreateParameters) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// SetSecret unused
func (KeyVaultAPIClient) SetSecret(ctx context.Context, vaultBaseURL string, secretName string, parameters keyvault.SecretSetParameters) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// SetStorageAccount unused
func (KeyVaultAPIClient) SetStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string, parameters keyvault.StorageAccountCreateParameters) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// Sign unused
func (KeyVaultAPIClient) Sign(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeySignParameters) (result keyvault.KeyOperationResult, err error) {
	panic("implement me")
}

// UnwrapKey unused
func (KeyVaultAPIClient) UnwrapKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	panic("implement me")
}

// UpdateCertificate unused
func (KeyVaultAPIClient) UpdateCertificate(ctx context.Context, vaultBaseURL string, certificateName string, certificateVersion string, parameters keyvault.CertificateUpdateParameters) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// UpdateCertificateIssuer unused
func (KeyVaultAPIClient) UpdateCertificateIssuer(ctx context.Context, vaultBaseURL string, issuerName string, parameter keyvault.CertificateIssuerUpdateParameters) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// UpdateCertificateOperation unused
func (KeyVaultAPIClient) UpdateCertificateOperation(ctx context.Context, vaultBaseURL string, certificateName string, certificateOperation keyvault.CertificateOperationUpdateParameter) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// UpdateCertificatePolicy unused
func (KeyVaultAPIClient) UpdateCertificatePolicy(ctx context.Context, vaultBaseURL string, certificateName string, certificatePolicy keyvault.CertificatePolicy) (result keyvault.CertificatePolicy, err error) {
	panic("implement me")
}

// UpdateKey unused
func (KeyVaultAPIClient) UpdateKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyUpdateParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// UpdateSasDefinition unused
func (KeyVaultAPIClient) UpdateSasDefinition(ctx context.Context, vaultBaseURL string, storageAccountName string, sasDefinitionName string, parameters keyvault.SasDefinitionUpdateParameters) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// UpdateSecret unused
func (KeyVaultAPIClient) UpdateSecret(ctx context.Context, vaultBaseURL string, secretName string, secretVersion string, parameters keyvault.SecretUpdateParameters) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// UpdateStorageAccount unused
func (KeyVaultAPIClient) UpdateStorageAccount(ctx context.Context, vaultBaseURL string, storageAccountName string, parameters keyvault.StorageAccountUpdateParameters) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// Verify unused
func (KeyVaultAPIClient) Verify(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyVerifyParameters) (result keyvault.KeyVerifyResult, err error) {
	panic("implement me")
}

// WrapKey unused
func (KeyVaultAPIClient) WrapKey(ctx context.Context, vaultBaseURL string, keyName string, keyVersion string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	panic("implement me")
}
