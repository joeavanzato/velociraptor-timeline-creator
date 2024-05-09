package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
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
		Headersize     int           `json:"Headersize"`
		LinkClsID      string        `json:"LinkClsID"`
		LinkFlags      []string      `json:"LinkFlags"`
		FileAttributes []interface{} `json:"FileAttributes"`
		FileSize       int           `json:"FileSize"`
		CreationTime   time.Time     `json:"CreationTime"`
		AccessTime     time.Time     `json:"AccessTime"`
		WriteTime      time.Time     `json:"WriteTime"`
		IconIndex      int           `json:"IconIndex"`
		ShowCommand    string        `json:"ShowCommand"`
		HotKey         string        `json:"HotKey"`
	} `json:"ShellLinkHeader"`
	LinkInfo   any `json:"LinkInfo"`
	LinkTarget any `json:"LinkTarget"`
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

func Process_Windows_Forensics_Lnk(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_Lnk{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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
