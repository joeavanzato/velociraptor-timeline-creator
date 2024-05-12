package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations struct {
	EventTime             string `json:"EventTime"`
	Computer              string `json:"Computer"`
	EID                   string `json:"EID"`
	Description           string `json:"Description"`
	Channel               string `json:"Channel"`
	Provider              string `json:"Provider"`
	SystemProcessID       string `json:"System Process ID"`
	Keywords              string `json:"Keywords"`
	SecurityUserID        string `json:"Security User ID"`
	CorrelationActivityID string `json:"Correlation ActivityID"`
	EventRecordID         string `json:"EventRecord ID"`
	TargetUserName        string `json:"TargetUserName"`
	TargetDomainName      string `json:"TargetDomainName"`
	TargetSid             string `json:"TargetSid"`
	OldTargetUserName     string `json:"OldTargetUserName"`
	NewTargetUserName     string `json:"NewTargetUserName"`
	SubjectUserSid        string `json:"SubjectUserSid"`
	SubjectUserName       string `json:"SubjectUserName"`
	SubjectDomainName     string `json:"SubjectDomainName"`
	SubjectLogonID        string `json:"SubjectLogonId"`
	PrivilegeList         string `json:"PrivilegeList"`
	MemberName            string `json:"MemberName"`
	MemberSid             string `json:"MemberSid"`
	SourceUserName        string `json:"SourceUserName"`
	SourceSid             string `json:"SourceSid"`
	SamAccountName        string `json:"SamAccountName"`
	DisplayName           string `json:"DisplayName"`
	UserPrincipalName     string `json:"UserPrincipalName"`
	HomeDirectory         string `json:"HomeDirectory"`
	HomePath              string `json:"HomePath"`
	ScriptPath            string `json:"ScriptPath"`
	ProfilePath           string `json:"ProfilePath"`
	UserWorkstations      string `json:"UserWorkstations"`
	PasswordLastSet       string `json:"PasswordLastSet"`
	AccountExpires        string `json:"AccountExpires"`
	PrimaryGroupID        string `json:"PrimaryGroupId"`
	AllowedToDelegateTo   string `json:"AllowedToDelegateTo"`
	OldUacValue           string `json:"OldUacValue"`
	NewUacValue           string `json:"NewUacValue"`
	UserAccountControl    string `json:"UserAccountControl"`
	UserParameters        string `json:"UserParameters"`
	SidHistorySidList     string `json:"SidHistory/SidList"`
	LogonHours            string `json:"LogonHours"`
	Status                string `json:"Status"`
	Source                string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations) StringArray() []string {
	return []string{s.EventTime, s.Computer, s.EID, s.Description, s.Channel, s.Provider, s.SystemProcessID, s.Keywords, s.SecurityUserID, s.CorrelationActivityID, s.EventRecordID, s.TargetUserName, s.TargetDomainName,
		s.TargetSid, s.OldTargetUserName, s.NewTargetUserName, s.SubjectUserSid, s.SubjectUserName, s.SubjectDomainName, s.SubjectLogonID, s.PrivilegeList, s.MemberName, s.MemberSid, s.SourceUserName, s.SourceSid,
		s.SamAccountName, s.DisplayName, s.UserPrincipalName, s.HomeDirectory, s.HomePath, s.ScriptPath, s.ProfilePath, s.UserWorkstations, s.PasswordLastSet, s.AccountExpires, s.PrimaryGroupID, s.AllowedToDelegateTo,
		s.OldUacValue, s.NewUacValue, s.UserAccountControl, s.UserParameters, s.SidHistorySidList, s.LogonHours, s.Status, s.Source}
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       tmp.SamAccountName,
			SourceHost:       tmp.Computer,
			DestinationUser:  tmp.TargetUserName,
			DestinationHost:  "",
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("Keywords: %v, Status: %v, EventID: %v", tmp.Keywords, tmp.Status, tmp.EID),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender struct {
	EventTime                   string `json:"EventTime"`
	Computer                    string `json:"Computer"`
	EID                         string `json:"EID"`
	Description                 string `json:"Description"`
	Channel                     string `json:"Channel"`
	Provider                    string `json:"Provider"`
	SystemProcessID             string `json:"System Process ID"`
	Keywords                    string `json:"Keywords"`
	SecurityUserID              string `json:"Security User ID"`
	CorrelationActivityID       string `json:"Correlation ActivityID"`
	EventRecordID               string `json:"EventRecord ID"`
	ProductName                 string `json:"Product Name"`
	ProductVersion              string `json:"Product Version"`
	ScanID                      string `json:"Scan ID"`
	DetectionID                 string `json:"Detection ID"`
	ThreatName                  string `json:"Threat Name"`
	ThreatID                    string `json:"Threat ID"`
	SeverityName                string `json:"Severity Name"`
	CategoryName                string `json:"Category Name"`
	ActionName                  string `json:"Action Name"`
	ProcessName                 string `json:"Process Name"`
	ProcessID                   string `json:"Process ID"`
	Path                        string `json:"Path"`
	ScanType                    string `json:"Scan Type"`
	ScanParameters              string `json:"Scan Parameters"`
	Domain                      string `json:"Domain"`
	User                        string `json:"User"`
	RemediationUser             string `json:"Remediation User"`
	SID                         string `json:"SID"`
	ScanResources               string `json:"Scan Resources"`
	Timestamp                   string `json:"Timestamp"`
	OldValue                    string `json:"Old Value"`
	NewValue                    string `json:"New Value"`
	ErrorCode                   string `json:"Error Code"`
	ErrorDescription            string `json:"Error Description"`
	DetectionSource             string `json:"Detection Source"`
	DetectionOrigin             string `json:"Detection Origin"`
	ExecutionStatus             string `json:"Execution Status"`
	DetectionType               string `json:"Detection Type"`
	SecurityIntelligenceVersion string `json:"Security intelligence Version"`
	EngineVersion               string `json:"Engine Version"`
	StatusCode                  string `json:"Status Code"`
	CleaningAction              string `json:"Cleaning Action"`
	SignatureVersion            string `json:"Signature Version"`
	PreviousSignatureVersion    string `json:"Previous Signature Version"`
	FidelityValue               string `json:"FidelityValue"`
	FidelityLabel               string `json:"FidelityLabel"`
	ImageFileHash               string `json:"Image File Hash"`
	TargetFileHash              string `json:"TargetFileHash"`
	TargetFileName              string `json:"TargetFileName"`
	State                       string `json:"State"`
	SourceName                  string `json:"Source Name"`
	OriginName                  string `json:"Origin Name"`
	ExecutionName               string `json:"Execution Name"`
	TypeName                    string `json:"Type Name"`
	PreExecutionStatus          string `json:"Pre Execution Status"`
	PostCleanStatus             string `json:"Post Clean Status"`
	AdditionalActionsString     string `json:"Additional Actions String"`
	TargetCommandline           string `json:"Target Commandline"`
	ParentCommandline           string `json:"Parent Commandline"`
	InvolvedFile                string `json:"Involved File"`
	InhertianceFlags            string `json:"Inhertiance Flags"`
	SourceApp                   string `json:"Source app"`
	TargetApp                   string `json:"Target app"`
	LastFullScanStartTime       string `json:"Last full scan start time"`
	LastFullScanEndTime         string `json:"Last full scan end time"`
	LastFullScanSource          string `json:"Last full scan source"`
	SecurityIntelligenceType    string `json:"Security intelligence Type"`
	UpdateType                  string `json:"Update Type"`
	UpdateSource                string `json:"Update Source"`
	SignatureType               string `json:"Signature Type"`
	UpdateState                 string `json:"Update State"`
	FeatureName                 string `json:"Feature Name"`
	Reason                      string `json:"Reason"`
	Configuration               string `json:"Configuration"`
	Resource                    string `json:"Resource"`
	FailureType                 string `json:"Failure Type"`
	ExceptionCode               string `json:"Exception Code"`
	ChangedType                 string `json:"Changed Type"`
	ScanTimeHours               string `json:"Scan Time Hours"`
	ScanTimeMinutes             string `json:"Scan Time Minutes"`
	ScanTimeSeconds             string `json:"Scan Time Seconds"`
	FWLink                      string `json:"FWLink"`
	Source                      string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender) StringArray() []string {
	return []string{s.EventTime, s.Computer, s.EID, s.Description, s.Channel, s.Provider, s.SystemProcessID, s.Keywords, s.SecurityUserID, s.CorrelationActivityID, s.EventRecordID,
		s.ProductName, s.ProductVersion, s.ScanID, s.DetectionID, s.ThreatName, s.ThreatID, s.SeverityName, s.CategoryName, s.ActionName, s.ProcessName, s.ProcessID,
		s.Path, s.ScanType, s.ScanParameters, s.Domain, s.User, s.RemediationUser, s.SID, s.ScanResources, s.Timestamp, s.OldValue, s.NewValue, s.ErrorCode, s.ErrorDescription,
		s.DetectionSource, s.DetectionOrigin, s.ExecutionStatus, s.DetectionType, s.SecurityIntelligenceVersion, s.EngineVersion, s.StatusCode, s.CleaningAction, s.SignatureVersion,
		s.PreviousSignatureVersion, s.FidelityValue, s.FidelityLabel, s.ImageFileHash, s.TargetFileHash, s.TargetFileName, s.State, s.SourceName, s.OriginName, s.ExecutionName,
		s.TypeName, s.PreExecutionStatus, s.PostCleanStatus, s.AdditionalActionsString, s.TargetCommandline, s.ParentCommandline, s.InvolvedFile, s.InhertianceFlags,
		s.SourceApp, s.TargetApp, s.LastFullScanStartTime, s.LastFullScanEndTime, s.LastFullScanSource, s.SecurityIntelligenceType, s.UpdateType, s.UpdateSource, s.SignatureType, s.UpdateState,
		s.FeatureName, s.Reason, s.Configuration, s.Resource, s.FailureType, s.ExceptionCode, s.ChangedType, s.ScanTimeHours, s.ScanTimeMinutes, s.ScanTimeSeconds, s.FWLink, s.Source}
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: fmt.Sprintf("Threat Name: %v, Severity: %v, Category: %v", tmp.ThreatName, tmp.SeverityName, tmp.CategoryName),
			SourceUser:       tmp.User,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.InvolvedFile,
			MetaData:         fmt.Sprintf("Command Line: %v, Target Hash: %v, Target File Name: %v, Cleaning Action: %v", tmp.TargetCommandline, tmp.TargetFileHash, tmp.TargetFileName, tmp.CleaningAction),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown struct {
	EventTime             string `json:"EventTime"`
	Computer              string `json:"Computer"`
	EID                   string `json:"EID"`
	Description           string `json:"Description"`
	Channel               string `json:"Channel"`
	Provider              string `json:"Provider"`
	SystemProcessID       string `json:"System Process ID"`
	Keywords              string `json:"Keywords"`
	SecurityUserID        string `json:"Security User ID"`
	CorrelationActivityID string `json:"Correlation ActivityID"`
	EventRecordID         string `json:"EventRecord ID"`
	SubjectUserName       string `json:"SubjectUserName"`
	SubjectUserSid        string `json:"SubjectUserSid"`
	SubjectDomainName     string `json:"SubjectDomainName"`
	SubjectLogonID        string `json:"SubjectLogonId"`
	ProcessName           string `json:"ProcessName"`
	Reason                string `json:"Reason"`
	ReasonCode            string `json:"ReasonCode"`
	Status                string `json:"Status"`
	Type                  string `json:"Type"`
	Comment               string `json:"Comment"`
	SourceComputer        string `json:"SourceComputer"`
	StartTime             string `json:"StartTime"`
	StopTime              string `json:"StopTime"`
	Uptime                string `json:"Uptime"`
	LastShutdownGood      string `json:"LastShutdownGood"`
	LastBootGood          string `json:"LastBootGood"`
	Bugcheck              string `json:"Bugcheck"`
	DumpPath              string `json:"DumpPath"`
	ReportID              string `json:"ReportID"`
	ObjectServer          string `json:"ObjectServer"`
	ObjectType            string `json:"ObjectType"`
	ObjectName            string `json:"ObjectName"`
	HandleID              string `json:"HandleId"`
	AccessMask            string `json:"AccessMask"`
	PrivilegeList         string `json:"PrivilegeList"`
	ProcessID             string `json:"ProcessId"`
	Source                string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown) StringArray() []string {
	return []string{s.EventTime, s.Computer, s.EID, s.Description, s.Channel, s.Provider, s.SystemProcessID, s.Keywords, s.SecurityUserID, s.CorrelationActivityID, s.EventRecordID,
		s.SubjectUserName, s.SubjectUserSid, s.SubjectDomainName, s.SubjectLogonID, s.ProcessName, s.Reason, s.ReasonCode, s.Status, s.Type, s.Comment, s.SourceComputer, s.StartTime, s.StopTime,
		s.Uptime, s.LastShutdownGood, s.LastBootGood, s.Bugcheck, s.DumpPath, s.ReportID, s.ObjectServer, s.ObjectType, s.ObjectName, s.HandleID, s.AccessMask, s.PrivilegeList, s.ProcessID, s.Source}
}

func (s Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       tmp.SubjectUserName,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, Process Name: %v, Reason: %v, Status: %v, Uptime: %v", tmp.EID, tmp.ProcessName, tmp.Reason, tmp.Status, tmp.Uptime),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_Logons struct {
	EventTime                            string `json:"EventTime"`
	Computer                             string `json:"Computer"`
	EID                                  string `json:"EID"`
	Description                          string `json:"Description"`
	Channel                              string `json:"Channel"`
	Provider                             string `json:"Provider"`
	SystemProcessID                      string `json:"System Process ID"`
	Keywords                             string `json:"Keywords"`
	SecurityUserID                       string `json:"Security User ID"`
	CorrelationActivityID                string `json:"Correlation ActivityID"`
	EventRecordID                        string `json:"EventRecord ID"`
	SubjectUserSid                       string `json:"SubjectUserSid"`
	SubjectUserName                      string `json:"SubjectUserName"`
	SubjectDomainName                    string `json:"SubjectDomainName"`
	SubjectLogonID                       string `json:"SubjectLogonId"`
	LogonGUID                            string `json:"LogonGuid"`
	TargetUserSid                        string `json:"TargetUserSid"`
	TargetUserName                       string `json:"TargetUserName"`
	TargetDomainName                     string `json:"TargetDomainName"`
	TargetLogonID                        string `json:"TargetLogonId"`
	LogonType                            string `json:"LogonType"`
	LogonTypeUseCases                    string `json:"LogonType (Use cases)"`
	TargetLogonGUID                      string `json:"TargetLogonGuid"`
	TargetServerName                     string `json:"TargetServerName"`
	TargetInfo                           string `json:"TargetInfo"`
	PrivilegeList                        string `json:"PrivilegeList"`
	Status                               string `json:"Status"`
	SubStatus                            string `json:"SubStatus"`
	FailureReasonFailureCode             string `json:"FailureReason/FailureCode"`
	LogonProcessName                     string `json:"LogonProcessName"`
	AuthenticationPackageName            string `json:"AuthenticationPackageName"`
	WorkstationName                      string `json:"WorkstationName"`
	IPAddress                            string `json:"IpAddress"`
	IPPort                               string `json:"IpPort"`
	ServiceName                          string `json:"ServiceName"`
	ServiceSid                           string `json:"ServiceSid"`
	SidList                              string `json:"SidList"`
	SessionName                          string `json:"SessionName"`
	SessionID                            string `json:"SessionId"`
	TransmittedServicesTransitedServices string `json:"TransmittedServices/TransitedServices"`
	LmPackageName                        string `json:"LmPackageName"`
	KeyLength                            string `json:"KeyLength"`
	ProcessID                            string `json:"ProcessId"`
	ProcessName                          string `json:"ProcessName"`
	ImpersonationLevel                   string `json:"ImpersonationLevel"`
	RestrictedAdminMode                  string `json:"RestrictedAdminMode"`
	TargetOutboundUserName               string `json:"TargetOutboundUserName"`
	TargetOutboundDomainName             string `json:"TargetOutboundDomainName"`
	VirtualAccount                       string `json:"VirtualAccount"`
	TargetLinkedLogonID                  string `json:"TargetLinkedLogonId"`
	ElevatedToken                        string `json:"ElevatedToken"`
	GroupMembership                      string `json:"GroupMembership"`
	RequestType                          string `json:"RequestType"`
	TicketOptions                        string `json:"TicketOptions"`
	TicketEncryptionType                 string `json:"TicketEncryptionType"`
	PreAuthType                          string `json:"PreAuthType"`
	CertIssuerName                       string `json:"CertIssuerName"`
	CertSerialNumber                     string `json:"CertSerialNumber"`
	CertThumbprint                       string `json:"CertThumbprint"`
	SiloName                             string `json:"SiloName"`
	PolicyName                           string `json:"PolicyName"`
	TGTLifetime                          string `json:"TGT Lifetime"`
	Source                               string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Logons) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Logons) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_Logons(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_Logons{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: fmt.Sprintf("Status: %v, Reason: %v", tmp.Status, tmp.FailureReasonFailureCode),
			SourceUser:       tmp.SubjectUserName,
			SourceHost:       fmt.Sprintf("Workstation: %v, IP: %v", tmp.WorkstationName, tmp.IPAddress),
			DestinationUser:  fmt.Sprintf("TargetUser: %v, TargetOutboundUser: %v", tmp.TargetUserName, tmp.TargetOutboundUserName),
			DestinationHost:  "",
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, LogonProcess: %v, LogonType: %v, Ticket Encryption: %v", tmp.EID, tmp.LogonProcessName, tmp.LogonType, tmp.TicketEncryptionType),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events struct {
	EventTime                         string `json:"EventTime"`
	Computer                          string `json:"Computer"`
	EID                               string `json:"EID"`
	Description                       string `json:"Description"`
	Channel                           string `json:"Channel"`
	Provider                          string `json:"Provider"`
	SystemProcessID                   string `json:"System Process ID"`
	Keywords                          string `json:"Keywords"`
	SecurityUserID                    string `json:"Security User ID"`
	CorrelationActivityID             string `json:"Correlation ActivityID"`
	EventRecordID                     string `json:"EventRecord ID"`
	HostApplication                   string `json:"HostApplication"`
	HostApplicationBase64Decoded      string `json:"HostApplication (Base64 decoded)"`
	ScriptName                        string `json:"ScriptName"`
	Payload                           string `json:"Payload"`
	HostName                          string `json:"HostName"`
	HostVersion                       string `json:"HostVersion"`
	EngineVersion                     string `json:"EngineVersion"`
	CommandInvocationParameterBinding string `json:"CommandInvocation/ParameterBinding"`
	CommandLine                       string `json:"CommandLine"`
	CommandName                       string `json:"CommandName"`
	CommandPath                       string `json:"CommandPath"`
	CommandType                       string `json:"CommandType"`
	RunspaceID                        string `json:"RunspaceId"`
	ProviderName                      string `json:"ProviderName"`
	ConnectedUser                     string `json:"Connected User"`
	User                              string `json:"User"`
	UserID                            string `json:"UserId"`
	DetailSequence                    string `json:"DetailSequence"`
	DetailTotal                       string `json:"DetailTotal"`
	ErrorMessage                      string `json:"ErrorMessage"`
	ErrorCode                         string `json:"ErrorCode"`
	FileName                          string `json:"FileName"`
	FullyQualifiedErrorID             string `json:"Fully Qualified Error ID"`
	HostID                            string `json:"HostId"`
	InstanceID                        string `json:"InstanceId"`
	NewCommandState                   string `json:"NewCommandState"`
	PreviousEngineState               string `json:"PreviousEngineState"`
	NewEngineState                    string `json:"NewEngineState"`
	NewProviderState                  string `json:"NewProviderState"`
	Path                              string `json:"Path"`
	PipelineID                        string `json:"PipelineId"`
	ScriptBlockID                     string `json:"ScriptBlockId"`
	ScriptBlockText                   string `json:"ScriptBlockText"`
	SequenceNumber                    string `json:"SequenceNumber"`
	SessionID                         string `json:"SessionId"`
	Severity                          string `json:"Severity"`
	ShellID                           string `json:"Shell ID"`
	Param1                            string `json:"param1"`
	Param2                            string `json:"param2"`
	MinRunspaces                      string `json:"MinRunspaces"`
	MaxRunspaces                      string `json:"MaxRunspaces"`
	Source                            string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, ScriptName: %v, CommandLine: %v, ScriptBlock: %v", tmp.EID, tmp.ScriptName, tmp.CommandLine, tmp.ScriptBlockText),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary struct {
	FinalName                  string `json:"FinalName"`
	UploadSHA256               string `json:"UploadSHA256"`
	Size                       int    `json:"Size"`
	HumanSize                  string `json:"Human Size"`
	First100CharactersOfScript string `json:"First 100 characters of script"`
	Source                     string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary{}
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

type Exchange_Windows_EventLogs_EvtxHussar_RDP struct {
	EventTime             string `json:"EventTime"`
	Computer              string `json:"Computer"`
	EID                   string `json:"EID"`
	Description           string `json:"Description"`
	Channel               string `json:"Channel"`
	Provider              string `json:"Provider"`
	SystemProcessID       string `json:"System Process ID"`
	Keywords              string `json:"Keywords"`
	SecurityUserID        string `json:"Security User ID"`
	CorrelationActivityID string `json:"Correlation ActivityID"`
	EventRecordID         string `json:"EventRecord ID"`
	SourceIP              string `json:"SourceIP"`
	SourceWorkstation     string `json:"SourceWorkstation"`
	ServerName            string `json:"ServerName"`
	TargetIP              string `json:"TargetIP"`
	User                  string `json:"User"`
	DomainName            string `json:"DomainName"`
	LogonType             string `json:"LogonType"`
	SessionID             string `json:"SessionID"`
	SourceSessionID       string `json:"SourceSessionID"`
	StatusCode            string `json:"Status Code"`
	MessageName           string `json:"MessageName"`
	ErrorCode             string `json:"ErrorCode"`
	ConnectionName        string `json:"ConnectionName"`
	State                 string `json:"State"`
	StateName             string `json:"StateName"`
	Event                 string `json:"Event"`
	EventName             string `json:"EventName"`
	Reason                string `json:"Reason"`
	ReasonCode            string `json:"ReasonCode"`
	TimezoneBiasHour      string `json:"TimezoneBiasHour"`
	ConnType              string `json:"ConnType"`
	MonitorWidth          string `json:"MonitorWidth"`
	MonitorHeight         string `json:"MonitorHeight"`
	MajorType             string `json:"MajorType"`
	MinorType             string `json:"MinorType"`
	Source                string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_RDP) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_RDP) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_RDP(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_RDP{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       tmp.User,
			SourceHost:       fmt.Sprintf("Workstation: %v, IP: %v", tmp.SourceWorkstation, tmp.SourceIP),
			DestinationUser:  "",
			DestinationHost:  fmt.Sprintf("Workstation: %v, IP: %v", tmp.Computer, tmp.TargetIP),
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, ConnType: %v, EventName: %v, Event: %v", tmp.EID, tmp.EID, tmp.ConnType, tmp.EventName, tmp.Event),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_Services struct {
	EventTime                 string `json:"EventTime"`
	Computer                  string `json:"Computer"`
	EID                       string `json:"EID"`
	Description               string `json:"Description"`
	Channel                   string `json:"Channel"`
	Provider                  string `json:"Provider"`
	SystemProcessID           string `json:"System Process ID"`
	Keywords                  string `json:"Keywords"`
	SecurityUserID            string `json:"Security User ID"`
	CorrelationActivityID     string `json:"Correlation ActivityID"`
	EventRecordID             string `json:"EventRecord ID"`
	ServiceName               string `json:"ServiceName"`
	ExtraServiceName          string `json:"ExtraServiceName"`
	ImagePathServiceFileName  string `json:"ImagePath/ServiceFileName"`
	ServiceType               string `json:"ServiceType"`
	ServiceStartType          string `json:"ServiceStartType"`
	ServiceAccountAccountName string `json:"ServiceAccount/AccountName"`
	SubjectUserSid            string `json:"SubjectUserSid"`
	SubjectUserName           string `json:"SubjectUserName"`
	SubjectDomainName         string `json:"SubjectDomainName"`
	SubjectLogonID            string `json:"SubjectLogonId"`
	ServiceStartTypeOld       string `json:"ServiceStartTypeOld"`
	ServiceStartTypeNew       string `json:"ServiceStartTypeNew"`
	ServiceReason             string `json:"ServiceReason"`
	ServiceReasonText         string `json:"ServiceReasonText"`
	State                     string `json:"State"`
	Status                    string `json:"Status"`
	Comment                   string `json:"Comment"`
	Error                     string `json:"Error"`
	ClientProcessStartKey     string `json:"ClientProcessStartKey"`
	ClientProcessID           string `json:"ClientProcessId"`
	ParentProcessID           string `json:"ParentProcessId"`
	Source                    string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Services) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_Services) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_Services(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_Services{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       tmp.SubjectUserName,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, ImagePathServiceFileName: %v, ServiceName: %v", tmp.EID, tmp.ImagePathServiceFileName, tmp.ServiceName),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations struct {
	EventTime             string `json:"EventTime"`
	Computer              string `json:"Computer"`
	EID                   string `json:"EID"`
	Description           string `json:"Description"`
	Channel               string `json:"Channel"`
	Provider              string `json:"Provider"`
	SystemProcessID       string `json:"System Process ID"`
	Keywords              string `json:"Keywords"`
	SecurityUserID        string `json:"Security User ID"`
	CorrelationActivityID string `json:"Correlation ActivityID"`
	EventRecordID         string `json:"EventRecord ID"`
	RemoteAddress         string `json:"RemoteAddress"`
	ServerName            string `json:"ServerName"`
	UserName              string `json:"UserName"`
	LogonID               string `json:"LogonId"`
	PrincipalName         string `json:"PrincipalName"`
	LocalAddress          string `json:"LocalAddress"`
	Reason                string `json:"Reason"`
	Status                string `json:"Status"`
	SecurityStatus        string `json:"SecurityStatus"`
	ShareName             string `json:"ShareName"`
	SecurityMode          string `json:"SecurityMode"`
	ObjectName            string `json:"ObjectName"`
	ConnectionType        string `json:"ConnectionType"`
	SessionID             string `json:"SessionId"`
	Smb2Command           string `json:"Smb2Command"`
	MessageID             string `json:"MessageId"`
	Object                string `json:"Object"`
	OldState              string `json:"OldState"`
	NewState              string `json:"NewState"`
	Capabilities          string `json:"Capabilities"`
	GUID                  string `json:"Guid"`
	TreeID                string `json:"TreeId"`
	InstanceName          string `json:"InstanceName"`
	Dialect               string `json:"Dialect"`
	Dialect2              string `json:"Dialect2"`
	SecurityMode2         string `json:"SecurityMode2"`
	Capabilities2         string `json:"Capabilities2"`
	GUID2                 string `json:"Guid2"`
	Source                string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       tmp.UserName,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  tmp.RemoteAddress,
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, ShareName: %v, ConnectionType: %v", tmp.EID, tmp.ShareName, tmp.ConnectionType),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit struct {
	EventTime               string `json:"EventTime"`
	Computer                string `json:"Computer"`
	EID                     string `json:"EID"`
	Description             string `json:"Description"`
	Channel                 string `json:"Channel"`
	Provider                string `json:"Provider"`
	SystemProcessID         string `json:"System Process ID"`
	Keywords                string `json:"Keywords"`
	SecurityUserID          string `json:"Security User ID"`
	CorrelationActivityID   string `json:"Correlation ActivityID"`
	EventRecordID           string `json:"EventRecord ID"`
	ShareName               string `json:"ShareName"`
	ShareLocalPath          string `json:"ShareLocalPath"`
	ClientName              string `json:"ClientName"`
	SubjectUserName         string `json:"SubjectUserName"`
	SubjectUserSid          string `json:"SubjectUserSid"`
	SubjectDomainName       string `json:"SubjectDomainName"`
	SubjectLogonID          string `json:"SubjectLogonId"`
	ClientAddress           string `json:"ClientAddress"`
	IPPort                  string `json:"IpPort"`
	Status                  string `json:"Status"`
	SessionGUID             string `json:"SessionGuid"`
	FileName                string `json:"FileName"`
	ComputerName            string `json:"ComputerName"`
	RelativeTargetName      string `json:"RelativeTargetName"`
	ObjectType              string `json:"ObjectType"`
	AccessMask              string `json:"AccessMask"`
	AccessList              string `json:"AccessList"`
	AccessReason            string `json:"AccessReason"`
	DurableHandle           string `json:"DurableHandle"`
	ResilientHandle         string `json:"ResilientHandle"`
	PersistentHandle        string `json:"PersistentHandle"`
	ResumeKey               string `json:"ResumeKey"`
	Reason                  string `json:"Reason"`
	PersistentFID           string `json:"PersistentFID"`
	VolatileFID             string `json:"VolatileFID"`
	Command                 string `json:"Command"`
	Duration                string `json:"Duration"`
	TranslatedStatus        string `json:"TranslatedStatus"`
	RKFStatus               string `json:"RKFStatus"`
	TranslatedRKFStatus     string `json:"TranslatedRKFStatus"`
	ConnectionGUID          string `json:"ConnectionGUID"`
	Threshold               string `json:"Threshold"`
	MappedAccess            string `json:"MappedAccess"`
	GrantedAccess           string `json:"GrantedAccess"`
	ShareSecurityDescriptor string `json:"ShareSecurityDescriptor"`
	SPN                     string `json:"SPN"`
	SpnName                 string `json:"SpnName"`
	SPNValidationPolicy     string `json:"SPNValidationPolicy"`
	ErrorCode               string `json:"ErrorCode"`
	ServerNames             string `json:"ServerNames"`
	ConfiguredNames         string `json:"ConfiguredNames"`
	Source                  string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       fmt.Sprintf("ClientAddress: %v, ClientName: %v", tmp.ClientAddress, tmp.ClientName),
			DestinationUser:  tmp.SubjectUserName,
			DestinationHost:  "",
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, FileName: %v, SPNName: %v, GrantedAccess: %v", tmp.EID, tmp.FileName, tmp.SpnName, tmp.GrantedAccess),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications struct {
	EventTime             string `json:"EventTime"`
	Computer              string `json:"Computer"`
	EID                   string `json:"EID"`
	Description           string `json:"Description"`
	Channel               string `json:"Channel"`
	Provider              string `json:"Provider"`
	SystemProcessID       string `json:"System Process ID"`
	Keywords              string `json:"Keywords"`
	SecurityUserID        string `json:"Security User ID"`
	CorrelationActivityID string `json:"Correlation ActivityID"`
	EventRecordID         string `json:"EventRecord ID"`
	SubjectUserSid        string `json:"SubjectUserSid"`
	SubjectUserName       string `json:"SubjectUserName"`
	SubjectDomainName     string `json:"SubjectDomainName"`
	SubjectLogonID        string `json:"SubjectLogonId"`
	ShareName             string `json:"ShareName"`
	ShareLocalPath        string `json:"ShareLocalPath"`
	ObjectType            string `json:"ObjectType"`
	OldRemark             string `json:"OldRemark"`
	NewRemark             string `json:"NewRemark"`
	OldMaxUsers           string `json:"OldMaxUsers"`
	NewMaxUsers           string `json:"NewMaxUsers"`
	OldShareFlags         string `json:"OldShareFlags"`
	NewShareFlags         string `json:"NewShareFlags"`
	OldSD                 string `json:"OldSD"`
	NewSD                 string `json:"NewSD"`
	Source                string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       tmp.Computer,
			DestinationUser:  tmp.SubjectUserName,
			DestinationHost:  "",
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, Keywords: %v, ShareName: %v, ShareLocalPath: %v", tmp.EID, tmp.Keywords, tmp.ShareName, tmp.ShareLocalPath),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall struct {
	EventTime                      string `json:"EventTime"`
	Computer                       string `json:"Computer"`
	EID                            string `json:"EID"`
	Description                    string `json:"Description"`
	Channel                        string `json:"Channel"`
	Provider                       string `json:"Provider"`
	SystemProcessID                string `json:"System Process ID"`
	Keywords                       string `json:"Keywords"`
	SecurityUserID                 string `json:"Security User ID"`
	CorrelationActivityID          string `json:"Correlation ActivityID"`
	EventRecordID                  string `json:"EventRecord ID"`
	RuleID                         string `json:"RuleId"`
	RuleName                       string `json:"RuleName"`
	RuleAttr                       string `json:"RuleAttr"`
	ApplicationPath                string `json:"ApplicationPath"`
	ModifyingApplication           string `json:"ModifyingApplication"`
	ModifyingUser                  string `json:"ModifyingUser"`
	ServiceName                    string `json:"ServiceName"`
	LocalSourceAddresses           string `json:"Local/SourceAddresses"`
	LocalPortsSourcePort           string `json:"LocalPorts/SourcePort"`
	RemoteDestAddresses            string `json:"Remote/DestAddresses"`
	RemotePortsDestPort            string `json:"RemotePorts/DestPort"`
	Direction                      string `json:"Direction"`
	Protocol                       string `json:"Protocol"`
	ProfilesNewProfile             string `json:"Profiles/NewProfile"`
	OldProfile                     string `json:"OldProfile"`
	Action                         string `json:"Action"`
	Active                         string `json:"Active"`
	IPVersion                      string `json:"IPVersion"`
	FilterRTID                     string `json:"FilterRTID"`
	LayerName                      string `json:"LayerName"`
	LayerRTID                      string `json:"LayerRTID"`
	ReasonCode                     string `json:"ReasonCode"`
	RemoteMachineID                string `json:"RemoteMachineID"`
	RemoteUserID                   string `json:"RemoteUserID"`
	RemoteMachineAuthorizationList string `json:"RemoteMachineAuthorizationList"`
	RemoteUserAuthorizationList    string `json:"RemoteUserAuthorizationList"`
	EmbeddedContext                string `json:"EmbeddedContext"`
	Flags                          string `json:"Flags"`
	EdgeTraversal                  string `json:"EdgeTraversal"`
	SecurityOptions                string `json:"SecurityOptions"`
	SchemaVersion                  string `json:"SchemaVersion"`
	SettingType                    string `json:"SettingType"`
	SettingValueText               string `json:"SettingValueText"`
	Origin                         string `json:"Origin"`
	ErrorCode                      string `json:"ErrorCode"`
	Reason                         string `json:"Reason"`
	ProcessID                      string `json:"ProcessId"`
	Publisher                      string `json:"Publisher"`
	CallerProcessName              string `json:"CallerProcessName"`
	InterfaceGUID                  string `json:"InterfaceGuid"`
	InterfaceName                  string `json:"InterfaceName"`
	StoreType                      string `json:"Store Type"`
	ProductName                    string `json:"ProductName"`
	Categories                     string `json:"Categories"`
	Source                         string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       fmt.Sprintf("LocalAddresses: %v, LocalPorts: %v", tmp.LocalSourceAddresses, tmp.LocalPortsSourcePort),
			DestinationUser:  "",
			DestinationHost:  fmt.Sprintf("RemoteAddresses: %v, RemotePorts: %v", tmp.RemoteDestAddresses, tmp.RemotePortsDestPort),
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, IPVersion: %v, Protocol: %v, Direction: %v", tmp.EID, tmp.IPVersion, tmp.Protocol, tmp.Direction),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_EventLogs_EvtxHussar_WinRM struct {
	EventTime                   string `json:"EventTime"`
	Computer                    string `json:"Computer"`
	EID                         string `json:"EID"`
	Description                 string `json:"Description"`
	Channel                     string `json:"Channel"`
	Provider                    string `json:"Provider"`
	SystemProcessID             string `json:"System Process ID"`
	Keywords                    string `json:"Keywords"`
	SecurityUserID              string `json:"Security User ID"`
	CorrelationActivityID       string `json:"Correlation ActivityID"`
	EventRecordID               string `json:"EventRecord ID"`
	Connection                  string `json:"connection"`
	ConnectionHostname          string `json:"connection (hostname)"`
	ConnectionPowershellVersion string `json:"connection (powershell version)"`
	Username                    string `json:"username"`
	Authentication              string `json:"authentication"`
	Destination                 string `json:"destination"`
	Errorcode                   string `json:"errorcode"`
	AuthFailureMessage          string `json:"authFailureMessage"`
	ResourceURI                 string `json:"resourceUri"`
	ShellID                     string `json:"shellId"`
	CommandID                   string `json:"commandId"`
	ApplicationID               string `json:"applicationID"`
	OperationType               string `json:"operationType"`
	NamespaceName               string `json:"namespaceName"`
	ClassName                   string `json:"className"`
	OperationName               string `json:"operationName"`
	EventPayload                string `json:"EventPayload"`
	Port                        string `json:"port"`
	Subject                     string `json:"subject"`
	AuthServer1                 string `json:"authServer1"`
	AuthServer2                 string `json:"authServer2"`
	AuthServer3                 string `json:"authServer3"`
	AuthServer4                 string `json:"authServer4"`
	AuthServer5                 string `json:"authServer5"`
	AuthProxy1                  string `json:"authProxy1"`
	AuthProxy2                  string `json:"authProxy2"`
	AuthProxy3                  string `json:"authProxy3"`
	AuthProxy4                  string `json:"authProxy4"`
	AuthProxy5                  string `json:"authProxy5"`
	Source                      string `json:"__Source"`
}

func (s Exchange_Windows_EventLogs_EvtxHussar_WinRM) StringArray() []string {
	parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", s.EventTime)
	if terr != nil {
		parsedTime = time.Now()
	}
	tmp := helpers.GetStructValuesAsStringSlice(s)
	tmp[0] = parsedTime.String()
	return tmp
}

func (s Exchange_Windows_EventLogs_EvtxHussar_WinRM) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_EventLogs_EvtxHussar_WinRM(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_EventLogs_EvtxHussar_WinRM{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		parsedTime, terr := time.Parse("2006.1.2 03:04:05.999", tmp.EventTime)
		if terr != nil {
			parsedTime = time.Now()
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime, clientIdentifier, tmp.Computer, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        parsedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Description,
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       tmp.ConnectionHostname,
			DestinationUser:  tmp.Username,
			DestinationHost:  tmp.Destination,
			SourceFile:       fmt.Sprintf("%v/%v", tmp.Provider, tmp.Channel),
			MetaData:         fmt.Sprintf("EventID: %v, EventPayload: %v, Keywords: %v", tmp.EID, tmp.EventPayload, tmp.Keywords),
		}
		outputChannel <- tmp2.StringArray()
	}
}
