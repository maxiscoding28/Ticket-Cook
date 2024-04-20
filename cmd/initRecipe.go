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
	// Prevent creating a default recipe
}
