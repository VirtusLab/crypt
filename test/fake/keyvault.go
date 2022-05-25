package fake

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

// KeyVaultAPIClient is a fake KeyVault API client provider
type KeyVaultAPIClient struct {
}

// GetCertificateIssuersComplete unused
func (c KeyVaultAPIClient) GetCertificateIssuersComplete(_ context.Context, _ string, _ *int32) (result keyvault.CertificateIssuerListResultIterator, err error) {
	panic("implement me")
}

// GetCertificatesComplete unused
func (c KeyVaultAPIClient) GetCertificatesComplete(_ context.Context, _ string, _ *int32) (result keyvault.CertificateListResultIterator, err error) {
	panic("implement me")
}

// GetCertificateVersionsComplete unused
func (c KeyVaultAPIClient) GetCertificateVersionsComplete(_ context.Context, _ string, _ string, _ *int32) (result keyvault.CertificateListResultIterator, err error) {
	panic("implement me")
}

// GetDeletedCertificatesComplete unused
func (c KeyVaultAPIClient) GetDeletedCertificatesComplete(_ context.Context, _ string, _ *int32) (result keyvault.DeletedCertificateListResultIterator, err error) {
	panic("implement me")
}

// GetDeletedKeysComplete unused
func (c KeyVaultAPIClient) GetDeletedKeysComplete(_ context.Context, _ string, _ *int32) (result keyvault.DeletedKeyListResultIterator, err error) {
	panic("implement me")
}

// GetDeletedSecretsComplete unused
func (c KeyVaultAPIClient) GetDeletedSecretsComplete(_ context.Context, _ string, _ *int32) (result keyvault.DeletedSecretListResultIterator, err error) {
	panic("implement me")
}

// GetKeysComplete unused
func (c KeyVaultAPIClient) GetKeysComplete(_ context.Context, _ string, _ *int32) (result keyvault.KeyListResultIterator, err error) {
	panic("implement me")
}

// GetKeyVersionsComplete unused
func (c KeyVaultAPIClient) GetKeyVersionsComplete(_ context.Context, _ string, _ string, _ *int32) (result keyvault.KeyListResultIterator, err error) {
	panic("implement me")
}

// GetSasDefinitionsComplete unused
func (c KeyVaultAPIClient) GetSasDefinitionsComplete(_ context.Context, _ string, _ string, _ *int32) (result keyvault.SasDefinitionListResultIterator, err error) {
	panic("implement me")
}

// GetSecretsComplete unused
func (c KeyVaultAPIClient) GetSecretsComplete(_ context.Context, _ string, _ *int32) (result keyvault.SecretListResultIterator, err error) {
	panic("implement me")
}

// GetSecretVersionsComplete unused
func (c KeyVaultAPIClient) GetSecretVersionsComplete(_ context.Context, _ string, _ string, _ *int32) (result keyvault.SecretListResultIterator, err error) {
	panic("implement me")
}

// GetStorageAccountsComplete unused
func (c KeyVaultAPIClient) GetStorageAccountsComplete(_ context.Context, _ string, _ *int32) (result keyvault.StorageListResultIterator, err error) {
	panic("implement me")
}

// BackupKey unused
func (KeyVaultAPIClient) BackupKey(_ context.Context, _ string, _ string) (result keyvault.BackupKeyResult, err error) {
	panic("implement me")
}

// BackupSecret unused
func (KeyVaultAPIClient) BackupSecret(_ context.Context, _ string, _ string) (result keyvault.BackupSecretResult, err error) {
	panic("implement me")
}

