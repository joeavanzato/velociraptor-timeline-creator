package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Windows_Registry_AppCompatCache struct {
	Position         int       `json:"Position"`
	ModificationTime time.Time `json:"ModificationTime"`
	Path             string    `json:"Path"`
	ControlSet       string    `json:"ControlSet"`
	Key              string    `json:"Key"`
}

func (s Windows_Registry_AppCompatCache) StringArray() []string {
	return []string{strconv.Itoa(s.Position), s.ModificationTime.String(), s.Path, s.ControlSet, s.Key}
}

func (s Windows_Registry_AppCompatCache) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Registry_AppCompatCache(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Registry_AppCompatCache{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.ModificationTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.ModificationTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Path,
			MetaData:         "",
		}
		outputChannel <- tmp2.StringArray()
	}
}
