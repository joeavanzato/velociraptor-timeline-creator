package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Applications_DefenderDHParser struct {
	Detection struct {
		GUID                     string `json:"GUID"`
		MagicVersion             string `json:"Magic.Version"`
		Trojan                   string `json:"Trojan"`
		ThreatStatusID           int    `json:"ThreatStatusID"`
		Containerfile            string `json:"containerfile"`
		File                     string `json:"file"`
		ThreatTrackingSha256     string `json:"ThreatTrackingSha256"`
		ThreatTrackingSigSeq     string `json:"ThreatTrackingSigSeq"`
		ThreatTrackingID         string `json:"ThreatTrackingId"`
		ThreatTrackingStartTime  string `json:"ThreatTrackingStartTime"`
		ThreatTrackingThreatName string `json:"ThreatTrackingThreatName"`
		ThreatTrackingSha1       string `json:"ThreatTrackingSha1"`
		ThreatTrackingSigSha     string `json:"ThreatTrackingSigSha"`
		ThreatTrackingSize       int    `json:"ThreatTrackingSize"`
		ThreatTrackingMD5        string `json:"ThreatTrackingMD5"`
		ThreatTrackingScanFlags  string `json:"ThreatTrackingScanFlags"`
		ThreatTrackingIsEsuSig   string `json:"ThreatTrackingIsEsuSig"`
		ThreatTrackingThreatID   int64  `json:"ThreatTrackingThreatId"`
		ThreatTrackingScanSource string `json:"ThreatTrackingScanSource"`
		ThreatTrackingScanType   string `json:"ThreatTrackingScanType"`
		Webfile                  string `json:"webfile"`
		User                     string `json:"User"`
		SpawningProcessName      string `json:"SpawningProcessName"`
		SecurityGroup            string `json:"SecurityGroup"`
	} `json:"Detection"`
	Hostname string `json:"Hostname"`
}

func (s Exchange_Windows_Applications_DefenderDHParser) StringArray() []string {
	base := helpers.GetStructValuesAsStringSlice(s.Detection)
	base = append(base, "Hostname")
	return base
}

func (s Exchange_Windows_Applications_DefenderDHParser) GetHeaders() []string {
	base := helpers.GetStructHeadersAsStringSlice(s.Detection)
	base = append(base, "Hostname")
	return base
}
func Process_Exchange_Windows_Applications_DefenderDHParser(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Applications_DefenderDHParser{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, timeErr := time.Parse("01-02-2006 15:04:05", tmp.Detection.ThreatTrackingStartTime)
		if timeErr != nil {
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
			EventDescription: tmp.Detection.ThreatTrackingThreatName,
			SourceUser:       tmp.Detection.User,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Detection.File,
			MetaData:         fmt.Sprintf("MD5: %v, SpawningProcess: %v, Trojan: %v", tmp.Detection.ThreatTrackingMD5, tmp.Detection.SpawningProcessName, tmp.Detection.Trojan),
		}
		outputChannel <- tmp2.StringArray()
	}
}
