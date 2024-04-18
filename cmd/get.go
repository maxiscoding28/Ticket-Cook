/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		nav, _ := cmd.Flags().GetBool("nav")
		var homeInfo HomeInfoStruct
		var ticket TicketStruct

		homeInfo.setHomeDirectory(envVars["TC_HOME_DIR"])
		if err := ticket.setTicketId(args, envVars["TC_ID"]); err != nil {
			fatalError(err)
		}
		path := ticket.getTicketDirectory(homeInfo.getTicketsPath())

		if err := openTicketDirectory(path, envVars["TC_EDITOR"]); err != nil {
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
			if err := openUrl(meta["url"]); err != nil {
				fatalError(err)
			}
		}
	},
}

func init() {
	getCmd.Flags().BoolP("nav", "n", false, "Navigate to the ticket ID in theb browser.")
	rootCmd.AddCommand(getCmd)
}
