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

type Windows_Timeline_MFT struct {
	EventTime time.Time `json:"event_time"`
	Hostname  string    `json:"hostname"`
	Parser    string    `json:"parser"`
	Source    string    `json:"source"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
	Optional1 struct {
		EntryNumber       int  `json:"EntryNumber"`
		ParentEntryNumber int  `json:"ParentEntryNumber"`
		FileSize          int  `json:"FileSize"`
		IsDir             bool `json:"IsDir"`
		InUse             bool `json:"InUse"`
	} `json:"optional_1"`
	Optional2 struct {
		FNCreatedShift bool `json:"FNCreatedShift"`
		USecZero       bool `json:"USecZero"`
		PossibleCopy   bool `json:"PossibleCopy"`
		VolumeCopy     bool `json:"VolumeCopy"`
	} `json:"optional_2"`
	Optional3 struct {
		LastModified0X10     time.Time `json:"LastModified0x10"`
		LastAccess0X10       time.Time `json:"LastAccess0x10"`
		LastRecordChange0X10 time.Time `json:"LastRecordChange0x10"`
		Created0X10          time.Time `json:"Created0x10"`
	} `json:"optional_3"`
	Optional4 struct {
		LastModified0X30     time.Time `json:"LastModified0x30"`
		LastAccess0X30       time.Time `json:"LastAccess0x30"`
		LastRecordChange0X30 time.Time `json:"LastRecordChange0x30"`
		Created0X30          time.Time `json:"Created0x30"`
	} `json:"optional_4"`
}

func (s Windows_Timeline_MFT) StringArray() []string {
	base := helpers.GetStructValuesAsStringSlice(s)
	base = base[:len(base)-4]
	base = append(base, helpers.GetStructValuesAsStringSlice(s.Optional1)...)
	base = append(base, helpers.GetStructValuesAsStringSlice(s.Optional2)...)
	base = append(base, helpers.GetStructValuesAsStringSlice(s.Optional3)...)
	base = append(base, helpers.GetStructValuesAsStringSlice(s.Optional4)...)
	return base
}

func (s Windows_Timeline_MFT) GetHeaders() []string {
	base := helpers.GetStructHeadersAsStringSlice(s)
	base = base[:len(base)-4]
	base = append(base, helpers.GetStructHeadersAsStringSlice(s.Optional1)...)
	base = append(base, helpers.GetStructHeadersAsStringSlice(s.Optional2)...)
	base = append(base, helpers.GetStructHeadersAsStringSlice(s.Optional3)...)
	base = append(base, helpers.GetStructHeadersAsStringSlice(s.Optional4)...)
	return base
}

func Process_Windows_Timeline_MFT(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Timeline_MFT{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		fileExtension := path.Ext(tmp.Path)
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
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, tmp.Hostname, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: tmp.Message,
			SourceUser:       "",
			SourceHost:       tmp.Hostname,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Path,
			MetaData:         fmt.Sprintf("SI_LastModified: %v, SI_LastAccess: %v, FileSize: %v", tmp.Optional3.LastModified0X10, tmp.Optional3.LastAccess0X10, tmp.Optional1.FileSize),
		}
		outputChannel <- tmp2.StringArray()
	}
}
