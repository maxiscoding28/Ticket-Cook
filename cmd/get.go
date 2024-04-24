/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func handleOpen(directoryPath string, envVar envVarStruct, nav bool) error {
	if err := fileOrDirectoryExists(directoryPath); err == nil {
		if err := openDirectory(directoryPath, envVar); err != nil {
			return err
		}
		log(fmt.Sprintf("Ticket directory opened: %s", directoryPath), "success")
		if nav {
			content, err := readFile(filepath.Join(directoryPath, "meta.json"))
			if err != nil {
				return err
			}
			meta, err := unMarshallMetadata(content)
			if err != nil {
				return err
			}
			log(fmt.Sprintf("The `-nav` flag was set. Opening url: %s", meta["url"]), "info")

			if err := openUrl(meta["url"]); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

var getCmd = &cobra.Command{
	Use:   "get [ticket]",
	Short: "Open the given ticket in the ticket/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		closed, _ := cmd.Flags().GetBool("closed")
		nav, _ := cmd.Flags().GetBool("nav")
		all, _ := cmd.Flags().GetBool("all")

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
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
		} else {
			var ticket TicketStruct
			directoryPath, err := ticket.setTicketId(args, envVars["TCK_ID"], rootPath)
			if err != nil {
				fatalError(err)
			}

			if err := handleOpen(directoryPath, envVars["TCK_EDITOR"], nav); err != nil {
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
