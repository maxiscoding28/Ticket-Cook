package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
)

func fileOrDirectoryExists(pathToFileOrDirectory string) error {
	if _, err := os.Stat(pathToFileOrDirectory); err != nil {
		return err
	}
	return nil
}

func fileDoesntExist(err error) bool {
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

func copyListOfFiles(filesToCopy []string, filesToCopyDir string, ticketDirectoryPath string) error {
	if len(filesToCopy) > 0 {
		if err := fileOrDirectoryExists(filesToCopyDir); err != nil {
			return err
		}
		var sourceFile string
		var destinationFile string
		for _, file := range filesToCopy {
			sourceFile = filepath.Join(filesToCopyDir, file)
			destinationFile = filepath.Join(ticketDirectoryPath, filepath.Base(file))
			source, err := os.Open(sourceFile)
			if err != nil {
				log("There was a problem opening the source file:", err)
				continue
			}
			defer source.Close()
			destination, err := os.Create(destinationFile)
			if err != nil {
				log("There was a problem opening the destination file:", err)
				continue
			}
			defer destination.Close()

			_, err = io.Copy(destination, source)
			if err != nil {
				log("There was a problem copying the file:", err)
				continue
			}
		}
		return nil
	}
	return nil
}

func executeListOfFiles(filesToExecute []string, filesToExecuteDir string, ticketDirectoryPath string) error {
	if len(filesToExecute) > 0 {
		if err := fileOrDirectoryExists(filesToExecuteDir); err != nil {
			return err
		}

		for _, file := range filesToExecute {
			cmd := exec.Command("sh", filepath.Join(filesToExecuteDir, file))
			cmd.Dir = ticketDirectoryPath
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	}
	return nil
}

func openFile(cmd []string, filepath string) error {
	cmdArg := cmd[0]
	if len(cmd) > 1 {
		args := cmd[1:]
		args = append(args, filepath)
		if err := exec.Command(cmdArg, args...).Start(); err != nil {
			return err
		}
	} else {
		if err := exec.Command(cmdArg, filepath).Start(); err != nil {
			return err
		}
	}

	return nil
}

func appendFilesToTable(files []fs.DirEntry, table table.Writer, homePath string) {

	for _, file := range files {
		if file.IsDir() {
			metaDataJson, err := readFile(filepath.Join(homePath, file.Name(), "meta.json"))
			if err != nil {
				log(fmt.Sprintf("error reading meta.json in %s: %v\n", file.Name(), err), err)
				continue
			}

			metaDataMap, err := unMarshallMetadata(metaDataJson)
			if err != nil {
				log(fmt.Sprintf("Error parsing meta.json in %s: %v\n", file.Name(), err), err)
				continue
			}

			description := metaDataMap["description"]
			urlFormat := metaDataMap["url"]
			table.AppendRow([]interface{}{file.Name(), description, urlFormat})
			table.AppendSeparator()
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
