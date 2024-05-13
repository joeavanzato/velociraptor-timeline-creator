package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type Generic_Client_Info_Users struct {
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	LastLogin   time.Time `json:"LastLogin"`
}

func (s Generic_Client_Info_Users) StringArray() []string {
	return []string{s.Name, s.Description, s.LastLogin.String()}
}

func (s Generic_Client_Info_Users) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Client_Info_Users(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Client_Info_Users{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastLogin.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastLogin,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "User Last Login",
			EventDescription: "",
			SourceUser:       tmp.Name,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Description: %v", tmp.Description),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Client_Info_BasicInformation struct {
	Name            string    `json:"Name"`
	BuildTime       time.Time `json:"BuildTime"`
	Version         string    `json:"Version"`
	BuildURL        string    `json:"build_url"`
	InstallTime     int       `json:"install_time"`
	Labels          any       `json:"Labels"`
	Hostname        string    `json:"Hostname"`
	OS              string    `json:"OS"`
	Architecture    string    `json:"Architecture"`
	Platform        string    `json:"Platform"`
	PlatformVersion string    `json:"PlatformVersion"`
	KernelVersion   string    `json:"KernelVersion"`
	Fqdn            string    `json:"Fqdn"`
	MACAddresses    []string  `json:"MACAddresses"`
}

func (s Generic_Client_Info_BasicInformation) StringArray() []string {
	return []string{s.Name, s.BuildTime.String(), s.Version, s.BuildURL, strconv.Itoa(s.InstallTime), fmt.Sprint(s.Labels),
		s.Hostname, s.OS, s.Architecture, s.Platform, s.PlatformVersion, s.KernelVersion, s.Fqdn, fmt.Sprint(s.MACAddresses)}
}

func (s Generic_Client_Info_BasicInformation) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Generic_Client_Info_BasicInformation(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Client_Info_BasicInformation{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.BuildTime.String(), clientIdentifier, tmp.Hostname, tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}

type Generic_Client_Info_WindowsInfo struct {
	ComputerInfo struct {
		DNSHostName         string `json:"DNSHostName"`
		Name                string `json:"Name"`
		Domain              string `json:"Domain"`
		TotalPhysicalMemory string `json:"TotalPhysicalMemory"`
		DomainRole          string `json:"DomainRole"`
	} `json:"Computer Info"`
	NetworkInfo []struct {
		Caption              string `json:"Caption"`
		IPAddresses          string `json:"IPAddresses"`
		IPSubnet             string `json:"IPSubnet"`
		MACAddress           string `json:"MACAddress"`
		DefaultIPGateway     string `json:"DefaultIPGateway"`
		DNSHostName          string `json:"DNSHostName"`
		DNSServerSearchOrder string `json:"DNSServerSearchOrder"`
	} `json:"Network Info"`
}

func (s Generic_Client_Info_WindowsInfo) StringArray() []string {
	networkInfo := make([]string, 0)
	for _, v := range s.NetworkInfo {
		networkInfo = append(networkInfo, fmt.Sprintf("Caption: %v, IPAddresses: %v, IPSubnet: %v, MACAddress: %v, DefaultIPGateway: %v, DNSHostName: %v, DNSServerSearchOrder: %v", v.Caption, v.IPAddresses, v.IPSubnet, v.MACAddress, v.DefaultIPGateway, v.DNSHostName, v.DNSServerSearchOrder))
	}
	return []string{s.ComputerInfo.DNSHostName, s.ComputerInfo.Name, s.ComputerInfo.Domain,
		s.ComputerInfo.TotalPhysicalMemory, s.ComputerInfo.DomainRole, fmt.Sprint(networkInfo)}
}

func (s Generic_Client_Info_WindowsInfo) GetHeaders() []string {
	return []string{"DNSHostName", "Name", "Domain", "TotalPhysicalMemory", "DomainRole", "NetworkInfo"}
}

func Process_Generic_Client_Info_WindowsInfo(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Client_Info_WindowsInfo{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, tmp.ComputerInfo.DNSHostName, tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
