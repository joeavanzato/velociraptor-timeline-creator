package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_EventLogs_AlternateLogon struct {
	EventTime        time.Time `json:"EventTime"`
	IPAddress        string    `json:"IpAddress"`
	Port             string    `json:"Port"`
	ProcessName      string    `json:"ProcessName"`
	SubjectUserSid   string    `json:"SubjectUserSid"`
	SubjectUserName  string    `json:"SubjectUserName"`
	TargetUserName   string    `json:"TargetUserName"`
	TargetServerName string    `json:"TargetServerName"`
	LogonTime        float64   `json:"LogonTime"`
}

func (s Windows_EventLogs_AlternateLogon) StringArray() []string {
	return []string{s.EventTime.String(), s.IPAddress, s.Port, s.ProcessName, s.SubjectUserSid, s.SubjectUserName,
		s.TargetUserName, s.TargetServerName, fmt.Sprint(s.LogonTime)}
}

func (s Windows_EventLogs_AlternateLogon) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_EventLogs_AlternateLogon(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_AlternateLogon{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.SubjectUserName,
			SourceHost:       fmt.Sprintf("%v:%v", tmp.IPAddress, tmp.Port),
			DestinationUser:  tmp.TargetUserName,
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("ProcessName: %v", tmp.ProcessName),
		}
		outputChannel <- tmp2.StringArray()
	}
}
