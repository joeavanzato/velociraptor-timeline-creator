package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"time"
)

type Exchange_Windows_System_WindowsErrorReporting struct {
	Timestamp                 time.Time `json:"timestamp"`
	ReportEventType           string    `json:"Report.EventType"`
	ReportFriendlyEventName   string    `json:"Report.FriendlyEventName"`
	ReportAppName             string    `json:"Report.AppName"`
	ReportAppPath             string    `json:"Report.AppPath"`
	ReportApplicationIdentity string    `json:"Report.ApplicationIdentity"`
	SHA1                      string    `json:"SHA1"`
	ReportOriginalFilename    string    `json:"Report.OriginalFilename"`
	Report                    struct {
		Version                    string `json:"Version"`
		EventType                  string `json:"EventType"`
		EventTime                  string `json:"EventTime"`
		ReportType                 string `json:"ReportType"`
		Consent                    string `json:"Consent"`
		UploadTime                 string `json:"UploadTime"`
		ReportStatus               string `json:"ReportStatus"`
		ReportIdentifier           string `json:"ReportIdentifier"`
		IntegratorReportIdentifier string `json:"IntegratorReportIdentifier"`
		Wow64Host                  string `json:"Wow64Host"`
		NsAppName                  string `json:"NsAppName"`
		AppSessionGuid             string `json:"AppSessionGuid"`
		TargetAppId                string `json:"TargetAppId"`
		TargetAppVer               string `json:"TargetAppVer"`
		BootId                     string `json:"BootId"`
		TargetAsId                 string `json:"TargetAsId"`
		IsFatal                    string `json:"IsFatal"`
		EtwNonCollectReason        string `json:"EtwNonCollectReason"`
		ResponseBucketId           string `json:"Response.BucketId"`
		ResponseBucketTable        string `json:"Response.BucketTable"`
		ResponseLegacyBucketId     string `json:"Response.LegacyBucketId"`
		ResponseType               string `json:"Response.type"`
		Sig0Name                   string `json:"Sig[0].Name"`
		Sig0Value                  string `json:"Sig[0].Value"`
		Sig1Name                   string `json:"Sig[1].Name"`
		Sig1Value                  string `json:"Sig[1].Value"`
		Sig2Name                   string `json:"Sig[2].Name"`
		Sig2Value                  string `json:"Sig[2].Value"`
		Sig3Name                   string `json:"Sig[3].Name"`
		Sig3Value                  string `json:"Sig[3].Value"`
		Sig4Name                   string `json:"Sig[4].Name"`
		Sig4Value                  string `json:"Sig[4].Value"`
		Sig5Name                   string `json:"Sig[5].Name"`
		Sig5Value                  string `json:"Sig[5].Value"`
		Sig6Name                   string `json:"Sig[6].Name"`
		Sig6Value                  string `json:"Sig[6].Value"`
		Sig7Name                   string `json:"Sig[7].Name"`
		Sig7Value                  string `json:"Sig[7].Value"`
		DynamicSig1Name            string `json:"DynamicSig[1].Name"`
		DynamicSig1Value           string `json:"DynamicSig[1].Value"`
		DynamicSig2Name            string `json:"DynamicSig[2].Name"`
		DynamicSig2Value           string `json:"DynamicSig[2].Value"`
		DynamicSig22Name           string `json:"DynamicSig[22].Name"`
		DynamicSig22Value          string `json:"DynamicSig[22].Value"`
		DynamicSig23Name           string `json:"DynamicSig[23].Name"`
		DynamicSig23Value          string `json:"DynamicSig[23].Value"`
		DynamicSig24Name           string `json:"DynamicSig[24].Name"`
		DynamicSig24Value          string `json:"DynamicSig[24].Value"`
		DynamicSig25Name           string `json:"DynamicSig[25].Name"`
		DynamicSig25Value          string `json:"DynamicSig[25].Value"`
		UI2                        string `json:"UI[2]"`
		LoadedModule0              string `json:"LoadedModule[0]"`
		LoadedModule1              string `json:"LoadedModule[1]"`
		LoadedModule2              string `json:"LoadedModule[2]"`
		LoadedModule3              string `json:"LoadedModule[3]"`
		LoadedModule4              string `json:"LoadedModule[4]"`
		LoadedModule5              string `json:"LoadedModule[5]"`
		LoadedModule6              string `json:"LoadedModule[6]"`
		LoadedModule7              string `json:"LoadedModule[7]"`
		LoadedModule8              string `json:"LoadedModule[8]"`
		LoadedModule9              string `json:"LoadedModule[9]"`
		LoadedModule10             string `json:"LoadedModule[10]"`
		LoadedModule11             string `json:"LoadedModule[11]"`
		LoadedModule12             string `json:"LoadedModule[12]"`
		LoadedModule13             string `json:"LoadedModule[13]"`
		LoadedModule14             string `json:"LoadedModule[14]"`
		LoadedModule15             string `json:"LoadedModule[15]"`
		LoadedModule16             string `json:"LoadedModule[16]"`
		LoadedModule17             string `json:"LoadedModule[17]"`
		LoadedModule18             string `json:"LoadedModule[18]"`
		LoadedModule19             string `json:"LoadedModule[19]"`
		LoadedModule20             string `json:"LoadedModule[20]"`
		LoadedModule21             string `json:"LoadedModule[21]"`
		LoadedModule22             string `json:"LoadedModule[22]"`
		LoadedModule23             string `json:"LoadedModule[23]"`
		LoadedModule24             string `json:"LoadedModule[24]"`
		LoadedModule25             string `json:"LoadedModule[25]"`
		LoadedModule26             string `json:"LoadedModule[26]"`
		LoadedModule27             string `json:"LoadedModule[27]"`
		LoadedModule28             string `json:"LoadedModule[28]"`
		LoadedModule29             string `json:"LoadedModule[29]"`
		LoadedModule30             string `json:"LoadedModule[30]"`
		LoadedModule31             string `json:"LoadedModule[31]"`
		LoadedModule32             string `json:"LoadedModule[32]"`
		LoadedModule33             string `json:"LoadedModule[33]"`
		LoadedModule34             string `json:"LoadedModule[34]"`
		LoadedModule35             string `json:"LoadedModule[35]"`
		LoadedModule36             string `json:"LoadedModule[36]"`
		LoadedModule37             string `json:"LoadedModule[37]"`
		LoadedModule38             string `json:"LoadedModule[38]"`
		LoadedModule39             string `json:"LoadedModule[39]"`
		LoadedModule40             string `json:"LoadedModule[40]"`
		LoadedModule41             string `json:"LoadedModule[41]"`
		LoadedModule42             string `json:"LoadedModule[42]"`
		LoadedModule43             string `json:"LoadedModule[43]"`
		LoadedModule44             string `json:"LoadedModule[44]"`
		LoadedModule45             string `json:"LoadedModule[45]"`
		LoadedModule46             string `json:"LoadedModule[46]"`
		LoadedModule47             string `json:"LoadedModule[47]"`
		LoadedModule48             string `json:"LoadedModule[48]"`
		LoadedModule49             string `json:"LoadedModule[49]"`
		LoadedModule50             string `json:"LoadedModule[50]"`
		LoadedModule51             string `json:"LoadedModule[51]"`
		LoadedModule52             string `json:"LoadedModule[52]"`
		LoadedModule53             string `json:"LoadedModule[53]"`
		LoadedModule54             string `json:"LoadedModule[54]"`
		LoadedModule55             string `json:"LoadedModule[55]"`
		LoadedModule56             string `json:"LoadedModule[56]"`
		LoadedModule57             string `json:"LoadedModule[57]"`
		LoadedModule58             string `json:"LoadedModule[58]"`
		LoadedModule59             string `json:"LoadedModule[59]"`
		LoadedModule60             string `json:"LoadedModule[60]"`
		LoadedModule61             string `json:"LoadedModule[61]"`
		LoadedModule62             string `json:"LoadedModule[62]"`
		LoadedModule63             string `json:"LoadedModule[63]"`
		LoadedModule64             string `json:"LoadedModule[64]"`
		LoadedModule65             string `json:"LoadedModule[65]"`
		LoadedModule66             string `json:"LoadedModule[66]"`
		LoadedModule67             string `json:"LoadedModule[67]"`
		LoadedModule68             string `json:"LoadedModule[68]"`
		LoadedModule69             string `json:"LoadedModule[69]"`
		LoadedModule70             string `json:"LoadedModule[70]"`
		LoadedModule71             string `json:"LoadedModule[71]"`
		LoadedModule72             string `json:"LoadedModule[72]"`
		LoadedModule73             string `json:"LoadedModule[73]"`
		LoadedModule74             string `json:"LoadedModule[74]"`
		LoadedModule75             string `json:"LoadedModule[75]"`
		LoadedModule76             string `json:"LoadedModule[76]"`
		LoadedModule77             string `json:"LoadedModule[77]"`
		LoadedModule78             string `json:"LoadedModule[78]"`
		LoadedModule79             string `json:"LoadedModule[79]"`
		LoadedModule80             string `json:"LoadedModule[80]"`
		LoadedModule81             string `json:"LoadedModule[81]"`
		LoadedModule82             string `json:"LoadedModule[82]"`
		LoadedModule83             string `json:"LoadedModule[83]"`
		LoadedModule84             string `json:"LoadedModule[84]"`
		LoadedModule85             string `json:"LoadedModule[85]"`
		LoadedModule86             string `json:"LoadedModule[86]"`
		LoadedModule87             string `json:"LoadedModule[87]"`
		LoadedModule88             string `json:"LoadedModule[88]"`
		LoadedModule89             string `json:"LoadedModule[89]"`
		LoadedModule90             string `json:"LoadedModule[90]"`
		LoadedModule91             string `json:"LoadedModule[91]"`
		LoadedModule92             string `json:"LoadedModule[92]"`
		LoadedModule93             string `json:"LoadedModule[93]"`
		LoadedModule94             string `json:"LoadedModule[94]"`
		LoadedModule95             string `json:"LoadedModule[95]"`
		LoadedModule96             string `json:"LoadedModule[96]"`
		LoadedModule97             string `json:"LoadedModule[97]"`
		LoadedModule98             string `json:"LoadedModule[98]"`
		LoadedModule99             string `json:"LoadedModule[99]"`
		LoadedModule100            string `json:"LoadedModule[100]"`
		LoadedModule101            string `json:"LoadedModule[101]"`
		LoadedModule102            string `json:"LoadedModule[102]"`
		LoadedModule103            string `json:"LoadedModule[103]"`
		LoadedModule104            string `json:"LoadedModule[104]"`
		LoadedModule105            string `json:"LoadedModule[105]"`
		LoadedModule106            string `json:"LoadedModule[106]"`
		LoadedModule107            string `json:"LoadedModule[107]"`
		LoadedModule108            string `json:"LoadedModule[108]"`
		LoadedModule109            string `json:"LoadedModule[109]"`
		LoadedModule110            string `json:"LoadedModule[110]"`
		LoadedModule111            string `json:"LoadedModule[111]"`
		LoadedModule112            string `json:"LoadedModule[112]"`
		LoadedModule113            string `json:"LoadedModule[113]"`
		LoadedModule114            string `json:"LoadedModule[114]"`
		State0Key                  string `json:"State[0].Key"`
		State0Value                string `json:"State[0].Value"`
		OsInfo0Key                 string `json:"OsInfo[0].Key"`
		OsInfo0Value               string `json:"OsInfo[0].Value"`
		OsInfo1Key                 string `json:"OsInfo[1].Key"`
		OsInfo1Value               string `json:"OsInfo[1].Value"`
		OsInfo2Key                 string `json:"OsInfo[2].Key"`
		OsInfo2Value               string `json:"OsInfo[2].Value"`
		OsInfo3Key                 string `json:"OsInfo[3].Key"`
		OsInfo3Value               string `json:"OsInfo[3].Value"`
		OsInfo4Key                 string `json:"OsInfo[4].Key"`
		OsInfo4Value               string `json:"OsInfo[4].Value"`
		OsInfo5Key                 string `json:"OsInfo[5].Key"`
		OsInfo5Value               string `json:"OsInfo[5].Value"`
		OsInfo6Key                 string `json:"OsInfo[6].Key"`
		OsInfo6Value               string `json:"OsInfo[6].Value"`
		OsInfo7Key                 string `json:"OsInfo[7].Key"`
		OsInfo7Value               string `json:"OsInfo[7].Value"`
		OsInfo8Key                 string `json:"OsInfo[8].Key"`
		OsInfo8Value               string `json:"OsInfo[8].Value"`
		OsInfo9Key                 string `json:"OsInfo[9].Key"`
		OsInfo9Value               string `json:"OsInfo[9].Value"`
		OsInfo10Key                string `json:"OsInfo[10].Key"`
		OsInfo10Value              string `json:"OsInfo[10].Value"`
		OsInfo11Key                string `json:"OsInfo[11].Key"`
		OsInfo11Value              string `json:"OsInfo[11].Value"`
		OsInfo12Key                string `json:"OsInfo[12].Key"`
		OsInfo12Value              string `json:"OsInfo[12].Value"`
		OsInfo13Key                string `json:"OsInfo[13].Key"`
		OsInfo13Value              string `json:"OsInfo[13].Value"`
		OsInfo14Key                string `json:"OsInfo[14].Key"`
		OsInfo14Value              string `json:"OsInfo[14].Value"`
		OsInfo15Key                string `json:"OsInfo[15].Key"`
		OsInfo15Value              string `json:"OsInfo[15].Value"`
		OsInfo16Key                string `json:"OsInfo[16].Key"`
		OsInfo16Value              string `json:"OsInfo[16].Value"`
		OsInfo17Key                string `json:"OsInfo[17].Key"`
		OsInfo17Value              string `json:"OsInfo[17].Value"`
		OsInfo18Key                string `json:"OsInfo[18].Key"`
		OsInfo18Value              string `json:"OsInfo[18].Value"`
		OsInfo19Key                string `json:"OsInfo[19].Key"`
		OsInfo19Value              string `json:"OsInfo[19].Value"`
		OsInfo20Key                string `json:"OsInfo[20].Key"`
		OsInfo20Value              string `json:"OsInfo[20].Value"`
		OsInfo21Key                string `json:"OsInfo[21].Key"`
		OsInfo21Value              string `json:"OsInfo[21].Value"`
		OsInfo22Key                string `json:"OsInfo[22].Key"`
		OsInfo22Value              string `json:"OsInfo[22].Value"`
		OsInfo23Key                string `json:"OsInfo[23].Key"`
		OsInfo23Value              string `json:"OsInfo[23].Value"`
		OsInfo24Key                string `json:"OsInfo[24].Key"`
		OsInfo24Value              string `json:"OsInfo[24].Value"`
		OsInfo25Key                string `json:"OsInfo[25].Key"`
		OsInfo25Value              string `json:"OsInfo[25].Value"`
		OsInfo26Key                string `json:"OsInfo[26].Key"`
		OsInfo26Value              string `json:"OsInfo[26].Value"`
		OsInfo27Key                string `json:"OsInfo[27].Key"`
		OsInfo27Value              string `json:"OsInfo[27].Value"`
		OsInfo28Key                string `json:"OsInfo[28].Key"`
		OsInfo28Value              string `json:"OsInfo[28].Value"`
		OsInfo29Key                string `json:"OsInfo[29].Key"`
		OsInfo29Value              string `json:"OsInfo[29].Value"`
		OsInfo30Key                string `json:"OsInfo[30].Key"`
		OsInfo31Key                string `json:"OsInfo[31].Key"`
		OsInfo31Value              string `json:"OsInfo[31].Value"`
		OsInfo32Key                string `json:"OsInfo[32].Key"`
		OsInfo32Value              string `json:"OsInfo[32].Value"`
		OsInfo33Key                string `json:"OsInfo[33].Key"`
		OsInfo34Key                string `json:"OsInfo[34].Key"`
		OsInfo35Key                string `json:"OsInfo[35].Key"`
		OsInfo36Key                string `json:"OsInfo[36].Key"`
		OsInfo37Key                string `json:"OsInfo[37].Key"`
		OsInfo37Value              string `json:"OsInfo[37].Value"`
		FriendlyEventName          string `json:"FriendlyEventName"`
		ConsentKey                 string `json:"ConsentKey"`
		AppName                    string `json:"AppName"`
		AppPath                    string `json:"AppPath"`
		NsPartner                  string `json:"NsPartner"`
		NsGroup                    string `json:"NsGroup"`
		ApplicationIdentity        string `json:"ApplicationIdentity"`
	} `json:"Report"`
	ReportFileName string `json:"ReportFileName"`
}

func (s Exchange_Windows_System_WindowsErrorReporting) StringArray() []string {
	baseSlice := helpers.GetStructValuesAsStringSlice(s)
	baseSlice = baseSlice[:len(baseSlice)-1]
	baseSlice = append(baseSlice, helpers.GetStructValuesAsStringSlice(s.Report)...)
	return baseSlice
}

func (s Exchange_Windows_System_WindowsErrorReporting) GetHeaders() []string {
	baseSlice := helpers.GetStructHeadersAsStringSlice(s)
	baseSlice = baseSlice[:len(baseSlice)-1]
	baseSlice = append(baseSlice, helpers.GetStructHeadersAsStringSlice(s.Report)...)
	return baseSlice
}

func Process_Exchange_Windows_System_WindowsErrorReporting(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_System_WindowsErrorReporting{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.Timestamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}

		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.Timestamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: tmp.ReportFriendlyEventName,
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.ReportAppPath,
			MetaData:         fmt.Sprintf("Event Type: %v, SHA1: %v", tmp.ReportEventType, tmp.SHA1),
		}
		outputChannel <- tmp2.StringArray()
	}
}
