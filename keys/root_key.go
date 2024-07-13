package keys

import (
	"bufio"
	"crypto/sha512"
	"errors"
	"fmt"
	"os"

	"github.com/lamanlu/tools/common"
	"golang.org/x/crypto/pbkdf2"
)

func creatRootKey() {
	_, err := os.Stat(RootKeyFile)
	if !os.IsNotExist(err) {
		fmt.Println("Root key file: " + RootKeyFile + " already exist")
		return
	}
	fd, err := os.Create(RootKeyFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fd.Close()

	for i := 0; i < KeyPartNum; i++ {
		key, err := common.GetRandKey(KeyLen)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fd.WriteString(common.TransByteToBase64(key) + "\n")
	}

	fmt.Println("Create Root Key Done")
}

func createKeySalt() error {
	_, err := os.Stat(RootKeySaltFile)
	if !os.IsNotExist(err) {
		fmt.Println("Root key salt file: " + RootKeySaltFile + " already exist, exit")
		return errors.New("root key salt already exist")
	}
	fd, err := os.Create(RootKeySaltFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer fd.Close()

	key, err := common.GetRandKey(KeyLen)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fd.WriteString(common.TransByteToBase64(key))

	fmt.Println("Create Root Key Salt Done")
	return nil
}

func GetRootKey() ([]byte, error) {
	var subs [][]byte

	fd, err := os.Open(RootKeyFile)
	if err != nil {
		return []byte{}, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		key := scanner.Text()
		sub, err := common.TransBase64ToByte(key)
		if err != nil {
			return []byte{}, err
		}
		subs = append(subs, sub)
	}

	//check byte len
	if len(subs) < 2 {
		fmt.Println("Rook key subs less than 2, not safe, exit")
		return []byte{}, errors.New("rook key subs less than 2, not safe, exit")
	}
	l := len(subs[0])
	for i := 1; i < len(subs); i++ {
		if len(subs[i]) != l {
			fmt.Println("Rook key subs length not same, exit")
			return []byte{}, errors.New("rook key subs length not same, exit")
		}
	}

	saltStr, err := os.ReadFile(RootKeySaltFile)
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
