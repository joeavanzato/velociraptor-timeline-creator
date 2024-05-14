package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_System_AuditPolicy struct {
	MachineName      string `json:"Machine Name"`
	PolicyTarget     string `json:"Policy Target"`
	Subcategory      string `json:"Subcategory"`
	SubcategoryGUID  string `json:"Subcategory GUID"`
	InclusionSetting string `json:"Inclusion Setting"`
	ExclusionSetting string `json:"Exclusion Setting"`
}

func (s Windows_System_AuditPolicy) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_System_AuditPolicy) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_System_AuditPolicy(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_AuditPolicy{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

	}
}
