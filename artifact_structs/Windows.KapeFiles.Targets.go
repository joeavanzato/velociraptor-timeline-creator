package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Windows_KapeFiles_Targets_AllFileMetadata struct {
	Created      time.Time `json:"Created"`
	LastAccessed time.Time `json:"LastAccessed"`
	Modified     time.Time `json:"Modified"`
	Size         int       `json:"Size"`
	SourceFile   string    `json:"SourceFile"`
	Source       string    `json:"_Source"`
}

func (s Windows_KapeFiles_Targets_AllFileMetadata) StringArray() []string {
	return []string{s.Created.String(), s.LastAccessed.String(), s.Modified.String(), strconv.Itoa(s.Size), s.SourceFile, s.Source}
}

func (s Windows_KapeFiles_Targets_AllFileMetadata) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Windows_KapeFiles_Targets_AllFileMetadata(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_KapeFiles_Targets_AllFileMetadata{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Created.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Created,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.SourceFile,
			MetaData:         fmt.Sprintf("Size: %v, LastAccess: %v, LastModified: %v", tmp.Size, tmp.LastAccessed, tmp.Modified),
		}
		outputChannel <- tmp2.StringArray()
	}
}
