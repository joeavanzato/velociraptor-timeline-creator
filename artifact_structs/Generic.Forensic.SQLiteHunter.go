package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
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

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Firefox_Cookies(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Firefox_Cookies{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Firefox_Form_History(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Firefox_Form_History{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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

func Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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
	WorkID                                int       `json:"WorkID"`
	SystemSearchRank                      int       `json:"System_Search_Rank"`
	SystemFileAttributes                  int       `json:"System_FileAttributes"`
	InvertedOnlyMD5                       string    `json:"InvertedOnlyMD5"`
	SystemIsFolder                        bool      `json:"System_IsFolder"`
	SystemFilePlaceholderStatus           int       `json:"System_FilePlaceholderStatus"`
	SystemSearchAccessCount               int       `json:"System_Search_AccessCount"`
	SystemItemFolderPathDisplay           string    `json:"System_ItemFolderPathDisplay"`
	SystemItemPathDisplay                 string    `json:"System_ItemPathDisplay"`
	SystemSearchLastIndexedTotalTime      float64   `json:"System_Search_LastIndexedTotalTime"`
	SystemItemURL                         string    `json:"System_ItemUrl"`
	SystemFileOwner                       string    `json:"System_FileOwner"`
	SystemDateImported                    string    `json:"System_DateImported"`
	SystemLinkTargetParsingPath           string    `json:"System_Link_TargetParsingPath"`
	SystemLinkTargetSFGAOFlags            int       `json:"System_Link_TargetSFGAOFlags"`
	SystemNotUserContent                  bool      `json:"System_NotUserContent"`
	SystemIsAttachment                    bool      `json:"System_IsAttachment"`
	SystemSearchAutoSummary               string    `json:"System_Search_AutoSummary"`
	SystemIsEncrypted                     bool      `json:"System_IsEncrypted"`
	SystemItemDate                        string    `json:"System_ItemDate"`
	SystemKind                            string    `json:"System_Kind"`
	SystemThumbnailCacheID                string    `json:"System_ThumbnailCacheId"`
	SystemVolumeID                        string    `json:"System_VolumeId"`
	SystemSearchStore                     string    `json:"System_Search_Store"`
	SystemItemFolderNameDisplay           string    `json:"System_ItemFolderNameDisplay"`
	SystemItemTypeText                    string    `json:"System_ItemTypeText"`
	SystemComment                         string    `json:"System_Comment"`
	SystemItemNameDisplay                 string    `json:"System_ItemNameDisplay"`
	SystemFileExtension                   string    `json:"System_FileExtension"`
	SystemDocumentDateCreated             string    `json:"System_Document_DateCreated"`
	SystemDocumentDateSaved               string    `json:"System_Document_DateSaved"`
	SystemItemName                        string    `json:"System_ItemName"`
	SystemKindText                        string    `json:"System_KindText"`
	SystemItemFolderPathDisplayNarrow     string    `json:"System_ItemFolderPathDisplayNarrow"`
	SystemItemNameDisplayWithoutExtension string    `json:"System_ItemNameDisplayWithoutExtension"`
	SystemComputerName                    string    `json:"System_ComputerName"`
	SystemItemPathDisplayNarrow           string    `json:"System_ItemPathDisplayNarrow"`
	SystemItemType                        string    `json:"System_ItemType"`
	SystemFileName                        string    `json:"System_FileName"`
	SystemParsingName                     string    `json:"System_ParsingName"`
	SystemSFGAOFlags                      int       `json:"System_SFGAOFlags"`
	InvertedOnlyPids                      string    `json:"InvertedOnlyPids"`
	SystemSearchGatherTime                time.Time `json:"System_Search_GatherTime"`
	SystemSize                            int       `json:"System_Size"`
	SystemDateModified                    time.Time `json:"System_DateModified"`
	SystemDateAccessed                    time.Time `json:"System_DateAccessed"`
	SystemDateCreated                     time.Time `json:"System_DateCreated"`
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
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.SystemDateAccessed,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Search Service Property Accessed",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.SystemItemFolderPathDisplay,
			MetaData:         fmt.Sprintf("Owner: %v, Created: %v", tmp.SystemFileOwner, tmp.SystemDateCreated),
		}
		outputChannel <- tmp2.StringArray()
	}
}
