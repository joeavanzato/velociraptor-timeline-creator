package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_Registry_ScheduledTasks struct {
	TaskID             string    `json:"TaskID"`
	Mtime              time.Time `json:"Mtime"`
	Path               string    `json:"Path"`
	Hash               string    `json:"Hash"`
	Version            string    `json:"Version"`
	SecurityDescriptor string    `json:"SecurityDescriptor"`
	TreeSD             string    `json:"TreeSD"`
	Source             string    `json:"Source"`
	Author             string    `json:"Author"`
	Description        string    `json:"Description"`
	URI                string    `json:"URI"`
	Triggers           string    `json:"Triggers"`
	Actions            string    `json:"Actions"`
	DynamicInfo        string    `json:"DynamicInfo"`
	Schema             string    `json:"Schema"`
	Date               any       `json:"Date"`
	OSPath             string    `json:"OSPath"`
	XMLEntry           struct {
		OSPath     string `json:"OSPath"`
		Command    any    `json:"Command"`
		Arguments  any    `json:"Arguments"`
		ComHandler string `json:"ComHandler"`
		UserID     string `json:"UserId"`
		XML        struct {
			Task struct {
				Attrxmlns        string `json:"Attrxmlns"`
				Attrversion      string `json:"Attrversion"`
				RegistrationInfo struct {
					Version            string `json:"Version"`
					SecurityDescriptor string `json:"SecurityDescriptor"`
					Source             string `json:"Source"`
					Author             string `json:"Author"`
					Description        string `json:"Description"`
					URI                string `json:"URI"`
				} `json:"RegistrationInfo"`
				Principals struct {
					Principal struct {
						Attrid  string `json:"Attrid"`
						GroupID string `json:"GroupId"`
						UserID  string `json:"UserID"`
					} `json:"Principal"`
				} `json:"Principals"`
				Settings struct {
					AllowHardTerminate         string `json:"StopIfGoingOnBatteries"`
					Priority                   string `json:"Priority"`
					StopIfGoingOnBatteries     string `json:"StopIfGoingOnBatteries"`
					ExecutionTimeLimit         string `json:"ExecutionTimeLimit"`
					StartWhenAvailable         string `json:"StartWhenAvailable"`
					RunOnlyIfNetworkAvailable  string `json:"RunOnlyIfNetworkAvailable"`
					UseUnifiedSchedulingEngine string `json:"UseUnifiedSchedulingEngine"`
					DisallowStartIfOnBatteries string `json:"DisallowStartIfOnBatteries"`
					MultipleInstancesPolicy    string `json:"MultipleInstancesPolicy"`
					Hidden                     string `json:"Hidden"`
					RestartOnFailure           struct {
						Count    string `json:"Count"`
						Interval string `json:"Interval"`
					} `json:"RestartOnFailure"`
					IdleSettings struct {
						StopOnIdleEnd string `json:"StopOnIdleEnd"`
						RestartOnIdle string `json:"RestartOnIdle"`
					} `json:"IdleSettings"`
				} `json:"Settings"`
				Triggers any `json:"Triggers"`
				Actions  struct {
					AttrContext string `json:"AttrContext"`
					ComHandler  struct {
						ClassID string `json:"ClassId"`
						Data    string `json:"Data"`
					} `json:"ComHandler"`
				} `json:"Actions"`
			} `json:"Task"`
		} `json:"_XML"`
	} `json:"XmlEntry"`
}

