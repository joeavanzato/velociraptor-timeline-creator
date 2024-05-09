package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Registry_UserAssist struct {
	KeyPath            string    `json:"_KeyPath"`
	Name               string    `json:"Name"`
	User               string    `json:"User"`
	LastExecution      time.Time `json:"LastExecution"`
	LastExecutionTS    int       `json:"LastExecutionTS"`
	NumberOfExecutions int       `json:"NumberOfExecutions"`
}

func Process_Windows_Registry_UserAssist(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_UserAssist{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastExecution,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Name,
			MetaData:         fmt.Sprintf("Number of Executions: %v", tmp.NumberOfExecutions),
		}
		outputChannel <- tmp2.StringArray()
	}
}
