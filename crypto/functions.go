package crypto

import (
	"text/template"

	"github.com/VirtusLab/crypt/crypto/render"
)

/*
TemplateFunctions provides template functions for render or the standard (text/template) template engine

  - encryptAWS - encrypts the data from inside of the template using AWS KMS, for best results use with gzip and b64enc
  - decryptAWS - decrypts the data from inside of the template using AWS KMS, for best results use with ungzip and b64dec
  - encryptGCP - encrypts the data from inside of the template using GCP KMS, for best results use with gzip and b64enc
  - decryptGCP - decrypts the data from inside of the template using GCP KMS, for best results use with ungzip and b64dec
  - encryptAzure - encrypts the data from inside of the template using Azure Key Vault, for best results use with gzip and b64enc
  - decryptAzure - decrypts the data from inside of the template using Azure Key Vault, for best results use with ungzip and b64dec

*/
func TemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"encryptAWS":   render.EncryptAWS,
		"decryptAWS":   render.DecryptAWS,
		"encryptGCP":   render.EncryptGCP,
		"decryptGCP":   render.DecryptGCP,
		"encryptAzure": render.EncryptAzure,
		"decryptAzure": render.DecryptAzure,
	}
}
