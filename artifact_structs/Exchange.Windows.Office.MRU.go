package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Exchange_Windows_Office_MRU struct {
	Timestamp      time.Time `json:"Timestamp"`
	SAMaccountname string    `json:"SAMaccountname"`
	FileType       string    `json:"FileType"`
	Path           string    `json:"Path"`
}

func Process_Exchange_Windows_Office_MRU(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Office_MRU{}
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
			EventDescription: "",
			SourceUser:       tmp.SAMaccountname,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Path,
			MetaData:         fmt.Sprintf(""),
		}
		outputChannel <- tmp2.StringArray()
	}
}