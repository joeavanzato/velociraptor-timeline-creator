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

type Exchange_Windows_System_Powershell_ISEAutoSave struct {
	OSPath  string    `json:"OSPath"`
	Size    int       `json:"Size"`
	Mtime   time.Time `json:"Mtime"`
	Btime   time.Time `json:"Btime"`
	Ctime   time.Time `json:"Ctime"`
	Atime   time.Time `json:"Atime"`
	Content string    `json:"Content"`
}

func (s Exchange_Windows_System_Powershell_ISEAutoSave) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_System_Powershell_ISEAutoSave) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_System_Powershell_ISEAutoSave(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_System_Powershell_ISEAutoSave{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "ISEAutoSave Modified",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Created: %v, Content: %v", tmp.Ctime, tmp.Content),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig struct {
	OSPath        string    `json:"OSPath"`
	Size          int       `json:"Size"`
	Mtime         time.Time `json:"Mtime"`
	Btime         time.Time `json:"Btime"`
	Ctime         time.Time `json:"Ctime"`
	Atime         time.Time `json:"Atime"`
	MRU           any       `json:"MRU"`
	Configuration struct {
		ConfigSections struct {
			SectionGroup struct {
				Attrname string `json:"Attrname"`
				Attrtype string `json:"Attrtype"`
				Section  struct {
					Attrname               string `json:"Attrname"`
					Attrtype               string `json:"Attrtype"`
					AttrallowExeDefinition string `json:"AttrallowExeDefinition"`
					AttrrequirePermission  string `json:"AttrrequirePermission"`
				} `json:"section"`
			} `json:"sectionGroup"`
		} `json:"configSections"`
		UserSettings struct {
			UserSettings struct {
				Setting []struct {
					Attrname        string `json:"Attrname"`
					AttrserializeAs string `json:"AttrserializeAs"`
					Value           any    `json:"value"`
				} `json:"setting"`
			} `json:"UserSettings"`
		} `json:"userSettings"`
	} `json:"Configuration"`
	RawXML string `json:"RawXml"`
}

func (s Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig) StringArray() []string {
	settings := make([]string, 0)
	for _, v := range s.Configuration.UserSettings.UserSettings.Setting {
		settings = append(settings, fmt.Sprintf("Attrname: %v, AttrserializeAs: %v, Value: %v", v.Attrname, v.AttrserializeAs, v.Value))
	}
	return []string{s.OSPath, strconv.Itoa(s.Size), s.Mtime.String(), s.Btime.String(), s.Ctime.String(), s.Atime.String(),
		fmt.Sprint(s.MRU), s.Configuration.ConfigSections.SectionGroup.Attrname, s.Configuration.ConfigSections.SectionGroup.Attrtype,
		s.Configuration.ConfigSections.SectionGroup.Section.Attrname, s.Configuration.ConfigSections.SectionGroup.Section.Attrtype,
		s.Configuration.ConfigSections.SectionGroup.Section.AttrallowExeDefinition, s.Configuration.ConfigSections.SectionGroup.Section.AttrrequirePermission,
		fmt.Sprint(settings)}
}

func (s Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig) GetHeaders() []string {
	return []string{"OSPath", "Size", "Mtime", "Btime", "Ctime", "Atime", "MRU", "Configuration_SectionGroup_Attrname", "Configuration_SectionGroup_Attrtype", "Configuration_SectionGroup_Section_Attrname",
		"Configuration_SectionGroup_Section_Attrtype", "Configuration_SectionGroup_Section_AttrallowExeDefinition", "Configuration_SectionGroup_Section_AttrrequirePermission",
		"UserSettings"}
}

func Process_Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "ISE UserConfig Modified",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Created: %v, Content: %v", tmp.Ctime),
		}
		outputChannel <- tmp2.StringArray()
	}
}
