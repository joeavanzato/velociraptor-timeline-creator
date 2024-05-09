package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Registry_RDP_Servers struct {
	Username      string    `json:"Username"`
	SID           string    `json:"SID"`
	HiveName      string    `json:"HiveName"`
	Key           string    `json:"Key"`
	LastWriteTime time.Time `json:"LastWriteTime"`
	Server        string    `json:"Server"`
	UsernameHint  string    `json:"UsernameHint"`
	CertHash      string    `json:"CertHash"`
}

func Process_Windows_Registry_RDP_Servers(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_RDP_Servers{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastWriteTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  tmp.UsernameHint,
			DestinationHost:  tmp.Server,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("CertHash: %v", tmp.CertHash),
		}
		outputChannel <- tmp2.StringArray()
	}
}
