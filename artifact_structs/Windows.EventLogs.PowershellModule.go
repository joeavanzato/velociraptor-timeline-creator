package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_EventLogs_PowershellModule struct {
	EventTime     time.Time `json:"EventTime"`
	EventID       int       `json:"EventID"`
	Computer      string    `json:"Computer"`
	SecurityID    string    `json:"SecurityID"`
	ContextInfo   string    `json:"ContextInfo"`
	Payload       string    `json:"Payload"`
	Message       string    `json:"Message"`
	EventRecordID int       `json:"EventRecordID"`
	Level         int       `json:"Level"`
	Opcode        int       `json:"Opcode"`
	Task          int       `json:"Task"`
	Source        string    `json:"Source"`
}

func (s Windows_EventLogs_PowershellModule) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_EventLogs_PowershellModule) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_EventLogs_PowershellModule(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_PowershellModule{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Source,
			MetaData:         fmt.Sprintf("Payload: %v", tmp.Payload),
		}
		outputChannel <- tmp2.StringArray()
	}
}
