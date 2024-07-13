/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package keys

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const RootKeyFile = "root.key"
const RootKeySaltFile = "root.salt"
const KeyLen = 32
const KeyPartNum = 2
const WorkKeyDir = "workKey"

// createKeysCmd represents the createKeys command
var cmd = &cobra.Command{
	Use:   "rootkey",
	Short: "",
	Long:  ``,
	Run:   runCmd,
}

var keyType string
var forceFlag bool
var workKeyFile string

func init() {
	cmd.Flags().StringVarP(&keyType, "type", "t", "work", "Key Type: root, work.")
	cmd.Flags().StringVarP(&workKeyFile, "name", "n", "work.key", "Work Key File Name. eg: work.key")
	cmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force Create RootKey, Ignore Exist key.")
}

func GetCmd() *cobra.Command {
	return cmd
}

func runCmd(cmd *cobra.Command, args []string) {
	switch keyType {
	case "root":
		if forceFlag {
			clearExistKeys()
		}
		err := createKeySalt()
		if err != nil {
			fmt.Println(err)
			break
		}
		creatRootKey()
	case "work":
		err := creatWorkKey(workKeyFile)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("unknow type: " + keyType)
	}

}

func clearExistKeys() {
	fmt.Println("Clear exist key files")
	os.Remove(RootKeyFile)
	os.Remove(RootKeySaltFile)
	os.RemoveAll(WorkKeyDir)
}
