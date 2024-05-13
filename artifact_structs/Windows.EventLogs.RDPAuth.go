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

type Windows_EventLogs_RDPAuth struct {
	EventTime     time.Time `json:"EventTime"`
	Computer      string    `json:"Computer"`
	Channel       string    `json:"Channel"`
	EventID       int       `json:"EventID"`
	DomainName    string    `json:"DomainName"`
	UserName      string    `json:"UserName"`
	LogonType     any       `json:"LogonType"`
	SourceIP      string    `json:"SourceIP"`
	Description   string    `json:"Description"`
	Message       string    `json:"Message"`
	EventRecordID int       `json:"EventRecordID"`
	OSPath        string    `json:"OSPath"`
}

func (s Windows_EventLogs_RDPAuth) StringArray() []string {
	return []string{s.EventTime.String(), s.Computer, s.Channel, strconv.Itoa(s.EventID), s.DomainName, s.UserName, fmt.Sprint(s.LogonType), s.SourceIP, s.Description, s.Message, strconv.Itoa(s.EventRecordID), s.OSPath}
}

func (s Windows_EventLogs_RDPAuth) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_EventLogs_RDPAuth(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_RDPAuth{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
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
			SourceUser:       tmp.UserName,
			SourceHost:       tmp.SourceIP,
			DestinationUser:  tmp.UserName,
			DestinationHost:  tmp.Computer,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Domain: %v, Description: %v, LogonType: %v", tmp.DomainName, tmp.Description, tmp.LogonType),
		}
		outputChannel <- tmp2.StringArray()
	}
}
