package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Exchange_Windows_EventLogs_Hayabusa struct {
	Timestamp      time.Time `json:"Timestamp"`
	RuleTitle      string    `json:"RuleTitle"`
	Level          string    `json:"Level"`
	Computer       string    `json:"Computer"`
	Channel        string    `json:"Channel"`
	EventID        int       `json:"EventID"`
	RecordID       int       `json:"RecordID"`
	Details        string    `json:"Details"`
	ExtraFieldInfo string    `json:"ExtraFieldInfo"`
	EventTime      time.Time `json:"EventTime"`
}

func Process_Exchange_Windows_EventLogs_Hayabusa(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_Hayabusa{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Timestamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: fmt.Sprintf("Rule: %v", tmp.RuleTitle),
			SourceUser:       "",
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Channel,
			MetaData:         fmt.Sprintf("Level: %v, Details: %v", tmp.Level, tmp.Details),
		}
		outputChannel <- tmp2.StringArray()
	}
}
