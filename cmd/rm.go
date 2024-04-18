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
	Short: "",
	Args:  cobra.MaximumNArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		var homeInfo HomeInfoStruct
		var ticket TicketStruct

		homeInfo.setHomeDirectory(envVars["TC_HOME_DIR"])
		if err := ticket.setTicketId(args, envVars["TC_ID"]); err != nil {
			fatalError(err)
		}

		// Remove confirm
		if err := removeDirectory(ticket.getTicketDirectory(homeInfo.getTicketsPath())); err != nil {
			fatalError(err)
		}
		success(fmt.Sprintf("Ticket directory removed (if it existed): %s", ticket.TicketId))

	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
