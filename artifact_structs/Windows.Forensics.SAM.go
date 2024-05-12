package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Windows_Forensics_SAM_CreateTimes struct {
	Username    string    `json:"Username"`
	CreatedTime time.Time `json:"CreatedTime"`
}

func (s Windows_Forensics_SAM_CreateTimes) StringArray() []string {
	return []string{s.Username, s.CreatedTime.String()}
}

func (s Windows_Forensics_SAM_CreateTimes) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Windows_Forensics_SAM_CreateTimes(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_SAM_CreateTimes{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.CreatedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.CreatedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "User Account Created (SAM)",
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf(""),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Windows_Forensics_SAM_Parsed struct {
	Key                     string `json:"Key"`
	Hive                    string `json:"Hive"`
	F                       string `json:"_F"`
	V                       string `json:"_V"`
	SupplementalCredentials string `json:"_SupplementalCredentials"`
	ParsedF                 struct {
		LastLoginDate     time.Time `json:"LastLoginDate"`
		PasswordResetDate time.Time `json:"PasswordResetDate"`
		PasswordFailDate  time.Time `json:"PasswordFailDate"`
		RID               int       `json:"RID"`
		Flags             []string  `json:"Flags"`
		FailedLoginCount  int       `json:"FailedLoginCount"`
		LoginCount        int       `json:"LoginCount"`
	} `json:"ParsedF"`
	ParsedV struct {
		AccountType string `json:"AccountType"`
		Username    string `json:"username"`
		Fullname    string `json:"fullname"`
		Comment     string `json:"comment"`
		Driveletter string `json:"driveletter"`
		LogonScript string `json:"logon_script"`
		ProfilePath string `json:"profile_path"`
		Workstation string `json:"workstation"`
		LmpwdHash   string `json:"lmpwd_hash"`
		NtpwdHash   string `json:"ntpwd_hash"`
	} `json:"ParsedV"`
}

func (s Windows_Forensics_SAM_Parsed) StringArray() []string {
	return []string{s.Key, s.Hive, s.F, s.V, s.SupplementalCredentials, s.ParsedF.LastLoginDate.String(), s.ParsedF.PasswordResetDate.String(),
		s.ParsedF.PasswordFailDate.String(), strconv.Itoa(s.ParsedF.RID), fmt.Sprint(s.ParsedF.Flags), strconv.Itoa(s.ParsedF.FailedLoginCount),
		strconv.Itoa(s.ParsedF.LoginCount), s.ParsedV.AccountType, s.ParsedV.Username, s.ParsedV.Fullname, s.ParsedV.Comment,
		s.ParsedV.Driveletter, s.ParsedV.LogonScript, s.ParsedV.ProfilePath, s.ParsedV.Workstation, s.ParsedV.LmpwdHash, s.ParsedV.NtpwdHash}
}

func (s Windows_Forensics_SAM_Parsed) GetHeaders() []string {
	return []string{"Key", "Hive", "F", "V", "SupplementalCredentials", "ParsedF_LastLogonDate", "ParsedF_PasswordResetDate", "ParsedF_PasswordFailDate", "ParsedF_RID", "ParsedF_Flags", "ParsedF_FailedLoginCount", "ParsedF_LoginCount",
		"ParsedV_AccountType", "ParsedV_UserName", "ParsedV_Fullname", "ParsedV_Comment", "ParsedV_Driveletter", "ParsedV_LogonScript", "ParsedV_ProfilePath", "ParsedV_Workstation", "ParsedV_LmpwdHash", "ParsedV_NtpwdHash"}
}

func Process_Windows_Forensics_SAM_Parsed(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_SAM_Parsed{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.ParsedF.LastLoginDate.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.ParsedF.LastLoginDate,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Last User Login (SAM)",
			EventDescription: "",
			SourceUser:       tmp.ParsedV.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("AccountType: %v, RID: %v, LoginCount: %v, FailedLoginCount: %v, Password Reset Date: %v", tmp.ParsedV.AccountType, tmp.ParsedF.RID, tmp.ParsedF.LoginCount, tmp.ParsedF.FailedLoginCount, tmp.ParsedF.PasswordResetDate),
		}
		outputChannel <- tmp2.StringArray()
	}
}
