/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [ticket]",
	Short: "Open the given ticket in the ticket/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		closed, _ := cmd.Flags().GetBool("closed")
		nav, _ := cmd.Flags().GetBool("nav")
		all, _ := cmd.Flags().GetBool("all")

		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}

		var rootPath string
		if closed {
			rootPath = homeInfo.getClosedPath()
		} else {
			rootPath = homeInfo.getTicketsPath()
		}
		if all {
			openDirectory(rootPath, envVars["TCK_EDITOR"])
			return
		} else {
			var ticket TicketStruct
			if err := ticket.setTicketId(args, envVars["TCK_ID"]); err != nil {
				fatalError(err)
			}
			ticketPath := ticket.getPath(rootPath)

			if err := fileOrDirectoryExists(ticketPath); err == nil {
				fmt.Println(ticketPath)
				if err := openDirectory(ticketPath, envVars["TCK_EDITOR"]); err != nil {
					fatalError(err)
				}
				log(fmt.Sprintf("Ticket directory opened: %s", ticket.TicketId), "success")
				if nav {
					content, err := readFile(filepath.Join(ticketPath, "meta.json"))
					if err != nil {
						fatalError(err)
					}
					meta, err := unMarshallMetadata(content)
					if err != nil {
						fatalError(err)
					}
					log(fmt.Sprintf("The `-nav` flag was set. Opening url: %s", meta["url"]), "info")

					if err := openUrl(meta["url"]); err != nil {
						fatalError(err)
					}
				}
			} else {
				fatalError(err)
			}

		}

	},
}

func init() {
	getCmd.Flags().BoolP("nav", "n", false, "Navigate to the ticket ID in theb browser.")
	getCmd.Flags().BoolP("all", "a", false, "Open the full tickets/ directory in the editor")
	getCmd.Flags().BoolP("closed", "c", false, "Open the given tickets in the .closed/ directory")

	rootCmd.AddCommand(getCmd)
}
