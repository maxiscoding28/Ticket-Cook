package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

type HomeInfoStruct struct {
	HomePath string
}

func (hi *HomeInfoStruct) setHomeDirectory(envVar envVarStruct) {
	if envVar.exists {
		hi.HomePath = envVar.value
	} else {
		hi.HomePath = filepath.Join(os.Getenv("HOME"), "tcf")
	}
	log(fmt.Sprintf("Home directory: %s", hi.HomePath), nil)

}

func (hi *HomeInfoStruct) getTicketsPath() string {
	return hi.HomePath + "/tickets"
}

func (hi *HomeInfoStruct) getTemplatesPath() string {
	return hi.HomePath + "/templates"
}
