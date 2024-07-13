package common

import (
	"crypto/rand"
	"encoding/base64"
	"io/fs"
	"os"
)

func GetRandKey(keyLen int) ([]byte, error) {
	key := make([]byte, keyLen)
	_, err := rand.Read(key)
	if err != nil {
		return []byte{}, err
	}
	return key, nil
}

func TransBase64ToByte(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func TransByteToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func TouchPath(path string, mode uint32) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(path, fs.FileMode(mode))
	return err
}
