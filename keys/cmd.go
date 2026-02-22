/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package keys

import (
	"fmt"

	"github.com/lamanlu/tools/common"
	"github.com/spf13/cobra"
)

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
			common.ClearAllKeys()
		}
		err := common.CreateRootKeySalt()
		if err != nil {
			return err
		}
		fmt.Println("Create Root Key Salt Done")
		err = common.CreateRootKeyParts()
		if err != nil {
			return err
		}
		fmt.Println("Create Root Key Done")
	case "work":
		var err error
		if forceFlag {
			err = common.ClearWorkKey(workKeyFile)
		}
		if err != nil {
			return err
		}
		err = common.CreateWorkKey(workKeyFile)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown type: %s", keyType)
	}

	return nil
}
