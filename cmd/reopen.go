/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reopenCmd represents the reopen command
var reopenCmd = &cobra.Command{
	Use:   "reopen",
	Short: "Move ticket out of the .closed/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}
		fmt.Print(homeInfo)

		// Check if ticket exists in .closed/ directory
		// If not throw error
		// If yes move to tickets/ directory
	},
}

func init() {
	rootCmd.AddCommand(reopenCmd)
}
