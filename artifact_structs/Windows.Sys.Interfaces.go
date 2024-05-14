package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Sys_Interfaces struct {
	Name    string `json:"Name"`
	Details struct {
		Description string `json:"Description"`
		MAC         string `json:"MAC"`
	} `json:"Details"`
}

func (s Windows_Sys_Interfaces) StringArray() []string {
	return []string{s.Name, s.Details.Description, s.Details.MAC}
}

func (s Windows_Sys_Interfaces) GetHeaders() []string {
	return []string{"Name", "Description", "MAC"}
}

func Process_Windows_Sys_Interfaces(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_Interfaces{}
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
