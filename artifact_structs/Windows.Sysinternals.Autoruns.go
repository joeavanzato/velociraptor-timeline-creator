package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Sysinternals_Autoruns struct {
	Time          string `json:"Time"`
	EntryLocation string `json:"Entry Location"`
	Entry         string `json:"Entry"`
	Enabled       string `json:"Enabled"`
	Category      string `json:"Category"`
	Profile       string `json:"Profile"`
	Description   string `json:"Description"`
	Signer        string `json:"Signer"`
	Company       string `json:"Company"`
	ImagePath     string `json:"Image Path"`
	Version       string `json:"Version"`
	LaunchString  string `json:"Launch String"`
	MD5           string `json:"MD5"`
	SHA1          string `json:"SHA-1"`
	PESHA1        string `json:"PESHA-1"`
	PESHA256      string `json:"PESHA-256"`
	SHA256        string `json:"SHA-256"`
	IMP           string `json:"IMP"`
}

func Process_Windows_Sysinternals_Autoruns(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sysinternals_Autoruns{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("20060102-150405", tmp.Time)
		if terr != nil {
			parsedTime = time.Now()
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Profile,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.ImagePath,
			MetaData:         fmt.Sprintf("Entry: %v, Enabled: %v, Entry Location: %v, Launch String: %v, MD5: %v", tmp.Entry, tmp.Enabled, tmp.EntryLocation, tmp.LaunchString, tmp.MD5),
		}
		outputChannel <- tmp2.StringArray()
	}
}