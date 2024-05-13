package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
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

func (s Exchange_Windows_EventLogs_LogonSessions) StringArray() []string {
	return []string{s.Start.String(), s.End.String(), fmt.Sprint(s.Duration), s.SourceHost, fmt.Sprint(s.SubjectUserSid), fmt.Sprint(s.SubjectUserName),
		fmt.Sprint(s.SubjectDomainName), fmt.Sprint(s.TargetUserName), fmt.Sprint(s.TargetDomainName), fmt.Sprint(s.TargetLogonID),
		fmt.Sprint(s.LogonType), fmt.Sprint(s.LogonProcessName), fmt.Sprint(s.ProcessName), fmt.Sprint(s.IPAddress)}
}

func (s Exchange_Windows_EventLogs_LogonSessions) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_LogonSessions(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_LogonSessions{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Start.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
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
