/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [ticket]",
	Args:  cobra.MaximumNArgs(2),
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		templateArg := cmd.Flag("template").Value.String()
		description := cmd.Flag("description").Value.String()
		urlFormatArg := cmd.Flag("url-format").Value.String()

		template := setConfigValue(templateArg, envVars["TC_TEMPLATE"], "default")
		urlFormat := setConfigValue(urlFormatArg, envVars["TC_URL_FORMAT"], "https://hashicorp.zendesk.com/agent/tickets/@")

		if err := urlFormatValidator(urlFormat); err != nil {
			fatalError(err)
		}

		var homeInfo HomeInfoStruct
		var newTicket TicketStruct

		homeInfo.setHomeDirectory(envVars["TC_HOME_DIR"])
		if err := newTicket.setTicketId(args, envVars["TC_ID"]); err != nil {
			fatalError(err)
		}

		templateMap, err := configureTemplate(homeInfo.getTemplatesPath(), template)
		if err != nil {
			fatalError(err)
		}

		ticketPath := newTicket.getTicketDirectory(homeInfo.getTicketsPath())

		// Overwrite check
		if err := createDirectory(ticketPath); err != nil {
			fatalError(err)
		}

		if err := createMetaJson(ticketPath, description, urlFormat, newTicket.TicketId); err != nil {
			fatalError(err)
		}

		if err := createListOfFiles(templateMap.FilesToCreate, ticketPath); err != nil {
			fatalError(err)
		}

		success(fmt.Sprintf("Ticket directory initialized: %s", newTicket.TicketId))

		if err := copyListOfFiles(templateMap.FilesToCopy, templateMap.FilesToCopyDir, ticketPath); err != nil {
			fatalError(err)
		}

		if err := executeListOfFiles(templateMap.FilesToExecute, templateMap.FilesToExecuteDir, ticketPath); err != nil {
			fatalError(err)
		}
	},
}

func init() {
	initCmd.Flags().StringP("description", "d", "", "Provide a short description for the ticket you are creating.")
	initCmd.Flags().StringP("template", "t", "", "Use the provided template file to initialize the ticket directory.")
	initCmd.Flags().StringP("url-format", "u", "", "Format for generating a URL with the ticket ID. Use @ for the ticket ID location")

	rootCmd.AddCommand(initCmd)
}
