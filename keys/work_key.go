package keys

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/lamanlu/tools/common"
)

func creatWorkKey(keyName string) error {
	err := common.TouchPath(WorkKeyDir, 0740)
	if err != nil {
		return err
	}

	fd, err := os.Create(filepath.Join(WorkKeyDir, keyName))
	if err != nil {
		return err
	}
	defer fd.Close()

	key, err := common.GetRandKey(KeyLen)
	if err != nil {
		return err
	}

	rootKey, err := GetRootKey()
	if err != nil {
		return err
	}

	encryptKey, err := common.Encrypt(key, rootKey)
	if err != nil {
		return err
	}

	_, err = fd.WriteString(common.TransByteToBase64(encryptKey))
	return err
}

func clearWorkKey(name string) error {
	target := filepath.Join(WorkKeyDir, name)
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return nil
	}
	return os.Remove(target)
}

func getWorkKey(keyFile string) ([]byte, error) {
	keyFile = filepath.Join(WorkKeyDir, keyFile)
	if !common.FileExist(keyFile) {
		return []byte{}, errors.New("work key file: " + keyFile + " is not exist.")
	}
	return os.ReadFile(keyFile)
}
