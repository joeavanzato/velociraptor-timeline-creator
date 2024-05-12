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

type DetectRaptor_Windows_Detection_MFT struct {
	Detection struct {
		Name         string `json:"Name"`
		KeywordRegex string `json:"KeywordRegex"`
		PathRegex    string `json:"PathRegex"`
		IgnoreRegex  string `json:"IgnoreRegex"`
		StringHit    string `json:"StringHit"`
		Criticality  string `json:"Criticality"`
	} `json:"Detection"`
	EntryNumber  int    `json:"EntryNumber"`
	InUse        bool   `json:"InUse"`
	OSPath       string `json:"OSPath"`
	FileSize     int    `json:"FileSize"`
	IsDir        bool   `json:"IsDir"`
	SITimestamps struct {
		Created0X10          time.Time `json:"Created0x10"`
		LastModified0X10     time.Time `json:"LastModified0x10"`
		LastRecordChange0X10 time.Time `json:"LastRecordChange0x10"`
		LastAccess0X10       time.Time `json:"LastAccess0x10"`
	} `json:"SITimestamps"`
	FNTimestamps struct {
		Created0X30          time.Time `json:"Created0x30"`
		LastModified0X30     time.Time `json:"LastModified0x30"`
		LastRecordChange0X30 time.Time `json:"LastRecordChange0x30"`
		LastAccess0X30       time.Time `json:"LastAccess0x30"`
	} `json:"FNTimestamps"`
}

func (s DetectRaptor_Windows_Detection_MFT) StringArray() []string {
	return []string{s.Detection.Name, s.Detection.KeywordRegex, s.Detection.PathRegex, s.Detection.IgnoreRegex, s.Detection.StringHit, s.Detection.Criticality,
		strconv.Itoa(s.EntryNumber), strconv.FormatBool(s.InUse), s.OSPath, strconv.Itoa(s.FileSize), strconv.FormatBool(s.IsDir), s.SITimestamps.Created0X10.String(),
		s.SITimestamps.LastModified0X10.String(), s.SITimestamps.LastRecordChange0X10.String(), s.SITimestamps.LastAccess0X10.String(),
		s.FNTimestamps.Created0X30.String(), s.FNTimestamps.LastModified0X30.String(), s.FNTimestamps.LastRecordChange0X30.String(), s.FNTimestamps.LastAccess0X30.String()}
}

func (s DetectRaptor_Windows_Detection_MFT) GetHeaders() []string {
	return []string{"Detection_Name", "Detection_KeywordRegex", "Detection_PathRegex", "Detection_IgnoreRegex", "Detection_StringHit", "Detection_Criticality",
		"EntryNumber", "InUse", "OSPath", "FileSize", "IsDir", "SI_Created", "SI_LastModified", "SI_LastRecordChange", "SI_LastAccess", "FN_Created", "FN_LastModified",
		"FN_LastRecordChange", "FN_LastAccess"}
}

