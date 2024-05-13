package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Exchange_Windows_Forensics_FileZilla_FileZilla struct {
	Attrselected  string `json:"Attrselected"`
	Port          string `json:"Port"`
	Type          string `json:"Type"`
	User          string `json:"User"`
	RemotePath    string `json:"RemotePath"`
	EncodingType  string `json:"EncodingType"`
	Site          string `json:"Site"`
	Host          string `json:"Host"`
	Attrconnected string `json:"Attrconnected"`
	Protocol      string `json:"Protocol"`
	Pass          struct {
		Attrencoding string `json:"Attrencoding"`
	} `json:"Pass"`
	Logontype      string `json:"Logontype"`
	BypassProxy    string `json:"BypassProxy"`
	LocalPath      string `json:"LocalPath"`
	SourceFilePath string `json:"SourceFilePath"`
}

func (s Exchange_Windows_Forensics_FileZilla_FileZilla) StringArray() []string {
	return []string{s.Type, s.User, s.EncodingType, s.Pass.Attrencoding, s.BypassProxy, s.Host, s.Protocol, s.Logontype, s.Port, s.Site, s.RemotePath, s.LocalPath, s.Attrselected, s.SourceFilePath}
}

func (s Exchange_Windows_Forensics_FileZilla_FileZilla) GetHeaders() []string {
	return []string{"Type", "User", "EncodingType", "AttrEncoding", "BypassProxy", "Host", "Protocol", "Logontype", "Port", "Site", "RemotePath", "LocalPath", "Attrselected", "SourceFilePath"}
}

func Process_Exchange_Windows_Forensics_FileZilla_FileZilla(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_FileZilla_FileZilla{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}

type Exchange_Windows_Forensics_FileZilla_RecentServers struct {
	Type         string `json:"Type"`
	User         string `json:"User"`
	EncodingType string `json:"EncodingType"`
	BypassProxy  string `json:"BypassProxy"`
	Host         string `json:"Host"`
	Protocol     string `json:"Protocol"`
	Logontype    string `json:"Logontype"`
	Port         string `json:"Port"`
	Pass         struct {
		Attrencoding string `json:"Attrencoding"`
	} `json:"Pass"`
	SourceFilePath string `json:"SourceFilePath"`
}

func (s Exchange_Windows_Forensics_FileZilla_RecentServers) StringArray() []string {
	return []string{s.Type, s.User, s.EncodingType, s.Pass.Attrencoding, s.BypassProxy, s.Host, s.Protocol, s.Logontype, s.Port, s.SourceFilePath}
}

func (s Exchange_Windows_Forensics_FileZilla_RecentServers) GetHeaders() []string {
	return []string{"Type", "User", "EncodingType", "AttrEncoding", "BypassProxy", "Host", "Protocol", "Logontype", "Port", "SourceFilePath"}
}

func Process_Exchange_Windows_Forensics_FileZilla_RecentServers(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_FileZilla_RecentServers{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
