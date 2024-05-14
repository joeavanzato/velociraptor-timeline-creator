package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Registry_CapabilityAccessManager struct {
	SourceLocation    string    `json:"SourceLocation"`
	Accessed          string    `json:"Accessed"`
	Program           string    `json:"Program"`
	LastUsedTimeStart time.Time `json:"LastUsedTimeStart"`
	LastUsedTimeStop  time.Time `json:"LastUsedTimeStop"`
	KeyPath           string    `json:"KeyPath"`
}

func (s Exchange_Windows_Registry_CapabilityAccessManager) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Registry_CapabilityAccessManager) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Registry_CapabilityAccessManager(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Registry_CapabilityAccessManager{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastUsedTimeStart.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastUsedTimeStart,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Program,
			MetaData:         fmt.Sprintf("Accessed: %v", tmp.Accessed),
		}
		outputChannel <- tmp2.StringArray()
	}
}
