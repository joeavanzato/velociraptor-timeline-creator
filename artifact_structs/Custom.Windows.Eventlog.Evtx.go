package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Custom_Windows_Eventlog_Evtx struct {
	PayloadData1    string    `json:"PayloadData1"`
	PayloadData2    string    `json:"PayloadData2"`
	PayloadData3    string    `json:"PayloadData3"`
	PayloadData4    string    `json:"PayloadData4"`
	PayloadData5    string    `json:"PayloadData5"`
	PayloadData6    string    `json:"PayloadData6"`
	UserName        string    `json:"UserName"`
	RemoteHost      string    `json:"RemoteHost"`
	ExecutableInfo  string    `json:"ExecutableInfo"`
	MapDescription  string    `json:"MapDescription"`
	ChunkNumber     int       `json:"ChunkNumber"`
	Computer        string    `json:"Computer"`
	Payload         string    `json:"Payload"`
	UserID          string    `json:"UserId"`
	Channel         string    `json:"Channel"`
	Provider        string    `json:"Provider"`
	EventID         int       `json:"EventId"`
	EventRecordID   string    `json:"EventRecordId"`
	ProcessID       int       `json:"ProcessId"`
	ThreadID        int       `json:"ThreadId"`
	Level           string    `json:"Level"`
	Keywords        string    `json:"Keywords"`
	SourceFile      string    `json:"SourceFile"`
	ExtraDataOffset int       `json:"ExtraDataOffset"`
	HiddenRecord    bool      `json:"HiddenRecord"`
	TimeCreated     time.Time `json:"TimeCreated"`
	RecordNumber    int       `json:"RecordNumber"`
}

func Process_Custom_Windows_Eventlog_Evtx(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Custom_Windows_Eventlog_Evtx{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimeCreated,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: tmp.MapDescription,
			SourceUser:       tmp.UserName,
			SourceHost:       tmp.RemoteHost,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Channel,
			MetaData:         fmt.Sprintf("EventID: %v, ExecutableInfo: %v, Data1: %v, Data2: %v, Data3: %v, Data4: %v, Data5: %v, Data6: %v", tmp.EventID, tmp.ExecutableInfo, tmp.PayloadData1, tmp.PayloadData2, tmp.PayloadData3, tmp.PayloadData4, tmp.PayloadData5, tmp.PayloadData6),
		}
		outputChannel <- tmp2.StringArray()
	}
}
