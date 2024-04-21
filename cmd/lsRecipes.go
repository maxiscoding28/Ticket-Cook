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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsRecipesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsRecipesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