func (s Exchange_Windows_Registry_ScheduledTasks) StringArray() []string {
	return []string{
		s.TaskID,
		s.Mtime.String(),
		s.Path,
		s.Hash,
		fmt.Sprint(s.Version),
		s.SecurityDescriptor,
		s.TreeSD,
		fmt.Sprint(s.Source),
		fmt.Sprint(s.Author),
		fmt.Sprint(s.Description),
		s.URI,
		s.Triggers,
		s.Actions,
		s.DynamicInfo,
		s.Schema,
		fmt.Sprint(s.Date),
		s.OSPath,
		fmt.Sprint(s.XMLEntry.Command),
		fmt.Sprint(s.XMLEntry.Arguments),
		s.XMLEntry.ComHandler,
		s.XMLEntry.UserID,
		s.XMLEntry.XML.Task.Actions.AttrContext,
		s.XMLEntry.XML.Task.Actions.ComHandler.ClassID,
		s.XMLEntry.XML.Task.Actions.ComHandler.Data,
		s.XMLEntry.XML.Task.Attrversion,
		s.XMLEntry.XML.Task.Attrxmlns,
		s.XMLEntry.XML.Task.RegistrationInfo.SecurityDescriptor,
		s.XMLEntry.XML.Task.RegistrationInfo.URI,
		s.XMLEntry.XML.Task.RegistrationInfo.Source,
		s.XMLEntry.XML.Task.RegistrationInfo.Version,
		s.XMLEntry.XML.Task.RegistrationInfo.Author,
		s.XMLEntry.XML.Task.RegistrationInfo.Description,
		s.XMLEntry.XML.Task.Principals.Principal.Attrid,
		s.XMLEntry.XML.Task.Principals.Principal.UserID,
		s.XMLEntry.XML.Task.Principals.Principal.GroupID,
		s.XMLEntry.XML.Task.Settings.AllowHardTerminate,
		s.XMLEntry.XML.Task.Settings.StartWhenAvailable,
		s.XMLEntry.XML.Task.Settings.Priority,
		s.XMLEntry.XML.Task.Settings.IdleSettings.StopOnIdleEnd,
		s.XMLEntry.XML.Task.Settings.IdleSettings.RestartOnIdle,
		s.XMLEntry.XML.Task.Settings.UseUnifiedSchedulingEngine,
		s.XMLEntry.XML.Task.Settings.DisallowStartIfOnBatteries,
		s.XMLEntry.XML.Task.Settings.StopIfGoingOnBatteries,
		s.XMLEntry.XML.Task.Settings.ExecutionTimeLimit,
		s.XMLEntry.XML.Task.Settings.Hidden,
		s.XMLEntry.XML.Task.Settings.MultipleInstancesPolicy,
		fmt.Sprint(s.XMLEntry.XML.Task.Triggers)}
}

func (s Exchange_Windows_Registry_ScheduledTasks) GetHeaders() []string {
	return []string{"TaskID", "Mtime", "Path", "Hash", "Version", "SecurityDescriptor", "TreeSD", "Source", "Author", "Description", "URI",
		"Triggers", "Actions", "DynamicInfo", "Schema", "Date", "OSPath", "XML_Command", "XML_Arguments", "XML_ComHandler", "XML_UserID", "Actions_AttrContext",
		"Action_COM_ClassID", "Actions_COM_Data", "Task_Attrversion", "Task_Attrxmlns", "Registration_SecurityDescriptor", "Registration_URI",
		"Registration_Source", "Registration_Version", "Registration_Author", "Registration_Description", "Principal_Attrid", "Principal_UserID",
		"Principal_GroupID", "AllowHardTerminate", "StartWhenAvailable", "Priority", "StopOnIdleEnd", "RestartOnIdle", "UseUnifiedSchedulingEngine",
		"DisallowStartIfOnBatteries", "StopIfGoingOnBatteries", "ExecutionTimeLimit", "Hidden", "MultipleInstancesPolicy", "Triggers"}
}

func Process_Exchange_Windows_Registry_ScheduledTasks(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Registry_ScheduledTasks{}
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
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       tmp.XMLEntry.XML.Task.RegistrationInfo.Author,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Path,
			MetaData:         fmt.Sprintf("Command: %v, Arguments: %v", tmp.XMLEntry.Command, tmp.XMLEntry.Arguments),
		}
		outputChannel <- tmp2.StringArray()
	}
}
