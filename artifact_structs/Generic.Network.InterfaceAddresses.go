package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Generic_Network_InterfaceAddresses struct {
	Index        int    `json:"Index"`
	MTU          int    `json:"MTU"`
	Name         string `json:"Name"`
	HardwareAddr any    `json:"HardwareAddr"`
	Flags        string `json:"Flags"`
	IP           string `json:"IP"`
	Mask         string `json:"Mask"`
}

func (s Generic_Network_InterfaceAddresses) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Generic_Network_InterfaceAddresses) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Network_InterfaceAddresses(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Network_InterfaceAddresses{}
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
