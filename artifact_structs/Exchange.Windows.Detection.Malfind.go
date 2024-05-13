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

type Exchange_Windows_Detection_Malfind struct {
	CreateTime   time.Time `json:"CreateTime"`
	Pid          int       `json:"Pid"`
	Name         string    `json:"Name"`
	Address      int64     `json:"_Address"`
	AddressRange string    `json:"AddressRange"`
	Protection   string    `json:"Protection"`
	SectionSize  int       `json:"SectionSize"`
	HexHeader    string    `json:"HexHeader"`
	DataMagic    string    `json:"DataMagic"`
	SectionData  string    `json:"SectionData"`
	YaraHit      []struct {
		Rule string `json:"Rule"`
		Meta struct {
			Author      string `json:"author"`
			Email       string `json:"email"`
			License     string `json:"license"`
			Copyright   string `json:"copyright"`
			Description string `json:"description"`
		} `json:"Meta"`
		Tags   any `json:"Tags"`
		String struct {
			Name    string   `json:"Name"`
			Offset  int      `json:"Offset"`
			HexData []string `json:"HexData"`
			Data    string   `json:"Data"`
		} `json:"String"`
	} `json:"YaraHit"`
	PathSpec     string `json:"_PathSpec"`
	ProcessChain []struct {
		Pid             int       `json:"Pid"`
		Ppid            int       `json:"Ppid"`
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
			ReadOperationCount  int   `json:"ReadOperationCount"`
			WriteOperationCount int   `json:"WriteOperationCount"`
			OtherOperationCount int   `json:"OtherOperationCount"`
			ReadTransferCount   int64 `json:"ReadTransferCount"`
			WriteTransferCount  int64 `json:"WriteTransferCount"`
			OtherTransferCount  int   `json:"OtherTransferCount"`
		} `json:"IoCounters"`
		Memory struct {
			PageFaultCount             int   `json:"PageFaultCount"`
			PeakWorkingSetSize         int64 `json:"PeakWorkingSetSize"`
			WorkingSetSize             int64 `json:"WorkingSetSize"`
			QuotaPeakPagedPoolUsage    int   `json:"QuotaPeakPagedPoolUsage"`
			QuotaPagedPoolUsage        int   `json:"QuotaPagedPoolUsage"`
			QuotaPeakNonPagedPoolUsage int   `json:"QuotaPeakNonPagedPoolUsage"`
			QuotaNonPagedPoolUsage     int   `json:"QuotaNonPagedPoolUsage"`
			PagefileUsage              int64 `json:"PagefileUsage"`
			PeakPagefileUsage          int64 `json:"PeakPagefileUsage"`
		} `json:"Memory"`
		PebBaseAddress int64 `json:"PebBaseAddress"`
		IsWow64        bool  `json:"IsWow64"`
	} `json:"ProcessChain"`
}

func (s Exchange_Windows_Detection_Malfind) StringArray() []string {
	yarahits := make([]string, 0)
	for _, v := range s.YaraHit {
		yarahits = append(yarahits, fmt.Sprintf("Rule: %v, Author: %v, Email: %v, License: %v, Copyright: %v, Description: %v, Tags: %v, Name: %v, OFfset: %v, HexData: %v, Data: %v", v.Rule, v.Meta.Author, v.Meta.Email, v.Meta.License, v.Meta.Copyright, v.Meta.Description, v.Tags, v.String.Name, v.String.Offset, v.String.HexData, v.String.Data))
	}
	processchains := make([]string, 0)
	for _, v := range s.ProcessChain {
		processchains = append(processchains, fmt.Sprintf("Pid: %v, Ppid: %v, Name: %v, Threads: %v, Username: %v, OwnerSid: %v, CommandLine: %v, Exe: %v, TokenIsElevated: %v, CreateTime: %v, User: %v, System: %v, ReadOperationCount: %v, WriteOperationCount: %v, OtherOperationCount: %v, ReadTransferCount: %v, WriteTransferCount: %v, OtherTransferCount: %v, PageFaultCount: %v, PeakWorkingSetSize: %v, WorkingSetSize: %v, QuotaPeakPagedPoolUsage: %v, QuotaPeakNonPagedPoolUsage: %v, QuotaNonPagedPoolUsage: %v, PagefileUsage: %v, PeakPagefileUsage: %v, PebBaseAddress: %v, IsWow64 :%v", v.Pid, v.Ppid, v.Name, v.Threads, v.Username, v.OwnerSid, v.Exe, v.TokenIsElevated, v.CreateTime, v.User, v.System, v.IoCounters.ReadOperationCount, v.IoCounters.WriteOperationCount, v.IoCounters.OtherOperationCount, v.IoCounters.ReadTransferCount, v.IoCounters.WriteTransferCount, v.IoCounters.OtherTransferCount, v.Memory.PageFaultCount, v.Memory.PeakWorkingSetSize, v.Memory.WorkingSetSize, v.Memory.QuotaPeakPagedPoolUsage, v.Memory.QuotaPagedPoolUsage, v.Memory.QuotaPeakNonPagedPoolUsage, v.Memory.QuotaNonPagedPoolUsage, v.Memory.PagefileUsage, v.Memory.PeakPagefileUsage, v.PebBaseAddress, v.IsWow64))
	}
	return []string{s.CreateTime.String(), strconv.Itoa(s.Pid), s.Name, strconv.FormatInt(s.Address, 10), s.AddressRange, s.Protection, strconv.Itoa(s.SectionSize), s.HexHeader,
		s.DataMagic, s.SectionData, fmt.Sprint(yarahits), s.PathSpec, fmt.Sprint(processchains)}
}

func (s Exchange_Windows_Detection_Malfind) GetHeaders() []string {
	return []string{"CreateTime", "Pid", "Name", "Address", "AddressRange", "Protection", "SectionSize", "HexHeader", "DataMagic", "SectionData",
		"YaraHits", "PathSpec", "ProcessChain"}
}

func Process_Exchange_Windows_Detection_Malfind(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Detection_Malfind{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.CreateTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.CreateTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "File Created",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Name,
			MetaData:         fmt.Sprintf("YARA: %v", tmp.YaraHit),
		}
		outputChannel <- tmp2.StringArray()
	}
}
