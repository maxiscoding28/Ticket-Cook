/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Move ticket to the .closed/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}
		fmt.Print(homeInfo)

		// Check if ticket exists in ticket directory
		// If not throw error
		// If yes move to .closed directory
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
