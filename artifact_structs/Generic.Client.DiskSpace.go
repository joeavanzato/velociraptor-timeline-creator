package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Generic_Client_DiskSpace struct {
	DeviceID           string `json:"DeviceID"`
	Description        string `json:"Description"`
	VolumeName         string `json:"VolumeName"`
	VolumeSerialNumber string `json:"VolumeSerialNumber"`
	Size               string `json:"Size"`
	FreeSpace          string `json:"FreeSpace"`
	Free               int    `json:"Free%"`
}

func (s Generic_Client_DiskSpace) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Generic_Client_DiskSpace) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Client_DiskSpace(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Client_DiskSpace{}
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
