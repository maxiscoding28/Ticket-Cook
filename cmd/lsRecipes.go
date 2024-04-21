/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// lsRecipesCmd represents the lsRecipes command
var lsRecipesCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all recipes in the recipe/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		homeInfo, err := getHomeDirectory(envVars["TCK_HOME_DIR"])
		if err != nil {
			fatalError(err)
		}

		renderRecipes(homeInfo.getRecipesPath())
	},
}

func init() {
	recipeCmd.AddCommand(lsRecipesCmd)
}
