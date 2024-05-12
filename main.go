package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/artifact_structs"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/joeavanzato/velo-timeline-creator/vars"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// If we are doing artifacts to CSVs instead of a super timeline (for ingestion to Splunk, browsing, etc) then we will want a single channel per artifact and to only close that channel once we are done parsing all possible artifacts
// This means we can have a single waiter for general writing and only close that once we have finished reading all possible artifacts

func parseArgs(logger zerolog.Logger) (map[string]any, error) {
	velodir := flag.String("velodir", "", "Provide the base path to your Velociraptor Datastore Directory")
	fullparse := flag.Bool("fullparse", false, "NOT IMPLEMENTED YET")
	maxgoperfile := flag.Int("maxgoperfile", 20, "Maximum number of goroutines to spawn on a per-file basis for concurrent processing of data.")
	batchsize := flag.Int("batchsize", 500, "Maximum number of lines to read at a time for processing within each spawned goroutine per file.")
	outputdir := flag.String("outputdir", "", "NOT IMPLEMENTED YET")
	writebuffer := flag.Int("writebuffer", 2000, "How many lines to queue at a time for writing to output CSV")
	mftlight := flag.Bool("mftlight", false, "Adds a subset of interesting extensions from $MFT to the super timeline")
	mftfull := flag.Bool("mftfull", false, "NOT IMPLEMENTED YET")
	artifactdump := flag.Bool("artifactdump", false, "")

	flag.Parse()

	arguments := map[string]any{
		"velodir":      *velodir,
		"fullparse":    *fullparse,
		"maxgoperfile": *maxgoperfile,
		"batchsize":    *batchsize,
		"outputdir":    *outputdir,
		"writebuffer":  *writebuffer,
		"mftlight":     *mftlight,
		"mftfull":      *mftfull,
		"artifactdump": *artifactdump,
	}

	if arguments["outputdir"] == "" {
		// If no output dir is provided, otherwise
		arguments["outputdir"], _ = os.Getwd()
	} else {
		doesOutputDirExist := helpers.DoesFileOrDirExist(arguments["outputdir"].(string))
		if !doesOutputDirExist {
			return arguments, fmt.Errorf("Output Directory Missing: %v ", arguments["outputdir"])
		}
	}

	if arguments["velodir"].(string) == "" {
		return arguments, fmt.Errorf("No Data Directory Specified - please provide the full path to the Velociraptor Data Storage Directory using -velodir")
	}

	return arguments, nil
}

func ValidateVelociraptorDirectory(path string) error {

	doesVeloDirExist := helpers.DoesFileOrDirExist(path)
	if !doesVeloDirExist {
		return fmt.Errorf("Velociraptor Directory Missing: %v ", path)
	}

	clientsPath := path + "\\clients"
	doesVeloDirExist = helpers.DoesFileOrDirExist(clientsPath)
	if !doesVeloDirExist {
		return fmt.Errorf("Velociraptor Clients Directory Missing: %v ", clientsPath)
	}

	return nil
}

func ProcessAllClients(arguments map[string]any) ([]string, error) {

	clientsPath := arguments["velodir"].(string) + "\\clients"
	items, _ := os.ReadDir(clientsPath)
	artifactPaths := make([]string, 0)
	for _, item := range items {
		if !item.IsDir() {
			continue
		}
		if item.Name() == "server" {
			continue
		}
		artifactPaths = append(artifactPaths, clientsPath+"\\"+item.Name()+"\\artifacts")
	}
	return artifactPaths, nil
}

func GetAllArtifactsInPath(path string) ([]string, error) {
	artifactPaths := make([]string, 0)
	items, _ := os.ReadDir(path)
	for _, item := range items {
		if !item.IsDir() {
			continue
		}
		artifactPaths = append(artifactPaths, path+"\\"+item.Name())
	}
	return artifactPaths, nil
}

func ProcessClientArtifactPath(path string, arguments map[string]any, clientWaiter *sync.WaitGroup, logger zerolog.Logger) error {
	defer clientWaiter.Done()
	artifactWaiter := sync.WaitGroup{}
	baseClientPath, _ := filepath.Split(path)
	clientIdentifier := filepath.Base(baseClientPath)
	outputFile := arguments["outputdir"].(string) + "\\" + clientIdentifier + ".csv"

	//fmt.Println(outputFile)

	artifactPaths, err := GetAllArtifactsInPath(path)
	if err != nil {
		return err
	}

	outputF, err := helpers.CreateOutput(outputFile)
	if err != nil {
		logger.Error().Msg(err.Error())
	}
	writer := csv.NewWriter(outputF)
	// Base CSV Headers - later on if we want more details/full parse need to adjust this
	writer.Write(vars.BaseHeaders)

	recordChannel := make(chan []string)
	// This waitgroup is purely for the csv writer and nothing else
	var writeWG sync.WaitGroup
	writeWG.Add(1)
	go helpers.ListenOnWriteChannel(recordChannel, writer, logger, outputF, arguments["writebuffer"].(int), &writeWG)
	for _, artifactDir := range artifactPaths {
		//go ProcessArtifact(artifactDir, arguments, &artifactWaiter, recordChannel, logger, clientIdentifier)
		paths := helpers.GetAllJSONFromDirectory(artifactDir)
		for _, artifactJSON := range paths {
			artifactWaiter.Add(1)
			go ProcessArtifactFile(artifactDir, arguments, &artifactWaiter, recordChannel, logger, clientIdentifier, artifactJSON)
		}
	}
	helpers.CloseChannelWhenDone(recordChannel, &artifactWaiter)
	writeWG.Wait()
	return nil
}

