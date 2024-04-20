/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getRecipeCmd represents the getRecipe command
var getRecipeCmd = &cobra.Command{
	Use:   "get",
	Short: "Open the given recipe in the recipes/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getRecipe called")
	},
}

func init() {
	recipeCmd.AddCommand(getRecipeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getRecipeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getRecipeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
