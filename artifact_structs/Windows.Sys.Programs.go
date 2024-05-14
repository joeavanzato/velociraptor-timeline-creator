package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Sys_Programs struct {
	KeyName               string    `json:"KeyName"`
	KeyLastWriteTimestamp time.Time `json:"KeyLastWriteTimestamp"`
	DisplayName           string    `json:"DisplayName"`
	DisplayVersion        any       `json:"DisplayVersion"`
	InstallLocation       string    `json:"InstallLocation"`
	InstallSource         string    `json:"InstallSource"`
	Language              any       `json:"Language"`
	Publisher             string    `json:"Publisher"`
	UninstallString       string    `json:"UninstallString"`
	InstallDate           any       `json:"InstallDate"`
	KeyPath               string    `json:"KeyPath"`
}

func (s Windows_Sys_Programs) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Sys_Programs) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Sys_Programs(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_Programs{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.KeyLastWriteTimestamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.KeyLastWriteTimestamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Name: %v, DisplayName: %v, InstallLocation: %v, Uninstall String: %v", tmp.KeyName, tmp.DisplayName, tmp.InstallLocation, tmp.UninstallString),
		}
		outputChannel <- tmp2.StringArray()
	}
}
