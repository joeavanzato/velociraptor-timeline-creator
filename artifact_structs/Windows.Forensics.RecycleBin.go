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

type Windows_Forensics_RecycleBin struct {
	DeletedTimestamp time.Time `json:"DeletedTimestamp"`
	Name             string    `json:"Name"`
	OriginalFilePath string    `json:"OriginalFilePath"`
	FileSize         int       `json:"FileSize"`
	OSPath           string    `json:"OSPath"`
	RecyclePath      string    `json:"RecyclePath"`
	Upload           any       `json:"Upload"`
}

func (s Windows_Forensics_RecycleBin) StringArray() []string {
	return []string{s.DeletedTimestamp.String(), s.Name, s.OriginalFilePath, strconv.Itoa(s.FileSize), s.OSPath, s.RecyclePath, fmt.Sprint(s.Upload)}
}

func (s Windows_Forensics_RecycleBin) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Forensics_RecycleBin(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_RecycleBin{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.DeletedTimestamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.DeletedTimestamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OriginalFilePath,
			MetaData:         fmt.Sprintf("FileSize: %v, RecyclePath: %v", tmp.FileSize, tmp.RecyclePath),
		}
		outputChannel <- tmp2.StringArray()
	}
}
