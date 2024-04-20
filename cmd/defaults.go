package cmd

import (
	"os"
	"path/filepath"
)

const DefaultUrlFormat string = "http:example.com/@"

// TODO: Change prior to releasing
var DefaultHomeDirectory string = filepath.Join(os.Getenv("HOME"), "dev/sandbox-go/tck/test")

const DefaultRecipeJson = `{
	"filesToCreate": [
		"scratch.sh",
		"notes.md"
	],
	"filesToCopy": ["*"]
}`

const StarMdFileContent = `
# Recipe Structure
- recipe.json file
- file for copying

# Recipe Fields
- files to create
- file to copy

# Recipe Management
git and github
cd recipes && git init
git remote add origin $REMOTE-URL

> ! Consider running a secret scanning tool on the repository to prevent exposing un wanted secrets
`
