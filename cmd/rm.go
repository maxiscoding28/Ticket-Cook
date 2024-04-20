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

		var ticket TicketStruct

		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}

		ticketPath := ticket.getTicketDirectory(homeInfo.getTicketsPath())

		if err := fileOrDirectoryExists(ticketPath); err == nil {
			removePrompt := fmt.Sprintf("Remove directory? %s?\nY to remove\nN to cancel", ticketPath)
			if err := confirmDirectoryRemove(removePrompt, "removed", ticketPath); err != nil {
				fatalError(err)
			}
		}
		if err := removeDirectory(ticket.getTicketDirectory(homeInfo.getTicketsPath())); err != nil {
			fatalError(err)
		}
		log(fmt.Sprintf("Ticket directory removed (if it existed): %s", ticket.TicketId), "success")

	},
}

func init() {
	rmCmd.Flags().StringP("closed", "c", "", "Remove the given ticket from the .closed/ directory")
	rootCmd.AddCommand(rmCmd)
}
