package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Exchange_Windows_EventLogs_Chainsaw struct {
	EventTime  time.Time `json:"EventTime"`
	Detection  string    `json:"Detection"`
	Severity   string    `json:"Severity"`
	Status     string    `json:"Status"`
	RuleGroup  string    `json:"Rule Group"`
	Computer   string    `json:"Computer"`
	Channel    string    `json:"Channel"`
	EventID    int       `json:"EventID"`
	User       string    `json:"_User"`
	SystemData any       `json:"SystemData"`
	EventData  struct {
		Domain              string `json:"Domain"`
		ProductName         string `json:"Product Name"`
		ProductVersion      string `json:"Product Version"`
		SID                 string `json:"SID"`
		ScanID              string `json:"Scan ID"`
		ScanParameters      string `json:"Scan Parameters"`
		ScanParametersIndex string `json:"Scan Parameters Index"`
		ScanType            string `json:"Scan Type"`
		ScanTypeIndex       string `json:"Scan Type Index"`
		User                string `json:"User"`
	} `json:"EventData"`
	Authors []string `json:"Authors"`
}

func (s Exchange_Windows_EventLogs_Chainsaw) StringArray() []string {
	return []string{s.EventTime.String(), s.Detection, s.Severity, s.Status, s.RuleGroup, s.Computer, s.Channel, strconv.Itoa(s.EventID),
		s.User, fmt.Sprint(s.SystemData), s.EventData.Domain, s.EventData.ProductName, s.EventData.ProductVersion, s.EventData.SID, s.EventData.ScanID, s.EventData.ScanParameters, s.EventData.ScanParametersIndex,
		s.EventData.ScanType, s.EventData.ScanTypeIndex, s.EventData.User, fmt.Sprint(s.Authors)}
}

func (s Exchange_Windows_EventLogs_Chainsaw) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_Chainsaw(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_Chainsaw{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: fmt.Sprintf("%v", tmp.Detection),
			SourceUser:       tmp.User,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Channel,
			MetaData:         fmt.Sprintf("Severity: %v, Rule Group: %v", tmp.Severity, tmp.RuleGroup),
		}
		outputChannel <- tmp2.StringArray()
	}
}
