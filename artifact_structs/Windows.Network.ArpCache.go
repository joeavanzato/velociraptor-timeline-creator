package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Network_ArpCache struct {
	AddressFamily    string `json:"AddressFamily"`
	Store            string `json:"Store"`
	State            string `json:"State"`
	InterfaceIndex   int    `json:"InterfaceIndex"`
	LocalAddress     string `json:"LocalAddress"`
	HardwareAddr     any    `json:"HardwareAddr"`
	RemoteAddress    string `json:"RemoteAddress"`
	InterfaceAlias   string `json:"InterfaceAlias"`
	RemoteMACAddress string `json:"RemoteMACAddress"`
}

func (s Windows_Network_ArpCache) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Network_ArpCache) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Network_ArpCache(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Network_ArpCache{}
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
