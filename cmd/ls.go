/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all tickets in the ticket/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		closed, _ := cmd.Flags().GetBool("closed")

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
		if err != nil {
			fatalError(err)
		}

		if closed {
			renderTickets(homeInfo.getClosedPath())
		} else {
			renderTickets(homeInfo.getTicketsPath())
		}
	},
}

func init() {
	lsCmd.Flags().BoolP("closed", "c", false, "List all tickets in the .closed/ directory")

	rootCmd.AddCommand(lsCmd)
}
