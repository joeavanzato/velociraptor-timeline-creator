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

func (s Windows_EventLogs_Evtx) StringArray() []string {
	return []string{s.System.Provider.Name, s.System.Provider.GUID, strconv.Itoa(s.System.EventID.Value), strconv.Itoa(s.System.Version), strconv.Itoa(s.System.Level),
		strconv.Itoa(s.System.Task), strconv.Itoa(s.System.Opcode), fmt.Sprint(s.System.Keywords), fmt.Sprint(s.System.TimeCreated.SystemTime),
		strconv.Itoa(s.System.EventRecordID), s.System.Correlation.ActivityID, strconv.Itoa(s.System.Execution.ProcessID),
		strconv.Itoa(s.System.Execution.ThreadID), s.System.Channel, s.System.Computer, s.System.Security.UserID,
		s.EventData.Message1, s.EventData.Message2, s.EventData.Message3, s.EventData.Message4, strconv.Itoa(s.EventData.HexInt1),
		strconv.Itoa(s.EventData.HexInt2), strconv.Itoa(s.EventData.HexInt3), s.Message, s.TimeCreated.String(), s.Channel,
		strconv.Itoa(s.EventRecordID), strconv.Itoa(s.EventID), s.OSPath}
}

func (s Windows_EventLogs_Evtx) GetHeaders() []string {
	return []string{"System_Provider_Name", "System_Provider_GUID", "System_EventID_Value", "System_Version", "System_Level", "System_Task",
		"System_Opcode", "System_Keywords", "System_TimeCreated_SystemTime", "System_EventRecordID", "System_Correlation_ActivityID",
		"System_Execution_ProcessID", "System_Execution_ThreadID", "System_Channel", "System_Computer", "System_Security_UserID", "System_EventData_Message1",
		"System_EventData_Message2", "System_EventData_Message3", "System_EventData_Message4", "System_EventData_HexInt1", "System_EventData_HexInt2", "System_EventData_HexInt3",
		"Message", "TimeCreated", "Channel", "EventRecordID", "EventID", "OSPath"}
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
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.TimeCreated.String(), clientIdentifier, tmp.System.Computer, tmp.StringArray(), outputChannel)
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
			SourceHost:       tmp.System.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Channel,
			MetaData:         fmt.Sprintf("Message: %v", tmp.Message),
		}
		outputChannel <- tmp2.StringArray()
	}
}
