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

// getRecipeCmd represents the getRecipe command
var getRecipeCmd = &cobra.Command{
	Use:   "get",
	Short: "Open the given recipe in the recipes/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()
		all, _ := cmd.Flags().GetBool("all")

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
		if err != nil {
			fatalError(err)
		}
		if all {
			openDirectory(homeInfo.getRecipesPath(), envVars["TCK_EDITOR"])
		} else {
			if len(args) != 1 {
				fatalError(errors.New("exactly 1 recipe name argument is required"))
			}
			recipe := args[0]
			recipePath := filepath.Join(homeInfo.getRecipesPath(), recipe)

			if err := fileOrDirectoryExists(recipePath); err == nil {
				if err := openDirectory(recipePath, envVars["TCK_EDITOR"]); err != nil {
					fatalError(err)
				}
				log(fmt.Sprintf("Recipe directory opened: %s", recipe), "success")
			} else {
				fatalError(err)
			}

		}
	},
}

func init() {
	getRecipeCmd.Flags().BoolP("all", "a", false, "Open the full tickets/ directory in the editor")

	recipeCmd.AddCommand(getRecipeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getRecipeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getRecipeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
