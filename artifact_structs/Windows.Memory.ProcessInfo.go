package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
	"strconv"
)

type Windows_Memory_ProcessInfo struct {
	Name             string `json:"Name"`
	PebBaseAddress   string `json:"PebBaseAddress"`
	Pid              int    `json:"Pid"`
	ImagePathName    string `json:"ImagePathName"`
	CommandLine      string `json:"CommandLine"`
	CurrentDirectory string `json:"CurrentDirectory"`
	Env              struct {
		ACSvcPort             string `json:"ACSvcPort"`
		ALLUSERSPROFILE       string `json:"ALLUSERSPROFILE"`
		APPDATA               string `json:"APPDATA"`
		CGOENABLED            string `json:"CGO_ENABLED"`
		ChocolateyInstall     string `json:"ChocolateyInstall"`
		CommonProgramFiles    string `json:"CommonProgramFiles"`
		CommonProgramFilesX86 string `json:"CommonProgramFiles(x86)"`
		CommonProgramW6432    string `json:"CommonProgramW6432"`
		COMPUTERNAME          string `json:"COMPUTERNAME"`
		ComSpec               string `json:"ComSpec"`
		DriverData            string `json:"DriverData"`
		ERLANGHOME            string `json:"ERLANG_HOME"`
		FLASKAPP              string `json:"FLASK_APP"`
		JAVAHOME              string `json:"JAVA_HOME"`
		LOCALAPPDATA          string `json:"LOCALAPPDATA"`
		NETBINS               string `json:"NETBINS"`
		NUMBEROFPROCESSORS    string `json:"NUMBER_OF_PROCESSORS"`
		OneDrive              string `json:"OneDrive"`
		OPENSSLCONF           string `json:"OPENSSL_CONF"`
		OS                    string `json:"OS"`
		Path                  string `json:"Path"`
		PATHEXT               string `json:"PATHEXT"`
		PROCESSORARCHITECTURE string `json:"PROCESSOR_ARCHITECTURE"`
		PROCESSORIDENTIFIER   string `json:"PROCESSOR_IDENTIFIER"`
		PROCESSORLEVEL        string `json:"PROCESSOR_LEVEL"`
		PROCESSORREVISION     string `json:"PROCESSOR_REVISION"`
		ProgramData           string `json:"ProgramData"`
		ProgramFiles          string `json:"ProgramFiles"`
		ProgramFilesX86       string `json:"ProgramFiles(x86)"`
		ProgramW6432          string `json:"ProgramW6432"`
		PSModulePath          string `json:"PSModulePath"`
		PUBLIC                string `json:"PUBLIC"`
		RlsSvcPort            string `json:"RlsSvcPort"`
		SystemDrive           string `json:"SystemDrive"`
		SystemRoot            string `json:"SystemRoot"`
		TEMP                  string `json:"TEMP"`
		TMP                   string `json:"TMP"`
		USERDOMAIN            string `json:"USERDOMAIN"`
		USERNAME              string `json:"USERNAME"`
		USERPROFILE           string `json:"USERPROFILE"`
		VBOXMSIINSTALLPATH    string `json:"VBOX_MSI_INSTALL_PATH"`
		Windir                string `json:"windir"`
	} `json:"Env"`
}

func (s Windows_Memory_ProcessInfo) StringArray() []string {
	baseValues := []string{s.Name, s.PebBaseAddress, strconv.Itoa(s.Pid), s.ImagePathName, s.CommandLine, s.CurrentDirectory}
	envValues := helpers.GetStructValuesAsStringSlice(s.Env)
	baseValues = append(baseValues, envValues...)
	return baseValues
}

func (s Windows_Memory_ProcessInfo) GetHeaders() []string {
	baseHeaders := helpers.GetStructHeadersAsStringSlice(s)
	envHeaders := helpers.GetStructHeadersAsStringSlice(s.Env)
	baseHeaders = baseHeaders[:len(baseHeaders)-1]
	baseHeaders = append(baseHeaders, envHeaders...)
	return baseHeaders
}

func Process_Windows_Memory_ProcessInfo(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Memory_ProcessInfo{}
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
