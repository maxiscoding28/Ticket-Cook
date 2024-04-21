package cmd

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func (hi *HomeInfoStruct) determineHomeDirectory(envVar envVarStruct) {
	if envVar.exists {
		hi.HomePath = envVar.value
	} else {

		hi.HomePath = DefaultHomeDirectory
		log("No `TCK_HOME_DIR` environment variable was set", "warn")
		log("The global default will be used: `$HOME/tck`", "warn")
	}
	log(fmt.Sprintf("Home directory: %s", hi.HomePath), "info")

}

func (hi HomeInfoStruct) getTicketsPath() string {
	return hi.HomePath + "/tickets"
}

func (hi HomeInfoStruct) getRecipesPath() string {
	return hi.HomePath + "/recipes"
}

func (hi HomeInfoStruct) getClosedPath() string {
	return hi.HomePath + "/.closed"
}

func getHomeDirectory(envVar envVarStruct) (*HomeInfoStruct, error) {
	var homeInfo HomeInfoStruct
	homeInfo.determineHomeDirectory(envVar)
	if err := fileOrDirectoryExists(homeInfo.HomePath); err != nil {
		if isFileNotFoundError(err) {
			message := fmt.Sprintf("%s Run `tck bootstrap` to set up a new tck home directory", emoji.HikingBoot)
			log(message, "info")
		} else {
			return nil, err
		}
	} else {
		logSnippet := []string{"tickets/", "recipes/", ".closed/"}
		directories := []string{homeInfo.getTicketsPath(), homeInfo.getRecipesPath(), homeInfo.getClosedPath()}
		for i := 0; i <= 2; i++ {
			if err := fileOrDirectoryExists(directories[i]); err != nil {
				if isFileNotFoundError(err) {
					message := fmt.Sprintf("%s No %s directory was found so a new one was created: %s", emoji.Hospital, logSnippet[i], directories[i])
					log(message, "success")
					createDirectory(directories[i])
				} else {
					return nil, err
				}
			}
		}
	}
	return &homeInfo, nil
}
