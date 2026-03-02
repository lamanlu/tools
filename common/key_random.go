package common

import (
	"errors"
	"os"
)

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
