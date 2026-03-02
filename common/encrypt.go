package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Core encryption/decryption.
func encrypt(text []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}

	cipherText := gcm.Seal(nonce, nonce, text, nil)

	return cipherText, nil
}

func decrypt(text []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce, cipherText := text[:gcm.NonceSize()], text[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
