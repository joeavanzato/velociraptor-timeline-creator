package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_System_Shares struct {
	Name           string `json:"Name"`
	Path           string `json:"Path"`
	Caption        string `json:"Caption"`
	Status         string `json:"Status"`
	MaximumAllowed any    `json:"MaximumAllowed"`
	AllowMaximum   bool   `json:"AllowMaximum"`
	InstallDate    any    `json:"InstallDate"`
}

func (s Windows_System_Shares) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_System_Shares) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_System_Shares(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Shares{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		// TODO - Investigate if a newly deployed share will have an actual timestamp for this artifact

	}
}
