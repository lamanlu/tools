package common

import (
	"crypto/sha512"
	"errors"
	"os"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	KeyLen     = 32
	KeyPartNum = 2
)

func CreateRootKeyParts() error {
	err := touchPath(RootKeyDir, 0740)
	if err != nil {
		return err
	}

	for i := 0; i < KeyPartNum; i++ {
		fileName := rootPartFileName(i)
		_, err := os.Stat(fileName)
		if !os.IsNotExist(err) {
			return errors.New("root key file: " + fileName + " already exist")
		}
	}

	for i := 0; i < KeyPartNum; i++ {
		fileName := rootPartFileName(i)
		fd, err := os.Create(fileName)
		if err != nil {
			return err
		}
		key, err := getRandKey(KeyLen)
		if err != nil {
			fd.Close()
			return err
		}
		_, err = fd.WriteString(transByteToBase64(key))
		fd.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateRootKeySalt() error {
	err := touchPath(RootKeyDir, 0740)
	if err != nil {
		return err
	}

	saltFile := rootSaltFileName()
	_, err = os.Stat(saltFile)
	if !os.IsNotExist(err) {
		return errors.New("root key salt file: " + saltFile + " already exist, exit")
	}
	fd, err := os.Create(saltFile)
	if err != nil {
		return err
	}
	defer fd.Close()

	key, err := getRandKey(KeyLen)
	if err != nil {
		return err
	}

	_, err = fd.WriteString(transByteToBase64(key))
	return err
}

func ClearAllKeys() {
	for i := 0; i < KeyPartNum; i++ {
		_ = os.Remove(rootPartFileName(i))
	}
	_ = os.Remove(rootSaltFileName())
	_ = os.RemoveAll(WorkKeyDir)
}

func getRootKey() ([]byte, error) {
	var subs [][]byte

	for i := 0; i < KeyPartNum; i++ {
		fileName := rootPartFileName(i)
		keyStr, err := os.ReadFile(fileName)
		if err != nil {
			return []byte{}, err
		}
		sub, err := transBase64ToByte(strings.TrimSpace(string(keyStr)))
		if err != nil {
			return []byte{}, err
		}
		subs = append(subs, sub)
	}

	if len(subs) < 2 {
		return []byte{}, errors.New("root key subs less than 2, not safe, exit")
	}
	length := len(subs[0])
	for i := 1; i < len(subs); i++ {
		if len(subs[i]) != length {
			return []byte{}, errors.New("root key subs length not same, exit")
		}
	}

	saltStr, err := os.ReadFile(rootSaltFileName())
	if err != nil {
		return []byte{}, err
	}
	salt, err := transBase64ToByte(strings.TrimSpace(string(saltStr)))
	if err != nil {
		return []byte{}, err
	}

	matrix := make([]byte, length)
	for i := 0; i < len(subs); i++ {
		for j := range matrix {
			matrix[j] ^= subs[i][j]
		}
	}

	rootKey := pbkdf2.Key(matrix, salt, 100000, KeyLen, sha512.New)

	return rootKey, nil
}
