package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_EventLogs_RDPClientActivity struct {
	Start                   time.Time `json:"Start"`
	End                     time.Time `json:"End"`
	Duration                float64   `json:"Duration"`
	SourceUserSID           string    `json:"SourceUserSID"`
	SourceUser              string    `json:"SourceUser"`
	SourceHost              string    `json:"SourceHost"`
	DestinationHost         string    `json:"DestinationHost"`
	ConnectedDomain         string    `json:"ConnectedDomain"`
	DestinationUsernameHash string    `json:"DestinationUsernameHash"`
	DisconnectReasonID      string    `json:"DisconnectReasonID"`
	DisconnectReason        string    `json:"DisconnectReason"`
}

func (s Exchange_Windows_EventLogs_RDPClientActivity) StringArray() []string {
	return []string{s.Start.String(), s.End.String(), fmt.Sprint(s.Duration), s.SourceUserSID, s.SourceUser, s.SourceHost,
		s.DestinationHost, s.ConnectedDomain, s.DestinationUsernameHash, s.DisconnectReasonID, s.DisconnectReason}
}

func (s Exchange_Windows_EventLogs_RDPClientActivity) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_RDPClientActivity(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_RDPClientActivity{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Start.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Start,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.SourceUser,
			SourceHost:       tmp.SourceHost,
			DestinationUser:  tmp.DestinationUsernameHash,
			DestinationHost:  tmp.DestinationHost,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Duration: %v, ConnectedDomain: %v, DisconnectReason: %v", tmp.Duration, tmp.ConnectedDomain, tmp.DisconnectReason),
		}
		outputChannel <- tmp2.StringArray()
	}
}
