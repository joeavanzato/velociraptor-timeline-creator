package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Detection_Amcache struct {
	HivePath         string    `json:"HivePath"`
	EntryKey         string    `json:"EntryKey"`
	KeyMTime         time.Time `json:"KeyMTime"`
	EntryType        string    `json:"EntryType"`
	SHA1             string    `json:"SHA1"`
	EntryName        string    `json:"EntryName"`
	EntryPath        string    `json:"EntryPath"`
	Publisher        string    `json:"Publisher"`
	OriginalFileName string    `json:"OriginalFileName"`
	BinaryType       string    `json:"BinaryType"`
}

func (s Windows_Detection_Amcache) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Detection_Amcache) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Detection_Amcache(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Detection_Amcache{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.KeyMTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.KeyMTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.EntryName,
			MetaData:         fmt.Sprintf("EntryName: %v, EntryPath: %v, SHA1: %v", tmp.EntryName, tmp.EntryPath, tmp.SHA1),
		}
		outputChannel <- tmp2.StringArray()
	}
}
