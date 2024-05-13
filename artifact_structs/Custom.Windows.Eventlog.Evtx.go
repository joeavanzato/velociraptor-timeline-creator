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

func (s Custom_Windows_Eventlog_Evtx) StringArray() []string {
	return []string{strconv.Itoa(s.RecordNumber), s.TimeCreated.String(), strconv.FormatBool(s.HiddenRecord), strconv.Itoa(s.ExtraDataOffset),
		s.SourceFile, s.Keywords, s.Level, strconv.Itoa(s.ThreadID), strconv.Itoa(s.ProcessID), s.EventRecordID,
		strconv.Itoa(s.EventID), s.Provider, s.Channel, s.UserID, s.Computer, s.MapDescription, s.UserName, s.RemoteHost, s.PayloadData1, s.PayloadData2, s.PayloadData3, s.PayloadData4, s.PayloadData5, s.PayloadData6,
		s.ExecutableInfo, strconv.Itoa(s.ChunkNumber), s.Payload}
}

func (s Custom_Windows_Eventlog_Evtx) GetHeaders() []string {
	return []string{"RecordNumber", "TimeCreated", "HiddenRecord", "ExtraDataOffset", "SourceFile", "Keywords", "Level", "ThreadID", "ProcessID", "EventRecordID", "EventID", "Provider", "Channel", "UserID", "Computer", "MapDescription", "UserName",
		"RemoteHost", "PayloadData1", "PayloadData2", "PayloadData3", "PayloadData4", "PayloadData5", "PayloadData6", "ExecutableInfo", "ChunkNumber", "Payload"}
}

func Process_Custom_Windows_Eventlog_Evtx(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Custom_Windows_Eventlog_Evtx{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.TimeCreated.String(), clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
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
