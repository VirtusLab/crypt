package fake

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"
)

// FakeKeyVaultAPIClient is a fake KeyVault API client provider
type FakeKeyVaultAPIClient struct {
}

// Encrypt unused
func (FakeKeyVaultAPIClient) Encrypt(ctx context.Context, name string, version string, parameters azkeys.KeyOperationsParameters, options *azkeys.EncryptOptions) (azkeys.EncryptResponse, error) {
	result := azkeys.EncryptResponse{
		KeyOperationResult: azkeys.KeyOperationResult{
			Result: parameters.Value,
		},
	}
	return result, nil
}

// Decrypt unused
func (FakeKeyVaultAPIClient) Decrypt(ctx context.Context, name string, version string, parameters azkeys.KeyOperationsParameters, options *azkeys.DecryptOptions) (azkeys.DecryptResponse, error) {
	result := azkeys.DecryptResponse{
		KeyOperationResult: azkeys.KeyOperationResult{
			Result: parameters.Value,
		},
	}
	return result, nil
}
