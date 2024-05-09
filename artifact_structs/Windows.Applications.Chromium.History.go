package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Applications_Chrome_History struct {
	User                            string      `json:"User"`
	UrlId                           int         `json:"url_id"`
	VisitTime                       time.Time   `json:"visit_time"`
	VisitedUrl                      string      `json:"visited_url"`
	Title                           string      `json:"title"`
	VisitCount                      int         `json:"visit_count"`
	TypedCount                      int         `json:"typed_count"`
	LastVisitTime                   time.Time   `json:"last_visit_time"`
	Hidden                          int         `json:"hidden"`
	FromUrlId                       int         `json:"from_url_id"`
	Source                          interface{} `json:"Source"`
	VisitDuration                   string      `json:"visit_duration"`
	Transition                      int         `json:"transition"`
	SourceLastModificationTimestamp time.Time   `json:"_SourceLastModificationTimestamp"`
	OSPath                          string      `json:"OSPath"`
}

type Windows_Applications_Edge_History struct {
	User                            string      `json:"User"`
	UrlId                           int         `json:"url_id"`
	VisitTime                       time.Time   `json:"visit_time"`
	VisitedUrl                      string      `json:"visited_url"`
	Title                           string      `json:"title"`
	VisitCount                      int         `json:"visit_count"`
	TypedCount                      int         `json:"typed_count"`
	LastVisitTime                   time.Time   `json:"last_visit_time"`
	Hidden                          int         `json:"hidden"`
	FromUrlId                       int         `json:"from_url_id"`
	Source                          interface{} `json:"Source"`
	VisitDuration                   string      `json:"visit_duration"`
	Transition                      int         `json:"transition"`
	SourceLastModificationTimestamp time.Time   `json:"_SourceLastModificationTimestamp"`
	OSPath                          string      `json:"OSPath"`
	_Source                         string      `json:"_Source"`
}

func Process_Windows_Applications_Chrome_History(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_Chrome_History{}
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
			DestinationHost:  tmp.VisitedUrl,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Title: %v, Duration: %v, ", tmp.Title, tmp.VisitDuration),
		}
		outputChannel <- tmp2.StringArray()
	}
}

func Process_Windows_Applications_Edge_History(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_Edge_History{}
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
			DestinationHost:  tmp.VisitedUrl,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Title: %v, Duration: %v, ", tmp.Title, tmp.VisitDuration),
		}
		outputChannel <- tmp2.StringArray()
	}
}
