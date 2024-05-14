package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Network_ExternalIpAddress struct {
	IP string `json:"IP"`
}

func (s Network_ExternalIpAddress) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Network_ExternalIpAddress) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Network_ExternalIpAddress(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Network_ExternalIpAddress{}
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
