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

func (s Windows_Applications_NirsoftBrowserViewer) StringArray() []string {
	return []string{s.URL, fmt.Sprint(s.Title), s.VisitTime, strconv.Itoa(s.VisitCount), s.VisitedFrom, s.VisitType, s.VisitDuration, s.WebBrowser, s.UserProfile, s.BrowserProfile, strconv.Itoa(s.URLLength), fmt.Sprint(s.TypedCount), s.HistoryFile, strconv.Itoa(s.RecordID), s.Visited.String()}
}

func (s Windows_Applications_NirsoftBrowserViewer) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Applications_NirsoftBrowserViewer(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_NirsoftBrowserViewer{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Visited.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
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
