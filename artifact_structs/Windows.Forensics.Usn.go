package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Forensics_Usn struct {
	Timestamp      time.Time `json:"Timestamp"`
	Filename       string    `json:"Filename"`
	Device         string    `json:"Device"`
	OSPath         string    `json:"OSPath"`
	Links          []string  `json:"_Links"`
	Reason         []string  `json:"Reason"`
	MFTId          int       `json:"MFTId"`
	Sequence       int       `json:"Sequence"`
	ParentMFTId    int       `json:"ParentMFTId"`
	ParentSequence int       `json:"ParentSequence"`
	FileAttributes []string  `json:"FileAttributes"`
	SourceInfo     []string  `json:"SourceInfo"`
	Usn            any       `json:"Usn"`
}

func (s Windows_Forensics_Usn) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Forensics_Usn) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Forensics_Usn(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_Usn{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Timestamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Timestamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Reason: %v, FileAttributes: %v, MFTId: %v", tmp.Reason, tmp.FileAttributes, tmp.MFTId),
		}
		outputChannel <- tmp2.StringArray()
	}
}
