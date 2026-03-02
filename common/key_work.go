package common

import (
	"errors"
	"os"
	"strings"
)

// Work key management.
func CreateWorkKey(keyName string) error {
	err := touchPath(WorkKeyDir, 0740)
	if err != nil {
		return err
	}

	target := workKeyFileName(keyName)
	fd, err := os.Create(target)
	if err != nil {
		return err
	}
	defer fd.Close()

	key, err := getRandKey(KeyLen)
	if err != nil {
		return err
	}

	rootKey, err := getRootKey()
	if err != nil {
		return err
	}

	encryptKey, err := encrypt(key, rootKey)
	if err != nil {
		return err
	}

	_, err = fd.WriteString(transByteToBase64(encryptKey))
	return err
}

func CreateRandomKeyFile(keyName string) error {
	err := touchPath(WorkKeyDir, 0740)
	if err != nil {
		return err
	}

	target := workKeyFileName(keyName)
	_, err = os.Stat(target)
	if !os.IsNotExist(err) {
		return errors.New("random key file: " + target + " already exist")
	}

	fd, err := os.Create(target)
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

func ClearWorkKey(name string) error {
	target := workKeyFileName(name)
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return nil
	}
	return os.Remove(target)
}

func loadWorkKey(keyFile string) ([]byte, error) {
	keyFile = workKeyFileName(keyFile)
	if !fileExist(keyFile) {
		return []byte{}, errors.New("work key file: " + keyFile + " is not exist.")
	}

	keyStr, err := os.ReadFile(keyFile)
	if err != nil {
		return []byte{}, err
	}

	workKey, err := transBase64ToByte(strings.TrimSpace(string(keyStr)))
	if err != nil {
		return []byte{}, err
	}

	rootKey, err := getRootKey()
	if err != nil {
		return []byte{}, err
	}

	return decrypt(workKey, rootKey)
}
