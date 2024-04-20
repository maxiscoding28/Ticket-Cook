package cmd

type envVarStruct struct {
	envVar string
	value  string
	exists bool
}
type HomeInfoStruct struct {
	HomePath string
}

type TicketStruct struct {
	TicketId      string
	DirectoryPath string
	MetaDataPath  string
	FilesToCreate []string
}

type RecipeMapStruct struct {
	FilesToCreate []string `json:"filesToCreate"`
	FilesToCopy   []string `json:"filesToCopy"`
}
