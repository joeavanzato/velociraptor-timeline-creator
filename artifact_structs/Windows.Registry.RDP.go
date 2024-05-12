package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Registry_RDP_Servers struct {
	Username      string    `json:"Username"`
	SID           string    `json:"SID"`
	HiveName      string    `json:"HiveName"`
	Key           string    `json:"Key"`
	LastWriteTime time.Time `json:"LastWriteTime"`
	Server        string    `json:"Server"`
	UsernameHint  string    `json:"UsernameHint"`
	CertHash      string    `json:"CertHash"`
}

func (s Windows_Registry_RDP_Servers) StringArray() []string {
	return []string{s.Username, s.SID, s.HiveName, s.Key, s.LastWriteTime.String(), s.Server, s.UsernameHint, s.CertHash}
}

func (s Windows_Registry_RDP_Servers) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Registry_RDP_Servers(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_RDP_Servers{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastWriteTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastWriteTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  tmp.UsernameHint,
			DestinationHost:  tmp.Server,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("CertHash: %v", tmp.CertHash),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Windows_Registry_RDP_Mru struct {
	Username      string    `json:"Username"`
	SID           string    `json:"SID"`
	HiveName      string    `json:"HiveName"`
	Key           string    `json:"Key"`
	LastWriteTime time.Time `json:"LastWriteTime"`
	Mru           []string  `json:"Mru"`
}

func (s Windows_Registry_RDP_Mru) StringArray() []string {
	return []string{s.Username, s.SID, s.HiveName, s.Key, s.LastWriteTime.String(), fmt.Sprint(s.Mru)}
}

func (s Windows_Registry_RDP_Mru) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Registry_RDP_Mru(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_RDP_Mru{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastWriteTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastWriteTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "RDP MRU Entry Modified",
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  fmt.Sprint(tmp.Mru),
			SourceFile:       tmp.HiveName,
			MetaData:         fmt.Sprintf("SID: %v", tmp.SID),
		}
		outputChannel <- tmp2.StringArray()
	}
}
