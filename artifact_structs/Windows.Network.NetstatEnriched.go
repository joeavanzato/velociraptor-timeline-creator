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

type Windows_Network_NetstatEnriched struct {
	Pid         int    `json:"Pid"`
	Ppid        int    `json:"Ppid"`
	Name        string `json:"Name"`
	Path        string `json:"Path"`
	CommandLine string `json:"CommandLine"`
	Hash        struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hash"`
	Username     string `json:"Username"`
	Authenticode struct {
		Filename      string `json:"Filename"`
		ProgramName   string `json:"ProgramName"`
		PublisherLink string `json:"PublisherLink"`
		MoreInfoLink  string `json:"MoreInfoLink"`
		SerialNumber  string `json:"SerialNumber"`
		IssuerName    string `json:"IssuerName"`
		SubjectName   string `json:"SubjectName"`
		Timestamp     string `json:"Timestamp"`
		Trusted       string `json:"Trusted"`
		ExtraInfo     struct {
			Catalog string `json:"Catalog"`
		} `json:"_ExtraInfo"`
	} `json:"Authenticode"`
	Family    string    `json:"Family"`
	Type      string    `json:"Type"`
	Status    string    `json:"Status"`
	SrcIP     string    `json:"SrcIP"`
	SrcPort   int       `json:"SrcPort"`
	DestIP    string    `json:"DestIP"`
	DestPort  int       `json:"DestPort"`
	Timestamp time.Time `json:"Timestamp"`
}

func (s Windows_Network_NetstatEnriched) StringArray() []string {
	return []string{strconv.Itoa(s.Pid), strconv.Itoa(s.Ppid), s.Name, s.Path, s.CommandLine, s.Hash.MD5, s.Hash.SHA1, s.Hash.SHA256, s.Username,
		s.Authenticode.Filename, s.Authenticode.ProgramName, s.Authenticode.PublisherLink, s.Authenticode.MoreInfoLink,
		s.Authenticode.SerialNumber, s.Authenticode.IssuerName, s.Authenticode.SubjectName, s.Authenticode.Timestamp,
		s.Authenticode.Trusted, s.Authenticode.ExtraInfo.Catalog, s.Family, s.Type, s.Status, s.SrcIP,
		strconv.Itoa(s.SrcPort), s.DestIP, strconv.Itoa(s.DestPort), s.Timestamp.String()}
}

func (s Windows_Network_NetstatEnriched) GetHeaders() []string {
	return []string{"Pid", "Ppid", "Name", "Path", "CommandLine", "Hash_MD5", "Hash_SHA1", "Hash_SHA256", "Username", "Authenticode_Filename",
		"Authenticode_ProgramName", "Authenticode_PublisherLink", "Authenticode_MoreInfoLink",
		"Authenticode_SerialNumber", "Authenticode_IssuerName", "Authenticode_SubjectName", "Authenticode_Timestamp",
		"Authenticode_Trusted", "Authenticode_Extra_Catalog", "Family", "Type", "Status", "SrcIP", "SrcPort", "DestIP", "DestPort", "Timestamp"}
}

func Process_Windows_Network_NetstatEnriched(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Network_NetstatEnriched{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Timestamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Timestamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       fmt.Sprintf("%v:%v", tmp.SrcIP, tmp.SrcPort),
			DestinationUser:  "",
			DestinationHost:  fmt.Sprintf("%v:%v", tmp.DestIP, tmp.DestPort),
			SourceFile:       tmp.Path,
			MetaData:         fmt.Sprintf("Type: %v, CommandLine: %v, MD5: %v", tmp.Type, tmp.CommandLine, tmp.Hash.MD5),
		}
		outputChannel <- tmp2.StringArray()
	}
}
