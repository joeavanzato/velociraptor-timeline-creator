package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"math/big"
	"time"
)

type Exchange_Windows_EventLogs_Bitsadmin struct {
	EventTime                         time.Time   `json:"EventTime"`
	Computer                          string      `json:"Computer"`
	EventId                           int         `json:"EventId"`
	UserId                            string      `json:"UserId"`
	TransferId                        string      `json:"TransferId"`
	Name                              string      `json:"Name"`
	Id                                interface{} `json:"Id"`
	Url                               string      `json:"Url"`
	TLD                               string      `json:"TLD"`
	Peer                              string      `json:"Peer"`
	FileTime                          time.Time   `json:"FileTime"`
	FileLength                        big.Int     `json:"fileLength"`
	BytesTotal                        big.Int     `json:"bytesTotal"`
	BytesTransferred                  int         `json:"bytesTransferred"`
	EventDataBytesTransferredFromPeer int         `json:"EventData.bytesTransferredFromPeer"`
}

func Process_Exchange_Windows_EventLogs_Bitsadmin(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_Bitsadmin{}
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
			SourceUser:       tmp.UserId,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Name: %v, URL: %v, Bytes Transferred: %v", tmp.Name, tmp.Url, tmp.BytesTransferred),
		}
		outputChannel <- tmp2.StringArray()
	}
}
