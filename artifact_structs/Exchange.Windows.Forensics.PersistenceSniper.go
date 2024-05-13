package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Exchange_Windows_Forensics_PersistenceSniper struct {
	Hostname        string `json:"Hostname"`
	Technique       string `json:"Technique"`
	Classification  string `json:"Classification"`
	Path            string `json:"Path"`
	Value           string `json:"Value"`
	AccessGained    string `json:"Access Gained"`
	Note            string `json:"Note"`
	Reference       string `json:"Reference"`
	Signature       string `json:"Signature"`
	IsBuiltinBinary string `json:"IsBuiltinBinary"`
	IsLolbin        string `json:"IsLolbin"`
	VTEntries       string `json:"VTEntries"`
}

func (s Exchange_Windows_Forensics_PersistenceSniper) StringArray() []string {
	return []string{s.Hostname, s.Technique, s.Classification, s.Path, s.Value, s.AccessGained, s.Note, s.Reference, s.Signature, s.IsBuiltinBinary, s.IsLolbin, s.VTEntries}
}

func (s Exchange_Windows_Forensics_PersistenceSniper) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Forensics_PersistenceSniper(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_PersistenceSniper{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, tmp.Hostname, tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
