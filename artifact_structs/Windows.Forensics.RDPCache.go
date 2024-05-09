package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Forensics_RDPCache struct {
	OSPath string    `json:"OSPath"`
	Size   int       `json:"Size"`
	Mtime  time.Time `json:"Mtime"`
	Atime  time.Time `json:"Atime"`
	Ctime  time.Time `json:"Ctime"`
	Btime  time.Time `json:"Btime"`
}

func Process_Windows_Forensics_RDPCache(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_RDPCache{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Atime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Modified: %v, Changed: %v, Birthed: %v", tmp.Mtime, tmp.Ctime, tmp.Btime),
		}
		outputChannel <- tmp2.StringArray()
	}
}
