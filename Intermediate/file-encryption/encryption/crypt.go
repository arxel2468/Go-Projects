package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

// EncryptFile encrypts the content of the input file and writes to the output file
func EncryptFile(inputFile, outputFile, key string) error {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, input, nil)
	return os.WriteFile(outputFile, ciphertext, 0644)
}

// DecryptFile decrypts the content of the input file and writes to the output file
func DecryptFile(inputFile, outputFile, key string) error {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(input) < nonceSize {
		return errors.New("invalid ciphertext")
	}

	nonce, ciphertext := input[:nonceSize], input[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, plaintext, 0644)
}
