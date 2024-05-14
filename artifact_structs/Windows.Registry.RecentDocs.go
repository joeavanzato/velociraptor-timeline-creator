package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Registry_RecentDocs struct {
	LastWriteTime time.Time `json:"LastWriteTime"`
	Type          string    `json:"Type"`
	MruEntries    []string  `json:"MruEntries"`
	Key           any       `json:"Key"`
	HiveName      string    `json:"HiveName"`
	Username      string    `json:"Username"`
	UUID          string    `json:"UUID"`
}

func (s Windows_Registry_RecentDocs) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Registry_RecentDocs) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Registry_RecentDocs(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_RecentDocs{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastWriteTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastWriteTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Recent Docs Entry Written",
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       fmt.Sprint(tmp.MruEntries),
			MetaData:         fmt.Sprintf("Type: %v", tmp.Type),
		}
		outputChannel <- tmp2.StringArray()
	}
}
