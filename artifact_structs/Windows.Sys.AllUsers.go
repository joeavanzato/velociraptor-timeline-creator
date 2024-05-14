package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Sys_AllUsers struct {
	UID          int    `json:"Uid"`
	Gid          int    `json:"Gid"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Directory    any    `json:"Directory"`
	UUID         string `json:"UUID"`
	Mtime        any    `json:"Mtime"`
	HomedirMtime any    `json:"HomedirMtime"`
	Data         any    `json:"Data"`
}

func (s Windows_Sys_AllUsers) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Sys_AllUsers) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Sys_AllUsers(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_AllUsers{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
	}
}
