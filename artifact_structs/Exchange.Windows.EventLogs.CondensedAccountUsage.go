package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type Exchange_Windows_EventLogs_CondensedAccountUsage struct {
	EventTime                 time.Time `json:"EventTime"`
	Computer                  string    `json:"Computer"`
	EventID                   int       `json:"EventID"`
	Description               string    `json:"Description"`
	DomainName                string    `json:"DomainName"`
	UserName                  string    `json:"UserName"`
	LogonID                   int       `json:"LogonId"`
	CredentialsUsedFor4648    string    `json:"CredentialsUsedFor4648"`
	LogonType                 any       `json:"LogonType"`
	LogonTypeDescription      string    `json:"LogonTypeDescription"`
	AuthenticationPackageName string    `json:"AuthenticationPackageName"`
	IPAddress                 string    `json:"IpAddress"`
	ClientName                string    `json:"ClientName"`
}

func (s Exchange_Windows_EventLogs_CondensedAccountUsage) StringArray() []string {
	return []string{s.EventTime.String(), s.Computer, strconv.Itoa(s.EventID), s.Description, s.DomainName, s.UserName, strconv.Itoa(s.LogonID), s.CredentialsUsedFor4648, fmt.Sprint(s.LogonType), s.LogonTypeDescription, s.AuthenticationPackageName, s.IPAddress, s.ClientName}
}

func (s Exchange_Windows_EventLogs_CondensedAccountUsage) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_CondensedAccountUsage(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_CondensedAccountUsage{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
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
			EventType:        fmt.Sprintf("Logon: %v", tmp.LogonTypeDescription),
			EventDescription: tmp.Description,
			SourceUser:       tmp.UserName,
			SourceHost:       fmt.Sprintf("Client: %v, IP: %v", tmp.ClientName, tmp.IPAddress),
			DestinationUser:  "",
			DestinationHost:  tmp.Computer,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("CredentialsUsedFor4648: %v, DomainName: %v", tmp.CredentialsUsedFor4648, tmp.DomainName),
		}
		outputChannel <- tmp2.StringArray()
	}
}
