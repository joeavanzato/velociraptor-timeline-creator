package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Forensics_PartitionTable struct {
	StartOffset       any      `json:"StartOffset"`
	EndOffset         any      `json:"EndOffset"`
	Size              string   `json:"Size"`
	Name              string   `json:"name"`
	TopLevelDirectory []string `json:"TopLevelDirectory"`
	Magic             string   `json:"Magic"`
	PartitionPath     string   `json:"_PartitionPath"`
}

func (s Windows_Forensics_PartitionTable) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Forensics_PartitionTable) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Forensics_PartitionTable(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_PartitionTable{}
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
