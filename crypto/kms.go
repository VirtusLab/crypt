package crypto

// KMS - Key Management Service interface abstracts common cryptographic operations.
// A KMS must be able to decrypt the data it encrypts.
type KMS interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

