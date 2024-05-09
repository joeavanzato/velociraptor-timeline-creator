package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
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

func Process_Windows_EventLogs_PowerShellScriptblock(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_PowerShellScriptblock{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Path,
			MetaData:         fmt.Sprintf("ScriptBlock: %v", tmp.ScriptBlockText),
		}
		outputChannel <- tmp2.StringArray()
	}
}