func ProcessArtifact(artifactDir string, arguments map[string]any, artifactWaiter *sync.WaitGroup, recordChannel chan []string, logger zerolog.Logger, clientIdentifier string) {

	// Here we will actually read the relevant artifact file (if it exists) and process records through the appropriate artifact helper func
	paths := helpers.GetAllJSONFromDirectory(artifactDir)
	defer artifactWaiter.Done()
	for _, artifactJSON := range paths {
		artifactWaiter.Add(1)
		go ProcessArtifactFile(artifactDir, arguments, artifactWaiter, recordChannel, logger, clientIdentifier, artifactJSON)
	}

	/*	items, _ := os.ReadDir(artifactDir)
		artifactJSON := ""
		for _, item := range items {
			if strings.HasSuffix(item.Name(), ".json") {
				artifactJSON = artifactDir + "\\" + item.Name()
			}
		}
		if artifactJSON == "" {
			return
		}*/
	// Checking if this artifact is implemented

}

func ProcessArtifactFile(artifactDir string, arguments map[string]any, artifactWaiter *sync.WaitGroup, recordOutputChannel chan []string, logger zerolog.Logger, clientIdentifier string, artifactJSON string) {
	defer artifactWaiter.Done()

	_, implemented := vars.ImplementedArtifacts[filepath.Base(artifactDir)]
	if !implemented {
		logger.Info().Msgf("Skipping (not implemented): %v", artifactJSON)
		return
	}
	logger.Info().Msgf("Processing: %v", artifactJSON)

	maxRoutinesPerFile := arguments["maxgoperfile"].(int)
	lineBatchSize := arguments["batchsize"].(int)
	//idx := 0
	jobTracker := vars.RunningJobs{
		//	JobCount: 0,
		//	Mw:       sync.RWMutex{},
	}
	records := make([]string, 0)

	// Open the Relevant File - error otherwise
	inputFile, readErr := os.Open(artifactJSON)
	if readErr != nil {
		logger.Error().Err(readErr)
	}
	scanner := bufio.NewScanner(inputFile)

	var readerWaitGroup sync.WaitGroup
	for scanner.Scan() {
		// We have each line of a file - we read them into a buffer of fixed size then send this buffer to the appropriate service func to handle unmarshalling/processing
		// We wait until that is finished to continue to the next buffered set to avoid over-reading into memory
		//fmt.Printf(scanner.Text() + "\n")
		records = append(records, scanner.Text())
		if len(records) <= lineBatchSize {
			continue
		} else {
			if jobTracker.GetJobs() >= maxRoutinesPerFile {
			waitForOthers:
				for {
					if jobTracker.GetJobs() >= maxRoutinesPerFile {
						continue
					} else {
						readerWaitGroup.Add(1)
						jobTracker.AddJob()
						go SendRecordsToAppropriateBus(logger, records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker, filepath.Base(artifactDir), clientIdentifier, artifactJSON)
						//go artifact_structs.Process_Windows_System_Service(records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker)
						//go helpers.ProcessRecords(logger, records, asnDB, cityDB, countryDB, domainDB, ipAddressColumn, jsonColumn, arguments["regex"].(bool), arguments["dns"].(bool), recordChannel, &fileWG, &jobTracker, tempArgs, dateindex)
						break waitForOthers
					}
				}
			} else {
				readerWaitGroup.Add(1)
				jobTracker.AddJob()
				go SendRecordsToAppropriateBus(logger, records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker, filepath.Base(artifactDir), clientIdentifier, artifactJSON)
				//go helpers.ProcessRecords(logger, records, asnDB, cityDB, countryDB, domainDB, ipAddressColumn, jsonColumn, arguments["regex"].(bool), arguments["dns"].(bool), recordChannel, &fileWG, &jobTracker, tempArgs, dateindex)
			}
			records = nil
		}
	}
	readerWaitGroup.Add(1)
	go SendRecordsToAppropriateBus(logger, records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker, filepath.Base(artifactDir), clientIdentifier, artifactJSON)
	readerWaitGroup.Wait()
}

