package common

import (
	"crypto/rand"
	"encoding/base64"
	"io/fs"
	"os"
)

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

func touchPath(path string, mode uint32) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(path, fs.FileMode(mode))
	return err
}

func fileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
