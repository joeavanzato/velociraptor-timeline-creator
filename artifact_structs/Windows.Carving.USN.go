package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Carving_USN struct {
	Offset      int64     `json:"Offset"`
	TimeStamp   time.Time `json:"TimeStamp"`
	Name        string    `json:"Name"`
	MFTID       int       `json:"MFTId"`
	OSPath      string    `json:"OSPath"`
	ParentMFTID int       `json:"ParentMFTId"`
	Reason      []string  `json:"Reason"`
}

func (s Windows_Carving_USN) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Carving_USN) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Carving_USN(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Carving_USN{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.TimeStamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimeStamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Reason: %v, MFTId: %v", tmp.Reason, tmp.MFTID),
		}
		outputChannel <- tmp2.StringArray()
	}
}