func SendRecordsToAppropriateBus(logger zerolog.Logger, records []string, recordOutputChannel chan<- []string, arguments map[string]any, wg *sync.WaitGroup, jobs *vars.RunningJobs, artifactName string, clientIdentifier string, artifactFile string) {

	defer wg.Done()
	defer jobs.SubJob()
	JSONFileName := filepath.Base(artifactFile)
	if artifactName == "Windows.System.Services" {
		artifact_structs.Process_Windows_System_Service("Windows.System.Services", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Timeline.Prefetch" {
		artifact_structs.Process_Windows_Timeline_Prefetch("Windows.Timeline.Prefetch", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Timeline.Registry.RunMRU" {
		artifact_structs.Process_Windows_Timeline_Registry_RunMRU("Windows.Timeline.Registry.RunMRU", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.System.Amcache" {
		if JSONFileName == "InventoryApplicationFile.json" {
			artifact_structs.Process_Windows_System_Amcache_InventoryApplicationFile("Windows.System.Amcache", clientIdentifier, records, recordOutputChannel, arguments)
		} else {
			logger.Error().Msgf("Amcache File Not Implemented: %v", JSONFileName)
		}
	} else if artifactName == "Windows.Sysinternals.Autoruns" {
		artifact_structs.Process_Windows_Sysinternals_Autoruns("Windows.Sysinternals.Autoruns", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Registry.UserAssist" {
		artifact_structs.Process_Windows_Registry_UserAssist("Windows.Registry.UserAssist", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Registry.Sysinternals.Eulacheck" {
		artifact_structs.Process_Windows_Registry_Sysinternals_Eulacheck("Windows.Registry.Sysinternals.Eulacheck", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Registry.RDP" {
		if JSONFileName == "Servers.json" {
			artifact_structs.Process_Windows_Registry_RDP_Servers("Windows.Registry.RDP", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Mru.json" {
			artifact_structs.Process_Windows_Registry_RDP_Mru("Windows.Registry.RDP", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if artifactName == "Windows.Registry.AppCompatCache" {
		artifact_structs.Process_Windows_Registry_AppCompatCache("Windows.Registry.AppCompatCache", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Network.NetstatEnriched" {
		artifact_structs.Process_Windows_Network_NetstatEnriched("Windows.Network.NetstatEnriched", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.KapeFiles.Targets" {
		if JSONFileName == "All File Metadata.json" {
			artifact_structs.Process_Windows_KapeFiles_Targets_AllFileMetadata("Windows.KapeFiles.Targets", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if artifactName == "Windows.Forensics.Timeline" {
		artifact_structs.Process_Windows_Forensics_Timeline("Windows.Forensics.Timeline", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Forensics.SRUM" {
		if JSONFileName == "Application Resource Usage.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_ApplicationResourceUsage("Windows.Forensics.SRUM", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Execution Stats.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_ExecutionStats("Windows.Forensics.SRUM", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Network Usage.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_NetworkUsage("Windows.Forensics.SRUM", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Network Connections.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_NetworkConnections("Windows.Forensics.SRUM", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if artifactName == "Windows.Forensics.Shellbags" {
		artifact_structs.Process_Windows_Forensics_Shellbags("Windows.Forensics.Shellbags", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Forensics.RecycleBin" {
		artifact_structs.Process_Windows_Forensics_RecycleBin("Windows.Forensics.RecycleBin", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Forensics.RDPCache" {
		artifact_structs.Process_Windows_Forensics_RDPCache("Windows.Forensics.RDPCache", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Forensics.Lnk" {
		artifact_structs.Process_Windows_Forensics_Lnk("Windows.Forensics.Lnk", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Forensics.CertUtil" {
		artifact_structs.Process_Windows_Forensics_CertUtil("Windows.Forensics.CertUtil", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Forensics.Bam" {
		artifact_structs.Process_Windows_Forensics_Bam("Windows.Forensics.Bam", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.EventLogs.AlternateLogon" {
		artifact_structs.Process_Windows_EventLogs_AlternateLogon("Windows.Forensics.AlternateLogon", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Exchange.Windows.Office.MRU" {
		artifact_structs.Process_Exchange_Windows_Office_MRU("Exchange.Windows.Office.MRU", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Exchange.Windows.EventLogs.RDPClientActivity" {
		artifact_structs.Process_Exchange_Windows_EventLogs_RDPClientActivity("Exchange.Windows.EventLogs.RDPClientActivity", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Exchange.Windows.EventLogs.LogonSessions" {
		artifact_structs.Process_Exchange_Windows_EventLogs_LogonSessions("Exchange.Windows.EventLogs.LogonSessions", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Exchange.Windows.EventLogs.Bitsadmin" {
		artifact_structs.Process_Exchange_Windows_EventLogs_Bitsadmin("Exchange.Windows.EventLogs.Bitsadmin", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Custom.Windows.Eventlog.Evtx" {
		artifact_structs.Process_Custom_Windows_Eventlog_Evtx("Custom.Windows.Eventlog.Evtx", clientIdentifier, records, recordOutputChannel, arguments)
	} else if strings.HasPrefix(artifactName, "Custom.Windows.Mft") {
		if arguments["mftlight"].(bool) || arguments["mftfull"].(bool) {
			artifact_structs.Process_Custom_Windows_MFT("Custom.Windows.MFT", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if strings.HasPrefix(artifactName, "Generic.Client.Info") {
		if JSONFileName == "Users.json" {
			artifact_structs.Process_Generic_Client_Info_Users("Generic.Client.Info", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "BasicInformation.json" {
			artifact_structs.Process_Generic_Client_Info_BasicInformation("Generic.Client.Info", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "WindowsInfo.json" {
			artifact_structs.Process_Generic_Client_Info_WindowsInfo("Generic.Client.Info", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if artifactName == "Windows.Applications.Chrome.History" {
		artifact_structs.Process_Windows_Applications_Chrome_History("Windows.Applications.Chrome.History", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Applications.Edge.History" {
		artifact_structs.Process_Windows_Applications_Edge_History("Windows.Applications.Edge.History", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Applications.Firefox.Downloads" {
		artifact_structs.Process_Windows_Applications_Firefox_Downloads("Windows.Applications.Firefox.Downloads", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Applications.Firefox.History" {
		artifact_structs.Process_Windows_Applications_Firefox_History("Windows.Applications.Firefox.History", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.Applications.NirsoftBrowserViewer" {
		artifact_structs.Process_Windows_Applications_NirsoftBrowserViewer("Windows.Applications.NirsoftBrowserViewer", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.EventLogs.PowershellScriptblock" {
		artifact_structs.Process_Windows_EventLogs_PowerShellScriptblock("Windows.EventLogs.PowershellScriptblock", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.EventLogs.RDPAuth" {
		artifact_structs.Process_Windows_EventLogs_RDPAuth("Windows.EventLogs.RDPAuth", clientIdentifier, records, recordOutputChannel, arguments)
	} else if strings.HasPrefix(artifactName, "Windows.Forensics.SAM") {
		if JSONFileName == "CreateTimes.json" {
			artifact_structs.Process_Windows_Forensics_SAM_CreateTimes("Windows.Forensics.SAM", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Parsed.json" {
			artifact_structs.Process_Windows_Forensics_SAM_Parsed("Windows.Forensics.SAM", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if strings.HasPrefix(artifactName, "Generic.Forensic.SQLiteHunter") {
		if JSONFileName == "Chromium Browser Bookmarks.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser Favicons.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser History_Downloads.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser History_Keywords.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser History_Visits.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser Shortcuts.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser Extensions.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Extensions("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser Network_Predictor.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Network_Predictor("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Chromium Browser Top Sites.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Top_Sites("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Firefox Cookies.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Cookies("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Firefox Favicons.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Favicons("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Firefox Form History.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Form_History("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Firefox Places.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Places("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Firefox Places_Downloads.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Places_Downloads("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Firefox Places_History.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Places_History("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "IE or Edge WebCacheV01_All Data.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "IE or Edge WebCacheV01_Highlights.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_Highlights("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Windows Search Service_SystemIndex_Gthr.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Windows Search Service_SystemIndex_PropertyStore.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Windows Search Service_SystemIndex_GthrPth.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if strings.HasPrefix(artifactName, "Windows.Sys.Drivers") {
		if JSONFileName == "SignedDrivers.json" {
			artifact_structs.Process_Windows_Sys_Drivers_SignedDrivers("Windows.Sys.Drivers", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "RunningDrivers.json" {
			artifact_structs.Process_Windows_Sys_Drivers_RunningDrivers("Windows.Sys.Drivers", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if artifactName == "Windows.System.Powershell.ModuleAnalysisCache" {
		artifact_structs.Process_Windows_System_Powershell_ModuleAnalysisCache("Windows.System.Powershell.ModuleAnalysisCache", clientIdentifier, records, recordOutputChannel, arguments)
	} else if strings.HasPrefix(artifactName, "Windows.Analysis.EvidenceOfExecution") {
		if JSONFileName == "Amcache.json" {
			artifact_structs.Process_Windows_Analysis_EvidenceOfExecution_Amcache("Windows.Analysis.EvidenceOfExecution", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "UserAssist.json" {
			artifact_structs.Process_Windows_Registry_UserAssist("Windows.Analysis.EvidenceOfExecution", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if artifactName == "Windows.Analysis.EvidenceOfDownload" {
		artifact_structs.Process_Windows_Analysis_EvidenceOfDownload("Windows.Analysis.EvidenceOfDownload", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Windows.EventLogs.Evtx" {
		artifact_structs.Process_Windows_EventLogs_Evtx("Windows.EventLogs.Evtx", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Exchange.Windows.EventLogs.Chainsaw" {
		artifact_structs.Process_Exchange_Windows_EventLogs_Chainsaw("Exchange.Windows.EventLogs.Chainsaw", clientIdentifier, records, recordOutputChannel, arguments)
	} else if strings.HasPrefix(artifactName, "Exchange.Windows.Memory.InjectedThreadEx") {
		if JSONFileName != "RawResults.json" {
			artifact_structs.Process_Exchange_Windows_Memory_InjectedThreadEx("Exchange.Windows.Memory.InjectedThreadEx", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "RawResults.json" {
			artifact_structs.Process_Exchange_Windows_Memory_InjectedThreadEx_RawResults("Exchange.Windows.Memory.InjectedThreadEx", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if strings.HasPrefix(artifactName, "Exchange.Windows.EventLogs.Hayabusa") {
		if JSONFileName == "Results.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_Hayabusa("Exchange.Windows.EventLogs.Hayabusa", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if artifactName == "Exchange.Windows.Forensics.Trawler" {
		artifact_structs.Process_Exchange_Windows_Forensics_Trawler("Exchange.Windows.Forensics.Trawler", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Exchange.Windows.Forensics.PersistenceSniper" {
		artifact_structs.Process_Exchange_Windows_Forensics_PersistenceSniper("Exchange.Windows.Forensics.PersistenceSniper", clientIdentifier, records, recordOutputChannel, arguments)
	} else if artifactName == "Exchange.Windows.EventLogs.CondensedAccountUsage" {
		artifact_structs.Process_Exchange_Windows_EventLogs_CondensedAccountUsage("Exchange.Windows.EventLogs.CondensedAccountUsage", clientIdentifier, records, recordOutputChannel, arguments)
	} else if strings.HasPrefix(artifactName, "Exchange.Windows.EventLogs.EvtxHussar") {
		if JSONFileName == "Accounts_UsersRelatedOperations.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Antivirus_WindowsDefender.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "BootupRestartShutdown.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Logons.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Logons("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Powershell_Events.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Powershell_ScriptblocksSummary.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "RDP.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_RDP("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Services.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Services("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "SMB_ClientDestinations.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "SMB_ServerAccessAudit.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "SMB_ServerModifications.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "WindowsFirewall.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "WinRM.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_WinRM("Exchange.Windows.EventLogs.EvtxHussar", clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	}
}

func SetupArtifactListenChannels(clientArtifactPaths []string, logger zerolog.Logger, arguments map[string]any) {
	// Receives slice containing paths to each detected artifact directory - builds out the maps required to orchestrate feeding artifacts to the appropriate channels
	// Also starts a listener for each channel

	var writeWG sync.WaitGroup
	artifactWaiter := sync.WaitGroup{}
	for _, path := range clientArtifactPaths {
		baseClientPath, _ := filepath.Split(path)
		clientIdentifier := filepath.Base(baseClientPath)
		artifactPaths, err := GetAllArtifactsInPath(path)
		if err != nil {
			continue
		}
		for _, artifactPath := range artifactPaths {
			artifactName := filepath.Base(artifactPath)
			_, implemented := vars.ImplementedArtifacts[artifactName]
			if !implemented {
				//logger.Info().Msgf("Skipping (not implemented): %v", artifactName)
				continue
			}
			JSONFilePaths := helpers.GetAllJSONFromDirectory(artifactPath)
			for _, JsonFile := range JSONFilePaths {
				os.Mkdir("artifact_output", 0777)
				outputDir := "artifact_output\\"

				outputFile := ""
				if strings.HasPrefix(filepath.Base(JsonFile), "F.") {
					// We are using base output artifact name for this since there are no 'special' files in this artifact
					outputFile += fmt.Sprintf("%v.csv", artifactName)
				} else {
					outputFile += fmt.Sprintf("%v.%v.csv", artifactName, strings.TrimSuffix(filepath.Base(JsonFile), filepath.Ext(filepath.Base(JsonFile))))
				}
				outputFile = strings.Replace(outputFile, " ", "_", -1)

				// Get CSV headers for current artifact

				// Setup CSV - now this is hard because each CSV will have separate headers - so realistically what we want to do is allow the artifact to handle the headers on this or have some helper function that perform this
				// Probably easiest way is to have a map of outputFile to column headers and just write this here then start handing off to the appropriate handler using artifactdump, easy
				// Base CSV Headers - later on if we want more details/full parse need to adjust this
				// Create handler to the destination output file
				outputF, Outputerr := helpers.CreateOutput(outputDir + outputFile)
				if Outputerr != nil {
					logger.Error().Msg(Outputerr.Error())
					continue
				}
				artifactHeaders, headerError := GetAppropriateHeaders(filepath.Base(outputFile))
				if headerError != nil {
					logger.Error().Msg(headerError.Error())
					continue
				}
				// Check if we have encountered this artifact output file name before
				// If not, initialize new channel specific to this file
				// TODO - I think this will be a race condition if the first client artifact finishes before the rest are queued up
				// Could be solved by queuing up and Done() a WG for this for loop to ensure it gets through everything before closing - or moving the go helper.closechannell... to iterate all channels and start up threads after for loop
				artifactWaiter.Add(1)
				artifactRecordChannel, exists := vars.ArtifactToChannelMap[filepath.Base(outputFile)]
				if !exists {
					vars.ArtifactToChannelMap[filepath.Base(outputFile)] = make(chan []string)
					artifactRecordChannel, _ = vars.ArtifactToChannelMap[filepath.Base(outputFile)]
					writer := csv.NewWriter(outputF)
					writer.Write(artifactHeaders)
					writeWG.Add(1)
					logger.Info().Msgf("Setting Up Channel: %v", outputDir+outputFile)
					go helpers.ListenOnWriteChannel(artifactRecordChannel, writer, logger, outputF, arguments["writebuffer"].(int), &writeWG)
					// We only want to close a channel once we are done processing all files in their entirety
					go helpers.CloseChannelWhenDone(artifactRecordChannel, &artifactWaiter)
				}
				go ProcessArtifactFile(artifactName, arguments, &artifactWaiter, artifactRecordChannel, logger, clientIdentifier, JsonFile)
			}
		}
	}
	writeWG.Wait()

}

func GetAppropriateHeaders(artifact string) ([]string, error) {
	// TODO - Reformat this to use a generic map of struct where each implements an appropriate interface func
	// TODO - Replace with switch
	headers := []string{"Time", "ClientID", "Hostname"}
	if artifact == "Windows.System.Services.csv" {
		headers = append(headers, artifact_structs.Windows_System_Service.GetHeaders(artifact_structs.Windows_System_Service{})...)
	} else if artifact == "Windows.Timeline.Prefetch.csv" {
		headers = append(headers, artifact_structs.Windows_Timeline_Prefetch.GetHeaders(artifact_structs.Windows_Timeline_Prefetch{})...)
	} else if artifact == "Windows.Timeline.Registry.RunMRU.csv" {
		headers = append(headers, artifact_structs.Windows_Timeline_Registry_RunMRU.GetHeaders(artifact_structs.Windows_Timeline_Registry_RunMRU{})...)
	} else if artifact == "Windows.System.Powershell.ModuleAnalysisCache.csv" {
		headers = append(headers, artifact_structs.Windows_System_Powershell_ModuleAnalysisCache.GetHeaders(artifact_structs.Windows_System_Powershell_ModuleAnalysisCache{})...)
	} else if artifact == "Windows.System.Amcache.InventoryApplicationFile.csv" {
		headers = append(headers, artifact_structs.Windows_System_Amcache_InventoryApplicationFile.GetHeaders(artifact_structs.Windows_System_Amcache_InventoryApplicationFile{})...)
	} else if artifact == "Windows.Analysis.EvidenceOfExecution.Amcache.csv" {
		headers = append(headers, artifact_structs.Windows_Analysis_EvidenceOfExecution_Amcache.GetHeaders(artifact_structs.Windows_Analysis_EvidenceOfExecution_Amcache{})...)
	} else if artifact == "Windows.Analysis.EvidenceOfExecution.UserAssist.csv" {
		headers = append(headers, artifact_structs.Windows_Analysis_EvidenceOfExecution_UserAssist.GetHeaders(artifact_structs.Windows_Analysis_EvidenceOfExecution_UserAssist{})...)
	} else if artifact == "Windows.Sysinternals.Autoruns.csv" {
		headers = append(headers, artifact_structs.Windows_Sysinternals_Autoruns.GetHeaders(artifact_structs.Windows_Sysinternals_Autoruns{})...)
	} else if artifact == "Windows.Sys.Drivers.SignedDrivers.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_Drivers_SignedDrivers.GetHeaders(artifact_structs.Windows_Sys_Drivers_SignedDrivers{})...)
	} else if artifact == "Windows.Sys.Drivers.RunningDrivers.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_Drivers_RunningDrivers.GetHeaders(artifact_structs.Windows_Sys_Drivers_RunningDrivers{})...)
	} else if artifact == "Windows.Registry.UserAssist.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_UserAssist.GetHeaders(artifact_structs.Windows_Registry_UserAssist{})...)
	} else if artifact == "Windows.Registry.Sysinternals.Eulacheck.RegistryAPI.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_Sysinternals_Eulacheck.GetHeaders(artifact_structs.Windows_Registry_Sysinternals_Eulacheck{})...)
	} else if artifact == "Windows.Registry.RDP.Servers.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_RDP_Servers.GetHeaders(artifact_structs.Windows_Registry_RDP_Servers{})...)
	} else if artifact == "Windows.Registry.RDP.Mru.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_RDP_Mru.GetHeaders(artifact_structs.Windows_Registry_RDP_Mru{})...)
	} else if artifact == "Windows.Registry.AppCompatCache.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_AppCompatCache.GetHeaders(artifact_structs.Windows_Registry_AppCompatCache{})...)
	} else if artifact == "Windows.Network.NetstatEnriched.Netstat.csv" {
		headers = append(headers, artifact_structs.Windows_Network_NetstatEnriched.GetHeaders(artifact_structs.Windows_Network_NetstatEnriched{})...)
	} else if artifact == "Windows.KapeFiles.Targets.All_File_Metadata.csv" {
		headers = append(headers, artifact_structs.Windows_KapeFiles_Targets_AllFileMetadata.GetHeaders(artifact_structs.Windows_KapeFiles_Targets_AllFileMetadata{})...)
	} else if artifact == "Windows.Forensics.Timeline.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_Timeline.GetHeaders(artifact_structs.Windows_Forensics_Timeline{})...)
	} else if artifact == "Windows.Forensics.SRUM.Application_Resource_Usage.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_SRUM_ApplicationResourceUsage.GetHeaders(artifact_structs.Windows_Forensics_SRUM_ApplicationResourceUsage{})...)
	} else if artifact == "Windows.Forensics.SRUM.Execution_Stats.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_SRUM_ExecutionStats.GetHeaders(artifact_structs.Windows_Forensics_SRUM_ExecutionStats{})...)
	} else if artifact == "Windows.Forensics.SRUM.Network_Usage.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_SRUM_NetworkUsage.GetHeaders(artifact_structs.Windows_Forensics_SRUM_NetworkUsage{})...)
	} else if artifact == "Windows.Forensics.SRUM.Network_Connections.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_SRUM_NetworkConnections.GetHeaders(artifact_structs.Windows_Forensics_SRUM_NetworkConnections{})...)
	} else if artifact == "Windows.Forensics.Shellbags.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_Shellbags.GetHeaders(artifact_structs.Windows_Forensics_Shellbags{})...)
	} else if artifact == "Windows.Forensics.SAM.CreateTimes.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_SAM_CreateTimes.GetHeaders(artifact_structs.Windows_Forensics_SAM_CreateTimes{})...)
	} else if artifact == "Windows.Forensics.SAM.Parsed.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_SAM_Parsed.GetHeaders(artifact_structs.Windows_Forensics_SAM_Parsed{})...)
	} else if artifact == "Windows.Forensics.RecycleBin.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_RecycleBin.GetHeaders(artifact_structs.Windows_Forensics_RecycleBin{})...)
	} else if artifact == "Windows.Forensics.RDPCache.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_RDPCache.GetHeaders(artifact_structs.Windows_Forensics_RDPCache{})...)
	} else if artifact == "Windows.Forensics.Lnk.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_Lnk.GetHeaders(artifact_structs.Windows_Forensics_Lnk{})...)
	} else if artifact == "Windows.Forensics.CertUtil.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_CertUtil.GetHeaders(artifact_structs.Windows_Forensics_CertUtil{})...)
	} else if artifact == "Windows.Forensics.Bam.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_Bam.GetHeaders(artifact_structs.Windows_Forensics_Bam{})...)
	} else if artifact == "Windows.EventLogs.AlternateLogon.csv" {
		headers = append(headers, artifact_structs.Windows_EventLogs_AlternateLogon.GetHeaders(artifact_structs.Windows_EventLogs_AlternateLogon{})...)
	} else if artifact == "Windows.EventLogs.RDPAuth.csv" {
		headers = append(headers, artifact_structs.Windows_EventLogs_RDPAuth.GetHeaders(artifact_structs.Windows_EventLogs_RDPAuth{})...)
	} else if artifact == "Windows.EventLogs.PowershellScriptblock.csv" {
		headers = append(headers, artifact_structs.Windows_EventLogs_PowerShellScriptblock.GetHeaders(artifact_structs.Windows_EventLogs_PowerShellScriptblock{})...)
	} else if artifact == "Windows.EventLogs.Evtx.csv" {
		headers = append(headers, artifact_structs.Windows_EventLogs_Evtx.GetHeaders(artifact_structs.Windows_EventLogs_Evtx{})...)
	} else if artifact == "Windows.Applications.NirsoftBrowserViewer.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_NirsoftBrowserViewer.GetHeaders(artifact_structs.Windows_Applications_NirsoftBrowserViewer{})...)
	} else if artifact == "Windows.Applications.Firefox.Downloads.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_Firefox_Downloads.GetHeaders(artifact_structs.Windows_Applications_Firefox_Downloads{})...)
	} else if artifact == "Windows.Applications.Firefox.History.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_Firefox_History.GetHeaders(artifact_structs.Windows_Applications_Firefox_History{})...)
	} else if artifact == "Windows.Applications.Edge.History.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_Edge_History.GetHeaders(artifact_structs.Windows_Applications_Edge_History{})...)
	} else if artifact == "Windows.Applications.Chrome.History.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_Chrome_History.GetHeaders(artifact_structs.Windows_Applications_Chrome_History{})...)
	} else if artifact == "Windows.Analysis.EvidenceOfDownload.csv" {
		headers = append(headers, artifact_structs.Windows_Analysis_EvidenceOfDownload.GetHeaders(artifact_structs.Windows_Analysis_EvidenceOfDownload{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_Bookmarks.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_Favicons.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_History_Downloads.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_History_Keywords.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_History_Visits.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_Shortcuts.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Firefox_Cookies.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Cookies.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Cookies{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Firefox_Form_History.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Form_History.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Form_History{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.IE_or_Edge_WebCacheV01_All_Data.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Windows_Search_Service_SystemIndex_Gthr.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Windows_Search_Service_SystemIndex_GthrPth.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Windows_Search_Service_SystemIndex_PropertyStore.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore{})...)
	} else if artifact == "Generic.Client.Info.Users.csv" {
		headers = append(headers, artifact_structs.Generic_Client_Info_Users.GetHeaders(artifact_structs.Generic_Client_Info_Users{})...)
	} else if artifact == "Generic.Client.Info.BasicInformation.csv" {
		headers = append(headers, artifact_structs.Generic_Client_Info_BasicInformation.GetHeaders(artifact_structs.Generic_Client_Info_BasicInformation{})...)
	} else if artifact == "Generic.Client.Info.WindowsInfo.csv" {
		headers = append(headers, artifact_structs.Generic_Client_Info_WindowsInfo.GetHeaders(artifact_structs.Generic_Client_Info_WindowsInfo{})...)
	} else if artifact == "Exchange.Windows.Office.MRU.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Office_MRU.GetHeaders(artifact_structs.Exchange_Windows_Office_MRU{})...)
	} else if artifact == "Exchange.Windows.Memory.InjectedThreadEx.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Memory_InjectedThreadEx.GetHeaders(artifact_structs.Exchange_Windows_Memory_InjectedThreadEx{})...)
	} else if artifact == "Exchange.Windows.EventLogs.RDPClientActivity.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_RDPClientActivity.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_RDPClientActivity{})...)
	} else if artifact == "Exchange.Windows.Memory.InjectedThreadEx.RawResults.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Memory_InjectedThreadEx_RawResults.GetHeaders(artifact_structs.Exchange_Windows_Memory_InjectedThreadEx_RawResults{})...)
	} else if artifact == "Exchange.Windows.EventLogs.LogonSessions.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_LogonSessions.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_LogonSessions{})...)
	} else if artifact == "Exchange.Windows.EventLogs.Hayabusa.Results.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_Hayabusa.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_Hayabusa{})...)
	} else if artifact == "Exchange.Windows.EventLogs.Chainsaw.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_Chainsaw.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_Chainsaw{})...)
	} else if artifact == "Exchange.Windows.EventLogs.Bitsadmin.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_Bitsadmin.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_Bitsadmin{})...)
	} else if strings.HasPrefix(artifact, "Custom.Windows.Mft") {
		headers = append(headers, artifact_structs.Custom_Windows_MFT.GetHeaders(artifact_structs.Custom_Windows_MFT{})...)
	} else if artifact == "Custom.Windows.Eventlog.Evtx.csv" {
		headers = append(headers, artifact_structs.Custom_Windows_Eventlog_Evtx.GetHeaders(artifact_structs.Custom_Windows_Eventlog_Evtx{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_Extensions.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Extensions.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Extensions{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_Network_Predictor.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Network_Predictor.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Network_Predictor{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Chromium_Browser_Top_Sites.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Top_Sites.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Chromium_Browser_Top_Sites{})...)
	} else if artifact == "Exchange.Windows.Forensics.Trawler.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_Trawler.GetHeaders(artifact_structs.Exchange_Windows_Forensics_Trawler{})...)
	} else if artifact == "Exchange.Windows.Forensics.PersistenceSniper.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_PersistenceSniper.GetHeaders(artifact_structs.Exchange_Windows_Forensics_PersistenceSniper{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Firefox_Favicons.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Favicons.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Favicons{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Firefox_Places.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Places.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Places{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Firefox_Places_Downloads.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Places_Downloads.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Places_Downloads{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.Firefox_Places_History.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Places_History.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_Firefox_Places_History{})...)
	} else if artifact == "Generic.Forensic.SQLiteHunter.IE_or_Edge_WebCacheV01_Highlights.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_Highlights.GetHeaders(artifact_structs.Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_Highlights{})...)
	} else if artifact == "Exchange.Windows.EventLogs.CondensedAccountUsage.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_CondensedAccountUsage.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_CondensedAccountUsage{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.Accounts_UsersRelatedOperations.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.Antivirus_WindowsDefender.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.BootupRestartShutdown.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.Logons.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Logons.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Logons{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.Powershell_Events.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.Powershell_ScriptblocksSummary.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.RDP.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_RDP.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_RDP{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.Services.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Services.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_Services{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.SMB_ClientDestinations.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.SMB_ServerAccessAudit.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.SMB_ServerModifications.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.WindowsFirewall.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall{})...)
	} else if artifact == "Exchange.Windows.EventLogs.EvtxHussar.WinRM.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_WinRM.GetHeaders(artifact_structs.Exchange_Windows_EventLogs_EvtxHussar_WinRM{})...)
	}

	if len(headers) == 3 {
		return headers, errors.New("Artifact Not Implemented: " + artifact)
	}
	return headers, nil

}

func main() {
	logger := helpers.SetupLogger()
	arguments, err := parseArgs(logger)
	if err != nil {
		logger.Error().Err(err)
		return
	}

	velo_dir_err := ValidateVelociraptorDirectory(arguments["velodir"].(string))
	if velo_dir_err != nil {
		logger.Error().Msgf(velo_dir_err.Error())
		return
	}
	clientArtifactPaths, velo_clients_err := ProcessAllClients(arguments)
	if velo_clients_err != nil {
		logger.Error().Msgf(velo_clients_err.Error())
		return
	}
	clientWaiter := sync.WaitGroup{}
	// We are either doing an artifact dump to individual CSVs (client-agnostic) or a super timeline per client - not both
	if arguments["artifactdump"].(bool) {
		// populate all possible artifact channels - each destination file path has it's own channel
		SetupArtifactListenChannels(clientArtifactPaths, logger, arguments)

	} else {
		for _, path := range clientArtifactPaths {
			//
			clientWaiter.Add(1)
			ProcessClientArtifactPath(path, arguments, &clientWaiter, logger)
		}
	}

	clientWaiter.Wait()
	logger.Info().Msgf("Done!")

}
