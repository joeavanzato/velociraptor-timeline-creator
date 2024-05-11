package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"strings"
	"time"
)

type Windows_Sys_Drivers_SignedDrivers struct {
	Caption                 string `json:"Caption"`
	ClassGUID               string `json:"ClassGuid"`
	CompatID                string `json:"CompatID"`
	CreationClassName       string `json:"CreationClassName"`
	Description             string `json:"Description"`
	DeviceClass             string `json:"DeviceClass"`
	DeviceID                string `json:"DeviceID"`
	DeviceName              string `json:"DeviceName"`
	DevLoader               string `json:"DevLoader"`
	DriverDate              string `json:"DriverDate"`
	DriverName              string `json:"DriverName"`
	DriverProviderName      string `json:"DriverProviderName"`
	DriverVersion           string `json:"DriverVersion"`
	FriendlyName            string `json:"FriendlyName"`
	HardWareID              string `json:"HardWareID"`
	InfName                 string `json:"InfName"`
	InstallDate             string `json:"InstallDate"`
	IsSigned                bool   `json:"IsSigned"`
	Location                string `json:"Location"`
	Manufacturer            string `json:"Manufacturer"`
	Name                    string `json:"Name"`
	PDO                     string `json:"PDO"`
	Signer                  string `json:"Signer"`
	Started                 string `json:"Started"`
	StartMode               string `json:"StartMode"`
	Status                  string `json:"Status"`
	SystemCreationClassName string `json:"SystemCreationClassName"`
	SystemName              string `json:"SystemName"`
}

func (s Windows_Sys_Drivers_SignedDrivers) StringArray() []string {
	dateComponent := strings.Split(s.DriverDate, ".")[0]
	parsedTime, terr := time.Parse("20060102150405", dateComponent)
	if terr != nil {
		parsedTime = time.Now()
	}
	return []string{s.Caption, s.ClassGUID, s.CompatID, s.CreationClassName, s.Description, s.DeviceClass, s.DeviceID,
		s.DeviceName, s.DevLoader, parsedTime.String(), s.DriverName, s.DriverProviderName, s.DriverVersion, s.FriendlyName,
		s.HardWareID, s.InfName, s.InstallDate, strconv.FormatBool(s.IsSigned), s.Location, s.Manufacturer, s.Name, s.PDO, s.Signer, s.Started,
		s.StartMode, s.Status, s.SystemCreationClassName, s.SystemName}
}

// Headers should match the array above
func (s Windows_Sys_Drivers_SignedDrivers) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
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
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(parsedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
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

type Windows_Sys_Drivers_RunningDrivers struct {
	AcceptPause             bool   `json:"AcceptPause"`
	AcceptStop              bool   `json:"AcceptStop"`
	Caption                 string `json:"Caption"`
	CreationClassName       string `json:"CreationClassName"`
	Description             string `json:"Description"`
	DesktopInteract         bool   `json:"DesktopInteract"`
	DisplayName             string `json:"DisplayName"`
	ErrorControl            string `json:"ErrorControl"`
	ExitCode                int    `json:"ExitCode"`
	InstallDate             string `json:"InstallDate"`
	Name                    string `json:"Name"`
	PathName                string `json:"PathName"`
	ServiceSpecificExitCode int    `json:"ServiceSpecificExitCode"`
	ServiceType             string `json:"ServiceType"`
	Started                 bool   `json:"Started"`
	StartMode               string `json:"StartMode"`
	StartName               string `json:"StartName"`
	State                   string `json:"State"`
	Status                  string `json:"Status"`
	SystemCreationClassName string `json:"SystemCreationClassName"`
	SystemName              string `json:"SystemName"`
	TagID                   int    `json:"TagId"`
	Authenticode            struct {
		Filename      string `json:"Filename"`
		ProgramName   string `json:"ProgramName"`
		PublisherLink string `json:"PublisherLink"`
		MoreInfoLink  string `json:"MoreInfoLink"`
		SerialNumber  string `json:"SerialNumber"`
		IssuerName    string `json:"IssuerName"`
		SubjectName   string `json:"SubjectName"`
		Timestamp     string `json:"Timestamp"`
		Trusted       string `json:"Trusted"`
		ExtraInfo     struct {
			Catalog string `json:"Catalog"`
		} `json:"_ExtraInfo"`
	} `json:"Authenticode"`
	Hashes struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hashes"`
}

func (s Windows_Sys_Drivers_RunningDrivers) StringArray() []string {
	dateComponent := strings.Split(s.InstallDate, ".")[0]
	parsedTime, terr := time.Parse("20060102150405", dateComponent)
	if terr != nil {
		parsedTime = time.Now()
	}
	return []string{strconv.FormatBool(s.AcceptPause), strconv.FormatBool(s.AcceptStop), s.Caption, s.CreationClassName,
		s.Description, strconv.FormatBool(s.DesktopInteract), s.DisplayName, s.ErrorControl, strconv.Itoa(s.ExitCode),
		parsedTime.String(), s.Name, s.PathName, strconv.Itoa(s.ServiceSpecificExitCode), s.ServiceType, strconv.FormatBool(s.Started),
		s.StartMode, s.StartName, s.State, s.Status, s.SystemCreationClassName, s.SystemName, strconv.Itoa(s.TagID),
		s.Authenticode.Filename, s.Authenticode.ProgramName, s.Authenticode.PublisherLink, s.Authenticode.MoreInfoLink,
		s.Authenticode.SerialNumber, s.Authenticode.IssuerName, s.Authenticode.SubjectName, s.Authenticode.Timestamp,
		s.Authenticode.Trusted, s.Authenticode.ExtraInfo.Catalog, s.Hashes.MD5, s.Hashes.SHA1, s.Hashes.SHA256}
}

// Headers should match the array above
func (s Windows_Sys_Drivers_RunningDrivers) GetHeaders() []string {
	//return helpers.GetStructAsStringSlice(s)
	return []string{"AcceptPause", "AcceptStop", "Caption", "CreationClassName", "Description", "DesktopInteract",
		"DisplayName", "ErrorControl", "ExitCode", "InstallDate", "Name", "PathName", "ServiceSpecificExitCode",
		"ServiceType", "Started", "StartMode", "StartName", "State", "Status", "SystemCreationClassName", "SystemName",
		"TagID", "Authenticode_Filename", "Authenticode_ProgramName", "Authenticode_PublisherLink", "Authenticode_MoreInfoLink",
		"Authenticode_SerialNumber", "Authenticode_IssuerName", "Authenticode_SubjectName", "Authenticode_Timestamp",
		"Authenticode_Trusted", "Authenticode_ExtraInfo_Catalog", "Hashes_MD5", "Hashes_SHA1", "Hashes_SHA256"}
}

func Process_Windows_Sys_Drivers_RunningDrivers(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_Drivers_RunningDrivers{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// time format: 20060621000000.******+***
		dateComponent := strings.Split(tmp.InstallDate, ".")[0]
		parsedTime, terr := time.Parse("20060102150405", dateComponent)
		if terr != nil {
			parsedTime = time.Now()
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(parsedTime.String(), clientIdentifier, tmp.SystemName, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Running Driver Installed",
			EventDescription: "Driver Install (Date Often Missing)",
			SourceUser:       "",
			SourceHost:       tmp.SystemName,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.PathName,
			MetaData:         fmt.Sprintf("Description: %v, MD5: %v, DeviceName: %v, Started: %v, DesktopInteract: %v", tmp.Description, tmp.Hashes.MD5, tmp.Started, tmp.DesktopInteract),
		}
		outputChannel <- tmp2.StringArray()
	}
}
