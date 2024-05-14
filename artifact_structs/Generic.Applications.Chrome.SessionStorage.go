package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Generic_Applications_Chrome_SessionStorage struct {
	OSPath  string `json:"OSPath"`
	GUID    string `json:"GUID"`
	URL     string `json:"URL"`
	Mapping any    `json:"Mapping"`
}

func (s Generic_Applications_Chrome_SessionStorage) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Generic_Applications_Chrome_SessionStorage) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Applications_Chrome_SessionStorage(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Applications_Chrome_SessionStorage{}
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
