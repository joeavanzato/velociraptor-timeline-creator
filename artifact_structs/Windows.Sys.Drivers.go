package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strings"
	"time"
)

type Windows_Sys_Drivers_SignedDrivers struct {
	Caption                 interface{} `json:"Caption"`
	ClassGUID               string      `json:"ClassGuid"`
	CompatID                interface{} `json:"CompatID"`
	CreationClassName       interface{} `json:"CreationClassName"`
	Description             string      `json:"Description"`
	DeviceClass             string      `json:"DeviceClass"`
	DeviceID                string      `json:"DeviceID"`
	DeviceName              string      `json:"DeviceName"`
	DevLoader               interface{} `json:"DevLoader"`
	DriverDate              string      `json:"DriverDate"`
	DriverName              interface{} `json:"DriverName"`
	DriverProviderName      string      `json:"DriverProviderName"`
	DriverVersion           string      `json:"DriverVersion"`
	FriendlyName            string      `json:"FriendlyName"`
	HardWareID              string      `json:"HardWareID"`
	InfName                 string      `json:"InfName"`
	InstallDate             interface{} `json:"InstallDate"`
	IsSigned                bool        `json:"IsSigned"`
	Location                interface{} `json:"Location"`
	Manufacturer            string      `json:"Manufacturer"`
	Name                    interface{} `json:"Name"`
	PDO                     string      `json:"PDO"`
	Signer                  string      `json:"Signer"`
	Started                 interface{} `json:"Started"`
	StartMode               interface{} `json:"StartMode"`
	Status                  interface{} `json:"Status"`
	SystemCreationClassName interface{} `json:"SystemCreationClassName"`
	SystemName              interface{} `json:"SystemName"`
}

func Process_Windows_Sys_Drivers_SignedDrivers(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_Drivers_SignedDrivers{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// time format: 20060621000000.******+***
		dateComponent := strings.Split(tmp.DriverDate, ".")[0]
		parsedTime, terr := time.Parse("20060102150405", dateComponent)
		if terr != nil {
			parsedTime = time.Now()
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.InfName,
			MetaData:         fmt.Sprintf("Description: %v, DeviceClass: %v, DeviceName: %v, ProviderName: %v, Version: %v", tmp.Description, tmp.DeviceClass, tmp.DeviceName, tmp.DriverProviderName, tmp.DriverVersion),
		}
		outputChannel <- tmp2.StringArray()
	}
}
