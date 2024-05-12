package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Exchange_Windows_Detection_PipeHunter struct {
	Stdout     string `json:"Stdout"`
	Stderr     string `json:"Stderr"`
	ReturnCode int    `json:"ReturnCode"`
	Complete   bool   `json:"Complete"`
}

func (s Exchange_Windows_Detection_PipeHunter) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Detection_PipeHunter) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Detection_PipeHunter(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Detection_PipeHunter{}
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
