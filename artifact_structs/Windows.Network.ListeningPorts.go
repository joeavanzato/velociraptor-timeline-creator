package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Network_ListeningPorts struct {
	Pid      int    `json:"Pid"`
	Name     string `json:"Name"`
	Port     int    `json:"Port"`
	Protocol string `json:"Protocol"`
	Family   string `json:"Family"`
	Address  string `json:"Address"`
}

func (s Windows_Network_ListeningPorts) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Network_ListeningPorts) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Network_ListeningPorts(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Network_ListeningPorts{}
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
