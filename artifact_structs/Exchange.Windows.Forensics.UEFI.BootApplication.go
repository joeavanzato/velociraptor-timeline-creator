package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Forensics_UEFI_BootApplication struct {
	OSPath          string    `json:"OSPath"`
	Size            int       `json:"Size"`
	Mtime           time.Time `json:"Mtime"`
	Ctime           time.Time `json:"Ctime"`
	Btime           time.Time `json:"Btime"`
	BootApplication string    `json:"BootApplication"`
}

func (s Exchange_Windows_Forensics_UEFI_BootApplication) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Forensics_UEFI_BootApplication) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Forensics_UEFI_BootApplication(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_UEFI_BootApplication{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Ctime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Ctime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("BootApplication: %v, Modified: %v", tmp.BootApplication, tmp.Mtime),
		}
		outputChannel <- tmp2.StringArray()
	}
}
