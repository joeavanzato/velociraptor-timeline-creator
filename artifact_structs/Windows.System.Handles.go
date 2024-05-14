package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_System_Handles struct {
	ProcPid  int    `json:"ProcPid"`
	ProcName string `json:"ProcName"`
	Exe      string `json:"Exe"`
	Type     string `json:"Type"`
	Name     string `json:"Name"`
	Handle   int    `json:"Handle"`
}

func (s Windows_System_Handles) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_System_Handles) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_System_Handles(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Handles{}
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
