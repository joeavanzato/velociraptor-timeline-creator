package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Applications_NirsoftBrowserViewer struct {
	URL            string    `json:"URL"`
	Title          any       `json:"Title"`
	VisitTime      string    `json:"Visit Time"`
	VisitCount     int       `json:"Visit Count"`
	VisitedFrom    string    `json:"Visited From"`
	VisitType      string    `json:"Visit Type"`
	VisitDuration  string    `json:"Visit Duration"`
	WebBrowser     string    `json:"Web Browser"`
	UserProfile    string    `json:"User Profile"`
	BrowserProfile string    `json:"Browser Profile"`
	URLLength      int       `json:"URL Length"`
	TypedCount     any       `json:"Typed Count"`
	HistoryFile    string    `json:"History File"`
	RecordID       int       `json:"Record ID"`
	Visited        time.Time `json:"Visited"`
}

func Process_Windows_Applications_NirsoftBrowserViewer(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_NirsoftBrowserViewer{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Visited,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        fmt.Sprintf("URL Visit (%v)", tmp.WebBrowser),
			EventDescription: "",
			SourceUser:       tmp.UserProfile,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URL,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Title: %v, Visited From: %v, Visit Count: %v", tmp.Title, tmp.VisitedFrom, tmp.VisitCount),
		}
		outputChannel <- tmp2.StringArray()
	}
}
