package artifact_structs

import (
	"encoding/json"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type Windows_System_Powershell_PSReadline struct {
	Stat struct {
		Mtime  time.Time `json:"Mtime"`
		Atime  time.Time `json:"Atime"`
		Ctime  time.Time `json:"Ctime"`
		Btime  time.Time `json:"Btime"`
		Size   int       `json:"Size"`
		OSPath string    `json:"OSPath"`
	} `json:"Stat"`
	LineNum  int    `json:"LineNum"`
	Line     string `json:"Line"`
	Username string `json:"Username"`
	OSPath   string `json:"OSPath"`
}

func (s Windows_System_Powershell_PSReadline) StringArray() []string {
	return []string{s.Stat.Mtime.String(), s.Stat.Atime.String(), s.Stat.Ctime.String(), s.Stat.Btime.String(), strconv.Itoa(s.Stat.Size), s.Stat.OSPath, strconv.Itoa(s.LineNum), s.Line, s.Username, s.OSPath}
}

func (s Windows_System_Powershell_PSReadline) GetHeaders() []string {
	return []string{"Mtime", "Atime", "Ctime", "Btime", "Size", "OSPath", "LineNum", "Line", "Username", "OSPath"}
}

func Process_Windows_System_Powershell_PSReadline(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Powershell_PSReadline{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
	}
}
