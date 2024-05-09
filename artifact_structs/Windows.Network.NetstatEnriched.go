package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
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
	Username     string    `json:"Username"`
	Authenticode any       `json:"Authenticode"`
	Family       string    `json:"Family"`
	Type         string    `json:"Type"`
	Status       string    `json:"Status"`
	SrcIP        string    `json:"SrcIP"`
	SrcPort      int       `json:"SrcPort"`
	DestIP       string    `json:"DestIP"`
	DestPort     int       `json:"DestPort"`
	Timestamp    time.Time `json:"Timestamp"`
}

func Process_Windows_Network_NetstatEnriched(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Network_NetstatEnriched{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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
