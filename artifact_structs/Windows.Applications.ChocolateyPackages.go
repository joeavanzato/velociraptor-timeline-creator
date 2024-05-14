package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Applications_ChocolateyPackages struct {
	OSPath  string `json:"OSPath"`
	Name    string `json:"Name"`
	Version string `json:"Version"`
	Summary string `json:"Summary"`
	Authors string `json:"Authors"`
	License string `json:"License"`
}

func (s Windows_Applications_ChocolateyPackages) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Applications_ChocolateyPackages) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Applications_ChocolateyPackages(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_ChocolateyPackages{}
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
