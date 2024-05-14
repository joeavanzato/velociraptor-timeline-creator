package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Windows_Network_Netstat struct {
	Pid       int       `json:"Pid"`
	Name      string    `json:"Name"`
	Family    string    `json:"Family"`
	Type      string    `json:"Type"`
	Status    string    `json:"Status"`
	LaddrIP   string    `json:"Laddr.IP"`
	LaddrPort int       `json:"Laddr.Port"`
	RaddrIP   string    `json:"Raddr.IP"`
	RaddrPort int       `json:"Raddr.Port"`
	Timestamp time.Time `json:"Timestamp"`
}

func (s Windows_Network_Netstat) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Network_Netstat) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Network_Netstat(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Network_Netstat{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
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
			SourceHost:       fmt.Sprintf("%v:%v", tmp.LaddrIP, tmp.LaddrPort),
			DestinationUser:  "",
			DestinationHost:  fmt.Sprintf("%v:%v", tmp.RaddrIP, tmp.RaddrPort),
			SourceFile:       "",
			MetaData:         fmt.Sprintf("PID: %v, Name: %v, Status: %v", tmp.Pid, tmp.Name, tmp.Status),
		}
		outputChannel <- tmp2.StringArray()
	}
}
