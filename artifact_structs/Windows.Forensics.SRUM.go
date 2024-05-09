package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Forensics_SRUM_ApplicationResourceUsage struct {
	SRUMID                       int       `json:"SRUMId"`
	TimeStamp                    time.Time `json:"TimeStamp"`
	App                          string    `json:"App"`
	UserSid                      string    `json:"UserSid"`
	User                         string    `json:"User"`
	ForegroundCycleTime          int64     `json:"ForegroundCycleTime"`
	BackgroundCycleTime          int       `json:"BackgroundCycleTime"`
	FaceTime                     int64     `json:"FaceTime"`
	ForegroundContextSwitches    int       `json:"ForegroundContextSwitches"`
	BackgroundContextSwitches    int       `json:"BackgroundContextSwitches"`
	ForegroundBytesRead          int       `json:"ForegroundBytesRead"`
	ForegroundBytesWritten       int       `json:"ForegroundBytesWritten"`
	ForegroundNumReadOperations  int       `json:"ForegroundNumReadOperations"`
	ForegroundNumWriteOperations int       `json:"ForegroundNumWriteOperations"`
	ForegroundNumberOfFlushes    int       `json:"ForegroundNumberOfFlushes"`
	BackgroundBytesRead          int       `json:"BackgroundBytesRead"`
	BackgroundBytesWritten       int       `json:"BackgroundBytesWritten"`
	BackgroundNumReadOperations  int       `json:"BackgroundNumReadOperations"`
	BackgroundNumWriteOperations int       `json:"BackgroundNumWriteOperations"`
	BackgroundNumberOfFlushes    int       `json:"BackgroundNumberOfFlushes"`
}

func Process_Windows_Forensics_SRUM_ApplicationResourceUsage(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_SRUM_ApplicationResourceUsage{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimeStamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "SRUM Application Resource Usage Entry",
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.App,
			MetaData:         fmt.Sprintf("BytesRead: %v, BytesWritten: %v", tmp.ForegroundBytesRead+tmp.BackgroundBytesRead, tmp.ForegroundBytesWritten+tmp.BackgroundBytesWritten),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Windows_Forensics_SRUM_ExecutionStats struct {
	ID              int       `json:"ID"`
	TimeStamp       time.Time `json:"TimeStamp"`
	App             string    `json:"App"`
	UserSid         string    `json:"UserSid"`
	User            string    `json:"User"`
	EndTime         time.Time `json:"EndTime"`
	DurationMS      int       `json:"DurationMS"`
	NetworkBytesRaw int64     `json:"NetworkBytesRaw"`
}

func Process_Windows_Forensics_SRUM_ExecutionStats(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_SRUM_ExecutionStats{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimeStamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "SRUM Execution Stats Entry",
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.App,
			MetaData:         fmt.Sprintf("Duration(MS): %v, EndTime: %v, Network Bytes: %v", tmp.DurationMS, tmp.EndTime, tmp.NetworkBytesRaw),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Windows_Forensics_SRUM_NetworkUsage struct {
	SRUMID         int       `json:"SRUMId"`
	TimeStamp      time.Time `json:"TimeStamp"`
	App            string    `json:"App"`
	UserSid        string    `json:"UserSid"`
	User           bool      `json:"User"`
	UserID         int       `json:"UserId"`
	BytesSent      int       `json:"BytesSent"`
	BytesRecvd     int       `json:"BytesRecvd"`
	InterfaceLuid  int64     `json:"InterfaceLuid"`
	L2ProfileID    int       `json:"L2ProfileId"`
	L2ProfileFlags int       `json:"L2ProfileFlags"`
}

func Process_Windows_Forensics_SRUM_NetworkUsage(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_SRUM_NetworkUsage{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimeStamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "SRUM Network Usage Entry",
			EventDescription: "",
			SourceUser:       tmp.UserSid,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.App,
			MetaData:         fmt.Sprintf("Bytes Sent: %v, Bytes Received: %v", tmp.BytesSent, tmp.BytesRecvd),
		}
		outputChannel <- tmp2.StringArray()
	}
}
