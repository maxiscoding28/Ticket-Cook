package cmd

import "fmt"

const systemDefaultHomePath = "/Users/maxwinslow/dev/sandbox-go/tcmd/test"

type HomeInfoStruct struct {
	HomePath string
}

func (hi *HomeInfoStruct) setHomeDirectory(envVar envVarStruct) {
	if envVar.exists {
		hi.HomePath = envVar.value
	} else {
		hi.HomePath = systemDefaultHomePath
	}
	log(fmt.Sprintf("Home directory: %s", hi.HomePath), nil)

}

func (hi *HomeInfoStruct) getTicketsPath() string {
	return hi.HomePath + "/tickets"
}

func (hi *HomeInfoStruct) getTemplatesPath() string {
	return hi.HomePath + "/templates"
}
