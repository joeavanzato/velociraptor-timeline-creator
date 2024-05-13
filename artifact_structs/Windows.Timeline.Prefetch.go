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

type Windows_Timeline_Prefetch struct {
	EventTime       time.Time `json:"event_time"`
	Hostname        string    `json:"hostname"`
	Parser          string    `json:"parser"`
	Message         string    `json:"message"`
	Source          string    `json:"source"`
	FileName        string    `json:"file_name"`
	PrefetchCtime   time.Time `json:"prefetch_ctime"`
	PrefetchMtime   time.Time `json:"prefetch_mtime"`
	PrefetchSize    int       `json:"prefetch_size"`
	PrefetchHash    string    `json:"prefetch_hash"`
	PrefetchVersion string    `json:"prefetch_version"`
	PrefetchFile    string    `json:"prefetch_file"`
	PrefetchCount   int       `json:"prefetch_count"`
}

func (s Windows_Timeline_Prefetch) StringArray() []string {
	return []string{s.EventTime.String(), s.Hostname, s.Parser, s.Message, s.Source,
		s.FileName, s.PrefetchCtime.String(), s.PrefetchMtime.String(), strconv.Itoa(s.PrefetchSize),
		s.PrefetchHash, s.PrefetchVersion, s.PrefetchFile, strconv.Itoa(s.PrefetchCount)}
}

// Headers should match the string array above
func (s Windows_Timeline_Prefetch) GetHeaders() []string {
	return []string{"EventTime", "Hostname", "Parser", "Message", "Source", "FileName", "PrefetchCtime", "PrefetchMtime",
		"PrefetchSize", "PrefetchHash", "PrefetchVersion", "PrefetchFile", "PrefetchCount"}
}

func Process_Windows_Timeline_Prefetch(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Timeline_Prefetch{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.PrefetchMtime.String(), clientIdentifier, tmp.Hostname, tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.PrefetchMtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       tmp.Hostname,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.FileName,
			MetaData:         fmt.Sprintf("Created Date: %v, Message: %v", tmp.PrefetchCtime, tmp.Message),
		}
		outputChannel <- tmp2.StringArray()
	}
}
