package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Timeline_Prefetch_Improved struct {
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

func (s Exchange_Windows_Timeline_Prefetch_Improved) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Timeline_Prefetch_Improved) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Timeline_Prefetch_Improved(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Timeline_Prefetch_Improved{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, tmp.Hostname, tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Message,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       tmp.Hostname,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.FileName,
			MetaData:         fmt.Sprintf("Prefetch Last Modified: %v, Execution Count: %v", tmp.PrefetchMtime, tmp.PrefetchCount),
		}
		outputChannel <- tmp2.StringArray()
	}
}
