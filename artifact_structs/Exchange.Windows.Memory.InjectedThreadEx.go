package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Exchange_Windows_Memory_InjectedThreadEx struct {
	ProcessID                     string `json:"ProcessId"`
	Wow64                         string `json:"Wow64"`
	PathMismatch                  string `json:"PathMismatch"`
	ProcessIntegrity              string `json:"ProcessIntegrity"`
	ProcessLogonID                string `json:"ProcessLogonId"`
	ProcessSecurityIdentifier     string `json:"ProcessSecurityIdentifier"`
	ProcessUserName               string `json:"ProcessUserName"`
	ProcessLogonType              string `json:"ProcessLogonType"`
	ProcessAuthenticationPackage  string `json:"ProcessAuthenticationPackage"`
	ThreadID                      string `json:"ThreadId"`
	ThreadStartTime               string `json:"ThreadStartTime"`
	BasePriority                  string `json:"BasePriority"`
	WaitReason                    string `json:"WaitReason"`
	IsUniqueThreadToken           string `json:"IsUniqueThreadToken"`
	AllocatedMemoryProtection     string `json:"AllocatedMemoryProtection"`
	MemoryProtection              string `json:"MemoryProtection"`
	MemoryState                   string `json:"MemoryState"`
	MemoryType                    string `json:"MemoryType"`
	Win32StartAddress             string `json:"Win32StartAddress"`
	Win32StartAddressModuleSigned string `json:"Win32StartAddressModuleSigned"`
	Win32StartAddressPrivate      string `json:"Win32StartAddressPrivate"`
	Size                          string `json:"Size"`
	TailBytes                     string `json:"TailBytes"`
	StartBytes                    string `json:"StartBytes"`
	Detections                    string `json:"Detections"`
}

func (s Exchange_Windows_Memory_InjectedThreadEx) StringArray() []string {
	parsedTime, terr := time.Parse("1/2/2006 03:04:05 PM", s.ThreadStartTime)
	if terr != nil {
		parsedTime = time.Now()
	}

	return []string{s.ProcessID, s.Wow64, s.PathMismatch, s.ProcessIntegrity, s.ProcessLogonID, s.ProcessSecurityIdentifier, s.ProcessUserName,
		s.ProcessLogonType, s.ProcessAuthenticationPackage, s.ThreadID, parsedTime.String(), s.BasePriority, s.WaitReason, s.IsUniqueThreadToken, s.AllocatedMemoryProtection,
		s.MemoryProtection, s.MemoryState, s.MemoryType, s.Win32StartAddress, s.Win32StartAddressModuleSigned, s.Win32StartAddressPrivate, s.Size, s.TailBytes, s.StartBytes, s.Detections}
}

func (s Exchange_Windows_Memory_InjectedThreadEx) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_Memory_InjectedThreadEx(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Memory_InjectedThreadEx{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// TODO - Maybe pull description for the most common ones or use zimmerman maps?
		// time format: 5/5/2024 10:08:41 PM
		parsedTime, terr := time.Parse("1/2/2006 03:04:05 PM", tmp.ThreadStartTime)
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
			SourceUser:       tmp.ProcessUserName,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("ProcessID: %v, Process SID: %v, ProcessLogonType: %v, Detections: %v, Path Mismatch: %v", tmp.ProcessID, tmp.ProcessSecurityIdentifier, tmp.ProcessLogonType, tmp.Detections, tmp.PathMismatch),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_Memory_InjectedThreadEx_RawResults struct {
	Stdout       string `json:"Stdout"`
	Stderr       string `json:"Stderr"`
	ReturnCode   int    `json:"ReturnCode"`
	Complete     bool   `json:"Complete"`
	ScanSettings struct {
		ScanType  string `json:"ScanType"`
		PidTarget string `json:"PidTarget"`
	} `json:"ScanSettings"`
}

func (s Exchange_Windows_Memory_InjectedThreadEx_RawResults) StringArray() []string {
	return []string{s.Stdout, s.Stderr, strconv.Itoa(s.ReturnCode), strconv.FormatBool(s.Complete), s.ScanSettings.ScanType, s.ScanSettings.PidTarget}
}

func (s Exchange_Windows_Memory_InjectedThreadEx_RawResults) GetHeaders() []string {
	return []string{"Stdout", "Stderr", "ReturnCode", "Complete", "ScanSettings_ScanType", "ScanSettings_PidTarget"}
}

func Process_Exchange_Windows_Memory_InjectedThreadEx_RawResults(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	// TODO There is some type of error here - I think due to the long lines maybe or weird whitespacing
	for _, line := range inputLines {
		tmp := Exchange_Windows_Memory_InjectedThreadEx_RawResults{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue

	}
}
