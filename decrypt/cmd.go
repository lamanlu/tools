package decrypt

import (
	"fmt"

	"github.com/lamanlu/tools/common"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a Base64 string with a work key",
	Long:  "Decrypt a Base64-encoded ciphertext using the specified work key file and output plaintext.",
	RunE:  runCmd,
}

func GetCmd() *cobra.Command {
	return cmd
}

var workKeyFile string
var keyBaseDir string

func init() {
	cmd.Flags().StringVarP(&workKeyFile, "work-key", "k", "", "Work key file name, using for decrypt input string.")
	cmd.Flags().StringVarP(&keyBaseDir, "key-dir", "d", "", "Key base directory containing rootKey/workKey.")
}

func runCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("invalid base64 input string")
	}
	common.SetKeyBaseDir(keyBaseDir)
	input := args[0]
	plain, err := common.DecryptInput(input, workKeyFile)
	if err != nil {
		return err
	}
	fmt.Print(plain)
	return nil
}
