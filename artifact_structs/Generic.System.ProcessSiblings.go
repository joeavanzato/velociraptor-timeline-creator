package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

// TODO - Include Parent Tree Structs

type Generic_System_ProcessSiblings struct {
	Pid         string      `json:"Pid"`
	Ppid        string      `json:"Ppid"`
	Name        string      `json:"Name"`
	ChildPid    string      `json:"ChildPid"`
	CommandLine string      `json:"CommandLine"`
	Username    string      `json:"Username"`
	StartTime   time.Time   `json:"StartTime"`
	EndTime     time.Time   `json:"EndTime"`
	ParentTree  interface{} `json:"ParentTree"`
}

func (s Generic_System_ProcessSiblings) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Generic_System_ProcessSiblings) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_System_ProcessSiblings(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_System_ProcessSiblings{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.StartTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.StartTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("CommandLine: %v, Name: %v, PID: %v", tmp.CommandLine, tmp.Name, tmp.Pid),
		}
		outputChannel <- tmp2.StringArray()

	}
}
