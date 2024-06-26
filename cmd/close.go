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

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Move ticket to the .closed/ directory",
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

		if err := fileOrDirectoryExists(ticketDirectoryPath); err != nil {
			fatalError(err)
		} else {
			err := os.Rename(ticketDirectoryPath, filepath.Join(homeInfo.getClosedPath(), ticket.TicketId))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			log(fmt.Sprintf("Ticket closed - %s", ticket.TicketId), "success")
		}
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