// CreateCertificate unused
func (KeyVaultAPIClient) CreateCertificate(_ context.Context, _ string, _ string, _ keyvault.CertificateCreateParameters) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// CreateKey unused
func (KeyVaultAPIClient) CreateKey(_ context.Context, _ string, _ string, _ keyvault.KeyCreateParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// Decrypt unused
func (KeyVaultAPIClient) Decrypt(_ context.Context, _ string, _ string, _ string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	result.Result = parameters.Value
	return
}

// DeleteCertificate unused
func (KeyVaultAPIClient) DeleteCertificate(_ context.Context, _ string, _ string) (result keyvault.DeletedCertificateBundle, err error) {
	panic("implement me")
}

// DeleteCertificateContacts unused
func (KeyVaultAPIClient) DeleteCertificateContacts(_ context.Context, _ string) (result keyvault.Contacts, err error) {
	panic("implement me")
}

// DeleteCertificateIssuer unused
func (KeyVaultAPIClient) DeleteCertificateIssuer(_ context.Context, _ string, _ string) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// DeleteCertificateOperation unused
func (KeyVaultAPIClient) DeleteCertificateOperation(_ context.Context, _ string, _ string) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// DeleteKey unused
func (KeyVaultAPIClient) DeleteKey(_ context.Context, _ string, _ string) (result keyvault.DeletedKeyBundle, err error) {
	panic("implement me")
}

// DeleteSasDefinition unused
func (KeyVaultAPIClient) DeleteSasDefinition(_ context.Context, _ string, _ string, _ string) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// DeleteSecret unused
func (KeyVaultAPIClient) DeleteSecret(_ context.Context, _ string, _ string) (result keyvault.DeletedSecretBundle, err error) {
	panic("implement me")
}

// DeleteStorageAccount unused
func (KeyVaultAPIClient) DeleteStorageAccount(_ context.Context, _ string, _ string) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// Encrypt unused
func (KeyVaultAPIClient) Encrypt(_ context.Context, _ string, _ string, _ string, parameters keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	result.Result = parameters.Value
	return
}

// GetCertificate unused
func (KeyVaultAPIClient) GetCertificate(_ context.Context, _ string, _ string, _ string) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// GetCertificateContacts unused
func (KeyVaultAPIClient) GetCertificateContacts(_ context.Context, _ string) (result keyvault.Contacts, err error) {
	panic("implement me")
}

// GetCertificateIssuer unused
func (KeyVaultAPIClient) GetCertificateIssuer(_ context.Context, _ string, _ string) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// GetCertificateIssuers unused
func (KeyVaultAPIClient) GetCertificateIssuers(_ context.Context, _ string, _ *int32) (result keyvault.CertificateIssuerListResultPage, err error) {
	panic("implement me")
}

// GetCertificateOperation unused
func (KeyVaultAPIClient) GetCertificateOperation(_ context.Context, _ string, _ string) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// GetCertificatePolicy unused
func (KeyVaultAPIClient) GetCertificatePolicy(_ context.Context, _ string, _ string) (result keyvault.CertificatePolicy, err error) {
	panic("implement me")
}

// GetCertificates unused
func (KeyVaultAPIClient) GetCertificates(_ context.Context, _ string, _ *int32) (result keyvault.CertificateListResultPage, err error) {
	panic("implement me")
}

// GetCertificateVersions unused
func (KeyVaultAPIClient) GetCertificateVersions(_ context.Context, _ string, _ string, _ *int32) (result keyvault.CertificateListResultPage, err error) {
	panic("implement me")
}

// GetDeletedCertificate unused
func (KeyVaultAPIClient) GetDeletedCertificate(_ context.Context, _ string, _ string) (result keyvault.DeletedCertificateBundle, err error) {
	panic("implement me")
}

// GetDeletedCertificates unused
func (KeyVaultAPIClient) GetDeletedCertificates(_ context.Context, _ string, _ *int32) (result keyvault.DeletedCertificateListResultPage, err error) {
	panic("implement me")
}

// GetDeletedKey unused
func (KeyVaultAPIClient) GetDeletedKey(_ context.Context, _ string, _ string) (result keyvault.DeletedKeyBundle, err error) {
	panic("implement me")
}

// GetDeletedKeys unused
func (KeyVaultAPIClient) GetDeletedKeys(_ context.Context, _ string, _ *int32) (result keyvault.DeletedKeyListResultPage, err error) {
	panic("implement me")
}

// GetDeletedSecret unused
func (KeyVaultAPIClient) GetDeletedSecret(_ context.Context, _ string, _ string) (result keyvault.DeletedSecretBundle, err error) {
	panic("implement me")
}

// GetDeletedSecrets unused
func (KeyVaultAPIClient) GetDeletedSecrets(_ context.Context, _ string, _ *int32) (result keyvault.DeletedSecretListResultPage, err error) {
	panic("implement me")
}

// GetKey unused
func (KeyVaultAPIClient) GetKey(_ context.Context, _ string, _ string, _ string) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// GetKeys unused
func (KeyVaultAPIClient) GetKeys(_ context.Context, _ string, _ *int32) (result keyvault.KeyListResultPage, err error) {
	panic("implement me")
}

// GetKeyVersions unused
func (KeyVaultAPIClient) GetKeyVersions(_ context.Context, _ string, _ string, _ *int32) (result keyvault.KeyListResultPage, err error) {
	panic("implement me")
}

// GetSasDefinition unused
func (KeyVaultAPIClient) GetSasDefinition(_ context.Context, _ string, _ string, _ string) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// GetSasDefinitions unused
func (KeyVaultAPIClient) GetSasDefinitions(_ context.Context, _ string, _ string, _ *int32) (result keyvault.SasDefinitionListResultPage, err error) {
	panic("implement me")
}

// GetSecret unused
func (KeyVaultAPIClient) GetSecret(_ context.Context, _ string, _ string, _ string) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// GetSecrets unused
func (KeyVaultAPIClient) GetSecrets(_ context.Context, _ string, _ *int32) (result keyvault.SecretListResultPage, err error) {
	panic("implement me")
}

// GetSecretVersions unused
func (KeyVaultAPIClient) GetSecretVersions(_ context.Context, _ string, _ string, _ *int32) (result keyvault.SecretListResultPage, err error) {
	panic("implement me")
}

// GetStorageAccount unused
func (KeyVaultAPIClient) GetStorageAccount(_ context.Context, _ string, _ string) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// GetStorageAccounts unused
func (KeyVaultAPIClient) GetStorageAccounts(_ context.Context, _ string, _ *int32) (result keyvault.StorageListResultPage, err error) {
	panic("implement me")
}

// ImportCertificate unused
func (KeyVaultAPIClient) ImportCertificate(_ context.Context, _ string, _ string, _ keyvault.CertificateImportParameters) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// ImportKey unused
func (KeyVaultAPIClient) ImportKey(_ context.Context, _ string, _ string, _ keyvault.KeyImportParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// MergeCertificate unused
func (KeyVaultAPIClient) MergeCertificate(_ context.Context, _ string, _ string, _ keyvault.CertificateMergeParameters) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// PurgeDeletedCertificate unused
func (KeyVaultAPIClient) PurgeDeletedCertificate(_ context.Context, _ string, _ string) (result autorest.Response, err error) {
	panic("implement me")
}

// PurgeDeletedKey unused
func (KeyVaultAPIClient) PurgeDeletedKey(_ context.Context, _ string, _ string) (result autorest.Response, err error) {
	panic("implement me")
}

// PurgeDeletedSecret unused
func (KeyVaultAPIClient) PurgeDeletedSecret(_ context.Context, _ string, _ string) (result autorest.Response, err error) {
	panic("implement me")
}

// RecoverDeletedCertificate unused
func (KeyVaultAPIClient) RecoverDeletedCertificate(_ context.Context, _ string, _ string) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// RecoverDeletedKey unused
func (KeyVaultAPIClient) RecoverDeletedKey(_ context.Context, _ string, _ string) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// RecoverDeletedSecret unused
func (KeyVaultAPIClient) RecoverDeletedSecret(_ context.Context, _ string, _ string) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// RegenerateStorageAccountKey unused
func (KeyVaultAPIClient) RegenerateStorageAccountKey(_ context.Context, _ string, _ string, _ keyvault.StorageAccountRegenerteKeyParameters) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// RestoreKey unused
func (KeyVaultAPIClient) RestoreKey(_ context.Context, _ string, _ keyvault.KeyRestoreParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// RestoreSecret unused
func (KeyVaultAPIClient) RestoreSecret(_ context.Context, _ string, _ keyvault.SecretRestoreParameters) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// SetCertificateContacts unused
func (KeyVaultAPIClient) SetCertificateContacts(_ context.Context, _ string, _ keyvault.Contacts) (result keyvault.Contacts, err error) {
	panic("implement me")
}

// SetCertificateIssuer unused
func (KeyVaultAPIClient) SetCertificateIssuer(_ context.Context, _ string, _ string, _ keyvault.CertificateIssuerSetParameters) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// SetSasDefinition unused
func (KeyVaultAPIClient) SetSasDefinition(_ context.Context, _ string, _ string, _ string, _ keyvault.SasDefinitionCreateParameters) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// SetSecret unused
func (KeyVaultAPIClient) SetSecret(_ context.Context, _ string, _ string, _ keyvault.SecretSetParameters) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// SetStorageAccount unused
func (KeyVaultAPIClient) SetStorageAccount(_ context.Context, _ string, _ string, _ keyvault.StorageAccountCreateParameters) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// Sign unused
func (KeyVaultAPIClient) Sign(_ context.Context, _ string, _ string, _ string, _ keyvault.KeySignParameters) (result keyvault.KeyOperationResult, err error) {
	panic("implement me")
}

// UnwrapKey unused
func (KeyVaultAPIClient) UnwrapKey(_ context.Context, _ string, _ string, _ string, _ keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	panic("implement me")
}

// UpdateCertificate unused
func (KeyVaultAPIClient) UpdateCertificate(_ context.Context, _ string, _ string, _ string, _ keyvault.CertificateUpdateParameters) (result keyvault.CertificateBundle, err error) {
	panic("implement me")
}

// UpdateCertificateIssuer unused
func (KeyVaultAPIClient) UpdateCertificateIssuer(_ context.Context, _ string, _ string, _ keyvault.CertificateIssuerUpdateParameters) (result keyvault.IssuerBundle, err error) {
	panic("implement me")
}

// UpdateCertificateOperation unused
func (KeyVaultAPIClient) UpdateCertificateOperation(_ context.Context, _ string, _ string, _ keyvault.CertificateOperationUpdateParameter) (result keyvault.CertificateOperation, err error) {
	panic("implement me")
}

// UpdateCertificatePolicy unused
func (KeyVaultAPIClient) UpdateCertificatePolicy(_ context.Context, _ string, _ string, _ keyvault.CertificatePolicy) (result keyvault.CertificatePolicy, err error) {
	panic("implement me")
}

// UpdateKey unused
func (KeyVaultAPIClient) UpdateKey(_ context.Context, _ string, _ string, _ string, _ keyvault.KeyUpdateParameters) (result keyvault.KeyBundle, err error) {
	panic("implement me")
}

// UpdateSasDefinition unused
func (KeyVaultAPIClient) UpdateSasDefinition(_ context.Context, _ string, _ string, _ string, _ keyvault.SasDefinitionUpdateParameters) (result keyvault.SasDefinitionBundle, err error) {
	panic("implement me")
}

// UpdateSecret unused
func (KeyVaultAPIClient) UpdateSecret(_ context.Context, _ string, _ string, _ string, _ keyvault.SecretUpdateParameters) (result keyvault.SecretBundle, err error) {
	panic("implement me")
}

// UpdateStorageAccount unused
func (KeyVaultAPIClient) UpdateStorageAccount(_ context.Context, _ string, _ string, _ keyvault.StorageAccountUpdateParameters) (result keyvault.StorageBundle, err error) {
	panic("implement me")
}

// Verify unused
func (KeyVaultAPIClient) Verify(_ context.Context, _ string, _ string, _ string, _ keyvault.KeyVerifyParameters) (result keyvault.KeyVerifyResult, err error) {
	panic("implement me")
}

// WrapKey unused
func (KeyVaultAPIClient) WrapKey(_ context.Context, _ string, _ string, _ string, _ keyvault.KeyOperationsParameters) (result keyvault.KeyOperationResult, err error) {
	panic("implement me")
}
