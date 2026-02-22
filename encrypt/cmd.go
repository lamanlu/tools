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

var workKeyFIle string

func init() {
	cmd.Flags().StringVarP(&workKeyFIle, "work-key", "key", "", "Work key file name, using for encrypt input string.")
}

func runCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid input string")
	}
	input := args[0]
	encryptStr, err := common.EncryptInput(input, workKeyFIle)
	if err != nil {
		return err
	}
	fmt.Print(encryptStr)
	return nil
}
