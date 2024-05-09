package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/vars"
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

func Process_Windows_System_Amcache_InventoryApplicationFile(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_System_Amcache_InventoryApplicationFile{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			fmt.Println(err.Error())
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
