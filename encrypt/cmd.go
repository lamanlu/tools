package encrypt

import (
	"fmt"

	"github.com/lamanlu/tools/keys"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "encrypt",
	Short: "",
	Long:  ``,
	Run:   runCmd,
}

func GetCmd() *cobra.Command {
	return cmd
}

var workKeyFIle string

func init() {
	cmd.Flags().StringVarP(&workKeyFIle, "work-key", "k", "", "Work key file name, using for encrypt input string.")
}

func runCmd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Invalid input string")
		return
	}
	input := args[0]
	fmt.Println("Input is: " + input)
	encryptStr, err := keys.EncryptInput(input, workKeyFIle)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Encrypt successfully.")
	fmt.Println(encryptStr)
}
