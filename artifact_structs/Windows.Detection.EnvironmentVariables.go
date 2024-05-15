package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Detection_EnvironmentVariables struct {
	Pid           int    `json:"Pid"`
	Name          string `json:"Name"`
	ImagePathName string `json:"ImagePathName"`
	CommandLine   string `json:"CommandLine"`
	Var           string `json:"Var"`
	Value         string `json:"Value"`
}

func (s Windows_Detection_EnvironmentVariables) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Detection_EnvironmentVariables) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Detection_EnvironmentVariables(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Detection_EnvironmentVariables{}
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
