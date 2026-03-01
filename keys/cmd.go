/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package keys

import (
	"fmt"

	"github.com/lamanlu/tools/common"
	"github.com/spf13/cobra"
)

var forceFlag bool
var workKeyFile string
var keyBaseDir string

func GetCmds() []*cobra.Command {
	return []*cobra.Command{
		newGenRootKeyCmd(),
		newGenWorkKeyCmd(),
		newGenRandomKeyCmd(),
	}
}

func applyBaseFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&keyBaseDir, "dir", "d", "", "Key base directory. Will create rootKey/workKey under it.")
}

func newGenRootKeyCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gen-root-key",
		Short: "Generate root key files",
		Long:  "Generate root keys (rootKey/root_part_*.key + rootKey/root.salt).",
		RunE: func(cmd *cobra.Command, args []string) error {
			common.SetKeyBaseDir(keyBaseDir)
			if forceFlag {
				common.ClearAllKeys()
			}
			if err := common.CreateRootKeySalt(); err != nil {
				return err
			}
			fmt.Println("Create Root Key Salt Done")
			if err := common.CreateRootKeyParts(); err != nil {
				return err
			}
			fmt.Println("Create Root Key Done")
			return nil
		},
	}
	applyBaseFlags(rootCmd)
	rootCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force Create RootKey, Ignore Exist key.")
	return rootCmd
}

func newGenWorkKeyCmd() *cobra.Command {
	workCmd := &cobra.Command{
		Use:   "gen-work-key",
		Short: "Generate a work key file",
		Long:  "Generate a work key encrypted by the root key.",
		RunE: func(cmd *cobra.Command, args []string) error {
			common.SetKeyBaseDir(keyBaseDir)
			if forceFlag {
				if err := common.ClearWorkKey(workKeyFile); err != nil {
					return err
				}
			}
			if err := common.CreateWorkKey(workKeyFile); err != nil {
				return err
			}
			fmt.Printf("Create Work Key: %s Done\n", workKeyFile)
			return nil
		},
	}
	applyBaseFlags(workCmd)
	workCmd.Flags().StringVarP(&workKeyFile, "name", "n", "work.key", "Work Key File Name. eg: work.key")
	workCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force Create WorkKey, Ignore Exist key.")
	return workCmd
}

func newGenRandomKeyCmd() *cobra.Command {
	randomCmd := &cobra.Command{
		Use:   "gen-random-key",
		Short: "Generate a random key file",
		Long:  "Generate a random key file (Base64, 32 bytes) in workKey directory.",
		RunE: func(cmd *cobra.Command, args []string) error {
			common.SetKeyBaseDir(keyBaseDir)
			if forceFlag {
				if err := common.ClearWorkKey(workKeyFile); err != nil {
					return err
				}
			}
			if err := common.CreateRandomKeyFile(workKeyFile); err != nil {
				return err
			}
			fmt.Printf("Create Random Key File: %s Done\n", workKeyFile)
			return nil
		},
	}
	applyBaseFlags(randomCmd)
	randomCmd.Flags().StringVarP(&workKeyFile, "name", "n", "random.key", "Random Key File Name. eg: random.key")
	randomCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force Create RandomKey, Ignore Exist key.")
	return randomCmd
}
