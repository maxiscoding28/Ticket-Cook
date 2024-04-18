/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		var homeInfo HomeInfoStruct

		homeInfo.setHomeDirectory(envVars["TC_HOME_DIR"])

		files, err := readDirectory(homeInfo.getTicketsPath())
		if err != nil {
			fatalError(err)
		}

		t := createTable()
		appendFilesToTable(files, t, homeInfo.getTicketsPath())

		t.SetStyle(table.StyleLight)
		t.Render()

	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
