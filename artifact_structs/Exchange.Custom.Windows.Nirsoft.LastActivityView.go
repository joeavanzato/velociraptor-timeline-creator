package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Custom_Windows_Nirsoft_LastActivityView struct {
	ActionTime      string `json:"Action Time"`
	Description     string `json:"Description"`
	Filename        string `json:"Filename"`
	FullPath        string `json:"Full Path"`
	MoreInformation string `json:"More Information"`
	FileExtension   string `json:"File Extension"`
	DataSource      string `json:"Data Source"`
}

func (s Exchange_Custom_Windows_Nirsoft_LastActivityView) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Custom_Windows_Nirsoft_LastActivityView) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Custom_Windows_Nirsoft_LastActivityView(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Custom_Windows_Nirsoft_LastActivityView{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("1/02/2006 15:04:05 PM", tmp.ActionTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(parsedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: tmp.Description,
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.FullPath,
			MetaData:         fmt.Sprintf("MoreInformation: %v, DataSource: %v", tmp.MoreInformation, tmp.DataSource),
		}
		outputChannel <- tmp2.StringArray()
	}
}
