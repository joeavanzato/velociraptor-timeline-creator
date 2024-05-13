package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Timeline_Registry_RunMRU struct {
	EventTime time.Time `json:"event_time"`
	Hostname  string    `json:"hostname"`
	Parser    string    `json:"parser"`
	Message   string    `json:"message"`
	Source    string    `json:"source"`
	User      string    `json:"user"`
	RegKey    string    `json:"reg_key"`
	RegMtime  time.Time `json:"reg_mtime"`
	RegName   string    `json:"reg_name"`
	RegValue  string    `json:"reg_value"`
	RegType   string    `json:"reg_type"`
}

func (s Windows_Timeline_Registry_RunMRU) StringArray() []string {
	return []string{s.EventTime.String(), s.Hostname, s.Parser, s.Message, s.Source,
		s.User, s.RegKey, s.RegMtime.String(), s.RegName, s.RegValue, s.RegType}
}

// Headers should match the string array above
func (s Windows_Timeline_Registry_RunMRU) GetHeaders() []string {
	return []string{"EventTime", "Hostname", "Parser", "Message", "Source", "User", "RegKey", "RegMtime", "RegName", "RegValue", "RegType"}
}

func Process_Windows_Timeline_Registry_RunMRU(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Timeline_Registry_RunMRU{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.RegMtime.String(), clientIdentifier, tmp.Hostname, tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.RegMtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       tmp.Hostname,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.RegValue,
			MetaData:         "",
		}
		outputChannel <- tmp2.StringArray()
	}
}
