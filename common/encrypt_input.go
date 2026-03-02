package common

// String input encryption/decryption helpers.
func EncryptInput(input string, workKey string) (string, error) {
	key, err := loadWorkKey(workKey)
	if err != nil {
		return "", err
	}
	res, err := encrypt([]byte(input), key)
	if err != nil {
		return "", err
	}
	return transByteToBase64(res), nil
}

func DecryptInput(inputBase64 string, workKey string) (string, error) {
	key, err := loadWorkKey(workKey)
	if err != nil {
		return "", err
	}
	cipherText, err := transBase64ToByte(inputBase64)
	if err != nil {
		return "", err
	}
	plain, err := decrypt(cipherText, key)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
