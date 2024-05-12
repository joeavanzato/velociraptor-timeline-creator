package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
)

type Exchange_Windows_Forensics_Trawler struct {
	Name      string `json:"Name"`
	Risk      string `json:"Risk"`
	Source    string `json:"Source"`
	Technique string `json:"Technique"`
	Meta      string `json:"Meta"`
	Upload    struct {
		Path       string   `json:"Path"`
		Size       int      `json:"Size"`
		StoredSize int      `json:"StoredSize"`
		Sha256     string   `json:"sha256"`
		Md5        string   `json:"md5"`
		StoredName string   `json:"StoredName"`
		Components []string `json:"Components"`
		Accessor   string   `json:"Accessor"`
	} `json:"Upload"`
}

func (s Exchange_Windows_Forensics_Trawler) StringArray() []string {
	return []string{s.Name, s.Risk, s.Source, s.Technique, s.Meta}
}

func (s Exchange_Windows_Forensics_Trawler) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Forensics_Trawler(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_Trawler{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
