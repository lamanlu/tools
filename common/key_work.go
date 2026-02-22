package common

import (
	"errors"
	"os"
)

func CreateWorkKey(keyName string) error {
	err := TouchPath(WorkKeyDir, 0740)
	if err != nil {
		return err
	}

	target := WorkKeyFileName(keyName)
	fd, err := os.Create(target)
	if err != nil {
		return err
	}
	defer fd.Close()

	key, err := GetRandKey(KeyLen)
	if err != nil {
		return err
	}

	rootKey, err := GetRootKey()
	if err != nil {
		return err
	}

	encryptKey, err := Encrypt(key, rootKey)
	if err != nil {
		return err
	}

	_, err = fd.WriteString(TransByteToBase64(encryptKey))
	return err
}

func ClearWorkKey(name string) error {
	target := WorkKeyFileName(name)
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return nil
	}
	return os.Remove(target)
}

func LoadWorkKey(keyFile string) ([]byte, error) {
	keyFile = WorkKeyFileName(keyFile)
	if !FileExist(keyFile) {
		return []byte{}, errors.New("work key file: " + keyFile + " is not exist.")
	}

	keyStr, err := os.ReadFile(keyFile)
	if err != nil {
		return []byte{}, err
	}

	workKey, err := TransBase64ToByte(string(keyStr))
	if err != nil {
		return []byte{}, err
	}

	rootKey, err := GetRootKey()
	if err != nil {
		return []byte{}, err
	}

	return Decrypt(workKey, rootKey)
}
