package cmd

import (
	"os"
	"path/filepath"
)

const DefaultUrlFormat string = "https://hashicorp.zendesk.com/agent/tickets/@"

var DefaultHomeDirectory string = filepath.Join(os.Getenv("HOME"), "dev/sandbox-go/tck/test")

const DefaultRecipeJson = `{
	"filesToCreate": [
		"scratch.sh",
		"notes.md"
	],
	"filesToCopy": []
}`
