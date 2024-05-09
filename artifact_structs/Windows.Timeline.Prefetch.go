package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Timeline_Prefetch struct {
	EventTime       time.Time `json:"event_time"`
	Hostname        string    `json:"hostname"`
	Parser          string    `json:"parser"`
	Message         string    `json:"message"`
	Source          string    `json:"source"`
	FileName        string    `json:"file_name"`
	PrefetchCtime   time.Time `json:"prefetch_ctime"`
	PrefetchMtime   time.Time `json:"prefetch_mtime"`
	PrefetchSize    int       `json:"prefetch_size"`
	PrefetchHash    string    `json:"prefetch_hash"`
	PrefetchVersion string    `json:"prefetch_version"`
	PrefetchFile    string    `json:"prefetch_file"`
	PrefetchCount   int       `json:"prefetch_count"`
}

func Process_Windows_Timeline_Prefetch(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Timeline_Prefetch{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.PrefetchMtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       clientIdentifier,
			DestinationUser:  "",
			DestinationHost:  clientIdentifier,
			SourceFile:       tmp.FileName,
			MetaData:         fmt.Sprintf("Created Date: %v, Message: %v", tmp.PrefetchCtime, tmp.Message),
		}
		outputChannel <- tmp2.StringArray()
	}
}
