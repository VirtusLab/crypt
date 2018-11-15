package crypto

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

type FakeFunc func(input []byte) ([]byte, error)

type FakeKMS struct {
	encrypt FakeFunc
	decrypt FakeFunc
}

func NewFakeKMS(encrypt FakeFunc, decrypt FakeFunc) *FakeKMS {
	return &FakeKMS{
		encrypt: encrypt,
		decrypt: decrypt,
	}
}

func (f *FakeKMS) Encrypt(plaintext []byte) ([]byte, error) {
	return f.encrypt(plaintext)
}

func (f *FakeKMS) Decrypt(ciphertext []byte) ([]byte, error) {
	return f.decrypt(ciphertext)
}

func TestCrypt(t *testing.T) {
	type TestCase struct {
		name    string
		f       func(TestCase)
		logHook *test.Hook
	}

	when := func(crypt *Crypt, inputPath string) (string, error) {
		defer os.Remove(inputPath + ".encrypted") // clean up
		defer os.Remove(inputPath + ".decrypted") // clean up

		err := crypt.EncryptFile(inputPath, inputPath+".encrypted")
		if err != nil {
			return "", err
		}

		err = crypt.DecryptFile(inputPath+".encrypted", inputPath+".decrypted")
		if err != nil {
			return "", err
		}

		result, err := ioutil.ReadFile(inputPath + ".decrypted")
		if err != nil {
			return "", err
		}

		return string(result), nil
	}

	cases := []TestCase{
		{
			name: "encrypt decrypt file",
			f: func(tc TestCase) {
				encrypt := func(plaintext []byte) ([]byte, error) {
					// do nothing
					return plaintext, nil
				}
				decrypt := func(ciphertext []byte) ([]byte, error) {
					// do nothing
					return ciphertext, nil
				}

				fake := NewFakeKMS(encrypt, decrypt)
				crypt := New(fake)

				inputFile := "test.txt"
				expected := "top secret token"
				err := ioutil.WriteFile(inputFile, []byte(expected), 0644)
				if err != nil {
					t.Fatal("Can't write plaintext file", err)
				}
				defer os.Remove(inputFile)

				actual, err := when(crypt, inputFile)

				assert.NoError(t, err, tc.name)
				assert.Equal(t, expected, string(actual))
			},
		},
		{
			name: "encrypt decrypt non-existing file",
			f: func(tc TestCase) {
				encrypt := func(plaintext []byte) ([]byte, error) {
					// do nothing
					return plaintext, nil
				}
				decrypt := func(ciphertext []byte) ([]byte, error) {
					// do nothing
					return ciphertext, nil
				}

				fakeKMS := NewFakeKMS(encrypt, decrypt)
				crypt := New(fakeKMS)

				inputFile := "test.txt"

				_, err := when(crypt, inputFile)

				assert.Error(t, err, tc.name)
			},
		},
	}

	logrus.SetLevel(logrus.DebugLevel)
	hook := test.NewGlobal()

	for i, c := range cases {
		c.logHook = hook
		t.Run(fmt.Sprintf("[%d] %s", i, c.name), func(t *testing.T) { c.f(c) })
		hook.Reset()
	}
}
