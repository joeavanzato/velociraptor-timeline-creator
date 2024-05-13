package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Detection_Mutants_Handles struct {
	ProcPid  int    `json:"ProcPid"`
	ProcName string `json:"ProcName"`
	Exe      string `json:"Exe"`
	Type     string `json:"Type"`
	Name     string `json:"Name"`
	Handle   int    `json:"Handle"`
}

func (s Windows_Detection_Mutants_Handles) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Detection_Mutants_Handles) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Detection_Mutants_Handles(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Detection_Mutants_Handles{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}

type Windows_Detection_Mutants_ObjectTree struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
}

func (s Windows_Detection_Mutants_ObjectTree) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Detection_Mutants_ObjectTree) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Detection_Mutants_ObjectTree(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Detection_Mutants_ObjectTree{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
