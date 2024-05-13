package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Detection_PrefetchHunter struct {
	Executable       string    `json:"Executable"`
	FileAccessed     string    `json:"FileAccessed"`
	FullPath         string    `json:"FullPath"`
	ModificationTime time.Time `json:"ModificationTime"`
	CreationTime     time.Time `json:"CreationTime"`
	Hash             string    `json:"Hash"`
	Binary           string    `json:"Binary"`
}

func (s Exchange_Windows_Detection_PrefetchHunter) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Detection_PrefetchHunter) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Detection_PrefetchHunter(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Detection_PrefetchHunter{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.CreationTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.CreationTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "Item Created",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Binary,
			MetaData:         fmt.Sprintf("Modified: %v", tmp.ModificationTime),
		}
		outputChannel <- tmp2.StringArray()
	}
}
