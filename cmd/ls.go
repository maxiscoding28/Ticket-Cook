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
	Short: "List all tickets in the ticket/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		envVars := getEnvVars()

		homeInfo, err := setHomeDirectory(envVars["TCK_HOME_DIR"], false)
		if err != nil {
			fatalError(err)
		}

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
	lsCmd.Flags().StringP("closed", "c", "", "List all tickets in the .closed/ directory")

	rootCmd.AddCommand(lsCmd)
}
