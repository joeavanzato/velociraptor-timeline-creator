package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"math/big"
	"time"
)

type Windows_EventLogs_Evtx struct {
	System struct {
		Provider struct {
			Name string `json:"Name"`
			GUID string `json:"Guid"`
		} `json:"Provider"`
		EventID struct {
			Value int `json:"Value"`
		} `json:"EventID"`
		Version     int     `json:"Version"`
		Level       int     `json:"Level"`
		Task        int     `json:"Task"`
		Opcode      int     `json:"Opcode"`
		Keywords    big.Int `json:"Keywords"`
		TimeCreated struct {
			SystemTime float64 `json:"SystemTime"`
		} `json:"TimeCreated"`
		EventRecordID int `json:"EventRecordID"`
		Correlation   struct {
			ActivityID string `json:"ActivityID"`
		} `json:"Correlation"`
		Execution struct {
			ProcessID int `json:"ProcessID"`
			ThreadID  int `json:"ThreadID"`
		} `json:"Execution"`
		Channel  string `json:"Channel"`
		Computer string `json:"Computer"`
		Security struct {
			UserID string `json:"UserID"`
		} `json:"Security"`
	} `json:"System"`
	EventData struct {
		Message1 string `json:"Message1"`
		Message2 string `json:"Message2"`
		Message3 string `json:"Message3"`
		Message4 string `json:"Message4"`
		HexInt1  int    `json:"HexInt1"`
		HexInt2  int    `json:"HexInt2"`
		HexInt3  int    `json:"HexInt3"`
	} `json:"EventData"`
	Message       string    `json:"Message"`
	TimeCreated   time.Time `json:"TimeCreated"`
	Channel       string    `json:"Channel"`
	EventRecordID int       `json:"EventRecordID"`
	EventID       int       `json:"EventID"`
	OSPath        string    `json:"OSPath"`
}

func Process_Windows_EventLogs_Evtx(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_Evtx{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// TODO - Maybe pull description for the most common ones or use zimmerman maps?
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimeCreated,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: fmt.Sprintf("Event ID: %v", tmp.EventID),
			SourceUser:       tmp.System.Security.UserID,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Channel,
			MetaData:         fmt.Sprintf("Message: %v", tmp.Message),
		}
		outputChannel <- tmp2.StringArray()
	}
}
