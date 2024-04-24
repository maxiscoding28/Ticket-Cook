package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func handleDefaultTemplate(recipeDirectory string) error {
	if err := fileOrDirectoryExists(recipeDirectory); err != nil {
		if isFileNotFoundError(err) {
			if err := createDirectory(recipeDirectory); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func configureRecipe(recipePath string, recipeArg string) error {
	recipeDirectory := getRecipeDirectory(recipePath, recipeArg)

	if recipeArgIsDefault(recipeArg) {
		if err := handleDefaultTemplate(recipeDirectory); err != nil {
			return err
		}
	}

	if err := fileOrDirectoryExists(recipeDirectory); err != nil {
		if isFileNotFoundError(err) {
			return fmt.Errorf("recipe directory doesn't exist: %s", recipeDirectory)
		}
		return err
	}
	return nil
}

func getRecipeDirectory(recipePath string, recipeArg string) string {
	return recipePath + "/" + recipeArg
}

func getRecipeJsonFilePath(recipeDirectory string) string {
	return recipeDirectory + "/recipe.json"
}

func recipeArgIsDefault(recipeArg string) bool {
	return recipeArg == "default"
}

var recipeCmd = &cobra.Command{
	Use:   "recipe",
	Short: "Manage recipe directories",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(recipeCmd)
}
