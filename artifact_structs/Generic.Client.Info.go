package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Generic_Client_Info_Users struct {
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	LastLogin   time.Time `json:"LastLogin"`
}

func Process_Generic_Client_Info_Users(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Client_Info_Users{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastLogin,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "User Last Login",
			EventDescription: "",
			SourceUser:       tmp.Name,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Description: %v", tmp.Description),
		}
		outputChannel <- tmp2.StringArray()
	}
}
