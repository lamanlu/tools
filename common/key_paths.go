package common

import (
	"fmt"
	"path/filepath"
)

const (
	RootKeyDir        = "rootKey"
	RootKeyPartPrefix = "root_part_"
	RootKeyPartSuffix = ".key"
	RootKeySaltFile   = "root.salt"
	WorkKeyDir        = "workKey"
)

func RootPartFileName(idx int) string {
	name := fmt.Sprintf("%s%d%s", RootKeyPartPrefix, idx+1, RootKeyPartSuffix)
	return filepath.Join(RootKeyDir, name)
}

func RootSaltFileName() string {
	return filepath.Join(RootKeyDir, RootKeySaltFile)
}

func WorkKeyFileName(name string) string {
	return filepath.Join(WorkKeyDir, name)
}
