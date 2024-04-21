/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [ticket]",
	Args:  cobra.MaximumNArgs(2),
	Short: "Initialize a new ticket directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		recipeFlag := cmd.Flag("recipe").Value.String()
		descriptionFlag := cmd.Flag("description").Value.String()
		urlFormatFlag := cmd.Flag("url-format").Value.String()

		recipe := setConfigValue(recipeFlag, envVars["TCK_RECIPE"], "default")
		urlFormat := setConfigValue(urlFormatFlag, envVars["TCK_URL_FORMAT"], DefaultUrlFormat)

		if err := urlFormatValidator(urlFormat); err != nil {
			fatalError(err)
		}

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
		if err != nil {
			fatalError(err)
		}

		recipeMap, err := configureRecipe(homeInfo.getRecipesPath(), recipe)
		if err != nil {
			fatalError(err)
		}

		var ticket TicketStruct
		ticketDirectoryPath, err := ticket.setTicketId(args, envVars["TCK_ID"], homeInfo.getTicketsPath())
		if err != nil {
			fatalError(err)
		}

		if err := fileOrDirectoryExists(ticketDirectoryPath); err == nil {
			overWritePrompt := fmt.Sprintf("Overwrite existing directory? %s?\nY to overwrite\nN to cancel", ticketDirectoryPath)
			if err := confirmDirectoryRemove(overWritePrompt, "cancelled", ticketDirectoryPath); err != nil {
				fatalError(err)
			}
		}

		if err := createDirectory(ticketDirectoryPath); err != nil {
			fatalError(err)
		}

		if err := createMetaJson(ticketDirectoryPath, descriptionFlag, urlFormat, ticket.TicketId); err != nil {
			fatalError(err)
		}

		if err := createListOfFiles(recipeMap.FilesToCreate, ticketDirectoryPath); err != nil {
			fatalError(err)
		}

		if err := copyListOfFiles(recipeMap.FilesToCopy, filepath.Join(homeInfo.getRecipesPath(), recipe), ticketDirectoryPath); err != nil {
			fatalError(err)
		}

		log(fmt.Sprintf("Ticket directory initialized: %s/", ticketDirectoryPath), "success")
	},
}

func init() {
	initCmd.Flags().StringP("description", "d", "", "Provide a short description for the ticket you are creating.")
	initCmd.Flags().StringP("recipe", "r", "", "Use the provided recipe file to initialize the ticket directory.")
	initCmd.Flags().StringP("url-format", "u", "", "Format for generating a URL with the ticket ID. Use @ for the ticket ID location")

	rootCmd.AddCommand(initCmd)
}
