/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const DefaultUrlFormat string = "https://hashicorp.zendesk.com/agent/tickets/@"

var DefaultHomeDirectory string = filepath.Join(os.Getenv("HOME"), "dev/sandbox-go/tck/test")

var rootCmd = &cobra.Command{
	Use:   "tck",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
