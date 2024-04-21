package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

func fileOrDirectoryExists(pathToFileOrDirectory string) error {
	if _, err := os.Stat(pathToFileOrDirectory); err != nil {
		return err
	}
	return nil
}

func isFileNotFoundError(err error) bool {
	return os.IsNotExist(err)
}

func createDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func removeDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func readDirectory(path string) ([]fs.DirEntry, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	files, err := dir.ReadDir(0)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func readFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	return content, err
}

func createListOfFiles(filesToCreate []string, ticketDirectoryPath string) error {
	for _, filename := range filesToCreate {
		filePath := filepath.Join(ticketDirectoryPath, filename)
		_, err := os.Create(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func globCopy(recipeDirectoryPath string, ticketDirectoryPath string) error {
	files, err := readDirectory(recipeDirectoryPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Println(file)

		if file.Name() == "recipe.json" {
			continue
		}
		if file.IsDir() {
			log(fmt.Sprintf("Can't copy directories: %s/", file.Name()), "error")
			continue
		}

		srcFile, err := os.Open(filepath.Join(recipeDirectoryPath, file.Name()))

		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(filepath.Join(ticketDirectoryPath, file.Name()))
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
		log(fmt.Sprintf("Copy successful - %s", file.Name()), "success")
	}

	return nil
}

func noFilesInTemplateDirectory(files []fs.DirEntry) bool {
	if len(files) == 1 && files[0].Name() == "recipe.json" {
		return true
	}
	return false
}

func copyListOfFiles(recipeDirectoryPath string, ticketDirectoryPath string) error {
	if err := fileOrDirectoryExists(recipeDirectoryPath); err != nil {
		return err
	}

	files, err := readDirectory(recipeDirectoryPath)
	if err != nil {
		return err
	}
	if noFilesInTemplateDirectory(files) {
		log(fmt.Sprintf("No files to copy: %s", recipeDirectoryPath), "warn")
		return nil
	}
	for _, file := range files {
		if file.Name() == "recipe.json" {
			continue
		}
		if file.IsDir() {
			// TODO, implement recursive copy
			log(fmt.Sprintf("Can't copy directories: %s/", file.Name()), "error")
			continue
		}
		srcFile, err := os.Open(filepath.Join(recipeDirectoryPath, file.Name()))

		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(filepath.Join(ticketDirectoryPath, file.Name()))
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
		log(fmt.Sprintf("Copy successful - %s", file.Name()), "success")

	}

	return nil
}

func openDirectory(directoryPath string, envVar envVarStruct) error {
	var command []string
	if envVar.exists {
		command = strings.Split(envVar.value, " ")
	} else {
		switch runtime.GOOS {
		case "darwin":
			command = append(command, "open")
		case "linux":
			command = append(command, "xdg-open")
		case "windows":
			command = append(command, "explorer")
		default:
			return fmt.Errorf("unsupported platform")
		}
		log("No `TCK_EDITOR` environment variable was set", "warn")
		log(fmt.Sprintf("The global default will be used: `%s`", command[0]), "warn")
	}

	rootCommand := command[0]

	if len(command) > 1 {
		args := command[1:]
		args = append(args, directoryPath)
		if err := exec.Command(rootCommand, args...).Start(); err != nil {
			return err
		}
	} else {
		if err := exec.Command(rootCommand, directoryPath).Start(); err != nil {
			return err
		}
	}

	return nil
}

func appendTicketsToTable(files []fs.DirEntry, t table.Writer, path string) {
	t.AppendHeader(table.Row{"Ticket ID", "Description", "Url"})

	for _, file := range files {
		if file.IsDir() {
			metaDataJson, err := readFile(filepath.Join(path, file.Name(), "meta.json"))
			if err != nil {
				log(fmt.Sprintf("error reading meta.json in %s: %v\n", file.Name(), err.Error()), "error")
				continue
			}

			metaDataMap, err := unMarshallMetadata(metaDataJson)
			if err != nil {
				log(fmt.Sprintf("Error parsing meta.json in %s: %v\n", file.Name(), err.Error()), "error")
				continue
			}

			description := metaDataMap["description"]
			urlFormat := metaDataMap["url"]
			t.AppendRow([]interface{}{file.Name(), description, urlFormat})
			t.AppendSeparator()
		}
	}
}
func appendRecipesToTable(files []fs.DirEntry, t table.Writer, path string) {
	t.AppendHeader(table.Row{"Recipes"})
	for i, file := range files {
		if files[i].IsDir() {
			if err := fileOrDirectoryExists(filepath.Join(path, files[i].Name(), "recipe.json")); err != nil {
				log(fmt.Sprintf("error reading recipe.json in %s: %v\n", file.Name(), err.Error()), "error")
				continue
			} else {
				data := fmt.Sprintf("- %s", file.Name())
				t.AppendRow([]interface{}{data})
			}
		}
	}
}

func createFileWithContent(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
