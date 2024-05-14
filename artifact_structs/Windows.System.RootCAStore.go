package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_System_RootCAStore struct {
	RegistryValue string    `json:"_RegistryValue"`
	ModTime       time.Time `json:"ModTime"`
	Name          any       `json:"Name"`
	FingerPrint   string    `json:"FingerPrint"`
	Certificate   struct {
		SerialNumber       string    `json:"SerialNumber"`
		SignatureAlgorithm string    `json:"SignatureAlgorithm"`
		Subject            string    `json:"Subject"`
		Issuer             string    `json:"Issuer"`
		NotBefore          time.Time `json:"NotBefore"`
		NotAfter           time.Time `json:"NotAfter"`
		PublicKey          string    `json:"PublicKey"`
		Extensions         any       `json:"Extensions"`
	} `json:"Certificate"`
}

func (s Windows_System_RootCAStore) StringArray() []string {
	base := helpers.GetStructValuesAsStringSlice(s)
	base = base[:len(base)-1]
	base = append(base, helpers.GetStructValuesAsStringSlice(s.Certificate)...)
	return base

}

func (s Windows_System_RootCAStore) GetHeaders() []string {
	base := helpers.GetStructHeadersAsStringSlice(s)
	base = base[:len(base)-1]
	base = append(base, helpers.GetStructHeadersAsStringSlice(s.Certificate)...)
	return base
}

func Process_Windows_System_RootCAStore(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_RootCAStore{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.ModTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.ModTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Name: %v, Fingerprint: %v, Certificate_Subject: %v, Certificate_Issuer: %v", tmp.Name, tmp.FingerPrint, tmp.Certificate.Subject, tmp.Certificate.Issuer),
		}
		outputChannel <- tmp2.StringArray()

	}
}
