package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Forensics_Jumplists_JLECmd struct {
	SourceFile           string `json:"SourceFile"`
	SourceCreated        string `json:"SourceCreated"`
	SourceModified       string `json:"SourceModified"`
	LocalPath            string `json:"LocalPath"`
	Arguments            string `json:"Arguments"`
	TargetCreated        string `json:"TargetCreated"`
	TargetModified       string `json:"TargetModified"`
	VolumeLabel          string `json:"VolumeLabel"`
	DriveType            string `json:"DriveType"`
	AppIDDescription     string `json:"AppIdDescription"`
	CommonPath           any    `json:"CommonPath"`
	VolumeSerialNumber   any    `json:"VolumeSerialNumber"`
	MachineID            string `json:"MachineID"`
	MachineMACAddress    string `json:"MachineMACAddress"`
	TargetMFTEntryNumber string `json:"TargetMFTEntryNumber"`
	TargetSequenceNumber int    `json:"TargetSequenceNumber"`
	TargetIDAbsolutePath string `json:"TargetIDAbsolutePath"`
	TrackerCreatedOn     string `json:"TrackerCreatedOn"`
	ExtraBlocksPresent   string `json:"ExtraBlocksPresent"`
	HeaderFlags          string `json:"HeaderFlags"`
	FileAttributes       any    `json:"FileAttributes"`
	FileSize             int    `json:"FileSize"`
}

func (s Exchange_Windows_Forensics_Jumplists_JLECmd) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_Forensics_Jumplists_JLECmd) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Forensics_Jumplists_JLECmd(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_Jumplists_JLECmd{}
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
			MetaData:         fmt.Sprintf("Arguments: %v, Modified: %v", tmp.Arguments, tmp.SourceModified),
		}
		outputChannel <- tmp2.StringArray()
	}
}
