package keys

import (
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

	_, err = fd.WriteString(common.TransByteToBase64(key))
	return err
}
