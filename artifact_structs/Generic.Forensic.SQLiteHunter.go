package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks struct {
	OSPath       string    `json:"OSPath"`
	Name         string    `json:"Name"`
	DateAdded    time.Time `json:"DateAdded"`
	DateLastUsed time.Time `json:"DateLastUsed"`
	URL          string    `json:"URL"`
	Type         string    `json:"Type"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks) StringArray() []string {
	return []string{s.OSPath, s.Name, s.DateAdded.String(), s.DateLastUsed.String(), s.URL, s.Type}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.DateAdded.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.DateAdded,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Chromium Bookmark Added",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URL,
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Name: %v, Last Used: %v, Type: %v", tmp.Name, tmp.DateLastUsed, tmp.Type),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons struct {
	ID          int       `json:"ID"`
	IconID      int       `json:"IconID"`
	LastUpdated time.Time `json:"LastUpdated"`
	PageURL     string    `json:"PageURL"`
	FaviconURL  string    `json:"FaviconURL"`
	Image       struct {
		Path       string   `json:"Path"`
		Size       int      `json:"Size"`
		StoredSize int      `json:"StoredSize"`
		Sha256     string   `json:"sha256"`
		Md5        string   `json:"md5"`
		StoredName string   `json:"StoredName"`
		Components []string `json:"Components"`
		Accessor   string   `json:"Accessor"`
	} `json:"Image"`
	OSPath string `json:"_OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons) StringArray() []string {
	return []string{strconv.Itoa(s.ID), strconv.Itoa(s.IconID), s.LastUpdated.String(), s.PageURL, s.FaviconURL, s.Image.Path, strconv.Itoa(s.Image.Size), strconv.Itoa(s.Image.StoredSize), s.Image.Sha256, s.Image.Md5, s.Image.StoredName, fmt.Sprint(s.Image.Components), s.Image.Accessor, s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons) GetHeaders() []string {
	return []string{"ID", "IconID", "LastUpdated", "PageURL", "FaviconURL", "Image_Path", "Image_Size", "Image_StoredSize", "Image_SHA256", "Image_MD5", "Image_StoredName", "Image_Components", "Image_Accessor", "OSPath"}
}

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastUpdated.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastUpdated,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Chromium Favicon Updated",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.PageURL,
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Favicon URL: %v", tmp.FaviconURL),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads struct {
	ID               int       `json:"ID"`
	GUID             string    `json:"GUID"`
	CurrentPath      string    `json:"CurrentPath"`
	TargetPath       string    `json:"TargetPath"`
	OriginalMIMEType string    `json:"OriginalMIMEType"`
	ReceivedBytes    int       `json:"ReceivedBytes"`
	TotalBytes       int       `json:"TotalBytes"`
	StartTime        time.Time `json:"StartTime"`
	EndTime          time.Time `json:"EndTime"`
	Opened           time.Time `json:"Opened"`
	LastAccessTime   time.Time `json:"LastAccessTime"`
	LastModified     time.Time `json:"LastModified"`
	State            string    `json:"State"`
	DangerType       string    `json:"DangerType"`
	InterruptReason  string    `json:"InterruptReason"`
	ReferrerURL      string    `json:"ReferrerURL"`
	SiteURL          string    `json:"SiteURL"`
	TabURL           string    `json:"TabURL"`
	TabReferrerURL   string    `json:"TabReferrerURL"`
	DownloadURL      string    `json:"DownloadURL"`
	OSPath           string    `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads) StringArray() []string {
	return []string{strconv.Itoa(s.ID), s.GUID, s.CurrentPath, s.TargetPath, s.OriginalMIMEType, strconv.Itoa(s.ReceivedBytes),
		strconv.Itoa(s.TotalBytes), s.StartTime.String(), s.EndTime.String(), s.Opened.String(), s.LastAccessTime.String(),
		s.LastModified.String(), s.State, s.DangerType, s.InterruptReason, s.ReferrerURL, s.SiteURL, s.TabURL,
		s.TabReferrerURL, s.DownloadURL, s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.StartTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.StartTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Chromium Download Started",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.DownloadURL,
			SourceFile:       tmp.TargetPath,
			MetaData:         fmt.Sprintf("Total Bytes: %v, State: %v, Danger Type: %v, Referer: %v", tmp.TotalBytes, tmp.State, tmp.DangerType, tmp.ReferrerURL),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords struct {
	KeywordID         int       `json:"KeywordID"`
	URLID             int       `json:"URLID"`
	LastVisitedTime   time.Time `json:"LastVisitedTime"`
	KeywordSearchTerm string    `json:"KeywordSearchTerm"`
	Title             string    `json:"Title"`
	URL               string    `json:"URL"`
	OSPath            string    `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords) StringArray() []string {
	return []string{strconv.Itoa(s.KeywordID), strconv.Itoa(s.URLID), s.LastVisitedTime.String(), s.KeywordSearchTerm, s.Title, s.URL, s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastVisitedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastVisitedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Chromium Keyword Last Visit",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URL,
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Search Term: %v, Title: %v", tmp.KeywordSearchTerm, tmp.Title),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits struct {
	ID                     int       `json:"ID"`
	VisitTime              time.Time `json:"VisitTime"`
	LastVisitedTime        time.Time `json:"LastVisitedTime"`
	URLTitle               string    `json:"URLTitle"`
	URL                    string    `json:"URL"`
	VisitCount             int       `json:"VisitCount"`
	TypedCount             int       `json:"TypedCount"`
	Hidden                 string    `json:"Hidden"`
	VisitID                int       `json:"VisitID"`
	FromVisitID            int       `json:"FromVisitID"`
	VisitDurationInSeconds float64   `json:"VisitDurationInSeconds"`
	OSPath                 string    `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits) StringArray() []string {
	return []string{strconv.Itoa(s.ID), s.VisitTime.String(), s.LastVisitedTime.String(), s.URLTitle, s.URL, strconv.Itoa(s.VisitCount), strconv.Itoa(s.TypedCount), s.Hidden, strconv.Itoa(s.VisitID), strconv.Itoa(s.FromVisitID), fmt.Sprint(s.VisitDurationInSeconds), s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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
			EventType:        "Chromium History Visit",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URL,
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Title: %v, Visit Count: %v, Visit Duration: %v, Hidden: %v", tmp.URLTitle, tmp.VisitCount, tmp.VisitDurationInSeconds, tmp.Hidden),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts struct {
	ID                  string    `json:"ID"`
	LastAccessTime      time.Time `json:"LastAccessTime"`
	TextTyped           string    `json:"TextTyped"`
	FillIntoEdit        string    `json:"FillIntoEdit"`
	URL                 string    `json:"URL"`
	Contents            string    `json:"Contents"`
	Description         string    `json:"Description"`
	Type                int       `json:"Type"`
	Keyword             string    `json:"Keyword"`
	TimesSelectedByUser int       `json:"TimesSelectedByUser"`
	OSPath              string    `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts) StringArray() []string {
	return []string{s.ID, s.LastAccessTime.String(), s.TextTyped, s.FillIntoEdit, s.URL, s.Contents, s.Description, strconv.Itoa(s.Type), s.Keyword, strconv.Itoa(s.TimesSelectedByUser), s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastAccessTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastAccessTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Chromium Shortcut Accessed",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URL,
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Text Typed: %v, Fill Content: %v, Times Selected: %v, Keyword: %v", tmp.TextTyped, tmp.FillIntoEdit, tmp.TimesSelectedByUser, tmp.Keyword),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Firefox_Cookies struct {
	ID               int       `json:"ID"`
	Host             string    `json:"Host"`
	Name             string    `json:"Name"`
	Value            string    `json:"Value"`
	CreationTime     time.Time `json:"CreationTime"`
	LastAccessedTime time.Time `json:"LastAccessedTime"`
	Expiration       time.Time `json:"Expiration"`
	IsSecure         string    `json:"IsSecure"`
	IsHTTPOnly       string    `json:"IsHTTPOnly"`
	OSPath           string    `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Firefox_Cookies) StringArray() []string {
	return []string{strconv.Itoa(s.ID), s.Host, s.Name, s.Value, s.CreationTime.String(), s.LastAccessedTime.String(), s.Expiration.String(), s.IsSecure, s.IsHTTPOnly, s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Firefox_Cookies) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Firefox_Cookies(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Firefox_Cookies{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastAccessedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastAccessedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Cookie Accessed (Firefox)",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.Host,
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Created: %v, Is Secure: %v, Is HTTP Only: %v, Name: %v", tmp.CreationTime, tmp.IsSecure, tmp.IsHTTPOnly, tmp.Name),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Firefox_Form_History struct {
	ID        int       `json:"ID"`
	FieldName string    `json:"FieldName"`
	Value     string    `json:"Value"`
	TimesUsed int       `json:"TimesUsed"`
	FirstUsed time.Time `json:"FirstUsed"`
	LastUsed  time.Time `json:"LastUsed"`
	GUID      string    `json:"GUID"`
	OSPath    string    `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Firefox_Form_History) StringArray() []string {
	return []string{strconv.Itoa(s.ID), s.FieldName, s.Value, strconv.Itoa(s.TimesUsed), s.FirstUsed.String(), s.LastUsed.String(), s.GUID, s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Firefox_Form_History) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Firefox_Form_History(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Firefox_Form_History{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastUsed.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastUsed,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Form Last Used (Firefox)",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Field Name: %v, First Used: %v, Value: %v", tmp.FieldName, tmp.FirstUsed, tmp.Value),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data struct {
	ExpiryTime      time.Time `json:"ExpiryTime"`
	ModifiedTime    time.Time `json:"ModifiedTime"`
	AccessedTime    time.Time `json:"AccessedTime"`
	URL             string    `json:"Url"`
	EntryID         int       `json:"EntryId"`
	ContainerID     int       `json:"ContainerId"`
	CacheID         int       `json:"CacheId"`
	URLHash         int64     `json:"UrlHash"`
	SecureDirectory int       `json:"SecureDirectory"`
	FileSize        int       `json:"FileSize"`
	Type            int       `json:"Type"`
	Flags           int       `json:"Flags"`
	AccessCount     int       `json:"AccessCount"`
	SyncTime        int64     `json:"SyncTime"`
	CreationTime    int64     `json:"CreationTime"`
	PostCheckTime   int       `json:"PostCheckTime"`
	SyncCount       int       `json:"SyncCount"`
	ExemptionDelta  int       `json:"ExemptionDelta"`
	Filename        string    `json:"Filename"`
	RequestHeaders  string    `json:"RequestHeaders"`
	ResponseHeaders string    `json:"ResponseHeaders"`
}

func (s Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data) StringArray() []string {
	return []string{s.ExpiryTime.String(), s.ModifiedTime.String(), s.AccessedTime.String(), s.URL, strconv.Itoa(s.EntryID),
		strconv.Itoa(s.ContainerID), strconv.Itoa(s.CacheID), strconv.FormatInt(s.URLHash, 10), strconv.Itoa(s.SecureDirectory), strconv.Itoa(s.FileSize),
		strconv.Itoa(s.Type), strconv.Itoa(s.Flags), strconv.Itoa(s.AccessCount), strconv.FormatInt(s.SyncTime, 10),
		strconv.FormatInt(s.CreationTime, 10), strconv.Itoa(s.PostCheckTime), strconv.Itoa(s.SyncCount), strconv.Itoa(s.ExemptionDelta), s.Filename, s.RequestHeaders, s.ResponseHeaders}
}

func (s Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.AccessedTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.AccessedTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "IE/Edge WebCache Access",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  tmp.URL,
			SourceFile:       "",
			MetaData:         fmt.Sprintf("File Name: %v, Creation Time: %v", tmp.Filename, tmp.CreationTime),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr struct {
	ScopeID      int       `json:"ScopeID"`
	DocumentID   int       `json:"DocumentID"`
	SDID         int       `json:"SDID"`
	LastModified time.Time `json:"LastModified"`
	FileName     string    `json:"FileName"`
}

func (s Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr) StringArray() []string {
	return []string{strconv.Itoa(s.ScopeID), strconv.Itoa(s.DocumentID), strconv.Itoa(s.SDID), s.LastModified.String(), s.FileName}
}

func (s Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.LastModified.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.LastModified,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Search Service Entry Modified",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.FileName,
			MetaData:         fmt.Sprintf(""),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore struct {
	WorkID                          int       `json:"WorkID"`
	SearchRank                      int       `json:"System_Search_Rank"`
	FileAttributes                  int       `json:"System_FileAttributes"`
	InvertedOnlyMD5                 string    `json:"InvertedOnlyMD5"`
	IsFolder                        bool      `json:"System_IsFolder"`
	FilePlaceholderStatus           int       `json:"System_FilePlaceholderStatus"`
	SearchAccessCount               int       `json:"System_Search_AccessCount"`
	ItemFolderPathDisplay           string    `json:"System_ItemFolderPathDisplay"`
	ItemPathDisplay                 string    `json:"System_ItemPathDisplay"`
	SearchLastIndexedTotalTime      float64   `json:"System_Search_LastIndexedTotalTime"`
	ItemURL                         string    `json:"System_ItemUrl"`
	FileOwner                       string    `json:"System_FileOwner"`
	DateImported                    string    `json:"System_DateImported"`
	LinkTargetParsingPath           string    `json:"System_Link_TargetParsingPath"`
	LinkTargetSFGAOFlags            int       `json:"System_Link_TargetSFGAOFlags"`
	NotUserContent                  bool      `json:"System_NotUserContent"`
	IsAttachment                    bool      `json:"System_IsAttachment"`
	SearchAutoSummary               string    `json:"System_Search_AutoSummary"`
	IsEncrypted                     bool      `json:"System_IsEncrypted"`
	ItemDate                        string    `json:"System_ItemDate"`
	Kind                            string    `json:"System_Kind"`
	ThumbnailCacheID                string    `json:"System_ThumbnailCacheId"`
	VolumeID                        string    `json:"System_VolumeId"`
	SearchStore                     string    `json:"System_Search_Store"`
	ItemFolderNameDisplay           string    `json:"System_ItemFolderNameDisplay"`
	ItemTypeText                    string    `json:"System_ItemTypeText"`
	Comment                         string    `json:"System_Comment"`
	ItemNameDisplay                 string    `json:"System_ItemNameDisplay"`
	FileExtension                   string    `json:"System_FileExtension"`
	DocumentDateCreated             string    `json:"System_Document_DateCreated"`
	DocumentDateSaved               string    `json:"System_Document_DateSaved"`
	ItemName                        string    `json:"System_ItemName"`
	KindText                        string    `json:"System_KindText"`
	ItemFolderPathDisplayNarrow     string    `json:"System_ItemFolderPathDisplayNarrow"`
	ItemNameDisplayWithoutExtension string    `json:"System_ItemNameDisplayWithoutExtension"`
	ComputerName                    string    `json:"System_ComputerName"`
	ItemPathDisplayNarrow           string    `json:"System_ItemPathDisplayNarrow"`
	ItemType                        string    `json:"System_ItemType"`
	FileName                        string    `json:"System_FileName"`
	ParsingName                     string    `json:"System_ParsingName"`
	SFGAOFlags                      int       `json:"System_SFGAOFlags"`
	InvertedOnlyPids                string    `json:"InvertedOnlyPids"`
	SearchGatherTime                time.Time `json:"System_Search_GatherTime"`
	Size                            int       `json:"System_Size"`
	DateModified                    time.Time `json:"System_DateModified"`
	DateAccessed                    time.Time `json:"System_DateAccessed"`
	DateCreated                     time.Time `json:"System_DateCreated"`
}

func (s Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore) StringArray() []string {
	return []string{strconv.Itoa(s.WorkID), strconv.Itoa(s.SearchRank), strconv.Itoa(s.FileAttributes), s.InvertedOnlyMD5, strconv.FormatBool(s.IsFolder), strconv.Itoa(s.FilePlaceholderStatus),
		strconv.Itoa(s.SearchAccessCount), s.ItemFolderPathDisplay, s.ItemPathDisplay, fmt.Sprint(s.SearchLastIndexedTotalTime),
		s.ItemURL, s.FileOwner, s.DateImported, s.LinkTargetParsingPath, strconv.Itoa(s.LinkTargetSFGAOFlags), strconv.FormatBool(s.NotUserContent),
		strconv.FormatBool(s.IsAttachment), s.SearchAutoSummary, strconv.FormatBool(s.IsEncrypted), s.ItemDate, s.Kind,
		s.ThumbnailCacheID, s.VolumeID, s.SearchStore, s.ItemFolderNameDisplay, s.ItemTypeText, s.Comment, s.ItemNameDisplay,
		s.FileExtension, s.DocumentDateCreated, s.DocumentDateSaved, s.ItemName, s.KindText, s.ItemFolderPathDisplayNarrow,
		s.ItemNameDisplayWithoutExtension, s.ComputerName, s.ItemPathDisplayNarrow, s.ItemType, s.FileName, s.ParsingName,
		strconv.Itoa(s.SFGAOFlags), s.InvertedOnlyPids, s.SearchGatherTime.String(), strconv.Itoa(s.Size), s.DateModified.String(), s.DateAccessed.String(), s.DateCreated.String()}
}

func (s Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.DateAccessed.String(), clientIdentifier, tmp.ComputerName, tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.DateAccessed,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Search Service Property Accessed",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.ItemFolderPathDisplay,
			MetaData:         fmt.Sprintf("Owner: %v, Created: %v", tmp.FileOwner, tmp.DateCreated),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth struct {
	Scope  int    `json:"Scope"`
	Parent int    `json:"Parent"`
	Name   string `json:"Name"`
}

func (s Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth) StringArray() []string {
	return []string{strconv.Itoa(s.Scope), strconv.Itoa(s.Parent), s.Name}
}

func (s Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_Extensions struct {
	OSPath      string   `json:"OSPath"`
	Email       string   `json:"Email"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Scopes      any      `json:"Scopes"`
	Permissions []string `json:"Permissions"`
	Key         string   `json:"Key"`
	Image       struct {
		Path       string   `json:"Path"`
		Size       int      `json:"Size"`
		StoredSize int      `json:"StoredSize"`
		Sha256     string   `json:"sha256"`
		Md5        string   `json:"md5"`
		StoredName string   `json:"StoredName"`
		Components []string `json:"Components"`
		Accessor   string   `json:"Accessor"`
	} `json:"Image"`
	Manifest struct {
		Author struct {
			Email string `json:"email"`
		} `json:"author"`
		Background struct {
			Persistent bool     `json:"persistent"`
			Scripts    []string `json:"scripts"`
		} `json:"background"`
		BrowserAction struct {
			DefaultIcon struct {
				Num19 string `json:"19"`
				Num38 string `json:"38"`
			} `json:"default_icon"`
			DefaultPopup string `json:"default_popup"`
			DefaultTitle string `json:"default_title"`
		} `json:"browser_action"`
		ContentScripts []struct {
			AllFrames bool     `json:"all_frames"`
			CSS       []string `json:"css"`
			Js        []string `json:"js"`
			Matches   []string `json:"matches"`
		} `json:"content_scripts"`
		ContentSecurityPolicy   string `json:"content_security_policy"`
		DefaultLocale           string `json:"default_locale"`
		Description             string `json:"description"`
		DifferentialFingerprint string `json:"differential_fingerprint"`
		Icons                   struct {
			Num16  string `json:"16"`
			Num19  string `json:"19"`
			Num32  string `json:"32"`
			Num38  string `json:"38"`
			Num48  string `json:"48"`
			Num128 string `json:"128"`
		} `json:"icons"`
		Key                    string   `json:"key"`
		ManifestVersion        int      `json:"manifest_version"`
		Name                   string   `json:"name"`
		OptionsPage            string   `json:"options_page"`
		Permissions            []string `json:"permissions"`
		UpdateURL              string   `json:"update_url"`
		Version                string   `json:"version"`
		WebAccessibleResources []string `json:"web_accessible_resources"`
	} `json:"_Manifest"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Extensions) StringArray() []string {
	contentScripts := make([]string, 0)
	for _, v := range s.Manifest.ContentScripts {
		contentScripts = append(contentScripts, fmt.Sprintf("AllFrames: %v, CSS: %v, Js: %v, Matches: %v", v.AllFrames, v.CSS, v.Js, v.Matches))
	}
	return []string{s.OSPath, s.Email, s.Name, s.Description, fmt.Sprint(s.Scopes), fmt.Sprint(s.Permissions),
		s.Key, s.Image.Path, strconv.Itoa(s.Image.Size), strconv.Itoa(s.Image.StoredSize), s.Image.Sha256, s.Image.Md5, s.Image.StoredName,
		fmt.Sprint(s.Image.Components), s.Image.Accessor, s.Manifest.Author.Email, fmt.Sprint(s.Manifest.Background.Persistent),
		fmt.Sprint(s.Manifest.Background.Scripts), s.Manifest.BrowserAction.DefaultPopup, s.Manifest.BrowserAction.DefaultTitle,
		fmt.Sprint(contentScripts), s.Manifest.ContentSecurityPolicy, s.Manifest.DefaultLocale, s.Manifest.Description, s.Manifest.DifferentialFingerprint,
		s.Manifest.Key, strconv.Itoa(s.Manifest.ManifestVersion), s.Manifest.Name, s.Manifest.OptionsPage, fmt.Sprint(s.Manifest.Permissions), s.Manifest.UpdateURL, s.Manifest.Version, fmt.Sprint(s.Manifest.WebAccessibleResources)}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Extensions) GetHeaders() []string {
	return []string{"OSPath", "Email", "Name", "Description", "Scopes", "Permissions", "Key", "Image_Path", "Image_Size",
		"Image_StoredSize", "Image_SHA256", "Image_MD5", "Image_StoredName", "Image_Components", "Image_Accessor",
		"Manifest_Author_Email", "Manifest_Background_Persistent", "Manifest_Background_Scripts",
		"Manifest_BrowserAction_DefaultPopup", "Manifest_BrowserAction_DefaultTitle", "Manifest_ContentScripts",
		"Manifest_ContentSecurityPolicy", "Manifest_DefaultLocale", "Manifest_Description", "Manifest_DifferentialFingerprint",
		"Manifest_Key", "Manifest_ManifestVersion", "Manifest_Name", "Manifest_OptionsPage", "Manifest_Permissions",
		"Manifest_UpdateURL", "Manifest_Version", "Manifest_WebAccessibleResources"}
}

func Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Extensions(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Extensions{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_Network_Predictor struct {
	ID             string `json:"ID"`
	UserText       string `json:"UserText"`
	URL            string `json:"URL"`
	NumberOfHits   int    `json:"NumberOfHits"`
	NumberOfMisses int    `json:"NumberOfMisses"`
	OSPath         string `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Network_Predictor) StringArray() []string {
	return []string{s.ID, s.UserText, s.URL, strconv.Itoa(s.NumberOfHits), strconv.Itoa(s.NumberOfMisses), s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Network_Predictor) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Network_Predictor(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	// TODO - This artifact by it's nature has lots of duplicate rows - we are removing these when processing it - probably need to not do that or have an option so it is 'per-artifact' deduplication
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Network_Predictor{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}

type Generic_Forensic_SQLiteHunter_Chromium_Browser_Top_Sites struct {
	URLRank int    `json:"URLRank"`
	URL     string `json:"URL"`
	Title   string `json:"Title"`
	OSPath  string `json:"OSPath"`
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Top_Sites) StringArray() []string {
	return []string{strconv.Itoa(s.URLRank), s.URL, s.Title, s.OSPath}
}

func (s Generic_Forensic_SQLiteHunter_Chromium_Browser_Top_Sites) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Top_Sites(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Top_Sites{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
