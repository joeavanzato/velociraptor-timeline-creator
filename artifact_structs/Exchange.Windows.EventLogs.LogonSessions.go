package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Exchange_Windows_EventLogs_LogonSessions struct {
	Start             time.Time `json:"Start"`
	End               time.Time `json:"End"`
	Duration          float64   `json:"Duration"`
	SourceHost        string    `json:"SourceHost"`
	SubjectUserSid    []string  `json:"SubjectUserSid"`
	SubjectUserName   []string  `json:"SubjectUserName"`
	SubjectDomainName []string  `json:"SubjectDomainName"`
	TargetUserName    []string  `json:"TargetUserName"`
	TargetDomainName  []string  `json:"TargetDomainName"`
	TargetLogonID     []int64   `json:"TargetLogonId"`
	LogonType         []int     `json:"LogonType"`
	LogonProcessName  []string  `json:"LogonProcessName"`
	ProcessName       []string  `json:"ProcessName"`
	IPAddress         []string  `json:"IpAddress"`
}

func Process_Exchange_Windows_EventLogs_LogonSessions(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_LogonSessions{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Start,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       fmt.Sprint(tmp.SubjectUserName),
			SourceHost:       fmt.Sprintf("%v, %v", tmp.SourceHost, tmp.IPAddress),
			DestinationUser:  fmt.Sprint(tmp.TargetUserName),
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("LogonProcess: %v, LogonType: %v, TargetDomain: %v", tmp.LogonProcessName, tmp.LogonType, tmp.TargetDomainName),
		}
		outputChannel <- tmp2.StringArray()
	}
}
