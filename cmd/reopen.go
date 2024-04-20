/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		var ticket TicketStruct
		if err := ticket.setTicketId(args, envVars["TCK_ID"]); err != nil {
			fatalError(err)
		}

		ticketDirectoryPath := filepath.Join(homeInfo.getClosedPath(), ticket.TicketId)

		if err := fileOrDirectoryExists(ticketDirectoryPath); err != nil {
			fatalError(err)
		} else {
			err := os.Rename(ticketDirectoryPath, filepath.Join(homeInfo.getTicketsPath(), ticket.TicketId))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			log(fmt.Sprintf("Ticket reopened - %s", ticket.TicketId), "success")
		}
	},
}

func init() {
	rootCmd.AddCommand(reopenCmd)
}
