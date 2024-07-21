package keys

import (
	"github.com/lamanlu/tools/common"
)

func EncryptInput(input string, workKey string) (string, error) {
	key, err := getWorkKey(workKey)
	if err != nil {
		return "", err
	}
	res, err := common.Encrypt([]byte(input), key)
	if err != nil {
		return "", err
	}
	return common.TransByteToBase64(res), nil
}
