package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"path"
	"slices"
	"strings"
	"time"
)

type Windows_NTFS_MFT struct {
	EntryNumber          int       `json:"EntryNumber"`
	InUse                bool      `json:"InUse"`
	ParentEntryNumber    int       `json:"ParentEntryNumber"`
	OSPath               string    `json:"OSPath"`
	Links                []string  `json:"_Links"`
	FileName             string    `json:"FileName"`
	FileSize             int       `json:"FileSize"`
	ReferenceCount       int       `json:"ReferenceCount"`
	IsDir                bool      `json:"IsDir"`
	Created0X10          time.Time `json:"Created0x10"`
	Created0X30          time.Time `json:"Created0x30"`
	LastModified0X10     time.Time `json:"LastModified0x10"`
	LastModified0X30     time.Time `json:"LastModified0x30"`
	LastRecordChange0X10 time.Time `json:"LastRecordChange0x10"`
	LastRecordChange0X30 time.Time `json:"LastRecordChange0x30"`
	LastAccess0X10       time.Time `json:"LastAccess0x10"`
	LastAccess0X30       time.Time `json:"LastAccess0x30"`
	HasADS               bool      `json:"HasADS"`
	SILtFN               bool      `json:"SI_Lt_FN"`
	USecZeros            bool      `json:"uSecZeros"`
	Copied               bool      `json:"Copied"`
	FileNames            []string  `json:"FileNames"`
	FileNameTypes        string    `json:"FileNameTypes"`
}

func (s Windows_NTFS_MFT) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_NTFS_MFT) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_NTFS_MFT(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_NTFS_MFT{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		fileExtension := path.Ext(tmp.FileName)
		if arguments["mftlight"].(bool) {
			if !slices.Contains(vars.LightMFTExtensionsOfInterest, strings.ToLower(fileExtension)) {
				continue
			}
		} else if arguments["mftfull"].(bool) {
			// process as normal
		} else {
			// we shouldn't even be here - a logic check failed
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Created0X10.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Created0X10,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Modified: %v, LastAccess: %v, HasADS: %v, IsDir: %v", tmp.LastModified0X10, tmp.LastAccess0X10, tmp.HasADS, tmp.IsDir),
		}
		outputChannel <- tmp2.StringArray()
	}
}
