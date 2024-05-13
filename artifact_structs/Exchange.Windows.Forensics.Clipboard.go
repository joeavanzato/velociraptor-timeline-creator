package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Forensics_Clipboard struct {
	CreatedTime          time.Time `json:"CreatedTime"`
	LastModifiedTime     time.Time `json:"LastModifiedTime"`
	LastModifiedOnClient time.Time `json:"LastModifiedOnClient"`
	StartTime            time.Time `json:"StartTime"`
	EndTime              time.Time `json:"EndTime"`
	Payload              string    `json:"Payload"`
	User                 string    `json:"User"`
	ClipboardPayload     string    `json:"ClipboardPayload"`
	Path                 string    `json:"Path"`
	Mtime                time.Time `json:"Mtime"`
}

func (s Exchange_Windows_Forensics_Clipboard) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Forensics_Clipboard) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Forensics_Clipboard(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_Clipboard{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.CreatedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.CreatedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Payload: %v, ClipboardPayload: %v, LastModified: %v", tmp.Payload, tmp.ClipboardPayload, tmp.LastModifiedTime, tmp.Path),
		}
		outputChannel <- tmp2.StringArray()
	}
}
