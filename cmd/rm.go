// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the given ticket from the ticket/ directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
		if err != nil {
			fatalError(err)
		}

		var ticket TicketStruct
		ticketDirectoryPath, err := ticket.setTicketId(args, envVars["TCK_ID"], homeInfo.getTicketsPath())
		if err != nil {
			fatalError(err)
		}

		if err := fileOrDirectoryExists(ticketDirectoryPath); err == nil {
			removePrompt := fmt.Sprintf("Remove directory? %s?\nY to remove\nN to cancel", ticketDirectoryPath)
			if err := confirmDirectoryRemove(removePrompt, "removed", ticketDirectoryPath); err != nil {
				fatalError(err)
			}
		} else {
			fatalError(err)
		}
	},
}

func init() {
	rmCmd.Flags().StringP("closed", "c", "", "Remove the given ticket from the .closed/ directory")
	rootCmd.AddCommand(rmCmd)
}
