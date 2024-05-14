package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Generic_Forensic_Timeline struct {
	Md5    int       `json:"Md5"`
	OSPath string    `json:"OSPath"`
	Inode  any       `json:"Inode"`
	Mode   string    `json:"Mode"`
	UID    any       `json:"Uid"`
	Gid    any       `json:"Gid"`
	Size   int       `json:"Size"`
	Atime  time.Time `json:"Atime"`
	Mtime  time.Time `json:"Mtime"`
	Ctime  time.Time `json:"Ctime"`
}

func (s Generic_Forensic_Timeline) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Generic_Forensic_Timeline) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Forensic_Timeline(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_Timeline{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("MD5: %v, Atime: %v, Ctime: %v, Mode: %v", tmp.Md5, tmp.Atime, tmp.Ctime, tmp.Mode),
		}
		outputChannel <- tmp2.StringArray()
	}
}
