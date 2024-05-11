package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Windows_Forensics_Lnk struct {
	SourceFile struct {
		OSPath string    `json:"OSPath"`
		Size   int       `json:"Size"`
		Mtime  time.Time `json:"Mtime"`
		Btime  time.Time `json:"Btime"`
	} `json:"SourceFile"`
	ShellLinkHeader struct {
		Headersize     int       `json:"Headersize"`
		LinkClsID      string    `json:"LinkClsID"`
		LinkFlags      []string  `json:"LinkFlags"`
		FileAttributes any       `json:"FileAttributes"`
		FileSize       int       `json:"FileSize"`
		CreationTime   time.Time `json:"CreationTime"`
		AccessTime     time.Time `json:"AccessTime"`
		WriteTime      time.Time `json:"WriteTime"`
		IconIndex      int       `json:"IconIndex"`
		ShowCommand    string    `json:"ShowCommand"`
		HotKey         string    `json:"HotKey"`
	} `json:"ShellLinkHeader"`
	LinkInfo struct {
		LinkInfoFlags []string `json:"LinkInfoFlags"`
		Target        struct {
			Path       string `json:"Path"`
			VolumeInfo struct {
				DriveType         string `json:"DriveType"`
				DriveSerialNumber int64  `json:"DriveSerialNumber"`
				VolumeLabel       string `json:"VolumeLabel"`
			} `json:"VolumeInfo"`
		} `json:"Target"`
	} `json:"LinkInfo"`
	LinkTarget struct {
		LinkTarget       string `json:"LinkTarget"`
		LinkTargetIDList struct {
			IDListSize int `json:"IDListSize"`
			IDList     []struct {
				ItemIDSize int `json:"ItemIDSize"`
				Offset     int `json:"Offset"`
				Type       int `json:"Type"`
				Subtype    int `json:"Subtype"`
				ShellBag   struct {
					Description struct {
						ShortName string `json:"ShortName"`
						Type      any    `json:"Type"`
					} `json:"Description"`
				} `json:"ShellBag"`
			} `json:"IDList"`
		} `json:"LinkTargetIDList"`
	} `json:"LinkTarget"`
	StringData struct {
		TargetPath   string `json:"TargetPath"`
		Name         string `json:"Name"`
		RelativePath string `json:"RelativePath"`
		WorkingDir   string `json:"WorkingDir"`
		Arguments    string `json:"Arguments"`
		IconLocation string `json:"IconLocation"`
	} `json:"StringData"`
	ExtraData  any `json:"ExtraData"`
	Suspicious any `json:"Suspicious"`
}

func (s Windows_Forensics_Lnk) StringArray() []string {
	idlist := make([]string, 0)
	for _, v := range s.LinkTarget.LinkTargetIDList.IDList {
		idlist = append(idlist, fmt.Sprintf("ItemIDSize: %v, Offset: %v, Type: %v, Subtype: %v, Shellbag_Description_ShortName: %v, ShellBag_Description_Type: %v", v.ItemIDSize, v.Offset, v.Type, v.Subtype, v.ShellBag.Description.ShortName, v.ShellBag.Description.Type))
	}
	return []string{s.SourceFile.OSPath, strconv.Itoa(s.SourceFile.Size), s.SourceFile.Mtime.String(), s.SourceFile.Btime.String(),
		strconv.Itoa(s.ShellLinkHeader.Headersize), s.ShellLinkHeader.LinkClsID, fmt.Sprint(s.ShellLinkHeader.LinkFlags), fmt.Sprint(s.ShellLinkHeader.FileAttributes),
		strconv.Itoa(s.ShellLinkHeader.FileSize), s.ShellLinkHeader.CreationTime.String(), s.ShellLinkHeader.AccessTime.String(), s.ShellLinkHeader.WriteTime.String(),
		strconv.Itoa(s.ShellLinkHeader.IconIndex), s.ShellLinkHeader.ShowCommand, s.ShellLinkHeader.HotKey, fmt.Sprint(s.LinkInfo.LinkInfoFlags), s.LinkInfo.Target.Path,
		s.LinkInfo.Target.VolumeInfo.DriveType, strconv.FormatInt(s.LinkInfo.Target.VolumeInfo.DriveSerialNumber, 10), s.LinkInfo.Target.VolumeInfo.VolumeLabel,
		s.LinkTarget.LinkTarget, fmt.Sprint(idlist), s.StringData.TargetPath, s.StringData.Name, s.StringData.RelativePath, s.StringData.WorkingDir, s.StringData.Arguments, s.StringData.IconLocation,
		fmt.Sprint(s.ExtraData), fmt.Sprint(s.Suspicious)}
}

func (s Windows_Forensics_Lnk) GetHeaders() []string {
	return []string{"SourceFile_OSPath", "SourceFile_Size", "SourceFile_Mtime", "SourceFile_Btime", "ShellLinkHeader_HeaderSize", "ShellLinkHeader_LinkClsID", "ShellLinkHeader_LinkFlags", "ShellLinkHeader_FileAttributes",
		"ShellLinkHeader_FileSize", "ShellLinkHeader_CreationTime", "ShellLinkHeader_AccessTime", "ShellLinkHeader_WriteTime", "ShellLinkHeader_IconIndex", "ShellLinkHeader_ShowCommand", "ShellLinkHeader_HotKey",
		"LinkInfo_LinkInfoFlags", "LinkInfo_Target_Path", "LinkInfo_Target_VolumeInfo_DriveType", "LinkInfo_Target_VolumeInfo_DriveSerialNumber", "LinkInfo_Target_VolumeInfo_VolumeLabel", "LinkTarget_LinkTarget", "LinkTarget_LinkTargetIDList_IDList",
		"StringData_TargetPath", "StringData_Name", "StringData_RelativePath", "StringData_WorkingDir", "StringData_Arguments", "StringData_IconLocation", "ExtraData", "Suspicious"}
}

func Process_Windows_Forensics_Lnk(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_Lnk{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.SourceFile.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.SourceFile.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.SourceFile.OSPath,
			MetaData:         fmt.Sprintf("LinkTarget: %v, Arguments: %v, Created: %v, Access: %v", tmp.StringData.TargetPath, tmp.StringData.Arguments, tmp.ShellLinkHeader.CreationTime, tmp.ShellLinkHeader.AccessTime),
		}
		outputChannel <- tmp2.StringArray()
	}
}
