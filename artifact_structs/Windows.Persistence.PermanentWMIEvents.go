package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
	"strconv"
)

type Windows_Persistence_PermanentWMIEvents struct {
	ConsumerDetails struct {
		Category                 int      `json:"Category"`
		CreatorSID               []int    `json:"CreatorSID"`
		EventID                  int      `json:"EventID"`
		EventType                int      `json:"EventType"`
		InsertionStringTemplates []string `json:"InsertionStringTemplates"`
		MachineName              any      `json:"MachineName"`
		MaximumQueueSize         any      `json:"MaximumQueueSize"`
		Name                     string   `json:"Name"`
		NameOfRawDataProperty    any      `json:"NameOfRawDataProperty"`
		NameOfUserSIDProperty    string   `json:"NameOfUserSIDProperty"`
		NumberOfInsertionStrings int      `json:"NumberOfInsertionStrings"`
		SourceName               string   `json:"SourceName"`
		UNCServerName            any      `json:"UNCServerName"`
	} `json:"ConsumerDetails"`
	FilterDetails struct {
		CreatorSID     []int  `json:"CreatorSID"`
		EventAccess    any    `json:"EventAccess"`
		EventNamespace string `json:"EventNamespace"`
		Name           string `json:"Name"`
		Query          string `json:"Query"`
		QueryLanguage  string `json:"QueryLanguage"`
	} `json:"FilterDetails"`
	Namespace string `json:"Namespace"`
}

func (s Windows_Persistence_PermanentWMIEvents) StringArray() []string {
	return []string{strconv.Itoa(s.ConsumerDetails.Category), fmt.Sprint(s.ConsumerDetails.CreatorSID), strconv.Itoa(s.ConsumerDetails.EventID), strconv.Itoa(s.ConsumerDetails.EventType), fmt.Sprint(s.ConsumerDetails.InsertionStringTemplates),
		fmt.Sprint(s.ConsumerDetails.MachineName), fmt.Sprint(s.ConsumerDetails.MaximumQueueSize), s.ConsumerDetails.Name, fmt.Sprint(s.ConsumerDetails.NameOfRawDataProperty), s.ConsumerDetails.NameOfUserSIDProperty,
		strconv.Itoa(s.ConsumerDetails.NumberOfInsertionStrings), s.ConsumerDetails.SourceName, fmt.Sprint(s.ConsumerDetails.UNCServerName), fmt.Sprint(s.FilterDetails.CreatorSID), fmt.Sprint(s.FilterDetails.EventAccess),
		s.FilterDetails.EventNamespace, s.FilterDetails.Name, s.FilterDetails.Query, s.FilterDetails.QueryLanguage, s.Namespace}
}

func (s Windows_Persistence_PermanentWMIEvents) GetHeaders() []string {
	return []string{"Consumer_Category", "Consumer_CreatorSID", "Consumer_EventID", "Consumer_EventType", "Consumer_InsertionStringTemplates", "Consumer_MachineName", "Consumer_MaximumQueueSize", "Consumer_Name",
		"Consumer_NameOfRawDataProperty", "Consumer_NameOfUserSIDProperty", "Consumer_SourceName", "Consumer_UNCServerName", "Filter_CreatorSID", "Filter_EventAccess", "Filter_EventNamespace", "Filter_Name", "Filter_Query",
		"Filter_QueryLanguage", "Namespace"}
}

func Process_Windows_Persistence_PermanentWMIEvents(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Persistence_PermanentWMIEvents{}
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
