package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Registry_Sysinternals_Eulacheck struct {
	ProgramName  string    `json:"ProgramName"`
	Key          string    `json:"Key"`
	TimeAccepted time.Time `json:"TimeAccepted"`
	User         string    `json:"User"`
	EulaAccepted int       `json:"EulaAccepted"`
}

func Process_Windows_Registry_Sysinternals_Eulacheck(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_Sysinternals_Eulacheck{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimeAccepted,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.ProgramName,
			MetaData:         "",
		}
		outputChannel <- tmp2.StringArray()
	}
}
