package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type Windows_Forensics_Shellbags struct {
	ModTime     time.Time `json:"ModTime"`
	OSPath      string    `json:"_OSPath"`
	Hive        string    `json:"Hive"`
	KeyPath     string    `json:"KeyPath"`
	Description struct {
		Type         any       `json:"Type"`
		Modified     time.Time `json:"Modified"`
		LastAccessed time.Time `json:"LastAccessed"`
		CreateDate   time.Time `json:"CreateDate"`
		ShortName    string    `json:"ShortName"`
		LongName     string    `json:"LongName"`
		MFTID        int       `json:"MFTID"`
		MFTSeq       int       `json:"MFTSeq"`
	} `json:"Description"`
	Path    string `json:"Path"`
	RawData string `json:"_RawData"`
	Parsed  struct {
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
			Description struct {
				Type         any       `json:"Type"`
				Modified     time.Time `json:"Modified"`
				LastAccessed time.Time `json:"LastAccessed"`
				CreateDate   time.Time `json:"CreateDate"`
				ShortName    string    `json:"ShortName"`
				LongName     string    `json:"LongName"`
				MFTID        int       `json:"MFTID"`
				MFTSeq       int       `json:"MFTSeq"`
			} `json:"Description"`
		} `json:"ShellBag"`
	} `json:"_Parsed"`
}

func (s Windows_Forensics_Shellbags) StringArray() []string {
	return []string{s.ModTime.String(), s.OSPath, s.Hive, s.KeyPath, fmt.Sprint(s.Description.Type),
		s.Description.Modified.String(), s.Description.LastAccessed.String(), s.Description.CreateDate.String(), s.Description.ShortName,
		s.Description.LongName, strconv.Itoa(s.Description.MFTID), strconv.Itoa(s.Description.MFTSeq), s.Path, s.RawData,
		strconv.Itoa(s.Parsed.ItemIDSize), strconv.Itoa(s.Parsed.Offset), strconv.Itoa(s.Parsed.Type), strconv.Itoa(s.Parsed.Subtype),
		strconv.Itoa(s.Parsed.ShellBag.Size), strconv.Itoa(s.Parsed.ShellBag.Type), fmt.Sprint(s.Parsed.ShellBag.SubType),
		s.Parsed.ShellBag.LastModificationTime.String(), s.Parsed.ShellBag.ShortName, strconv.Itoa(s.Parsed.ShellBag.Extension.Size),
		strconv.Itoa(s.Parsed.ShellBag.Extension.Version), s.Parsed.ShellBag.Extension.Signature,
		s.Parsed.ShellBag.Extension.CreateDate.String(), s.Parsed.ShellBag.Extension.LastAccessed.String(),
		strconv.Itoa(s.Parsed.ShellBag.Extension.MFTReference.MFTID), strconv.Itoa(s.Parsed.ShellBag.Extension.MFTReference.SequenceNumber),
		s.Parsed.ShellBag.Extension.LongName, fmt.Sprint(s.Parsed.ShellBag.Description.Type), s.Parsed.ShellBag.Description.Modified.String(),
		s.Parsed.ShellBag.Description.LastAccessed.String(), s.Parsed.ShellBag.Description.CreateDate.String(),
		s.Parsed.ShellBag.Description.ShortName, s.Parsed.ShellBag.Description.LongName, strconv.Itoa(s.Parsed.ShellBag.Description.MFTID), strconv.Itoa(s.Parsed.ShellBag.Description.MFTSeq)}
}

func (s Windows_Forensics_Shellbags) GetHeaders() []string {
	return []string{"ModTime", "OSPath", "Hive", "KeyPath", "Description_Type", "Description_Modified", "Description_LastAccessed", "Description_CreateDate", "Description_ShortName", "Description_LongName",
		"Description_MFTID", "Description_MFTSeq", "Path", "RawData", "Parsed_ItemIDSize", "Parsed_Offset", "Parsed_Type",
		"Parsed_Subtype", "Parsed_ShellBag_Size", "Parsed_ShellBag_Type", "Parsed_ShellBag_SubType", "Parsed_ShellBag_LastModificationTime",
		"Parsed_ShellBag_ShortName", "Parsed_ShellBag_Extension_Size", "Parsed_ShellBag_Extension_Version",
		"Parsed_ShellBag_Extension_Signature", "Parsed_ShellBag_Extension_CreateDate", "Parsed_ShellBag_Extension_LastAccessed",
		"Parsed_ShellBag_Extension_MFTReference_MFTID", "Parsed_ShellBag_Extension_MFTReference_SequenceNumber",
		"Parsed_ShellBag_Extension_LongName", "Parsed_ShellBag_Description_Type", "Parsed_ShellBag_Description_Modified", "Parsed_ShellBag_Description_LastAccessed",
		"Parsed_ShellBag_Description_CreateDate", "Parsed_ShellBag_Description_ShortName", "Parsed_ShellBag_Description_LongName", "Parsed_ShellBag_Description_MFTID", "Parsed_ShellBag_Description_MFTSeq"}
}

func Process_Windows_Forensics_Shellbags(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_Shellbags{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Parsed.ShellBag.Extension.LastAccessed.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
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
