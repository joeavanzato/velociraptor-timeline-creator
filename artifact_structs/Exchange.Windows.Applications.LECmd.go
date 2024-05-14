package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Applications_LECmd struct {
	SourceFile       string `json:"SourceFile"`
	LocalPath        string `json:"LocalPath"`
	Arguments        any    `json:"Arguments"`
	SourceCreated    string `json:"SourceCreated"`
	SourceModified   string `json:"SourceModified"`
	WorkingDirectory string `json:"WorkingDirectory"`
	RelativePath     string `json:"RelativePath"`
	TargetCreated    string `json:"TargetCreated"`
	TargetModified   string `json:"TargetModified"`
	DriveType        string `json:"DriveType"`
	VolumeLabel      string `json:"VolumeLabel"`
}

func (s Exchange_Windows_Applications_LECmd) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Applications_LECmd) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Applications_LECmd(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Applications_LECmd{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006-01-02 15:04:05", tmp.SourceCreated)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(parsedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.SourceFile,
			MetaData:         fmt.Sprintf("LocalPath: %v, Modified: %v, TargetCreated: %v", tmp.LocalPath, tmp.SourceModified, tmp.TargetCreated),
		}
		outputChannel <- tmp2.StringArray()
	}
}
