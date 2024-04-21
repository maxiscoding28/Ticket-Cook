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

// rmRecipeCmd represents the rmRecipes command
var rmRecipeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the given recipe from the recipes/ directory",

	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
		if err != nil {
			fatalError(err)
		}

		if len(args) != 1 {
			fatalError(errors.New("exactly 1 recipe name argument is required"))
		}

		recipe := args[0]
		recipePath := filepath.Join(homeInfo.getRecipesPath(), recipe)

		if recipe == "default" {
			fatalError(errors.New("default recipe cannot be removed"))
		}

		if err := fileOrDirectoryExists(recipePath); err == nil {
			removePrompt := fmt.Sprintf("Remove directory? %s?\nY to remove\nN to cancel", recipePath)
			if err := confirmDirectoryRemove(removePrompt, "removed", recipePath); err != nil {
				fatalError(err)
			}
		} else {
			fatalError(err)
		}
	},
}

func init() {
	recipeCmd.AddCommand(rmRecipeCmd)
}
