/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

func bootstrapDirectories(homeInfo HomeInfoStruct) error {
	logSnippet := []string{"Home", "tickets/", "recipes/", ".closed/"}
	directories := []string{homeInfo.HomePath, homeInfo.getTicketsPath(), homeInfo.getRecipesPath(), homeInfo.getClosedPath()}

	for i := 0; i <= 3; i++ {
		if err := createDirectory(directories[i]); err != nil {
			return err
		} else {
			log(fmt.Sprintf("%s directory created - %s", logSnippet[i], directories[i]), "success")
		}
	}

	return nil
}

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Boostrap the tck home directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], true)
		if err != nil {
			fatalError(err)
		}
		if err := fileOrDirectoryExists(homeInfo.HomePath); err != nil {
			if isFileNotFoundError(err) {
				bootstrapDirectories(*homeInfo)
			} else {
				fatalError(err)
			}
		} else {
			message := fmt.Sprintf("%s This action will permanently remove your existing home directory. Are you sure you'd like to proceed?", emoji.Skull)
			overWritePrompt := fmt.Sprintf("Overwrite existing directory? %s?\nY to overwrite\nN to cancel", homeInfo.HomePath)

			log(message, "error")

			if err := confirmDirectoryRemove(overWritePrompt, "cancelled", homeInfo.HomePath); err != nil {
				fatalError(err)
			} else {
				bootstrapDirectories(*homeInfo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}
