package cmd

import (
	"os"
	"path/filepath"
)

const DefaultUrlFormat string = "http:example.com/@"

// TODO: Change prior to releasing
var DefaultHomeDirectory string = filepath.Join(os.Getenv("HOME"), "tck")

const StarMdFileContent = `
# Recipe Structure
Recipe Name/
	Files.txt
	To.md
	Copy.sh

# Recipe Management
git and github
cd recipes && git init
git remote add origin $REMOTE-URL

> ! Consider running a secret scanning tool on the repository to prevent exposing un wanted secrets
`
