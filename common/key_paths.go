package common

import (
	"fmt"
	"path/filepath"
)

const (
	RootKeyPartPrefix = "root_part_"
	RootKeyPartSuffix = ".key"
	RootKeySaltFile   = "root.salt"
)

var (
	RootKeyDir = "rootKey"
	WorkKeyDir = "workKey"
)

func SetKeyBaseDir(base string) {
	if base == "" {
		return
	}
	base = filepath.Clean(base)
	if base == "." {
		RootKeyDir = "rootKey"
		WorkKeyDir = "workKey"
		return
	}
	RootKeyDir = filepath.Join(base, "rootKey")
	WorkKeyDir = filepath.Join(base, "workKey")
}

func rootPartFileName(idx int) string {
	name := fmt.Sprintf("%s%d%s", RootKeyPartPrefix, idx+1, RootKeyPartSuffix)
	return filepath.Join(RootKeyDir, name)
}

func rootSaltFileName() string {
	return filepath.Join(RootKeyDir, RootKeySaltFile)
}

func workKeyFileName(name string) string {
	return filepath.Join(WorkKeyDir, name)
}
