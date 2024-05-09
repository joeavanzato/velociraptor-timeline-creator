package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Timeline_Registry_RunMRU struct {
	EventTime time.Time `json:"event_time"`
	Hostname  string    `json:"hostname"`
	Parser    string    `json:"parser"`
	Message   string    `json:"message"`
	Source    string    `json:"source"`
	User      string    `json:"user"`
	RegKey    string    `json:"reg_key"`
	RegMtime  time.Time `json:"reg_mtime"`
	RegName   string    `json:"reg_name"`
	RegValue  string    `json:"reg_value"`
	RegType   string    `json:"reg_type"`
}

func Process_Windows_Timeline_Registry_RunMRU(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Timeline_Registry_RunMRU{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.RegMtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       clientIdentifier,
			DestinationUser:  tmp.User,
			DestinationHost:  clientIdentifier,
			SourceFile:       tmp.RegValue,
			MetaData:         "",
		}
		outputChannel <- tmp2.StringArray()
	}
}
