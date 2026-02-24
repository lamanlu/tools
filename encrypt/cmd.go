package encrypt

import (
	"fmt"

	"github.com/lamanlu/tools/common"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a string with a work key",
	Long:  "Encrypt a plaintext string using the specified work key file and output Base64.",
	RunE:  runCmd,
}

func GetCmd() *cobra.Command {
	return cmd
}

var workKeyFile string
var keyBaseDir string

func init() {
	cmd.Flags().StringVarP(&workKeyFile, "work-key", "k", "", "Work key file name, using for encrypt input string.")
	cmd.Flags().StringVarP(&keyBaseDir, "key-dir", "d", "", "Key base directory containing rootKey/workKey.")
}

func runCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid input string")
	}
	common.SetKeyBaseDir(keyBaseDir)
	input := args[0]
	encryptStr, err := common.EncryptInput(input, workKeyFile)
	if err != nil {
		return err
	}
	fmt.Print(encryptStr)
	return nil
}
