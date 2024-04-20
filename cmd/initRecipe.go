/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initRecipeCmd represents the initRecipe command
var initRecipeCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new recipe directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initRecipe called")
	},
}

func init() {
	recipeCmd.AddCommand(initRecipeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initRecipeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initRecipeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
