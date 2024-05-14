package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Applications_Chrome_Extensions struct {
	UID         string `json:"Uid"`
	User        string `json:"User"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Identifier  string `json:"Identifier"`
	Version     string `json:"Version"`
	Author      any    `json:"Author"`
	Persistent  any    `json:"Persistent"`
	Path        string `json:"Path"`
	Scopes      any    `json:"Scopes"`
	Permissions any    `json:"Permissions"`
	Key         string `json:"Key"`
}

func (s Windows_Applications_Chrome_Extensions) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Applications_Chrome_Extensions) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Applications_Chrome_Extensions(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_Chrome_Extensions{}
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
