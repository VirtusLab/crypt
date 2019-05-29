// Package fake is a fake provider for testing purposes
package fake

// Operation represents a Fake operation
type Operation func(input []byte) ([]byte, error)

// Fake is a fake KMS provider
type Fake struct {
	encrypt Operation
	decrypt Operation
}

// Empty creates an empty (as in does nothing) fake provider
func Empty() *Fake {
	encrypt := func(plaintext []byte) ([]byte, error) {
		// do nothing
		return plaintext, nil
	}
	decrypt := func(ciphertext []byte) ([]byte, error) {
		// do nothing
		return ciphertext, nil
	}
	return New(encrypt, decrypt)
}

// New creates a custom fake provider
func New(encrypt Operation, decrypt Operation) *Fake {
	return &Fake{
		encrypt: encrypt,
		decrypt: decrypt,
	}
}

// Encrypt is a fake encryption operation
func (f *Fake) Encrypt(plaintext []byte) ([]byte, error) {
	return f.encrypt(plaintext)
}

// Decrypt is a fake decryption operation
func (f *Fake) Decrypt(ciphertext []byte) ([]byte, error) {
	return f.decrypt(ciphertext)
}
