// /*
// Copyright © 2024 NAME HERE <EMAIL ADDRESS>
// */
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the given ticket from the ticket/ directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}

		var ticket TicketStruct
		if err := ticket.setTicketId(args, envVars["TCK_ID"]); err != nil {
			fatalError(err)
		}
		ticketPath := filepath.Join(homeInfo.getTicketsPath(), ticket.TicketId)

		if err := fileOrDirectoryExists(ticketPath); err == nil {
			removePrompt := fmt.Sprintf("Remove directory? %s?\nY to remove\nN to cancel", ticketPath)
			if err := confirmDirectoryRemove(removePrompt, "removed", ticketPath); err != nil {
				fatalError(err)
			}
		}
		if err := removeDirectory(filepath.Join(homeInfo.getTicketsPath(), ticket.TicketId)); err != nil {
			fatalError(err)
		}
		log(fmt.Sprintf("Ticket directory removed (if it existed): %s", ticket.TicketId), "success")

	},
}

func init() {
	rmCmd.Flags().StringP("closed", "c", "", "Remove the given ticket from the .closed/ directory")
	rootCmd.AddCommand(rmCmd)
}
