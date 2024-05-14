package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Generic_Applications_Office_Keywords struct {
	OfficePath    string    `json:"OfficePath"`
	OfficeMtime   time.Time `json:"OfficeMtime"`
	OfficeSize    int       `json:"OfficeSize"`
	InternalMtime time.Time `json:"InternalMtime"`
	HexContext    []string  `json:"HexContext"`
	OSPath        string    `json:"OSPath"`
}

func (s Generic_Applications_Office_Keywords) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Generic_Applications_Office_Keywords) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Applications_Office_Keywords(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Applications_Office_Keywords{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.OfficeMtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.OfficeMtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OfficePath,
			MetaData:         fmt.Sprintf(""),
		}
		outputChannel <- tmp2.StringArray()
	}
}
