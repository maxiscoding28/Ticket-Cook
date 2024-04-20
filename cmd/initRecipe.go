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

// initRecipeCmd represents the initRecipe command
var initRecipeCmd = &cobra.Command{
	Use:   "init [recipe]",
	Short: "Initialize a new recipe directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initRecipe called")
		envVars := getEnvVars()

		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}

		// var recipeMap RecipeMapStruct

		if len(args) != 1 {
			fatalError(errors.New("exactly 1 recipe name argument is required"))
		}

		recipePath := filepath.Join(homeInfo.getRecipesPath(), args[0])

		if err := fileOrDirectoryExists(recipePath); err != nil {
			if isFileNotFoundError(err) {
				// Create recipe
				// Create starter files
				// start.md - documentation on how to write recipes
			} else {
				fatalError(err)
			}
		} else {
			// Overwrite flow
		}
	},
}

func init() {
	recipeCmd.AddCommand(initRecipeCmd)
	// Prevent creating a default recipe
}
