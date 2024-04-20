package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const DefaultRecipeJson = `{
	"filesToCreate": [
		"scratch.sh",
		"notes.md"
	],
	"filesToCopy": []
}`

type RecipeMapStruct struct {
	FilesToCreate []string `json:"filesToCreate"`
	FilesToCopy   []string `json:"filesToCopy"`
}

func globCopy(files []string) bool {
	return len(files) == 1 && files[0] == "*"
}

func handleDefaultTemplate(recipeDirectory string, recipeJsonFilePath string) error {
	if err := fileOrDirectoryExists(recipeDirectory); err != nil {
		if fileOrDirectoryDoesNotExist(err) {
			if err := createDirectory(recipeDirectory); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if err := fileOrDirectoryExists(recipeJsonFilePath); err != nil {
		if fileOrDirectoryDoesNotExist(err) {
			if err := createDefaultRecipe(recipeJsonFilePath); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func configureRecipe(recipePath string, recipeArg string) (*RecipeMapStruct, error) {
	recipeDirectory := getRecipeDirectory(recipePath, recipeArg)
	recipeJsonFilePath := getRecipeJsonFilePath(recipeDirectory)

	if recipeArgIsDefault(recipeArg) {
		if err := handleDefaultTemplate(recipeDirectory, recipeJsonFilePath); err != nil {
			return nil, err
		}
	}

	if err := fileOrDirectoryExists(recipeDirectory); err != nil {
		if fileOrDirectoryDoesNotExist(err) {
			return nil, fmt.Errorf("recipe directory doesn't exist: %s", recipeDirectory)
		}
		return nil, err
	}
	if err := fileOrDirectoryExists(recipeJsonFilePath); err != nil {
		if fileOrDirectoryDoesNotExist(err) {
			return nil, fmt.Errorf("directory has no `recipe.json` file: %s", recipeDirectory)
		}
		return nil, err
	}

	recipeMap, err := recipeIsValid(recipeJsonFilePath)
	if err != nil {
		return nil, err
	}

	return recipeMap, nil
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

func createDefaultRecipe(recipeFilePath string) error {
	if err := createFileWithContent(recipeFilePath, DefaultRecipeJson); err != nil {
		return err
	}
	return nil
}

func recipeIsValid(pathToRecipe string) (*RecipeMapStruct, error) {
	content, err := os.ReadFile(pathToRecipe)
	if err != nil {
		return nil, err
	}
	var data RecipeMapStruct
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return &data, nil
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
