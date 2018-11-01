package crypto

// Key Management Service interface abstracts common cryptographic operations.
type KMS interface {
	Encrypt(plaintext []byte, params map[string]interface{}) (error, []byte)
	Decrypt(ciphertext []byte, params map[string]interface{}) (error, []byte)
}

