package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"slices"
	"strings"
	"time"
)

type Custom_Windows_MFT struct {
	EntryNumber           int       `json:"EntryNumber"`
	SequenceNumber        int       `json:"SequenceNumber"`
	ParentEntryNumber     int       `json:"ParentEntryNumber"`
	ParentSequenceNumber  int       `json:"ParentSequenceNumber"`
	InUse                 bool      `json:"InUse"`
	ParentPath            string    `json:"ParentPath"`
	FileName              string    `json:"FileName"`
	Extension             string    `json:"Extension"`
	IsDirectory           bool      `json:"IsDirectory"`
	HasAds                bool      `json:"HasAds"`
	IsAds                 bool      `json:"IsAds"`
	FileSize              int       `json:"FileSize"`
	Created0X10           time.Time `json:"Created0x10"`
	LastModified0X10      time.Time `json:"LastModified0x10"`
	LastModified0X30      time.Time `json:"LastModified0x30"`
	LastRecordChange0X10  time.Time `json:"LastRecordChange0x10"`
	LastRecordChange0X30  time.Time `json:"LastRecordChange0x30"`
	LastAccess0X10        time.Time `json:"LastAccess0x10"`
	LastAccess0X30        time.Time `json:"LastAccess0x30"`
	UpdateSequenceNumber  int64     `json:"UpdateSequenceNumber"`
	LogfileSequenceNumber int64     `json:"LogfileSequenceNumber"`
	SecurityID            int       `json:"SecurityId"`
	SiFlags               int       `json:"SiFlags"`
	ReferenceCount        int       `json:"ReferenceCount"`
	NameType              int       `json:"NameType"`
	Timestomped           bool      `json:"Timestomped"`
	USecZeros             bool      `json:"uSecZeros"`
	Copied                bool      `json:"Copied"`
	FnAttributeID         int       `json:"FnAttributeId"`
	OtherAttributeID      int       `json:"OtherAttributeId"`
}

func Process_Custom_Windows_MFT(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Custom_Windows_MFT{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["mftlight"].(bool) {
			if !slices.Contains(vars.LightMFTExtensionsOfInterest, strings.ToLower(tmp.Extension)) {
				continue
			}
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
			SourceFile:       fmt.Sprintf("%v\\%v", tmp.ParentPath, tmp.FileName),
			MetaData:         fmt.Sprintf("LastModified: %v, LastAccess: %v, Timestomped: %v, Copied: %v", tmp.LastModified0X10, tmp.LastAccess0X10, tmp.Timestomped, tmp.Copied),
		}
		outputChannel <- tmp2.StringArray()
	}
}
