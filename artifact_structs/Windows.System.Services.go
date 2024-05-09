package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_System_Service struct {
	State              string    `json:"State"`
	Name               string    `json:"Name"`
	DisplayName        string    `json:"DisplayName"`
	Status             string    `json:"Status"`
	Pid                int       `json:"Pid"`
	ExitCode           int       `json:"ExitCode"`
	StartMode          string    `json:"StartMode"`
	PathName           string    `json:"PathName"`
	ServiceType        string    `json:"ServiceType"`
	UserAccount        string    `json:"UserAccount"`
	Created            time.Time `json:"Created"`
	ServiceDll         string    `json:"ServiceDll"`
	FailureCommand     string    `json:"FailureCommand"`
	FailureActions     any       `json:"FailureActions"`
	AbsoluteExePath    string    `json:"AbsoluteExePath"`
	HashServiceExe     string    `json:"HashServiceExe"`
	CertinfoServiceExe string    `json:"CertinfoServiceExe"`
	HashServiceDll     string    `json:"HashServiceDll"`
	CertinfoServiceDll string    `json:"CertinfoServiceDll"`
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
