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

func DecryptInput(inputBase64 string, workKey string) (string, error) {
	key, err := LoadWorkKey(workKey)
	if err != nil {
		return "", err
	}
	cipherText, err := TransBase64ToByte(inputBase64)
	if err != nil {
		return "", err
	}
	plain, err := Decrypt(cipherText, key)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
