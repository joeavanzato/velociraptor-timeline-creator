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

type Windows_System_Amcache_InventoryApplicationFile struct {
	FileID         string    `json:"FileId"`
	Key            string    `json:"Key"`
	Hive           string    `json:"Hive"`
	LastModified   time.Time `json:"LastModified"`
	Binary         string    `json:"Binary"`
	Name           string    `json:"Name"`
	Size           int       `json:"Size"`
	ProductName    string    `json:"ProductName"`
	Publisher      string    `json:"Publisher"`
	Version        string    `json:"Version"`
	BinFileVersion string    `json:"BinFileVersion"`
}

func (s Windows_System_Amcache_InventoryApplicationFile) StringArray() []string {
	return []string{s.FileID, s.Key, s.Hive, s.LastModified.String(), s.Binary, s.Name, strconv.Itoa(s.Size), s.ProductName, s.Publisher, s.Version, s.BinFileVersion}
}

// Headers should match the array above
func (s Windows_System_Amcache_InventoryApplicationFile) GetHeaders() []string {
	return []string{"FileId", "Key", "Hive", "LastModified", "Binary", "Name", "Size", "ProductName", "Publisher", "Version", "BinFileVersion"}
}

func Process_Windows_System_Amcache_InventoryApplicationFile(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Amcache_InventoryApplicationFile{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
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
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Name,
			MetaData:         fmt.Sprintf("Binary: %v, Product Name: %v, Publisher: %v", tmp.Binary, tmp.ProductName, tmp.Publisher),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Windows_Analysis_EvidenceOfExecution_Amcache struct {
	HivePath         string    `json:"HivePath"`
	EntryKey         string    `json:"EntryKey"`
	KeyMTime         time.Time `json:"KeyMTime"`
	EntryType        string    `json:"EntryType"`
	SHA1             string    `json:"SHA1"`
	EntryName        string    `json:"EntryName"`
	EntryPath        string    `json:"EntryPath"`
	Publisher        string    `json:"Publisher"`
	OriginalFileName string    `json:"OriginalFileName"`
	BinaryType       string    `json:"BinaryType"`
	Source           string    `json:"_Source"`
}

func (s Windows_Analysis_EvidenceOfExecution_Amcache) StringArray() []string {
	return []string{s.HivePath, s.EntryKey, s.KeyMTime.String(), s.EntryType, s.SHA1, s.EntryName, s.EntryPath, s.Publisher, s.OriginalFileName, s.BinaryType, s.Source}
}

// Headers should match the array above
func (s Windows_Analysis_EvidenceOfExecution_Amcache) GetHeaders() []string {
	return []string{"HivePath", "EntryKey", "KeyMTime", "EntryType", "SHA1", "EntryName", "EntryPath", "Publisher", "OriginalFileName", "BinaryType", "Source"}
}

func Process_Windows_Analysis_EvidenceOfExecution_Amcache(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Analysis_EvidenceOfExecution_Amcache{}
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
			EventType:        "Amcache Entry Modified",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.EntryPath,
			MetaData:         fmt.Sprintf("Entry Type: %v, Publisher: %v, SHA1: %v", tmp.EntryType, tmp.Publisher, tmp.SHA1),
		}
		outputChannel <- tmp2.StringArray()
	}
}
