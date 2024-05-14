package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Sys_FirewallRules struct {
	Value       string `json:"Value"`
	Description string `json:"Description"`
	App         string `json:"App"`
	Active      string `json:"Active"`
	Action      string `json:"Action"`
	Dir         string `json:"Dir"`
	Protocol    string `json:"Protocol"`
	LPort       string `json:"LPort"`
	Name        string `json:"Name"`
}

func (s Windows_Sys_FirewallRules) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Sys_FirewallRules) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Sys_FirewallRules(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_FirewallRules{}
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
