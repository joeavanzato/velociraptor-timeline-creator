package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_Sys_DiskInfo struct {
	Partitions   int    `json:"Partitions"`
	DiskIndex    int    `json:"DiskIndex"`
	Type         string `json:"Type"`
	PNPDeviceID  string `json:"PNPDeviceID"`
	DeviceID     string `json:"DeviceID"`
	Size         string `json:"Size"`
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"Model"`
	Name         string `json:"Name"`
	SerialNumber string `json:"SerialNumber"`
	Description  string `json:"Description"`
}

func (s Windows_Sys_DiskInfo) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Windows_Sys_DiskInfo) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Sys_DiskInfo(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_DiskInfo{}
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
