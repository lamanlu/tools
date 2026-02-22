/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package keys

import (
	"fmt"

	"github.com/spf13/cobra"
)

const RootKeyDir = "rootKey"
const RootKeyPartPrefix = "root_part_"
const RootKeyPartSuffix = ".key"
const RootKeySaltFile = "root.salt"
const KeyLen = 32
const KeyPartNum = 2
const WorkKeyDir = "workKey"

// createKeysCmd represents the createKeys command
var cmd = &cobra.Command{
	Use:   "key-gen",
	Short: "Generate root or work keys",
	Long:  "Generate root keys (rootKey/root_part_*.key + rootKey/root.salt) or a work key encrypted by the root key.",
	RunE:  runCmd,
}

var keyType string
var forceFlag bool
var workKeyFile string

func init() {
	cmd.Flags().StringVarP(&keyType, "type", "t", "", "Key Type: root, work.")
	cmd.Flags().StringVarP(&workKeyFile, "name", "n", "work.key", "Work Key File Name. eg: work.key")
	cmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force Create RootKey, Ignore Exist key.")
	if err := cmd.MarkFlagRequired("type"); err != nil {
		panic(err)
	}
}

func GetCmd() *cobra.Command {
	return cmd
}

func runCmd(cmd *cobra.Command, args []string) error {
	if keyType == "" {
		return fmt.Errorf("type is required")
	}
	switch keyType {
	case "root":
		if forceFlag {
			clearExistKeys()
		}
		err := createKeySalt()
		if err != nil {
			return err
		}
		fmt.Println("Create Root Key Salt Done")
		err = creatRootKey()
		if err != nil {
			return err
		}
		fmt.Println("Create Root Key Done")
	case "work":
		var err error
		if forceFlag {
			err = clearWorkKey(workKeyFile)
		}
		if err != nil {
			return err
		}
		err = creatWorkKey(workKeyFile)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown type: %s", keyType)
	}

	return nil
}
