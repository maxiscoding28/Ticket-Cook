/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func initializeTemplate(recipePath string) {
	if err := createDirectory(recipePath); err != nil {
		fatalError(err)
	}
	if err := createFileWithContent(getRecipeJsonFilePath(recipePath), DefaultRecipeJson); err != nil {
		fatalError(err)
	}
	if err := createFileWithContent(filepath.Join(recipePath, "start.md"), StarMdFileContent); err != nil {
		fatalError(err)
	}
	log(fmt.Sprintf("Recipe directory initialized: %s/", filepath.Base(recipePath)), "success")
}

var initRecipeCmd = &cobra.Command{
	Use:   "init [recipe]",
	Short: "Initialize a new recipe directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
		if err != nil {
			fatalError(err)
		}

		if len(args) != 1 {
			fatalError(errors.New("exactly 1 recipe name argument is required"))
		}

		//  TODO: Make sure space doesn't exist

		recipe := args[0]
		recipePath := filepath.Join(homeInfo.getRecipesPath(), recipe)

		if err := fileOrDirectoryExists(recipePath); err != nil {
			if isFileNotFoundError(err) {
				initializeTemplate(recipePath)
			} else {
				fatalError(err)
			}
		} else {
			removePrompt := fmt.Sprintf("Remov existing recipe? %s?\nY to remove\nN to cancel", recipePath)
			if err := confirmDirectoryRemove(removePrompt, "removed", recipePath); err != nil {
				fatalError(err)
			}
			initializeTemplate(recipePath)
		}
	},
}

func init() {
	recipeCmd.AddCommand(initRecipeCmd)
}
