package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Forensics_Shellbags struct {
	ModTime     time.Time   `json:"ModTime"`
	OSPath      string      `json:"_OSPath"`
	Hive        interface{} `json:"Hive"`
	KeyPath     interface{} `json:"KeyPath"`
	Description any         `json:"Description"`
	Path        string      `json:"Path"`
	RawData     string      `json:"_RawData"`
	Parsed      struct {
		ItemIDSize int `json:"ItemIDSize"`
		Offset     int `json:"Offset"`
		Type       int `json:"Type"`
		Subtype    int `json:"Subtype"`
		ShellBag   struct {
			Size                 int       `json:"Size"`
			Type                 int       `json:"Type"`
			SubType              []string  `json:"SubType"`
			LastModificationTime time.Time `json:"LastModificationTime"`
			ShortName            string    `json:"ShortName"`
			Extension            struct {
				Size         int       `json:"Size"`
				Version      int       `json:"Version"`
				Signature    string    `json:"Signature"`
				CreateDate   time.Time `json:"CreateDate"`
				LastAccessed time.Time `json:"LastAccessed"`
				MFTReference struct {
					MFTID          int `json:"MFTID"`
					SequenceNumber int `json:"SequenceNumber"`
				} `json:"MFTReference"`
				LongName string `json:"LongName"`
			} `json:"Extension"`
			Description any `json:"Description"`
		} `json:"ShellBag"`
	} `json:"_Parsed"`
}

func Process_Windows_Forensics_Shellbags(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_Shellbags{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Parsed.ShellBag.Extension.LastAccessed,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Parsed.ShellBag.Extension.LongName,
			MetaData:         fmt.Sprintf("LastModified: %v, LastAccessed: %v, Created: %v, ", tmp.ModTime, tmp.Parsed.ShellBag.Extension.LastAccessed, tmp.Parsed.ShellBag.Extension.CreateDate),
		}
		outputChannel <- tmp2.StringArray()
	}
}
