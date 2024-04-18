package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const DefaultTemplateJson = `{
	"filesToCreate": [
		"scratch.sh",
		"notes.md"
	],
	"filesToCopyDir": "",
	"filesToCopy": [],
	"filesToExecuteDir": "",
	"filesToExecute": []
}`

type TemplateMapStruct struct {
	FilesToCreate     []string `json:"filesToCreate"`
	FilesToCopy       []string `json:"filesToCopy"`
	FilesToCopyDir    string   `json:"filesToCopyDir"`
	FilesToExecute    []string `json:"filesToExecute"`
	FilesToExecuteDir string   `json:"filesToExecuteDir"`
}

func configureTemplate(templatePath string, templateArg string) (*TemplateMapStruct, error) {
	if err := fileOrDirectoryExists(templatePath); err != nil {
		if fileDoesntExist(err) {
			if err := createDirectory(templatePath); err != nil {
				return nil, err
			}
		}
	}
	templateFilePath := getTemplateFilePath(templatePath, templateArg)

	if templateArgIsDefault(templateArg) && fileOrDirectoryExists(templateFilePath) != nil {
		if err := createDefaultTemplate(templateFilePath); err != nil {
			return nil, err
		}
	}

	if err := fileOrDirectoryExists(templateFilePath); err != nil {
		if fileDoesntExist(err) {
			return nil, fmt.Errorf("template doesn't exist: %s", templateArg)
		}
		return nil, err
	}

	templateMap, err := templateisValid(templateFilePath)
	if err != nil {
		return nil, err
	}

	return templateMap, nil
}

func getTemplateFilePath(templatePath string, templateArg string) string {
	return templatePath + "/" + templateArg + ".json"
}

func templateArgIsDefault(templateArg string) bool {
	return templateArg == "default"
}

func createDefaultTemplate(templateFilePath string) error {
	if err := createFileWithContent(templateFilePath, DefaultTemplateJson); err != nil {
		return err
	}
	return nil
}

func templateisValid(pathToTemplate string) (*TemplateMapStruct, error) {
	content, err := os.ReadFile(pathToTemplate)
	if err != nil {
		return nil, err
	}
	var data TemplateMapStruct
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	if len(data.FilesToCopy) > 0 && len(data.FilesToCopyDir) == 0 {
		return nil, errors.New("filesToCopyDir must be set if copying files")
	}
	if len(data.FilesToExecute) > 0 && len(data.FilesToExecuteDir) == 0 {
		return nil, errors.New("filesToExecuteDir must be set if copying files")
	}

	return &data, nil
}
