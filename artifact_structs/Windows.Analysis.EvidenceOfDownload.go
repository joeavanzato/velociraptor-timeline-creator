package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"time"
)

type Windows_Analysis_EvidenceOfDownload struct {
	DownloadedFilePath    string    `json:"DownloadedFilePath"`
	Mtime                 time.Time `json:"Mtime"`
	ZoneIdentifierContent string    `json:"_ZoneIdentifierContent"`
	FileHash              struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"FileHash"`
	ZoneID      string `json:"ZoneId"`
	HostURL     string `json:"HostUrl"`
	ReferrerURL string `json:"ReferrerUrl"`
}

func (s Windows_Analysis_EvidenceOfDownload) StringArray() []string {
	return []string{s.DownloadedFilePath, s.Mtime.String(), s.ZoneIdentifierContent, s.FileHash.MD5, s.FileHash.SHA1, s.FileHash.SHA256, s.ZoneID, s.HostURL, s.ReferrerURL}
}

func (s Windows_Analysis_EvidenceOfDownload) GetHeaders() []string {
	return []string{"DownloadedFilePath", "Mtime", "ZoneIdentifierContent", "FileHash_MD5", "FileHash_SHA1", "FileHash_SHA256", "ZoneID", "HostURL", "ReferrerURL"}
}

func Process_Windows_Analysis_EvidenceOfDownload(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Analysis_EvidenceOfDownload{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.DownloadedFilePath,
			MetaData:         fmt.Sprintf("Referer URL: %v, Host URL: %v, MD5: %v, Zone ID: %v", tmp.ReferrerURL, tmp.HostURL, tmp.FileHash.MD5, tmp.ZoneID),
		}
		outputChannel <- tmp2.StringArray()
	}
}
