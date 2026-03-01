/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/lamanlu/tools/decrypt"
	"github.com/lamanlu/tools/encrypt"
	"github.com/lamanlu/tools/keys"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "CLI utilities for key generation and encryption",
	Long:  "CLI utilities for generating root/work keys and encrypting strings using work keys.",
}

func setCmds() {
	for _, cmd := range keys.GetCmds() {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.AddCommand(encrypt.GetCmd())
	rootCmd.AddCommand(decrypt.GetCmd())
}

func main() {

	setCmds()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("check", "c", false, "Check cmd run conditions")
}
