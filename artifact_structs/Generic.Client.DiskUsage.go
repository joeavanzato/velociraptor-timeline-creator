package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Generic_Client_DiskUsage struct {
	DirPath        string `json:"DirPath"`
	TotalSize      int64  `json:"TotalSize"`
	TotalSizeHuman string `json:"TotalSizeHuman"`
}

func (s Generic_Client_DiskUsage) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Generic_Client_DiskUsage) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Client_DiskUsage(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Client_DiskUsage{}
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
