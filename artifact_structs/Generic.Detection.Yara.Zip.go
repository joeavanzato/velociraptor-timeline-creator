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

type Generic_Detection_Yara_Zip struct {
	Container     string `json:"Container"`
	ExtractedPath string `json:"ExtractedPath"`
	FilePath      string `json:"FilePath"`
	Hash          struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hash"`
	Size  int       `json:"Size"`
	Mtime time.Time `json:"Mtime"`
	Atime time.Time `json:"Atime"`
	Ctime time.Time `json:"Ctime"`
	Btime time.Time `json:"Btime"`
	Rule  string    `json:"Rule"`
	Tags  []string  `json:"Tags"`
	Meta  struct {
		Author      string `json:"author"`
		Date        string `json:"date"`
		Description string `json:"description"`
	} `json:"Meta"`
	YaraString any `json:"YaraString"`
	HitOffset  any `json:"HitOffset"`
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

func (s Generic_Detection_Yara_Zip) StringArray() []string {
	return []string{s.Container, s.ExtractedPath, s.FilePath, s.Hash.MD5, s.Hash.SHA1, s.Hash.SHA256, strconv.Itoa(s.Size),
		s.Mtime.String(), s.Atime.String(), s.Ctime.String(), s.Btime.String(), s.Rule, fmt.Sprint(s.Tags), s.Meta.Author, s.Meta.Date, s.Meta.Description,
		fmt.Sprint(s.YaraString), fmt.Sprint(s.HitOffset), s.HitContext.Path, strconv.Itoa(s.HitContext.Size), strconv.Itoa(s.HitContext.StoredSize), s.HitContext.Sha256, s.HitContext.Md5,
		s.HitContext.StoredName, fmt.Sprint(s.HitContext.Components), s.HitContext.Accessor}
}

func (s Generic_Detection_Yara_Zip) GetHeaders() []string {
	return []string{"Container", "ExtractedPath", "FilePath", "Hash_MD5", "Hash_SHA1", "Hash_SHA256", "Size", "Mtime", "Atime", "Ctime", "Btime", "Rule", "Tags", "Author", "Date", "Description", "YaraString", "HitOffset",
		"Hit_Path", "Hit_Size", "Hit_StoredSize", "Hit_SHA256", "Hit_MD5", "Hit_StoredName", "Hit_Components", "Hit_Accessor"}
}

func Process_Generic_Detection_Yara_Zip(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Generic_Detection_Yara_Zip{}
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
			EventDescription: fmt.Sprintf("Rule: %v, Description: %v", tmp.Rule, tmp.Meta.Description),
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.ExtractedPath,
			MetaData:         fmt.Sprintf("Container: %v, MD5: %v", tmp.Container, tmp.Hash.MD5),
		}
		outputChannel <- tmp2.StringArray()
	}
}
