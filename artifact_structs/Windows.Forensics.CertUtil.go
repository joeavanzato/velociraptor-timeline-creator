package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"strconv"
	"time"
)

type Windows_Forensics_CertUtil struct {
	MetadataFile       string    `json:"_MetadataFile"`
	ContentPath        string    `json:"_ContentPath"`
	MetdataUpload      any       `json:"_MetdataUpload"`
	Upload             any       `json:"_Upload"`
	URL                string    `json:"URL"`
	URLTLD             string    `json:"UrlTLD"`
	FileSize           int       `json:"FileSize"`
	Hash               string    `json:"Hash"`
	DownloadTime       time.Time `json:"DownloadTime"`
	VersionInformation string    `json:"VersionInformation"`
	Authenticode       any       `json:"Authenticode"`
}

func (s Windows_Forensics_CertUtil) StringArray() []string {
	// TODO - Probably revisit this and format with proper authenticode fields - this is only triggered if a cert-as-exe is detected (MZ)
	return []string{s.MetadataFile, s.ContentPath, fmt.Sprint(s.MetdataUpload), fmt.Sprint(s.Upload), s.URL, s.URLTLD, strconv.Itoa(s.FileSize), s.Hash, s.DownloadTime.String(), s.VersionInformation, fmt.Sprint(s.Authenticode)}
}

func (s Windows_Forensics_CertUtil) GetHeaders() []string {
	return helpers.GetStructAsStringSlice(s)
}

func Process_Windows_Forensics_CertUtil(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Forensics_CertUtil{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.DownloadTime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.DownloadTime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.MetadataFile,
			MetaData:         fmt.Sprintf("URL: %v", tmp.URL),
		}
		outputChannel <- tmp2.StringArray()
	}
}
