package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Windows_System_Service struct {
	State          string    `json:"State"`
	Name           string    `json:"Name"`
	DisplayName    string    `json:"DisplayName"`
	Status         string    `json:"Status"`
	Pid            int       `json:"Pid"`
	ExitCode       int       `json:"ExitCode"`
	StartMode      string    `json:"StartMode"`
	PathName       string    `json:"PathName"`
	ServiceType    string    `json:"ServiceType"`
	UserAccount    string    `json:"UserAccount"`
	Created        time.Time `json:"Created"`
	ServiceDll     string    `json:"ServiceDll"`
	FailureCommand string    `json:"FailureCommand"`
	FailureActions struct {
		ResetPeriod   int `json:"ResetPeriod"`
		FailureAction []struct {
			Type  string  `json:"Type"`
			Delay float32 `json:"Delay"`
		} `json:"FailureAction"`
	} `json:"FailureActions"`
	AbsoluteExePath    string `json:"AbsoluteExePath"`
	HashServiceExe     string `json:"HashServiceExe"`
	CertinfoServiceExe string `json:"CertinfoServiceExe"`
	HashServiceDll     string `json:"HashServiceDll"`
	CertinfoServiceDll string `json:"CertinfoServiceDll"`
}

func (s Windows_System_Service) StringArray() []string {
	return []string{s.State, s.Name, s.DisplayName, s.Status, strconv.Itoa(s.Pid), strconv.Itoa(s.ExitCode), s.StartMode,
		s.PathName, s.ServiceType, s.UserAccount, s.Created.String(), s.ServiceDll, s.FailureCommand,
		strconv.Itoa(s.FailureActions.ResetPeriod), fmt.Sprint(s.FailureActions.FailureAction),
		s.AbsoluteExePath, s.HashServiceExe, s.CertinfoServiceExe, s.HashServiceDll, s.CertinfoServiceDll}
}

// Headers should match the string array above
func (s Windows_System_Service) GetHeaders() []string {
	return []string{"State", "Name", "DisplayName", "Status", "Pid", "ExitCode", "StartMode", "PathName", "ServiceType",
		"UserAccount", "Created", "ServiceDll", "FailureCommand", "FailureActions_ResetPeriod", "FailureAction",
		"AbsoluteExePath", "HashServiceExe", "CertinfoServiceExe", "HashServiceDll", "CertinfoServiceDll"}
}

func Process_Windows_System_Service(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Service{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Created.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Created,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.UserAccount,
			SourceHost:       clientIdentifier,
			DestinationUser:  tmp.UserAccount,
			DestinationHost:  clientIdentifier,
			SourceFile:       tmp.AbsoluteExePath,
			MetaData:         fmt.Sprintf("Name: %v, State: %v, StartMode: %v, ExePath: %v, ServiceDLL: %v", tmp.Name, tmp.State, tmp.StartMode, tmp.AbsoluteExePath, tmp.ServiceDll),
		}
		outputChannel <- tmp2.StringArray()

	}
}
