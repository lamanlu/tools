package common

import (
	"crypto/rand"
	"encoding/base64"
)

// Encoding and random key utilities.
func getRandKey(keyLen int) ([]byte, error) {
	key := make([]byte, keyLen)
	_, err := rand.Read(key)
	if err != nil {
		return []byte{}, err
	}
	return key, nil
}

func transBase64ToByte(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func transByteToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
