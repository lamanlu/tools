package common

import (
	"errors"
	"os"
)

func CreateRandomKeyFile(keyName string) error {
	err := TouchPath(WorkKeyDir, 0740)
	if err != nil {
		return err
	}

	target := WorkKeyFileName(keyName)
	_, err = os.Stat(target)
	if !os.IsNotExist(err) {
		return errors.New("random key file: " + target + " already exist")
	}

	fd, err := os.Create(target)
	if err != nil {
		return err
	}
	defer fd.Close()

	key, err := GetRandKey(KeyLen)
	if err != nil {
		return err
	}

	_, err = fd.WriteString(TransByteToBase64(key))
	return err
}
