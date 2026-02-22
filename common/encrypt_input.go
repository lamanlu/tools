package common

func EncryptInput(input string, workKey string) (string, error) {
	key, err := LoadWorkKey(workKey)
	if err != nil {
		return "", err
	}
	res, err := Encrypt([]byte(input), key)
	if err != nil {
		return "", err
	}
	return TransByteToBase64(res), nil
}
