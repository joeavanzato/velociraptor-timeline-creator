package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Windows_Forensics_Timeline struct {
	Application      string    `json:"Application"`
	User             string    `json:"User"`
	LastModifiedTime time.Time `json:"LastModifiedTime"`
	LastExecutionTS  int       `json:"LastExecutionTS"`
}

func (s Windows_Forensics_Timeline) StringArray() []string {
	return []string{s.Application, s.User, s.LastModifiedTime.String()}
}

func (s Windows_Forensics_Timeline) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Forensics_Timeline(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_Timeline{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		timeField := time.Now()
		i, err := strconv.ParseInt(strconv.Itoa(tmp.LastExecutionTS), 10, 64)
		tmpUnix := time.Unix(i, 0)
		timeField = time.Unix(i, 0)
		if i == 0 {
			timeField = tmp.LastModifiedTime
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(timeField.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        timeField,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Application,
			MetaData:         fmt.Sprintf("LastModified: %v, LastExecuted: %v", tmp.LastModifiedTime, tmpUnix),
		}
		outputChannel <- tmp2.StringArray()
	}
}
