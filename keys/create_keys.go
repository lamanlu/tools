/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package keys

import (
	"fmt"

	"github.com/spf13/cobra"
)

const RootKeyFile = "root.key"
const RootKeySaltFile = "root.salt"
const KeyLen = 32
const KeyPartNum = 2
const WorkKeyDir = "workKey"

// createKeysCmd represents the createKeys command
var cmd = &cobra.Command{
	Use:   "key-gen",
	Short: "",
	Long:  ``,
	Run:   runCmd,
}

var keyType string
var forceFlag bool
var workKeyFile string

func init() {
	cmd.Flags().StringVarP(&keyType, "type", "t", "root", "Key Type: root, work.")
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
		fmt.Println("Create Root Key Salt Done")
		err = creatRootKey()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Create Root Key Done")
	case "work":
		var err error
		if forceFlag {
			err = clearWorkKey(workKeyFile)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		err = creatWorkKey(workKeyFile)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("unknow type: " + keyType)
	}

}
