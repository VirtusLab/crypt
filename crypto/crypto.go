package crypto

type Crypt interface {
	Encrypt(inputPath, outputPath string, params map[string]interface{}) error
	Decrypt(inputPath, outputPath string, params map[string]interface{}) error
}

