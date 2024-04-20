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
		nav, _ := cmd.Flags().GetBool("nav")
		all, _ := cmd.Flags().GetBool("all")

		var ticket TicketStruct

		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}

		// Handle all for tickets/ vs .closed/
		if all {
			openDirectory(homeInfo.getTicketsPath(), envVars["TCK_EDITOR"])
			return
		}

		if err := ticket.setTicketId(args, envVars["TCK_ID"]); err != nil {
			fatalError(err)
		}
		path := ticket.getTicketDirectory(homeInfo.getTicketsPath())

		if err := openDirectory(path, envVars["TCK_EDITOR"]); err != nil {
			fatalError(err)
		}
		if nav {
			content, err := readFile(filepath.Join(path, "meta.json"))
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
	},
}

func init() {
	getCmd.Flags().BoolP("nav", "n", false, "Navigate to the ticket ID in theb browser.")
	getCmd.Flags().BoolP("all", "a", false, "Open the full tickets/ directory in the editor")
	getCmd.Flags().StringP("closed", "c", "", "Open the given tickets in the .closed/ directory")

	rootCmd.AddCommand(getCmd)
}
