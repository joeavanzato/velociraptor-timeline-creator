package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Registry_NTUser struct {
	OSPath      string    `json:"OSPath"`
	Data        any       `json:"Data"`
	Mtime       time.Time `json:"Mtime"`
	Username    string    `json:"Username"`
	Description string    `json:"Description"`
	UID         string    `json:"Uid"`
	Gid         string    `json:"Gid"`
	UUID        string    `json:"UUID"`
	Directory   string    `json:"Directory"`
}

func (s Windows_Registry_NTUser) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Registry_NTUser) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}
func Process_Windows_Registry_NTUser(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_NTUser{}
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
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Directory: %v", tmp.Directory),
		}
		outputChannel <- tmp2.StringArray()
	}
}
