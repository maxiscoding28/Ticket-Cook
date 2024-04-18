package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"

	"github.com/jedib0t/go-pretty/v6/table"
)

type envVarStruct struct {
	envVar string
	value  string
	exists bool
}

func log(message string, err error) {
	colorRed := "\033[31m"
	colorYellow := "\033[33m"
	colorReset := "\033[0m"

	var color string

	if err != nil {
		color = colorRed
	} else {
		color = colorYellow
	}

	fmt.Printf("%s%s%s\n", color, message, colorReset)
}

func success(message string) {
	colorGreen := "\033[32m"
	colorReset := "\033[0m"

	fmt.Printf("%s%s%s\n", colorGreen, message, colorReset)
}

func getEnvVars() map[string]envVarStruct {
	envVarNames := []string{"TC_ID", "TC_TEMPLATE", "TC_HOME_DIR", "TC_EDITOR", "TC_URL_FORMAT"}

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
	log(fmt.Sprintf("Fatal error: %s", err.Error()), err)
	os.Exit(1)
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
