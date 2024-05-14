package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Exchange_HashRunKeys struct {
	FullPath string `json:"FullPath"`
	Name     string `json:"Name"`
	Value    string `json:"Value"`
	RealPath string `json:"RealPath"`
	Hash     string `json:"Hash"`
}

func (s Exchange_HashRunKeys) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_HashRunKeys) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_HashRunKeys(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_HashRunKeys{}
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
