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

type Windows_Applications_Chrome_History struct {
	User                            string    `json:"User"`
	UrlId                           int       `json:"url_id"`
	VisitTime                       time.Time `json:"visit_time"`
	VisitedUrl                      string    `json:"visited_url"`
	Title                           string    `json:"title"`
	VisitCount                      int       `json:"visit_count"`
	TypedCount                      int       `json:"typed_count"`
	LastVisitTime                   time.Time `json:"last_visit_time"`
	Hidden                          int       `json:"hidden"`
	FromUrlId                       int       `json:"from_url_id"`
	Source                          any       `json:"Source"`
	VisitDuration                   string    `json:"visit_duration"`
	Transition                      int       `json:"transition"`
	SourceLastModificationTimestamp time.Time `json:"_SourceLastModificationTimestamp"`
	OSPath                          string    `json:"OSPath"`
}

func (s Windows_Applications_Chrome_History) StringArray() []string {
	return []string{s.User, strconv.Itoa(s.UrlId), s.VisitTime.String(), s.VisitedUrl, s.Title, strconv.Itoa(s.VisitCount), strconv.Itoa(s.TypedCount),
		s.LastVisitTime.String(), strconv.Itoa(s.Hidden), strconv.Itoa(s.FromUrlId), fmt.Sprint(s.Source), s.VisitDuration, strconv.Itoa(s.Transition), s.SourceLastModificationTimestamp.String(), s.OSPath}
}

func (s Windows_Applications_Chrome_History) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Applications_Chrome_History(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_Chrome_History{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.VisitTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
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

func (s Windows_Applications_Edge_History) StringArray() []string {
	return []string{s.User, strconv.Itoa(s.UrlId), s.VisitTime.String(), s.VisitedUrl, s.Title, strconv.Itoa(s.VisitCount), strconv.Itoa(s.TypedCount),
		s.LastVisitTime.String(), strconv.Itoa(s.Hidden), strconv.Itoa(s.FromUrlId), fmt.Sprint(s.Source), s.VisitDuration, strconv.Itoa(s.Transition), s.SourceLastModificationTimestamp.String(), s.OSPath, s._Source}
}

func (s Windows_Applications_Edge_History) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Applications_Edge_History(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Applications_Edge_History{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.VisitTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
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
