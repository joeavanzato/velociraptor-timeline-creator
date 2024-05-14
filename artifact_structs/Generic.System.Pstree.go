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

type Generic_System_Pstree struct {
	Pid         string    `json:"Pid"`
	Ppid        string    `json:"Ppid"`
	Name        string    `json:"Name"`
	Username    string    `json:"Username"`
	Exe         string    `json:"Exe"`
	CommandLine string    `json:"CommandLine"`
	StartTime   time.Time `json:"StartTime"`
	EndTime     time.Time `json:"EndTime"`
	CallChain   string    `json:"CallChain"`
	PSTree      struct {
		Name      string    `json:"name"`
		ID        string    `json:"id"`
		StartTime time.Time `json:"start_time"`
		Data      struct {
			Pid             string    `json:"Pid"`
			Ppid            string    `json:"Ppid"`
			Name            string    `json:"Name"`
			Threads         int       `json:"Threads"`
			Username        string    `json:"Username"`
			OwnerSid        string    `json:"OwnerSid"`
			CommandLine     string    `json:"CommandLine"`
			Exe             string    `json:"Exe"`
			TokenIsElevated bool      `json:"TokenIsElevated"`
			CreateTime      time.Time `json:"CreateTime"`
			User            float64   `json:"User"`
			System          float64   `json:"System"`
			IoCounters      struct {
				ReadOperationCount  int `json:"ReadOperationCount"`
				WriteOperationCount int `json:"WriteOperationCount"`
				OtherOperationCount int `json:"OtherOperationCount"`
				ReadTransferCount   int `json:"ReadTransferCount"`
				WriteTransferCount  int `json:"WriteTransferCount"`
				OtherTransferCount  int `json:"OtherTransferCount"`
			} `json:"IoCounters"`
			Memory struct {
				PageFaultCount             int `json:"PageFaultCount"`
				PeakWorkingSetSize         int `json:"PeakWorkingSetSize"`
				WorkingSetSize             int `json:"WorkingSetSize"`
				QuotaPeakPagedPoolUsage    int `json:"QuotaPeakPagedPoolUsage"`
				QuotaPagedPoolUsage        int `json:"QuotaPagedPoolUsage"`
				QuotaPeakNonPagedPoolUsage int `json:"QuotaPeakNonPagedPoolUsage"`
				QuotaNonPagedPoolUsage     int `json:"QuotaNonPagedPoolUsage"`
				PagefileUsage              int `json:"PagefileUsage"`
				PeakPagefileUsage          int `json:"PeakPagefileUsage"`
			} `json:"Memory"`
			PebBaseAddress int       `json:"PebBaseAddress"`
			IsWow64        bool      `json:"IsWow64"`
			StartTime      time.Time `json:"StartTime"`
			EndTime        time.Time `json:"EndTime"`
		} `json:"data"`
		Children []struct {
			Name      string    `json:"name"`
			ID        string    `json:"id"`
			StartTime time.Time `json:"start_time"`
			Data      struct {
				Pid             any       `json:"Pid"`
				Ppid            any       `json:"Ppid"`
				Name            string    `json:"Name"`
				Threads         int       `json:"Threads"`
				Username        string    `json:"Username"`
				OwnerSid        string    `json:"OwnerSid"`
				CommandLine     string    `json:"CommandLine"`
				Exe             string    `json:"Exe"`
				TokenIsElevated bool      `json:"TokenIsElevated"`
				CreateTime      time.Time `json:"CreateTime"`
				User            float64   `json:"User"`
				System          float64   `json:"System"`
				IoCounters      struct {
					ReadOperationCount  int `json:"ReadOperationCount"`
					WriteOperationCount int `json:"WriteOperationCount"`
					OtherOperationCount int `json:"OtherOperationCount"`
					ReadTransferCount   int `json:"ReadTransferCount"`
					WriteTransferCount  int `json:"WriteTransferCount"`
					OtherTransferCount  int `json:"OtherTransferCount"`
				} `json:"IoCounters"`
				Memory struct {
					PageFaultCount             int `json:"PageFaultCount"`
					PeakWorkingSetSize         int `json:"PeakWorkingSetSize"`
					WorkingSetSize             int `json:"WorkingSetSize"`
					QuotaPeakPagedPoolUsage    int `json:"QuotaPeakPagedPoolUsage"`
					QuotaPagedPoolUsage        int `json:"QuotaPagedPoolUsage"`
					QuotaPeakNonPagedPoolUsage int `json:"QuotaPeakNonPagedPoolUsage"`
					QuotaNonPagedPoolUsage     int `json:"QuotaNonPagedPoolUsage"`
					PagefileUsage              int `json:"PagefileUsage"`
					PeakPagefileUsage          int `json:"PeakPagefileUsage"`
				} `json:"Memory"`
				PebBaseAddress int64     `json:"PebBaseAddress"`
				IsWow64        bool      `json:"IsWow64"`
				StartTime      time.Time `json:"StartTime"`
				EndTime        time.Time `json:"EndTime"`
			} `json:"data"`
		} `json:"children"`
	} `json:"PSTree"`
}

