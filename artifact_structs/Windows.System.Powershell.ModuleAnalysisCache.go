package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_System_Powershell_ModuleAnalysisCache struct {
	OSPath     string    `json:"OSPath"`
	ModuleName string    `json:"ModuleName"`
	Timestamp  time.Time `json:"Timestamp"`
	Functions  []string  `json:"Functions"`
}

func (s Windows_System_Powershell_ModuleAnalysisCache) StringArray() []string {
	return []string{s.OSPath, s.ModuleName, s.Timestamp.String(), fmt.Sprint(s.Functions)}
}

// Headers should match the array above
func (s Windows_System_Powershell_ModuleAnalysisCache) GetHeaders() []string {
	return []string{"OSPath", "ModuleName", "Timestamp", "Functions"}
}

func Process_Windows_System_Powershell_ModuleAnalysisCache(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Powershell_ModuleAnalysisCache{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.ModuleName,
			MetaData:         fmt.Sprintf("Functions: %v", tmp.Functions),
		}
		outputChannel <- tmp2.StringArray()
	}
}
