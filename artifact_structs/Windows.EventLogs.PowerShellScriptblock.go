package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type Windows_EventLogs_PowerShellScriptblock struct {
	EventTime       time.Time `json:"EventTime"`
	Computer        string    `json:"Computer"`
	Channel         string    `json:"Channel"`
	EventID         int       `json:"EventID"`
	SecurityID      string    `json:"SecurityID"`
	Path            string    `json:"Path"`
	ScriptBlockID   string    `json:"ScriptBlockId"`
	ScriptBlockText string    `json:"ScriptBlockText"`
	Message         string    `json:"Message"`
	EventRecordID   int       `json:"EventRecordID"`
	Level           int       `json:"Level"`
	Opcode          int       `json:"Opcode"`
	Task            int       `json:"Task"`
	OSPath          string    `json:"OSPath"`
}

func (s Windows_EventLogs_PowerShellScriptblock) StringArray() []string {
	return []string{s.EventTime.String(), s.Computer, s.Channel, strconv.Itoa(s.EventID), s.SecurityID, s.Path,
		s.ScriptBlockID, s.ScriptBlockText, s.Message, strconv.Itoa(s.EventRecordID), strconv.Itoa(s.Level), strconv.Itoa(s.Opcode), strconv.Itoa(s.Task), s.OSPath}
}

func (s Windows_EventLogs_PowerShellScriptblock) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_EventLogs_PowerShellScriptblock(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_PowerShellScriptblock{}
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
			SourceFile:       tmp.Path,
			MetaData:         fmt.Sprintf("ScriptBlock: %v", tmp.ScriptBlockText),
		}
		outputChannel <- tmp2.StringArray()
	}
}
