package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/enescakir/emoji"
	"github.com/jedib0t/go-pretty/v6/table"
)

type envVarStruct struct {
	envVar string
	value  string
	exists bool
}

func log(message string, level string) {
	colorRed := "\033[31m"
	colorYellow := "\033[33m"
	colorGreen := "\033[32m"
	colorBlue := "\033[34m"
	colorReset := "\033[0m"

	var color string

	if level == "error" {
		color = colorRed
	} else if level == "warn" {
		color = colorYellow
	} else if level == "success" {
		color = colorGreen
	} else if level == "info" {
		color = colorBlue
	} else {
		color = colorReset
	}

	messageString := fmt.Sprintf("%s: %s%s%s", emoji.ManCook.String(), color, message, colorReset)
	fmt.Println(messageString)
}

func tckBanner() {
	blue := "\033[1;34m%s\033[0m\n"
	reset := "\033[0m"
	fmt.Printf(blue, "=============================")
	fmt.Printf(blue, fmt.Sprintf("%s This is Ticket Cook %s", emoji.ManCook.String(), emoji.ManCook.String()))
	fmt.Printf(blue, "=============================")
	fmt.Print(reset)
}

func getEnvVars() map[string]envVarStruct {
	envVarNames := []string{"TCK_ID", "TCK_RECIPE", "TCK_HOME_DIR", "TCK_EDITOR", "TCK_URL_FORMAT"}

	envVarMap := make(map[string]envVarStruct)

	for _, envVarName := range envVarNames {
		envValue, exists := os.LookupEnv(envVarName)
		envVar := envVarStruct{
			envVar: envVarName,
			value:  envValue,
			exists: exists,
		}
		envVarMap[envVarName] = envVar
	}
	return envVarMap
}

func fatalError(err error) {
	log(fmt.Sprintf("Fatal error: %s", err.Error()), "error")
	os.Exit(1)
}

func confirmDirectoryRemove(promptMessage string, verb string, directory string) error {
	var userInput string
	log(promptMessage, "error")
	fmt.Scanln(&userInput)
	userInput = strings.ToUpper(userInput)
	switch userInput {
	case "Y":
		removeDirectory(directory)
		log(fmt.Sprintf("Diretory %s: %s", verb, directory), "info")
	case "N":
		return errors.New("operation cancelled")
	default:
		return errors.New("invalid input - operation cancelled")
	}

	return nil
}

func createTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Ticket ID", "Description", "Url"})

	return t
}

func openUrl(urlString string) error {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return err
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return errors.New("URL scheme must be http or https")
	}
	if err := exec.Command("open", urlString).Start(); err != nil {
		return err
	}
	return nil
}

func setConfigValue(arg string, envVar envVarStruct, globalDefault string) string {
	if len(arg) > 0 {
		return arg
	} else if envVar.exists {
		return envVar.value
	} else {
		return globalDefault
	}
}