func (s Generic_System_Pstree) StringArray() []string {
	children := make([]string, 0)
	for _, v := range s.PSTree.Children {
		children = append(children, fmt.Sprintf("Name: %v, ID: %v, StartTime: %v, PID: %v, ProcessName: %v, OwnerSid: %v, Username: %v, CommandLine: %v, Exe: %v, TokenIsElevated: %v, EndTime: %v", v.Name, v.ID, v.StartTime, v.Data.Pid, v.Data.Name, v.Data.OwnerSid, v.Data.Username, v.Data.CommandLine, v.Data.Exe, v.Data.TokenIsElevated, v.Data.EndTime))
	}
	return []string{s.Pid, s.Ppid, s.Name, s.Username, s.Exe, s.CommandLine, s.StartTime.String(), s.EndTime.String(),
		s.PSTree.Name, s.PSTree.ID, s.PSTree.StartTime.String(), s.PSTree.Data.Pid, s.PSTree.Data.Ppid, s.PSTree.Data.Name,
		strconv.Itoa(s.PSTree.Data.Threads), s.PSTree.Data.Username, s.PSTree.Data.OwnerSid, s.PSTree.Data.CommandLine, s.PSTree.Data.Exe,
		strconv.FormatBool(s.PSTree.Data.TokenIsElevated), s.PSTree.Data.CreateTime.String(), fmt.Sprint(s.PSTree.Data.User), fmt.Sprint(s.PSTree.Data.System),
		strconv.Itoa(s.PSTree.Data.IoCounters.ReadOperationCount), strconv.Itoa(s.PSTree.Data.IoCounters.WriteOperationCount), strconv.Itoa(s.PSTree.Data.IoCounters.OtherOperationCount),
		strconv.Itoa(s.PSTree.Data.IoCounters.ReadTransferCount), strconv.Itoa(s.PSTree.Data.IoCounters.WriteTransferCount), strconv.Itoa(s.PSTree.Data.IoCounters.OtherTransferCount),
		strconv.Itoa(s.PSTree.Data.Memory.PageFaultCount), strconv.Itoa(s.PSTree.Data.Memory.PeakWorkingSetSize), strconv.Itoa(s.PSTree.Data.Memory.WorkingSetSize), strconv.Itoa(s.PSTree.Data.Memory.QuotaPeakPagedPoolUsage),
		strconv.Itoa(s.PSTree.Data.Memory.QuotaPagedPoolUsage), strconv.Itoa(s.PSTree.Data.Memory.QuotaPeakNonPagedPoolUsage), strconv.Itoa(s.PSTree.Data.Memory.QuotaNonPagedPoolUsage),
		strconv.Itoa(s.PSTree.Data.Memory.PagefileUsage), strconv.Itoa(s.PSTree.Data.Memory.PeakPagefileUsage), strconv.Itoa(s.PSTree.Data.PebBaseAddress), strconv.FormatBool(s.PSTree.Data.IsWow64),
		s.PSTree.Data.StartTime.String(), s.PSTree.Data.EndTime.String(), fmt.Sprint(children)}
}

func (s Generic_System_Pstree) GetHeaders() []string {
	return []string{"Pid", "Ppid", "Name", "Username", "Exe", "CommandLine", "StartTime", "EndTime", "TreeName", "TreeID", "TreeStartTime", "TreeData_PID", "TreeData_PPID", "TreeData_Name",
		"TreeData_Threads", "TreeData_Username", "TreeData_OwnerSid", "TreeData_CommandLine", "TreeData_Exe", "TreeData_TokenIsElevated", "TreeData_CreateTime", "TreeData_User", "TreeData_System",
		"TreeData_IOCounter_ReadOperationCount", "TreeData_IOCounter_WriteOperationCount", "TreeData_IOCounter_OtherOperationCount", "TreeData_IOCounter_ReadTransferCount", "TreeData_IOCounter_WriteTransferCount", "TreeData_IOCounter_OtherTransferCount",
		"TreeData_Memory_PageFaultCount", "TreeData_Memory_PeakWorkingSetSize", "TreeData_Memory_WorkingSetSize", "TreeData_Memory_QuotaPeakPagedPoolUsage", "TreeData_Memory_QuotaPagedPoolUsage", "TreeData_Memory_QuotaPeakNonPagedPoolUsage", "TreeData_Memory_QuotaNonPagedPoolUsage",
		"TreeData_Memory_PagefileUsage", "TreeData_Memory_PeakPagefileUsage", "TreeData_PebBaseAddress", "TreeData_IsWow64", "TreeData_StartTime", "TreeData_EndTime", "ChildrenProcs"}
}

func Process_Generic_System_Pstree(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_System_Pstree{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.StartTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.StartTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Exe,
			MetaData:         fmt.Sprintf("CommandLine: %v, CallChain: %v", tmp.CommandLine, tmp.CallChain),
		}
		outputChannel <- tmp2.StringArray()
	}
}
