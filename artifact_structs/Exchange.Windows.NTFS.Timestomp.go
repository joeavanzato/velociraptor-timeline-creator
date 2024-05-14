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

type Exchange_Windows_NTFS_Timestomp struct {
	OSPath            string `json:"OSPath"`
	CreatedTimestamps struct {
		Created0X10 time.Time `json:"Created0x10"`
		Created0X30 time.Time `json:"Created0x30"`
	} `json:"CreatedTimestamps"`
	InUse                 bool `json:"InUse"`
	SIFN                  bool `json:"SI<FN"`
	USecZeros             bool `json:"uSecZeros"`
	SuspiciousCompileTime any  `json:"SuspiciousCompileTime"`
	NtfsMetadata          struct {
		FullPath       string `json:"FullPath"`
		MFTID          int    `json:"MFTID"`
		SequenceNumber int    `json:"SequenceNumber"`
		Size           int64  `json:"Size"`
		Allocated      bool   `json:"Allocated"`
		IsDir          bool   `json:"IsDir"`
		SITimes        struct {
			CreateTime       time.Time `json:"CreateTime"`
			FileModifiedTime time.Time `json:"FileModifiedTime"`
			MFTModifiedTime  time.Time `json:"MFTModifiedTime"`
			AccessedTime     time.Time `json:"AccessedTime"`
		} `json:"SI_Times"`
		Filenames []struct {
			Times struct {
				CreateTime       time.Time `json:"CreateTime"`
				FileModifiedTime time.Time `json:"FileModifiedTime"`
				MFTModifiedTime  time.Time `json:"MFTModifiedTime"`
				AccessedTime     time.Time `json:"AccessedTime"`
			} `json:"Times"`
			Type                 string `json:"Type"`
			Name                 string `json:"Name"`
			ParentEntryNumber    int    `json:"ParentEntryNumber"`
			ParentSequenceNumber int    `json:"ParentSequenceNumber"`
		} `json:"Filenames"`
		Attributes []struct {
			Type   string `json:"Type"`
			TypeID int    `json:"TypeId"`
			ID     int    `json:"Id"`
			Inode  string `json:"Inode"`
			Size   int    `json:"Size"`
			Name   string `json:"Name"`
		} `json:"Attributes"`
		Hardlinks []string `json:"Hardlinks"`
		Device    string   `json:"Device"`
		OSPath    string   `json:"OSPath"`
	} `json:"NtfsMetadata"`
	Magic string `json:"Magic"`
}

func (s Exchange_Windows_NTFS_Timestomp) StringArray() []string {

	filenames := make([]string, 0)
	for _, v := range s.NtfsMetadata.Filenames {
		filenames = append(filenames, fmt.Sprintf("| Created: %v, FileModified: %v, MFTModified: %v, Accessed: %v, Type: %v, Name: %v, ParentEntry: %v, ParentSequence: %v", v.Times.CreateTime, v.Times.FileModifiedTime, v.Times.MFTModifiedTime, v.Times.AccessedTime, v.Type, v.Name, v.ParentEntryNumber, v.ParentSequenceNumber))
	}
	attributes := make([]string, 0)
	for _, v := range s.NtfsMetadata.Attributes {
		attributes = append(attributes, fmt.Sprintf("| Type: %v, TypeID: %v, ID: %v, Inode: %v, Size: %v, Name: %v", v.Type, v.TypeID, v.ID, v.Inode, v.Size, v.Name))
	}
	return []string{s.OSPath, s.CreatedTimestamps.Created0X10.String(), s.CreatedTimestamps.Created0X30.String(), strconv.FormatBool(s.InUse), strconv.FormatBool(s.SIFN),
		strconv.FormatBool(s.USecZeros), fmt.Sprint(s.SuspiciousCompileTime), s.NtfsMetadata.FullPath, strconv.Itoa(s.NtfsMetadata.MFTID), strconv.Itoa(s.NtfsMetadata.SequenceNumber), strconv.FormatInt(s.NtfsMetadata.Size, 10),
		strconv.FormatBool(s.NtfsMetadata.Allocated), strconv.FormatBool(s.NtfsMetadata.IsDir), s.NtfsMetadata.SITimes.CreateTime.String(), s.NtfsMetadata.SITimes.FileModifiedTime.String(), s.NtfsMetadata.SITimes.MFTModifiedTime.String(), s.NtfsMetadata.SITimes.AccessedTime.String(),
		fmt.Sprint(filenames), fmt.Sprint(attributes), fmt.Sprint(s.NtfsMetadata.Hardlinks), s.NtfsMetadata.Device, s.NtfsMetadata.OSPath, s.Magic}
}

func (s Exchange_Windows_NTFS_Timestomp) GetHeaders() []string {
	return []string{"OSPath", "Created_0x10", " Created_0x30", "InUse", "SIFN", "USecZeros", "SuspiciousCompileTime",
		"FullPath", "MFTID", "SequenceNumber", "Size", "Allocated", "IsDir", "SI_Create", "SI_FileModified", "SI_MFTModified", "SI_Accessed",
		"Filenames", "Attributes", "Hardlinks", "Device", "Metadata_OSPath", "Magic"}
}

func Process_Exchange_Windows_NTFS_Timestomp(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_NTFS_Timestomp{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.CreatedTimestamps.Created0X10.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.CreatedTimestamps.Created0X10,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Magic: %v, FileModified: %v, MFTModified: %v, Accessed: %v", tmp.Magic, tmp.NtfsMetadata.SITimes.FileModifiedTime, tmp.NtfsMetadata.SITimes.MFTModifiedTime, tmp.NtfsMetadata.SITimes.AccessedTime),
		}
		outputChannel <- tmp2.StringArray()
	}
}
