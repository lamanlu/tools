package keys

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lamanlu/tools/common"
	"golang.org/x/crypto/pbkdf2"
)

func rootPartFileName(idx int) string {
	name := fmt.Sprintf("%s%d%s", RootKeyPartPrefix, idx+1, RootKeyPartSuffix)
	return filepath.Join(RootKeyDir, name)
}

func rootSaltFileName() string {
	return filepath.Join(RootKeyDir, RootKeySaltFile)
}

func creatRootKey() error {
	err := common.TouchPath(RootKeyDir, 0740)
	if err != nil {
		return err
	}

	for i := 0; i < KeyPartNum; i++ {
		fileName := rootPartFileName(i)
		_, err := os.Stat(fileName)
		if !os.IsNotExist(err) {
			msg := "root key file: " + fileName + " already exist"
			return errors.New(msg)
		}
	}

	for i := 0; i < KeyPartNum; i++ {
		fileName := rootPartFileName(i)
		fd, err := os.Create(fileName)
		if err != nil {
			return err
		}
		key, err := common.GetRandKey(KeyLen)
		if err != nil {
			fd.Close()
			return err
		}
		_, err = fd.WriteString(common.TransByteToBase64(key))
		fd.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func createKeySalt() error {
	err := common.TouchPath(RootKeyDir, 0740)
	if err != nil {
		return err
	}

	saltFile := rootSaltFileName()
	_, err = os.Stat(saltFile)
	if !os.IsNotExist(err) {
		msg := "root key salt file: " + saltFile + " already exist, exit"
		return errors.New(msg)
	}
	fd, err := os.Create(saltFile)
	if err != nil {
		return err
	}

	defer fd.Close()

	key, err := common.GetRandKey(KeyLen)
	if err != nil {
		return err
	}

	fd.WriteString(common.TransByteToBase64(key))

	return nil
}

func clearExistKeys() {
	fmt.Println("Clear exist key files")
	for i := 0; i < KeyPartNum; i++ {
		os.Remove(rootPartFileName(i))
	}
	os.Remove(rootSaltFileName())
	os.RemoveAll(WorkKeyDir)
}

func getRootKey() ([]byte, error) {
	var subs [][]byte

	for i := 0; i < KeyPartNum; i++ {
		fileName := rootPartFileName(i)
		keyStr, err := os.ReadFile(fileName)
		if err != nil {
			return []byte{}, err
		}
		sub, err := common.TransBase64ToByte(string(keyStr))
		if err != nil {
			return []byte{}, err
		}
		subs = append(subs, sub)
	}

	//check byte len
	if len(subs) < 2 {
		return []byte{}, errors.New("root key subs less than 2, not safe, exit")
	}
	l := len(subs[0])
	for i := 1; i < len(subs); i++ {
		if len(subs[i]) != l {
			return []byte{}, errors.New("root key subs length not same, exit")
		}
	}

	saltStr, err := os.ReadFile(rootSaltFileName())
	if err != nil {
		return []byte{}, err
	}
	salt, err := common.TransBase64ToByte(string(saltStr))
	if err != nil {
		return []byte{}, err
	}

	matrix := subs[0]
	for i := 1; i < len(subs); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[i] = matrix[i] ^ subs[i][j]
		}
	}

	rootKey := pbkdf2.Key(matrix, salt, 100000, KeyLen, sha512.New)

	return rootKey, nil

}
