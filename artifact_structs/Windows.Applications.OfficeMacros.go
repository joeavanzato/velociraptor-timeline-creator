package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Applications_OfficeMacros struct {
	Code       string `json:"Code"`
	ModuleName string `json:"ModuleName"`
	StreamName string `json:"StreamName"`
	Type       string `json:"Type"`
	Filename   string `json:"filename"`
}

func (s Windows_Applications_OfficeMacros) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Applications_OfficeMacros) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Applications_OfficeMacros(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_OfficeMacros{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
