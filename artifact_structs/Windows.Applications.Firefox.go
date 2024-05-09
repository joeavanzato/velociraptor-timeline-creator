package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Applications_Firefox_Downloads struct {
	User           string      `json:"User"`
	StartTime      time.Time   `json:"startTime"`
	EndTime        interface{} `json:"endTime"`
	LastModified   time.Time   `json:"last_modified"`
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	URL            string      `json:"url"`
	PlaceID        int         `json:"place_id"`
	FileSize       interface{} `json:"fileSize"`
	State          interface{} `json:"state"`
	LocalDirectory string      `json:"localDirectory"`
	Flags          int         `json:"flags"`
	Expiration     int         `json:"expiration"`
	Type           int         `json:"type"`
}

func Process_Windows_Applications_Firefox_Downloads(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_Firefox_Downloads{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.StartTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URL,
			SourceFile:       tmp.LocalDirectory,
			MetaData:         fmt.Sprintf("File Size: %v", tmp.FileSize),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Windows_Applications_Firefox_History struct {
	User        string    `json:"User"`
	OSPath      string    `json:"OSPath"`
	VisitTime   time.Time `json:"visit_time"`
	PlaceID     int       `json:"place_id"`
	URLVisited  string    `json:"url_visited"`
	Title       string    `json:"title"`
	RevHost     string    `json:"rev_host"`
	VisitCount  int       `json:"visit_count"`
	Hidden      int       `json:"hidden"`
	Typed       int       `json:"typed"`
	Description string    `json:"description"`
}

func Process_Windows_Applications_Firefox_History(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_Firefox_History{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.VisitTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URLVisited,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Title: %v, Hidden: %v, Typed: %v, Visit Count: %v", tmp.Title, tmp.Hidden, tmp.Typed, tmp.VisitCount),
		}
		outputChannel <- tmp2.StringArray()
	}
}
