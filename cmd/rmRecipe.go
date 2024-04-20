/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmRecipeCmd represents the rmRecipes command
var rmRecipeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the given recipe from the recipes/ directory",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmRecipes called")
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
