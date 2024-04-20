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

		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}

		if len(args) != 1 {
			fatalError(errors.New("exactly 1 recipe name argument is required"))
		}

		recipePath := filepath.Join(homeInfo.getRecipesPath(), args[0])

		// Make sure space doesn't exist

		if args[0] == "default" {
			fatalError(errors.New("default recipe cannot be removed"))
		}

		if err := fileOrDirectoryExists(recipePath); err == nil {
			removePrompt := fmt.Sprintf("Remove directory? %s?\nY to remove\nN to cancel", recipePath)
			if err := confirmDirectoryRemove(removePrompt, "removed", recipePath); err != nil {
				fatalError(err)
			}
		}
		if err := removeDirectory(recipePath); err != nil {
			fatalError(err)
		}
		log(fmt.Sprintf("Ticket directory removed (if it existed): %s", args[0]), "success")
	},
}

func init() {
	recipeCmd.AddCommand(rmRecipeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmRecipeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmRecipeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
