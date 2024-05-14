package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
)

type Windows_System_TaskScheduler struct {
	OSPath     string `json:"OSPath"`
	Command    any    `json:"Command"`
	Arguments  any    `json:"Arguments"`
	ComHandler any    `json:"ComHandler"`
	UserID     string `json:"UserId"`
	XML        struct {
		Task struct {
			Actions struct {
				AttrContext string `json:"AttrContext"`
				Exec        any    `json:"Exec"`
			} `json:"Actions"`
			Attrversion string `json:"Attrversion"`
			Attrxmlns   string `json:"Attrxmlns"`
			Principals  struct {
				Principal struct {
					Attrid    string `json:"Attrid"`
					LogonType string `json:"LogonType"`
					RunLevel  string `json:"RunLevel"`
					UserID    string `json:"UserId"`
				} `json:"Principal"`
			} `json:"Principals"`
			RegistrationInfo struct {
				Author string `json:"Author"`
				URI    string `json:"URI"`
			} `json:"RegistrationInfo"`
			Settings struct {
				AllowHardTerminate              string `json:"AllowHardTerminate"`
				AllowStartOnDemand              string `json:"AllowStartOnDemand"`
				DisallowStartIfOnBatteries      string `json:"DisallowStartIfOnBatteries"`
				DisallowStartOnRemoteAppSession string `json:"DisallowStartOnRemoteAppSession"`
				Enabled                         string `json:"Enabled"`
				ExecutionTimeLimit              string `json:"ExecutionTimeLimit"`
				Hidden                          string `json:"Hidden"`
				IdleSettings                    struct {
					Duration      string `json:"Duration"`
					RestartOnIdle string `json:"RestartOnIdle"`
					StopOnIdleEnd string `json:"StopOnIdleEnd"`
					WaitTimeout   string `json:"WaitTimeout"`
				} `json:"IdleSettings"`
				MultipleInstancesPolicy    string `json:"MultipleInstancesPolicy"`
				Priority                   string `json:"Priority"`
				RunOnlyIfIdle              string `json:"RunOnlyIfIdle"`
				RunOnlyIfNetworkAvailable  string `json:"RunOnlyIfNetworkAvailable"`
				StartWhenAvailable         string `json:"StartWhenAvailable"`
				StopIfGoingOnBatteries     string `json:"StopIfGoingOnBatteries"`
				UseUnifiedSchedulingEngine string `json:"UseUnifiedSchedulingEngine"`
				WakeToRun                  string `json:"WakeToRun"`
			} `json:"Settings"`
			Triggers any `json:"Triggers"`
		} `json:"Task"`
	} `json:"_XML"`
}

func (s Windows_System_TaskScheduler) StringArray() []string {
	return []string{s.OSPath, fmt.Sprint(s.Command), fmt.Sprint(s.Arguments), fmt.Sprint(s.ComHandler), s.UserID, s.XML.Task.Actions.AttrContext,
		fmt.Sprint(s.XML.Task.Actions.Exec),
		s.XML.Task.Attrversion, s.XML.Task.Attrxmlns,
		s.XML.Task.Principals.Principal.Attrid, s.XML.Task.Principals.Principal.LogonType,
		s.XML.Task.Principals.Principal.RunLevel, s.XML.Task.Principals.Principal.UserID,
		s.XML.Task.RegistrationInfo.Author, s.XML.Task.RegistrationInfo.URI,
		s.XML.Task.Settings.AllowHardTerminate, s.XML.Task.Settings.AllowStartOnDemand, s.XML.Task.Settings.DisallowStartIfOnBatteries,
		s.XML.Task.Settings.DisallowStartOnRemoteAppSession, s.XML.Task.Settings.Enabled, s.XML.Task.Settings.ExecutionTimeLimit,
		s.XML.Task.Settings.Hidden, s.XML.Task.Settings.IdleSettings.Duration,
		s.XML.Task.Settings.IdleSettings.RestartOnIdle, s.XML.Task.Settings.IdleSettings.StopOnIdleEnd, s.XML.Task.Settings.IdleSettings.WaitTimeout,
		s.XML.Task.Settings.MultipleInstancesPolicy, s.XML.Task.Settings.Priority, s.XML.Task.Settings.RunOnlyIfIdle, s.XML.Task.Settings.RunOnlyIfNetworkAvailable,
		s.XML.Task.Settings.StartWhenAvailable, s.XML.Task.Settings.StopIfGoingOnBatteries, s.XML.Task.Settings.UseUnifiedSchedulingEngine, s.XML.Task.Settings.WakeToRun,
		fmt.Sprint(s.XML.Task.Triggers)}
}

func (s Windows_System_TaskScheduler) GetHeaders() []string {
	return []string{"OSPath", "Command", "Arguments", "ComHandler", "UserID", "AttrContext", "ExecutionComponents", "Attrversion", "Attrxmlns", "Principal_Attrid", "Principal_LogonType",
		"Principal_RunLevel", "Principal_UserID", "Author", "URI", "AllowHardTerminate", "AllowStartOnDemand", "DisallowStartIfOnBatteries", "DisallowStartOnRemoteAppSession", "Enabled",
		"ExecutionTimeLimit", "Hidden", "Idle_Duration", "RestartOnIdle", "StopOnIdleEnd", "WaitTimeout", "MultipleInstancesPolicy", "Priority", "RunOnlyIfIdle", "RunOnlyIfNetworkAvailable", "StartWhenAvailable",
		"StopIfGoingOnBatteries", "UseUnifiedSchedulingEngine", "WakeToRun", "TriggerData"}
}

func Process_Windows_System_TaskScheduler(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_TaskScheduler{}
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
