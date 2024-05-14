package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Sys_StartupItems struct {
	Name    string `json:"Name"`
	OSPath  string `json:"OSPath"`
	Details any    `json:"Details"`
	Enabled string `json:"Enabled"`
	Upload  string `json:"Upload"`
}

func (s Windows_Sys_StartupItems) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Sys_StartupItems) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Sys_StartupItems(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_StartupItems{}
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