func Process_DetectRaptor_Windows_Detection_MFT(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_MFT{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.SITimestamps.Created0X10.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.SITimestamps.Created0X10,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Detection.Name,
			EventDescription: fmt.Sprintf("Criticality: %v, Hit: %v", tmp.Detection.Criticality, tmp.Detection.StringHit),
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("InUse: %v", tmp.InUse),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Generic_Detection_WebshellYara struct {
	OSPath string    `json:"OSPath"`
	Size   int       `json:"Size"`
	Mtime  time.Time `json:"Mtime"`
	Atime  time.Time `json:"Atime"`
	Ctime  time.Time `json:"Ctime"`
	Btime  time.Time `json:"Btime"`
	Rule   string    `json:"Rule"`
	Tags   any       `json:"Tags"`
	Meta   struct {
		Description string `json:"description"`
		License     string `json:"license"`
		Author      string `json:"author"`
		Reference   string `json:"reference"`
		Date        string `json:"date"`
	} `json:"Meta"`
	YaraString string `json:"YaraString"`
	HitOffset  int    `json:"HitOffset"`
	HitContext struct {
		Path       string   `json:"Path"`
		Size       int      `json:"Size"`
		StoredSize int      `json:"StoredSize"`
		Sha256     string   `json:"sha256"`
		Md5        string   `json:"md5"`
		StoredName string   `json:"StoredName"`
		Components []string `json:"Components"`
		Accessor   string   `json:"Accessor"`
	} `json:"HitContext"`
}

func (s DetectRaptor_Generic_Detection_WebshellYara) StringArray() []string {
	return []string{s.OSPath, strconv.Itoa(s.Size), s.Mtime.String(), s.Atime.String(), s.Ctime.String(), s.Btime.String(), s.Rule, fmt.Sprint(s.Tags), s.Meta.Description, s.Meta.License, s.Meta.Author, s.Meta.Reference, s.Meta.Date,
		s.YaraString, strconv.Itoa(s.HitOffset), s.HitContext.Path, strconv.Itoa(s.HitContext.Size), strconv.Itoa(s.HitContext.StoredSize),
		s.HitContext.Sha256, s.HitContext.Md5, s.HitContext.StoredName, fmt.Sprint(s.HitContext.Components), s.HitContext.Accessor}
}

func (s DetectRaptor_Generic_Detection_WebshellYara) GetHeaders() []string {
	return []string{"OSPath", "Size", "Mtime", "Atime", "Ctime", "Btime", "Rule", "Tags", "Meta_Description", "Meta_License", "Meta_Author",
		"Meta_Reference", "Meta_Date", "YaraString", "HitOffset", "HitContext_Path", "HitContext_Size", "HitContext_StoredSize", "HitContext_SHA256",
		"HitContext_MD5", "HitContext_StoredName", "HitContext_Components", "HitContext_Accessor"}
}

func Process_DetectRaptor_Generic_Detection_WebshellYara(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Generic_Detection_WebshellYara{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Ctime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Ctime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Rule,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("YaraString: %v, MD5: %v, MetaDescription: %v, Mtime: %v", tmp.YaraString, tmp.HitContext.Md5, tmp.Meta.Description, tmp.Mtime),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_Amcache struct {
	Detection struct {
		Name         string `json:"Name"`
		KeywordRegex string `json:"KeywordRegex"`
		PathName     string `json:"PathName"`
		Reference    string `json:"Reference"`
		Criticality  string `json:"Criticality"`
	} `json:"Detection"`
	KeyMTime         time.Time `json:"KeyMTime"`
	EntryName        string    `json:"EntryName"`
	EntryPath        string    `json:"EntryPath"`
	Publisher        string    `json:"Publisher"`
	OriginalFileName string    `json:"OriginalFileName"`
	SHA1             string    `json:"SHA1"`
}

func (s DetectRaptor_Windows_Detection_Amcache) StringArray() []string {
	return []string{s.Detection.Name, s.Detection.KeywordRegex, s.Detection.PathName, s.Detection.Reference, s.Detection.Criticality,
		s.KeyMTime.String(), s.EntryName, s.EntryPath, s.Publisher, s.OriginalFileName, s.SHA1}
}

func (s DetectRaptor_Windows_Detection_Amcache) GetHeaders() []string {
	return []string{"Detection_Name", "Detection_KeywordRegex", "Detection_PathName", "Detection_Reference", "Detection_Criticality",
		"KeyMTime", "EntryName", "EntryPath", "Publisher", "OriginalFileName", "SHA1"}
}

func Process_DetectRaptor_Windows_Detection_Amcache(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_Amcache{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.KeyMTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.KeyMTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Detection.Name,
			EventDescription: fmt.Sprintf("Criticality: %v, Hit: %v", tmp.Detection.Criticality, tmp.Detection.PathName),
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.EntryPath,
			MetaData:         fmt.Sprintf("Publisher: %v, SHA1: %v", tmp.Publisher, tmp.SHA1),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_Applications struct {
	Category              string    `json:"Category"`
	KeyName               string    `json:"KeyName"`
	KeyLastWriteTimestamp time.Time `json:"KeyLastWriteTimestamp"`
	DisplayName           string    `json:"DisplayName"`
	DisplayVersion        string    `json:"DisplayVersion"`
	InstallLocation       string    `json:"InstallLocation"`
	InstallSource         string    `json:"InstallSource"`
	Language              int       `json:"Language"`
	Publisher             string    `json:"Publisher"`
	UninstallString       string    `json:"UninstallString"`
	InstallDate           string    `json:"InstallDate"`
	KeyPath               string    `json:"KeyPath"`
	Regex                 string    `json:"Regex"`
	Comment               string    `json:"Comment"`
}

func (s DetectRaptor_Windows_Detection_Applications) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s DetectRaptor_Windows_Detection_Applications) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_DetectRaptor_Windows_Detection_Applications(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_Applications{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.KeyLastWriteTimestamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.KeyLastWriteTimestamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Category,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       "",
			MetaData:         fmt.Sprintf("InstallDate: %v, InstallSource: %v, DisplayName: %v, InstallLocation: %v", tmp.InstallDate, tmp.InstallSource, tmp.DisplayName, tmp.InstallLocation),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_BinaryRename struct {
	OSPath             string `json:"OSPath"`
	Name               string `json:"Name"`
	Size               int    `json:"Size"`
	VersionInformation struct {
		CompanyName      string `json:"CompanyName"`
		FileDescription  string `json:"FileDescription"`
		FileVersion      string `json:"FileVersion"`
		InternalName     string `json:"InternalName"`
		LegalCopyright   string `json:"LegalCopyright"`
		OriginalFilename string `json:"OriginalFilename"`
		ProductName      string `json:"ProductName"`
		ProductVersion   string `json:"ProductVersion"`
	} `json:"VersionInformation"`
	Hash struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hash"`
	Mtime time.Time `json:"Mtime"`
	Atime time.Time `json:"Atime"`
	Ctime time.Time `json:"Ctime"`
	Btime time.Time `json:"Btime"`
}

func (s DetectRaptor_Windows_Detection_BinaryRename) StringArray() []string {
	return []string{s.OSPath, s.Name, strconv.Itoa(s.Size), s.VersionInformation.CompanyName, s.VersionInformation.FileDescription, s.VersionInformation.FileVersion,
		s.VersionInformation.InternalName, s.VersionInformation.LegalCopyright, s.VersionInformation.OriginalFilename, s.VersionInformation.ProductName,
		s.VersionInformation.ProductVersion, s.Hash.MD5, s.Hash.SHA1, s.Hash.SHA256, s.Mtime.String(), s.Atime.String(), s.Ctime.String(), s.Btime.String()}
}

func (s DetectRaptor_Windows_Detection_BinaryRename) GetHeaders() []string {
	return []string{"OSPath", "Name", "Size", "Version_CompanyName", "Version_FileDescription", "Version_FileVersion", "Version_InternalName", "Version_LegalCopyright",
		"Version_OriginalFileName", "Version_ProductName", "Version_ProductVersion", "Hash_MD5", "Hash_SHA1", "Hash_SHA256", "Mtime", "Atime", "Ctime", "Btime"}
}

func Process_DetectRaptor_Windows_Detection_BinaryRename(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_BinaryRename{}
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
			EventType:        "",
			EventDescription: fmt.Sprintf("Renamed File: %v", tmp.OSPath),
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.VersionInformation.OriginalFilename,
			MetaData:         fmt.Sprintf("ProductName: %v, CompanyName: %v, SHA1: %v", tmp.VersionInformation.ProductName, tmp.VersionInformation.CompanyName, tmp.Hash.SHA1),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_Webhistory struct {
	Category        string    `json:"Category"`
	Domain          string    `json:"Domain"`
	VisitedURL      string    `json:"VisitedURL"`
	VisitTime       time.Time `json:"VisitTime"`
	Title           string    `json:"Title"`
	VisitCount      int       `json:"VisitCount"`
	BrowserArtifact string    `json:"BrowserArtifact"`
	DomainRegex     string    `json:"DomainRegex"`
	AllowRegex      string    `json:"AllowRegex"`
	Comment         string    `json:"Comment"`
}

func (s DetectRaptor_Windows_Detection_Webhistory) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s DetectRaptor_Windows_Detection_Webhistory) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_DetectRaptor_Windows_Detection_Webhistory(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_Webhistory{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.VisitTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.VisitTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Category,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.VisitedURL,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("Title: %v, Domain: %v, BrowserArtifact: %v", tmp.Title, tmp.Domain, tmp.BrowserArtifact),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_Evtx struct {
	EventTime time.Time `json:"EventTime"`
	Computer  string    `json:"Computer"`
	Detection struct {
		Name    string `json:"Name"`
		EventID string `json:"EventId"`
		Regex   string `json:"Regex"`
		Ignore  string `json:"Ignore"`
	} `json:"Detection"`
	Channel   string `json:"Channel"`
	EventID   int    `json:"EventID"`
	UserSID   string `json:"UserSID"`
	Username  string `json:"Username"`
	EventData struct {
		Status        int    `json:"Status"`
		VhdFile       string `json:"VhdFile"`
		VMID          string `json:"VmId"`
		VhdType       int    `json:"VhdType"`
		Version       int    `json:"Version"`
		Flags         int64  `json:"Flags"`
		AccessMask    int    `json:"AccessMask"`
		WriteDepth    int    `json:"WriteDepth"`
		GetInfoOnly   bool   `json:"GetInfoOnly"`
		ReadOnly      bool   `json:"ReadOnly"`
		HandleContext any    `json:"HandleContext"`
		VirtualDisk   any    `json:"VirtualDisk"`
		FileObject    any    `json:"FileObject"`
	} `json:"EventData"`
	Message string `json:"Message"`
	OSPath  string `json:"OSPath"`
}

func (s DetectRaptor_Windows_Detection_Evtx) StringArray() []string {
	return []string{s.EventTime.String(), s.Computer, s.Detection.Name, s.Detection.EventID, s.Detection.Regex, s.Detection.Ignore, s.Channel, strconv.Itoa(s.EventID), s.UserSID, s.Username,
		strconv.Itoa(s.EventData.Status), s.EventData.VhdFile, strconv.Itoa(s.EventData.VhdType), strconv.Itoa(s.EventData.Version), strconv.FormatInt(s.EventData.Flags, 10),
		strconv.Itoa(s.EventData.AccessMask), strconv.Itoa(s.EventData.WriteDepth), fmt.Sprint(s.EventData.GetInfoOnly),
		fmt.Sprint(s.EventData.ReadOnly), fmt.Sprint(s.EventData.HandleContext, 10), fmt.Sprint(s.EventData.VirtualDisk, 10), fmt.Sprint(s.EventData.FileObject, 10), s.Message, s.OSPath}
}

func (s DetectRaptor_Windows_Detection_Evtx) GetHeaders() []string {
	return []string{"EventTime", "Computer", "Detection_Name", "Detection_EventID", "Detection_Regex", "Detection_Ignore", "Channel", "EventID", "UserSID", "Username", "EventData_Status",
		"EventData_VhdFile", "EventData_VMID", "EventData_VhdType", "EventData_Version", "EventData_Flags", "EventData_AccessMask", "EventData_WriteDepth", "EventData_GetInfoOnly",
		"EventData_ReadOnly", "EventData_HandleContext", "EventData_VirtualDisk", "EventData_FileObject", "Message", "OSPath"}
}

func Process_DetectRaptor_Windows_Detection_Evtx(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_Evtx{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Detection.Name,
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       tmp.Computer,
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("EventID: %v, Message: %v", tmp.EventID, tmp.Message),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave struct {
	Detection struct {
		ID          string `json:"ID"`
		Name        string `json:"Name"`
		Regex       string `json:"Regex"`
		IgnoreRegex string `json:"IgnoreRegex"`
		HitString   string `json:"HitString"`
	} `json:"Detection"`
	FileInfo struct {
		OSPath string    `json:"OSPath"`
		Size   int       `json:"Size"`
		Mtime  time.Time `json:"Mtime"`
		Atime  time.Time `json:"Atime"`
		Ctime  time.Time `json:"Ctime"`
		Btime  time.Time `json:"Btime"`
	} `json:"FileInfo"`
	Content string `json:"Content"`
}

func (s DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave) StringArray() []string {
	return []string{s.Detection.ID, s.Detection.Name, s.Detection.Regex, s.Detection.IgnoreRegex, s.Detection.HitString, s.FileInfo.OSPath, strconv.Itoa(s.FileInfo.Size), s.FileInfo.Mtime.String(),
		s.FileInfo.Atime.String(), s.FileInfo.Ctime.String(), s.FileInfo.Btime.String(), s.Content}
}

func (s DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave) GetHeaders() []string {
	return []string{"Detection_ID", "Detection_Name", "Detection_Regex", "Detection_IgnoreRegex", "Detection_HitString", "OSPath", "Size", "Mtime", "Atime", "Ctime", "Btime", "Content"}
}

func Process_DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.FileInfo.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.FileInfo.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Detection.Name,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.FileInfo.OSPath,
			MetaData:         fmt.Sprintf("HitString: %v, Ctime: %v", tmp.Detection.HitString, tmp.FileInfo.Ctime),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_PowerShell_PSReadline struct {
	RuleID    string `json:"RuleID"`
	RuleName  string `json:"RuleName"`
	Username  string `json:"Username"`
	LineNum   int    `json:"LineNum"`
	Line      string `json:"Line"`
	RuleRegex string `json:"RuleRegex"`
	FileInfo  struct {
		FullPath string    `json:"FullPath"`
		Size     int       `json:"Size"`
		Mtime    time.Time `json:"Mtime"`
		Ctime    time.Time `json:"Ctime"`
		Btime    time.Time `json:"Btime"`
	} `json:"FileInfo"`
}

func (s DetectRaptor_Windows_Detection_PowerShell_PSReadline) StringArray() []string {
	return []string{s.RuleID, s.RuleName, s.Username, strconv.Itoa(s.LineNum), s.Line, s.RuleRegex, s.FileInfo.FullPath, strconv.Itoa(s.FileInfo.Size), s.FileInfo.Mtime.String(), s.FileInfo.Ctime.String(), s.FileInfo.Btime.String()}
}

func (s DetectRaptor_Windows_Detection_PowerShell_PSReadline) GetHeaders() []string {
	return []string{"Detection_ID", "Detection_Name", "Username", "LineNumber", "Line", "RuleRegex", "FullPath", "Size", "Mtime", "Ctime", "Atime", "Btime"}
}

func Process_DetectRaptor_Windows_Detection_PowerShell_PSReadline(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_PowerShell_PSReadline{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.FileInfo.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.FileInfo.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.RuleName,
			EventDescription: "",
			SourceUser:       tmp.Username,
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.FileInfo.FullPath,
			MetaData:         fmt.Sprintf("Regex: %v, Line: %v", tmp.RuleRegex, tmp.Line),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_ZoneIdentifier struct {
	Category         string `json:"Category"`
	OSPath           string `json:"OSPath"`
	AdsName          string `json:"AdsName"`
	Inode            string `json:"Inode"`
	Size             int    `json:"Size"`
	HostObject       string `json:"HostObject"`
	HostTimestampsSI struct {
		Mtime time.Time `json:"Mtime"`
		Atime time.Time `json:"Atime"`
		Ctime time.Time `json:"Ctime"`
		Btime time.Time `json:"Btime"`
	} `json:"HostTimestampsSI"`
	AdsContent struct {
		ZoneID    string `json:"ZoneId"`
		HostURL   string `json:"HostUrl"`
		URLDomain string `json:"UrlDomain"`
	} `json:"AdsContent"`
	DomainRegex string `json:"DomainRegex"`
	AllowRegex  string `json:"AllowRegex"`
	Comment     string `json:"Comment"`
}

func (s DetectRaptor_Windows_Detection_ZoneIdentifier) StringArray() []string {
	return []string{s.Category, s.OSPath, s.AdsName, s.Inode, strconv.Itoa(s.Size), s.HostObject, s.HostTimestampsSI.Mtime.String(),
		s.HostTimestampsSI.Atime.String(), s.HostTimestampsSI.Ctime.String(), s.HostTimestampsSI.Btime.String(),
		s.AdsContent.ZoneID, s.AdsContent.HostURL, s.AdsContent.URLDomain, s.DomainRegex, s.AllowRegex, s.Comment}
}

func (s DetectRaptor_Windows_Detection_ZoneIdentifier) GetHeaders() []string {
	return []string{"Category", "OSPath", "ADSName", "Inode", "Size", "HostObject", "SI_Mtime", "SI_Atime", "SI_Ctime", "SI_Btime", "ADS_ZoneID", "ADS_HostURL", "ADS_URLDomain", "DomainRegex", "AllowRegex", "Comment"}
}

func Process_DetectRaptor_Windows_Detection_ZoneIdentifier(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_ZoneIdentifier{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.HostTimestampsSI.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.HostTimestampsSI.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Category,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.AdsContent.HostURL,
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("ADSName: %v, Ctime: %v, ZoneID: %v, URLDomain: %v", tmp.AdsName, tmp.HostTimestampsSI.Ctime, tmp.AdsContent.ZoneID, tmp.AdsContent.URLDomain),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath struct {
	OSPath       string `json:"OSPath"`
	DllName      string `json:"DllName"`
	FileSize     int    `json:"FileSize"`
	InUse        bool   `json:"InUse"`
	SILtFN       bool   `json:"SI_Lt_FN"`
	USecZeros    bool   `json:"uSecZeros"`
	TimestampsSI struct {
		LastModified0X10     time.Time `json:"LastModified0x10"`
		LastAccess0X10       time.Time `json:"LastAccess0x10"`
		LastRecordChange0X10 time.Time `json:"LastRecordChange0x10"`
		Created0X10          time.Time `json:"Created0x10"`
	} `json:"TimestampsSI"`
	TimestampsFN struct {
		LastModified0X30     time.Time `json:"LastModified0x30"`
		LastAccess0X30       time.Time `json:"LastAccess0x30"`
		LastRecordChange0X30 time.Time `json:"LastRecordChange0x30"`
		Created0X30          time.Time `json:"Created0x30"`
	} `json:"TimestampsFN"`
	DllInfo         any    `json:"DllInfo"`
	DllAuthenticode any    `json:"DllAuthenticode"`
	DllHash         string `json:"DllHash"`
	HijackLib       struct {
		DllName          string `json:"DllName"`
		Vendor           string `json:"Vendor"`
		ExpectedLocation string `json:"ExpectedLocation"`
		ExecutablePath   string `json:"ExecutablePath"`
		Type             string `json:"Type"`
		ExecutableSHA256 string `json:"ExecutableSHA256"`
		URL              string `json:"Url"`
	} `json:"HijackLib"`
	Folder any `json:"Folder"`
}

func (s DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath) StringArray() []string {
	return []string{s.OSPath, s.DllName, strconv.Itoa(s.FileSize), strconv.FormatBool(s.InUse), strconv.FormatBool(s.SILtFN),
		strconv.FormatBool(s.USecZeros), s.TimestampsSI.LastModified0X10.String(), s.TimestampsSI.LastAccess0X10.String(), s.TimestampsSI.LastRecordChange0X10.String(),
		s.TimestampsSI.Created0X10.String(), s.TimestampsFN.LastModified0X30.String(), s.TimestampsFN.LastAccess0X30.String(),
		s.TimestampsFN.LastRecordChange0X30.String(), s.TimestampsFN.Created0X30.String(),
		fmt.Sprint(s.DllInfo), fmt.Sprint(s.DllAuthenticode), s.DllHash, s.HijackLib.DllName, s.HijackLib.Vendor, s.HijackLib.ExpectedLocation,
		s.HijackLib.ExecutablePath, s.HijackLib.Type, s.HijackLib.ExecutableSHA256, s.HijackLib.URL, fmt.Sprint(s.Folder)}
}

func (s DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath) GetHeaders() []string {
	return []string{"OSPath", "DllName", "FileSize", "InUse", "SILtFN", "USecZeros", "SI_LastModified", "SI_LastAccess", "SI_LastRecordChange",
		"SI_Created", "FN_LastModified", "FN_LastAccess", "FN_LastRecordChange", "FN_Created", "DLLInfo", "DLLAuthenticode", "DllHash", "HijackLib_DllName", "HijackLib_Vendor",
		"HijackLib_ExpectedLocation", "HijackLib_ExecutablePath", "HijackLib_Type", "HijackLib_ExecutableSHA256", "HijackLib_URL", "Folder"}
}

func Process_DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.TimestampsSI.LastModified0X10.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimestampsSI.LastModified0X10,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Suspicious DLL Location",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("SI_Created: %v, DllInfo: %v, ExeSHA256: %v, Expected Location: %v, Exe Path: %v", tmp.TimestampsSI.Created0X10, tmp.DllInfo, tmp.HijackLib.ExecutableSHA256, tmp.HijackLib.ExpectedLocation, tmp.HijackLib.ExecutablePath),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_HijackLibsMFT struct {
	HijackLibInfo struct {
		DllName          string `json:"DllName"`
		Vendor           string `json:"Vendor"`
		ExpectedLocation string `json:"ExpectedLocation"`
		ExecutablePath   string `json:"ExecutablePath"`
		Type             string `json:"Type"`
		ExecutableSHA256 string `json:"ExecutableSHA256"`
		URL              string `json:"Url"`
	} `json:"HijackLibInfo"`
	OSPath            string `json:"OSPath"`
	FileName          string `json:"FileName"`
	FileSize          int    `json:"FileSize"`
	ReferenceCount    int    `json:"ReferenceCount"`
	SILtFN            bool   `json:"SI_Lt_FN"`
	USecZeros         bool   `json:"uSecZeros"`
	EntryNumber       int    `json:"EntryNumber"`
	InUse             bool   `json:"InUse"`
	ParentEntryNumber int    `json:"ParentEntryNumber"`
	TimestampsSI      struct {
		LastModified0X10     time.Time `json:"LastModified0x10"`
		LastAccess0X10       time.Time `json:"LastAccess0x10"`
		LastRecordChange0X10 time.Time `json:"LastRecordChange0x10"`
		Created0X10          time.Time `json:"Created0x10"`
	} `json:"TimestampsSI"`
	TimestampsFN struct {
		LastModified0X30     time.Time `json:"LastModified0x30"`
		LastAccess0X30       time.Time `json:"LastAccess0x30"`
		LastRecordChange0X30 time.Time `json:"LastRecordChange0x30"`
		Created0X30          time.Time `json:"Created0x30"`
	} `json:"TimestampsFN"`
	DllInfo struct {
		FileHeader struct {
			Machine          string    `json:"Machine"`
			TimeDateStamp    time.Time `json:"TimeDateStamp"`
			TimeDateStampRaw int       `json:"TimeDateStampRaw"`
			Characteristics  int       `json:"Characteristics"`
			ImageBase        int       `json:"ImageBase"`
		} `json:"FileHeader"`
		GUIDAge     string `json:"GUIDAge"`
		PDB         string `json:"PDB"`
		Directories struct {
			BaseRelocationDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Base_Relocation_Directory"`
			DebugDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Debug_Directory"`
			ExportDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Export_Directory"`
			IATDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"IAT_Directory"`
			ImportDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Import_Directory"`
			LoadConfigDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Load_Config_Directory"`
			ResourceDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int64     `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Resource_Directory"`
			SecurityDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Security_Directory"`
		} `json:"Directories"`
		Sections []struct {
			Perm       string `json:"Perm"`
			Name       string `json:"Name"`
			FileOffset int    `json:"FileOffset"`
			VMA        int    `json:"VMA"`
			RVA        int    `json:"RVA"`
			Size       int    `json:"Size"`
		} `json:"Sections"`
		Resources []struct {
			Type       string `json:"Type"`
			TypeID     int    `json:"TypeId"`
			FileOffset int    `json:"FileOffset"`
			DataSize   int    `json:"DataSize"`
			CodePage   int    `json:"CodePage"`
		} `json:"Resources"`
		VersionInformation struct {
			CompanyName      string `json:"CompanyName"`
			FileDescription  string `json:"FileDescription"`
			FileVersion      string `json:"FileVersion"`
			InternalName     string `json:"InternalName"`
			LegalCopyright   string `json:"LegalCopyright"`
			OriginalFilename string `json:"OriginalFilename"`
			ProductName      string `json:"ProductName"`
			ProductVersion   string `json:"ProductVersion"`
		} `json:"VersionInformation"`
		Imports      []string `json:"Imports"`
		Exports      []string `json:"Exports"`
		Forwards     []string `json:"Forwards"`
		ImpHash      string   `json:"ImpHash"`
		Authenticode struct {
			Signer struct {
				IssuerName              string `json:"IssuerName"`
				SerialNumber            string `json:"SerialNumber"`
				DigestAlgorithm         string `json:"DigestAlgorithm"`
				AuthenticatedAttributes struct {
					ContentType      string `json:"ContentType"`
					MessageDigest    string `json:"MessageDigest"`
					MessageDigestHex string `json:"MessageDigestHex"`
					ProgramName      string `json:"ProgramName"`
					MoreInfo         string `json:"MoreInfo"`
				} `json:"AuthenticatedAttributes"`
				UnauthenticatedAttributes any    `json:"UnauthenticatedAttributes"`
				Subject                   string `json:"Subject"`
			} `json:"Signer"`
			Certificates []struct {
				SerialNumber       string    `json:"SerialNumber"`
				SignatureAlgorithm string    `json:"SignatureAlgorithm"`
				Subject            string    `json:"Subject"`
				Issuer             string    `json:"Issuer"`
				NotBefore          time.Time `json:"NotBefore"`
				NotAfter           time.Time `json:"NotAfter"`
				PublicKey          string    `json:"PublicKey"`
				Extensions         struct {
					BasicConstraints struct {
						Critical   bool `json:"Critical"`
						IsCA       bool `json:"IsCA"`
						MaxPathLen int  `json:"MaxPathLen"`
					} `json:"BasicConstraints"`
					SubjectKeyID struct {
						Critical bool   `json:"Critical"`
						Value    string `json:"Value"`
					} `json:"SubjectKeyId"`
					AuthorityKeyIdentifier struct {
						Critical bool   `json:"Critical"`
						KeyID    string `json:"KeyId"`
					} `json:"AuthorityKeyIdentifier"`
					KeyUsage struct {
						Critical bool     `json:"Critical"`
						KeyUsage []string `json:"KeyUsage"`
					} `json:"KeyUsage"`
					ExtendedKeyUsage struct {
						Critical bool     `json:"Critical"`
						KeyUsage []string `json:"KeyUsage"`
					} `json:"Extended Key Usage"`
					CRLDistributionPoints struct {
						Critical bool     `json:"Critical"`
						URI      []string `json:"URI"`
					} `json:"CRLDistributionPoints"`
					CertificatePolicies struct {
						Critical bool     `json:"Critical"`
						Policy   []string `json:"Policy"`
					} `json:"CertificatePolicies"`
				} `json:"Extensions"`
			} `json:"Certificates"`
			HashType        string `json:"HashType"`
			ExpectedHash    string `json:"ExpectedHash"`
			ExpectedHashHex string `json:"ExpectedHashHex"`
		} `json:"Authenticode"`
		AuthenticodeHash struct {
			MD5         string `json:"MD5"`
			SHA1        string `json:"SHA1"`
			SHA256      string `json:"SHA256"`
			HashMatches bool   `json:"HashMatches"`
		} `json:"AuthenticodeHash"`
	} `json:"DllInfo"`
	DllAuthenticode struct {
		Filename      string `json:"Filename"`
		ProgramName   string `json:"ProgramName"`
		PublisherLink string `json:"PublisherLink"`
		MoreInfoLink  string `json:"MoreInfoLink"`
		SerialNumber  string `json:"SerialNumber"`
		IssuerName    string `json:"IssuerName"`
		SubjectName   string `json:"SubjectName"`
		Timestamp     string `json:"Timestamp"`
		Trusted       string `json:"Trusted"`
		ExtraInfo     any    `json:"_ExtraInfo"`
	} `json:"DllAuthenticode"`
	DllHash struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"DllHash"`
	Folder any `json:"Folder"`
}

func (s DetectRaptor_Windows_Detection_HijackLibsMFT) StringArray() []string {
	return []string{s.HijackLibInfo.DllName, s.HijackLibInfo.Vendor, s.HijackLibInfo.ExpectedLocation,
		s.HijackLibInfo.ExecutablePath, s.HijackLibInfo.Type, s.HijackLibInfo.ExecutableSHA256, s.HijackLibInfo.URL,
		s.OSPath, s.FileName, strconv.Itoa(s.FileSize), strconv.Itoa(s.ReferenceCount), strconv.FormatBool(s.SILtFN),
		strconv.FormatBool(s.USecZeros), strconv.Itoa(s.EntryNumber), strconv.FormatBool(s.InUse), strconv.Itoa(s.ParentEntryNumber),
		s.TimestampsSI.LastModified0X10.String(), s.TimestampsSI.LastAccess0X10.String(),
		s.TimestampsSI.LastRecordChange0X10.String(), s.TimestampsSI.Created0X10.String(),
		s.TimestampsFN.LastModified0X30.String(), s.TimestampsFN.LastAccess0X30.String(),
		s.TimestampsFN.LastRecordChange0X30.String(), s.TimestampsFN.Created0X30.String(),
		s.DllInfo.FileHeader.Machine, s.DllInfo.FileHeader.TimeDateStamp.String(), strconv.Itoa(s.DllInfo.FileHeader.TimeDateStampRaw),
		strconv.Itoa(s.DllInfo.FileHeader.Characteristics), strconv.Itoa(s.DllInfo.FileHeader.ImageBase),
		s.DllInfo.GUIDAge, s.DllInfo.PDB, s.DllInfo.Directories.BaseRelocationDirectory.Timestamp.String(),
		strconv.Itoa(s.DllInfo.Directories.BaseRelocationDirectory.TimestampRaw), strconv.Itoa(s.DllInfo.Directories.BaseRelocationDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.BaseRelocationDirectory.FileAddress), s.DllInfo.Directories.BaseRelocationDirectory.SectionName,
		s.DllInfo.Directories.DebugDirectory.Timestamp.String(),
		strconv.Itoa(s.DllInfo.Directories.DebugDirectory.TimestampRaw), strconv.Itoa(s.DllInfo.Directories.DebugDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.DebugDirectory.FileAddress), s.DllInfo.Directories.DebugDirectory.SectionName,
		s.DllInfo.Directories.ExportDirectory.Timestamp.String(),
		strconv.Itoa(s.DllInfo.Directories.ExportDirectory.TimestampRaw), strconv.Itoa(s.DllInfo.Directories.ExportDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.ExportDirectory.FileAddress), s.DllInfo.Directories.ExportDirectory.SectionName,
		s.DllInfo.Directories.IATDirectory.Timestamp.String(),
		strconv.Itoa(s.DllInfo.Directories.IATDirectory.TimestampRaw), strconv.Itoa(s.DllInfo.Directories.IATDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.IATDirectory.FileAddress), s.DllInfo.Directories.IATDirectory.SectionName,
		s.DllInfo.Directories.ImportDirectory.Timestamp.String(),
		strconv.Itoa(s.DllInfo.Directories.ImportDirectory.TimestampRaw), strconv.Itoa(s.DllInfo.Directories.ImportDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.ImportDirectory.FileAddress), s.DllInfo.Directories.ImportDirectory.SectionName,
		s.DllInfo.Directories.LoadConfigDirectory.Timestamp.String(),
		strconv.Itoa(s.DllInfo.Directories.LoadConfigDirectory.TimestampRaw), strconv.Itoa(s.DllInfo.Directories.LoadConfigDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.LoadConfigDirectory.FileAddress), s.DllInfo.Directories.LoadConfigDirectory.SectionName,
		s.DllInfo.Directories.ResourceDirectory.Timestamp.String(),
		strconv.Itoa(int(s.DllInfo.Directories.ResourceDirectory.TimestampRaw)), strconv.Itoa(s.DllInfo.Directories.ResourceDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.ResourceDirectory.FileAddress), s.DllInfo.Directories.ResourceDirectory.SectionName,
		s.DllInfo.Directories.SecurityDirectory.Timestamp.String(),
		strconv.Itoa(s.DllInfo.Directories.SecurityDirectory.TimestampRaw), strconv.Itoa(s.DllInfo.Directories.SecurityDirectory.Size),
		strconv.Itoa(s.DllInfo.Directories.SecurityDirectory.FileAddress), s.DllInfo.Directories.SecurityDirectory.SectionName,
		s.DllInfo.VersionInformation.CompanyName, s.DllInfo.VersionInformation.FileDescription, s.DllInfo.VersionInformation.FileVersion,
		s.DllInfo.VersionInformation.InternalName, s.DllInfo.VersionInformation.LegalCopyright, s.DllInfo.VersionInformation.OriginalFilename,
		s.DllInfo.VersionInformation.ProductName, s.DllInfo.VersionInformation.ProductVersion, fmt.Sprint(s.DllInfo.Imports), fmt.Sprint(s.DllInfo.Exports),
		fmt.Sprint(s.DllInfo.Forwards), s.DllInfo.Authenticode.Signer.IssuerName, s.DllInfo.Authenticode.Signer.SerialNumber, s.DllInfo.Authenticode.Signer.DigestAlgorithm,
		s.DllInfo.Authenticode.Signer.AuthenticatedAttributes.ContentType, s.DllInfo.Authenticode.Signer.AuthenticatedAttributes.MessageDigest, s.DllInfo.Authenticode.Signer.AuthenticatedAttributes.MessageDigestHex,
		s.DllInfo.Authenticode.Signer.AuthenticatedAttributes.ProgramName, s.DllInfo.Authenticode.Signer.AuthenticatedAttributes.MoreInfo,
		fmt.Sprint(s.DllInfo.Authenticode.Signer.UnauthenticatedAttributes), s.DllInfo.Authenticode.Signer.Subject,
		fmt.Sprint(s.DllInfo.Authenticode.Certificates), s.DllInfo.Authenticode.HashType, s.DllInfo.Authenticode.ExpectedHash, s.DllInfo.Authenticode.ExpectedHashHex,
		s.DllInfo.AuthenticodeHash.MD5, s.DllInfo.AuthenticodeHash.SHA1, s.DllInfo.AuthenticodeHash.SHA256, strconv.FormatBool(s.DllInfo.AuthenticodeHash.HashMatches),
		s.DllAuthenticode.Filename, s.DllAuthenticode.ProgramName, s.DllAuthenticode.PublisherLink, s.DllAuthenticode.MoreInfoLink, s.DllAuthenticode.SerialNumber,
		s.DllAuthenticode.IssuerName, s.DllAuthenticode.SubjectName, s.DllAuthenticode.Timestamp, s.DllAuthenticode.Trusted, fmt.Sprint(s.DllAuthenticode.ExtraInfo),
		s.DllHash.MD5, s.DllHash.SHA1, s.DllHash.SHA256, fmt.Sprint(s.Folder)}
}

func (s DetectRaptor_Windows_Detection_HijackLibsMFT) GetHeaders() []string {
	return []string{"HijackLib_DllName", "HijackLib_Vendor", "HijackLib_ExpectedLocation", "HijackLib_ExecutablePath",
		"HijackLib_Type", "HijackLib_ExecutableSHA256", "HijackLib_URL", "OSPath", "FileName", "FileSize", "ReferenceCount", "SILtFN",
		"USecZeros", "EntryNumber", "InUse", "ParentEntryNumber", "SI_LastModified", "SI_LastAccess", "SI_LastRecordChange",
		"SI_Created", "FN_LastModified", "FN_LastAccess", "FN_LastRecordChange", "FN_Created", "FileHeader_Machine",
		"FileHeader_TimeDateStamp", "FileHeader_TimeDateStampRaw", "FileHeader_Characteristics", "FileHeader_ImageBase",
		"GUIDAge", "PDB",
		"BaseRelocationDir_Timestamp", "BaseRelocationDir_TimestampRaw", "BaseRelocationDir_Size", "BaseRelocationDir_FileAddress",
		"BaseRelocationDir_SectionName",
		"DebugDirectory_Timestamp", "DebugDirectory_TimestampRaw", "DebugDirectory_Size", "DebugDirectory_FileAddress",
		"DebugDirectory_SectionName",
		"ExportDirectory_Timestamp", "ExportDirectory_TimestampRaw", "ExportDirectory_Size", "ExportDirectory_FileAddress",
		"ExportDirectory_SectionName",
		"IATDirectory_Timestamp", "IATDirectory_TimestampRaw", "IATDirectory_Size", "IATDirectory_FileAddress",
		"IATDirectory_SectionName",
		"ImportDirectory_Timestamp", "ImportDirectory_TimestampRaw", "ImportDirectory_Size", "ImportDirectory_FileAddress",
		"ImportDirectory_SectionName",
		"LoadConfigDirectory_Timestamp", "LoadConfigDirectory_TimestampRaw", "LoadConfigDirectory_Size", "LoadConfigDirectory_FileAddress",
		"LoadConfigDirectory_SectionName",
		"ResourceDirectory_Timestamp", "ResourceDirectory_TimestampRaw", "ResourceDirectory_Size", "ResourceDirectory_FileAddress",
		"ResourceDirectory_SectionName",
		"SecurityDirectory_Timestamp", "SecurityDirectory_TimestampRaw", "SecurityDirectory_Size", "SecurityDirectory_FileAddress",
		"SecurityDirectory_SectionName",
		"Version_CompanyName", "Version_FileDescription", "Version_FileVersion", "Version_InternalName", "Version_LegalCopyright",
		"Version_OriginalFileName", "Version_ProductName", "Version_ProductVersion",
		"Imports", "Exports", "Forwards", "Authenticode_AA_IssuerName", "Authenticode_AA_SerialNumber", "Authenticode_AA_DigestAlgorithm",
		"Authenticode_AA_ContentType", "Authenticode_AA_MessageDigest", "Authenticode_AA_MessageDigestHex", "Authenticode_AA_ProgramName", "Authenticode_MoreInfo",
		"Authenticode_UnauthenticatedAttributes", "Authenticode_Subject", "Authenticode_Certificates", "Authenticde_HashType", "Authenticode_ExpectedHash",
		"Authenticode_ExpectedHashHex", "Authenticode_MD5", "Authenticode_SHA1", "Authenticode_SHA256", "Authenticode_HashMatches",
		"Authenticode_Filename", "Authenticode_ProgramName", "Authenticode_PublisherLink", "Authenticode_MoreInfoLink", "Authenticode_IssuerName",
		"Authenticode_SubjectName", "Authenticode_Timestamp", "Authenticode_Trusted", "Authenticode_ExtraInfo", "Hash_MD5", "Hash_SHA1", "Hash_SHA256", "Folder"}
}

func Process_DetectRaptor_Windows_Detection_HijackLibsMFT(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_HijackLibsMFT{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.TimestampsSI.LastModified0X10.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.TimestampsSI.LastModified0X10,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Suspicious DLL Location",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("SI_Created: %v, DllInfo: %v, SHA256: %v, Expected Location: %v, Exe Path: %v", tmp.TimestampsSI.Created0X10, tmp.DllInfo, tmp.HijackLibInfo.ExecutableSHA256, tmp.HijackLibInfo.ExpectedLocation, tmp.HijackLibInfo.ExecutablePath),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_LolDriversVulnerable struct {
	Name           string    `json:"Name"`
	SHA1           string    `json:"SHA1"`
	HivePath       string    `json:"HivePath"`
	EntryKey       string    `json:"EntryKey"`
	KeyMTime       time.Time `json:"KeyMTime"`
	EntryType      string    `json:"EntryType"`
	Product        string    `json:"Product"`
	Description    string    `json:"Description"`
	ProductVersion string    `json:"ProductVersion"`
	FileVersion    string    `json:"FileVersion"`
	MachineType    string    `json:"MachineType"`
	Category       string    `json:"Category"`
	Usecase        string    `json:"Usecase"`
	LolDriversURL  string    `json:"LolDriversUrl"`
}

func (s DetectRaptor_Windows_Detection_LolDriversVulnerable) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s DetectRaptor_Windows_Detection_LolDriversVulnerable) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_DetectRaptor_Windows_Detection_LolDriversVulnerable(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_LolDriversVulnerable{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.KeyMTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.KeyMTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Category,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Name,
			MetaData:         fmt.Sprintf("LolDriversUrl: %v, Usecase: %v, SHA1: %v", tmp.LolDriversURL, tmp.Usecase, tmp.SHA1),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type DetectRaptor_Windows_Detection_NamedPipes struct {
	EventTime time.Time `json:"EventTime"`
	Detection string    `json:"Detection"`
	ProcPid   int       `json:"ProcPid"`
	ProcName  string    `json:"ProcName"`
	Exe       string    `json:"Exe"`
	PipeName  string    `json:"PipeName"`
	Type      string    `json:"Type"`
	Regex     struct {
		PipeRegex      string `json:"PipeRegex"`
		IgnoreExeRegex string `json:"IgnoreExeRegex"`
	} `json:"Regex"`
	Reference string `json:"Reference"`
}

func (s DetectRaptor_Windows_Detection_NamedPipes) StringArray() []string {
	return []string{s.EventTime.String(), s.Detection, strconv.Itoa(s.ProcPid), s.ProcName, s.Exe, s.PipeName, s.Type, s.Regex.PipeRegex, s.Regex.IgnoreExeRegex, s.Reference}
}

func (s DetectRaptor_Windows_Detection_NamedPipes) GetHeaders() []string {
	return []string{"EventTime", "Detection", "ProcPid", "ProcName", "Exe", "PipeName", "Type", "PipeRegex", "IgnoreExeRegex", "Reference"}
}

func Process_DetectRaptor_Windows_Detection_NamedPipes(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := DetectRaptor_Windows_Detection_NamedPipes{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.EventTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.EventTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        tmp.Detection,
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Exe,
			MetaData:         fmt.Sprintf("PipeName: %v, ProcName: %v, PipeRegex: %v, Reference: %v, Type: %v, ProcPid: %v", tmp.PipeName, tmp.ProcName, tmp.Regex.PipeRegex, tmp.Reference, tmp.Type, tmp.ProcPid),
		}
		outputChannel <- tmp2.StringArray()
	}
}
