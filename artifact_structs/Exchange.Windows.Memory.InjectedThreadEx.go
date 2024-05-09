package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
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
