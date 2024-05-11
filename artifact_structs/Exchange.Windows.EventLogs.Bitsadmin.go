package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"math/big"
	"strconv"
	"time"
)

type Exchange_Windows_EventLogs_Bitsadmin struct {
	EventTime                         time.Time `json:"EventTime"`
	Computer                          string    `json:"Computer"`
	EventId                           int       `json:"EventId"`
	UserId                            string    `json:"UserId"`
	TransferId                        string    `json:"TransferId"`
	Name                              string    `json:"Name"`
	Id                                any       `json:"Id"`
	Url                               string    `json:"Url"`
	TLD                               string    `json:"TLD"`
	Peer                              string    `json:"Peer"`
	FileTime                          time.Time `json:"FileTime"`
	FileLength                        big.Int   `json:"fileLength"`
	BytesTotal                        big.Int   `json:"bytesTotal"`
	BytesTransferred                  int       `json:"bytesTransferred"`
	EventDataBytesTransferredFromPeer int       `json:"EventData.bytesTransferredFromPeer"`
}

func (s Exchange_Windows_EventLogs_Bitsadmin) StringArray() []string {
	return []string{s.EventTime.String(), s.Computer, strconv.Itoa(s.EventId), s.UserId, s.TransferId, s.Name,
		fmt.Sprint(s.Id), s.Url, s.TLD, s.Peer, s.FileTime.String(), fmt.Sprint(s.FileLength),
		fmt.Sprint(s.BytesTotal), strconv.Itoa(s.BytesTransferred), strconv.Itoa(s.EventDataBytesTransferredFromPeer)}
}

func (s Exchange_Windows_EventLogs_Bitsadmin) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
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
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
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
