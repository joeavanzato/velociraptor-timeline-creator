package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Sys_Users struct {
	UID          string    `json:"Uid"`
	Gid          string    `json:"Gid"`
	Name         string    `json:"Name"`
	Description  string    `json:"Description"`
	Directory    string    `json:"Directory"`
	UUID         string    `json:"UUID"`
	Mtime        time.Time `json:"Mtime"`
	HomedirMtime time.Time `json:"HomedirMtime"`
	Data         struct {
		ProfileLoadTime   time.Time `json:"ProfileLoadTime"`
		ProfileUnloadTime time.Time `json:"ProfileUnloadTime"`
	} `json:"Data"`
}

func (s Windows_Sys_Users) StringArray() []string {
	base := helpers.GetStructValuesAsStringSlice(s)
	base = base[:len(base)-1]
	base = append(base, helpers.GetStructValuesAsStringSlice(s.Data)...)
	return base
}

func (s Windows_Sys_Users) GetHeaders() []string {
	base := helpers.GetStructHeadersAsStringSlice(s)
	base = base[:len(base)-1]
	base = append(base, helpers.GetStructHeadersAsStringSlice(s.Data)...)
	return base
}

func Process_Windows_Sys_Users(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_Users{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.HomedirMtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.HomedirMtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Name,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Directory,
			MetaData:         fmt.Sprintf("LastProfileLoadTime:%v", tmp.Data.ProfileLoadTime),
		}
		outputChannel <- tmp2.StringArray()
	}
}
