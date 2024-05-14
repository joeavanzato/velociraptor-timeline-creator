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

type Windows_EventLogs_Modifications_Channels struct {
	Mtime           time.Time `json:"Mtime"`
	ChannelName     string    `json:"ChannelName"`
	Key             string    `json:"_Key"`
	OwningPublisher string    `json:"OwningPublisher"`
	Enabled         int       `json:"Enabled"`
}

func (s Windows_EventLogs_Modifications_Channels) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_EventLogs_Modifications_Channels) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_EventLogs_Modifications_Channels(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_Modifications_Channels{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Event Log Channel Modified",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Key,
			MetaData:         fmt.Sprintf("Enabled: %v, Channel: %v", tmp.Enabled, tmp.ChannelName),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Windows_EventLogs_Modifications_Providers struct {
	Mtime        time.Time `json:"Mtime"`
	GUID         string    `json:"GUID"`
	RegKey       string    `json:"_RegKey"`
	ProviderName string    `json:"ProviderName"`
	Enabled      int       `json:"Enabled"`
	Content      struct {
		EnableLevel     int    `json:"EnableLevel"`
		EnableProperty  int    `json:"EnableProperty"`
		Enabled         int    `json:"Enabled"`
		LoggerName      string `json:"LoggerName"`
		MatchAllKeyword int    `json:"MatchAllKeyword"`
		MatchAnyKeyword any    `json:"MatchAnyKeyword"`
		Status          int    `json:"Status"`
	} `json:"Content"`
}

func (s Windows_EventLogs_Modifications_Providers) StringArray() []string {
	return []string{s.Mtime.String(), s.GUID, s.RegKey, s.ProviderName, strconv.Itoa(s.Enabled), strconv.Itoa(s.Content.EnableLevel),
		strconv.Itoa(s.Content.EnableProperty), strconv.Itoa(s.Content.Enabled), s.Content.LoggerName, strconv.Itoa(s.Content.MatchAllKeyword), fmt.Sprint(s.Content.MatchAnyKeyword, 10), strconv.Itoa(s.Content.Status)}
}

func (s Windows_EventLogs_Modifications_Providers) GetHeaders() []string {
	return []string{"Mtime", "GUID", "RegKey", "ProviderName", "Enabled", "Content_EnableLevel", "Content_EnableProperty", "Content_Enabled", "Content_LoggerName", "Content_MatchAllKeyword", "Content_MatchAnyKeyword", "Content_Status"}
}

func Process_Windows_EventLogs_Modifications_Providers(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_EventLogs_Modifications_Providers{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Event Log Provider Modified",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.ProviderName,
			MetaData:         fmt.Sprintf("Enabled: %v", tmp.Enabled),
		}
		outputChannel <- tmp2.StringArray()
	}
}
