package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
	"strconv"
)

type Windows_Detection_Impersonation struct {
	ProcPid            int    `json:"ProcPid"`
	ProcName           string `json:"ProcName"`
	Username           string `json:"Username"`
	OwnerSid           string `json:"OwnerSid"`
	TokenIsElevated    bool   `json:"TokenIsElevated"`
	CommandLine        string `json:"CommandLine"`
	Exe                string `json:"Exe"`
	ImpersonationToken struct {
		IsElevated       bool     `json:"IsElevated"`
		User             string   `json:"User"`
		Username         string   `json:"Username"`
		ProfileDir       string   `json:"ProfileDir"`
		PrimaryGroup     string   `json:"PrimaryGroup"`
		PrimaryGroupName string   `json:"PrimaryGroupName"`
		Groups           []string `json:"Groups"`
	} `json:"ImpersonationToken"`
}

func (s Windows_Detection_Impersonation) StringArray() []string {
	return []string{strconv.Itoa(s.ProcPid), s.ProcName, s.Username, s.OwnerSid, strconv.FormatBool(s.TokenIsElevated),
		s.CommandLine, s.Exe, strconv.FormatBool(s.ImpersonationToken.IsElevated), s.ImpersonationToken.User,
		s.ImpersonationToken.Username, s.ImpersonationToken.ProfileDir, s.ImpersonationToken.PrimaryGroup,
		s.ImpersonationToken.PrimaryGroupName, fmt.Sprint(s.ImpersonationToken.Groups)}
}

func (s Windows_Detection_Impersonation) GetHeaders() []string {
	return []string{"ProcPid", "ProcName", "Username", "OwnerSid", "TokenIsElevated", "CommandLine", "Exe", "Impersonation_IsElevated", "Impersonation_User",
		"Impersonation_Username", "Impersonation_ProfileDir", "Impersonation_PrimaryGroup", "Impersonation_PrimaryGroupName", "Impersonation_Groups"}
}

func Process_Windows_Detection_Impersonation(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Detection_Impersonation{}
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
