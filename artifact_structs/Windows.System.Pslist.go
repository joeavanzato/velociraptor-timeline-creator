package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
	"strconv"
)

type Windows_System_Pslist struct {
	Pid             int    `json:"Pid"`
	Ppid            int    `json:"Ppid"`
	TokenIsElevated bool   `json:"TokenIsElevated"`
	Name            string `json:"Name"`
	CommandLine     string `json:"CommandLine"`
	Exe             string `json:"Exe"`
	TokenInfo       any    `json:"TokenInfo"`
	Hash            struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hash"`
	Authenticode struct {
		Filename      string `json:"Filename"`
		ProgramName   string `json:"ProgramName"`
		PublisherLink string `json:"PublisherLink"`
		MoreInfoLink  string `json:"MoreInfoLink"`
		SerialNumber  string `json:"SerialNumber"`
		IssuerName    string `json:"IssuerName"`
		SubjectName   string `json:"SubjectName"`
		Timestamp     any    `json:"Timestamp"`
		Trusted       string `json:"Trusted"`
		ExtraInfo     any    `json:"_ExtraInfo"`
	} `json:"Authenticode"`
	Username       string `json:"Username"`
	WorkingSetSize int    `json:"WorkingSetSize"`
}

func (s Windows_System_Pslist) StringArray() []string {
	return []string{strconv.Itoa(s.Pid), strconv.Itoa(s.Ppid), strconv.FormatBool(s.TokenIsElevated), s.Name, s.CommandLine, s.Exe, fmt.Sprint(s.TokenInfo), s.Hash.MD5, s.Hash.SHA1, s.Hash.SHA256, s.Authenticode.Filename, s.Authenticode.ProgramName, s.Authenticode.PublisherLink,
		s.Authenticode.MoreInfoLink, s.Authenticode.SerialNumber, s.Authenticode.IssuerName, s.Authenticode.SubjectName, fmt.Sprint(s.Authenticode.Timestamp), s.Authenticode.Trusted, fmt.Sprint(s.Authenticode.ExtraInfo), s.Username, strconv.Itoa(s.WorkingSetSize)}
}

func (s Windows_System_Pslist) GetHeaders() []string {
	return []string{"Pid", "Ppid", "TokenIsElevated", "Name", "CommandLine", "Exe", "TokenInfo", "MD5", "SHA1", "SHA256", "Authenticode_Filename", "Authenticode_ProgramName", "Authenticode_PublisherLink", "Authenticode_MoreInfoLink",
		"Authenticode_SerialNumber", "Authenticode_IssuerName", "Authenticode_SubjectName", "Authenticode_Timestamp", "Authenticode_Trusted", "Authenticode_ExtraInfo", "Username", "WorkingSetSize"}
}

func Process_Windows_System_Pslist(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Pslist{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
	}
}
