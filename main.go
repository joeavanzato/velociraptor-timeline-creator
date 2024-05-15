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
	mftfull := flag.Bool("mftfull", false, "Do not restrict parsing of MFT-related artifacts (Windows.Timeline.MFT, Windows.NTFS.MFT, etc)")
	artifactdump := flag.Bool("artifactdump", false, "Instead of creating per-client super-timelines, dump all artifacts present into aggregated CSV per-artifact across all clients.")

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

	if arguments["mftlight"].(bool) && arguments["mftfull"].(bool) {
		return arguments, fmt.Errorf(" Both -mftlight and -mftfull enabled - pick one!")
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
	emptyHeaders := make([]string, 0)
	go helpers.ListenOnWriteChannel(recordChannel, writer, logger, outputF, arguments["writebuffer"].(int), &writeWG)
	for _, artifactDir := range artifactPaths {
		//go ProcessArtifact(artifactDir, arguments, &artifactWaiter, recordChannel, logger, clientIdentifier)
		paths := helpers.GetAllJSONFromDirectory(artifactDir)
		for _, artifactJSON := range paths {
			artifactWaiter.Add(1)
			go ProcessArtifactFile(artifactDir, arguments, &artifactWaiter, recordChannel, logger, clientIdentifier, artifactJSON, emptyHeaders)
		}
	}
	helpers.CloseChannelWhenDone(recordChannel, &artifactWaiter)
	writeWG.Wait()
	return nil
}

func ProcessArtifact(artifactDir string, arguments map[string]any, artifactWaiter *sync.WaitGroup, recordChannel chan []string, logger zerolog.Logger, clientIdentifier string) {

	// Here we will actually read the relevant artifact file (if it exists) and process records through the appropriate artifact helper func
	/*	paths := helpers.GetAllJSONFromDirectory(artifactDir)
		defer artifactWaiter.Done()
		for _, artifactJSON := range paths {
			artifactWaiter.Add(1)
			go ProcessArtifactFile(artifactDir, arguments, artifactWaiter, recordChannel, logger, clientIdentifier, artifactJSON)
		}*/

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

func ProcessArtifactFile(artifactDir string, arguments map[string]any, artifactWaiter *sync.WaitGroup, recordOutputChannel chan []string, logger zerolog.Logger, clientIdentifier string, artifactJSON string, genericHeaders []string) {
	defer artifactWaiter.Done()

	_, implemented := vars.ImplementedArtifacts[filepath.Base(artifactDir)]
	if !implemented && !arguments["artifactdump"].(bool) {
		logger.Info().Msgf("Skipping (not implemented): %v", artifactJSON)
		return
	} else if !implemented && arguments["artifactdump"].(bool) {
		logger.Info().Msgf("Processing (Generic): %v", artifactJSON)
	} else {
		logger.Info().Msgf("Processing: %v", artifactJSON)
	}

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
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	// TODO - instead of just maxing buffer, let's just wait until we find a long-line error then expand accordingly or read-as-slices and combine for that specific line

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
						go SendRecordsToAppropriateBus(logger, records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker, filepath.Base(artifactDir), clientIdentifier, artifactJSON, genericHeaders)
						//go artifact_structs.Process_Windows_System_Service(records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker)
						//go helpers.ProcessRecords(logger, records, asnDB, cityDB, countryDB, domainDB, ipAddressColumn, jsonColumn, arguments["regex"].(bool), arguments["dns"].(bool), recordChannel, &fileWG, &jobTracker, tempArgs, dateindex)
						break waitForOthers
					}
				}
			} else {
				readerWaitGroup.Add(1)
				jobTracker.AddJob()
				go SendRecordsToAppropriateBus(logger, records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker, filepath.Base(artifactDir), clientIdentifier, artifactJSON, genericHeaders)
				//go helpers.ProcessRecords(logger, records, asnDB, cityDB, countryDB, domainDB, ipAddressColumn, jsonColumn, arguments["regex"].(bool), arguments["dns"].(bool), recordChannel, &fileWG, &jobTracker, tempArgs, dateindex)
			}
			records = nil
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Error().Err(err)
	}

	readerWaitGroup.Add(1)
	go SendRecordsToAppropriateBus(logger, records, recordOutputChannel, arguments, &readerWaitGroup, &jobTracker, filepath.Base(artifactDir), clientIdentifier, artifactJSON, genericHeaders)
	readerWaitGroup.Wait()
}

func SendRecordsToAppropriateBus(logger zerolog.Logger, records []string, recordOutputChannel chan<- []string, arguments map[string]any, wg *sync.WaitGroup, jobs *vars.RunningJobs, artifactName string, clientIdentifier string, artifactFile string, genericHeaders []string) {
	defer wg.Done()
	defer jobs.SubJob()
	JSONFileName := filepath.Base(artifactFile)
	if artifactName == "Windows.System.Services" {
		artifact_structs.Process_Windows_System_Service(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Timeline.Prefetch" {
		artifact_structs.Process_Windows_Timeline_Prefetch(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Timeline.Registry.RunMRU" {
		artifact_structs.Process_Windows_Timeline_Registry_RunMRU(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.Amcache" {
		if JSONFileName == "InventoryApplicationFile.json" {
			artifact_structs.Process_Windows_System_Amcache_InventoryApplicationFile(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else {
			logger.Error().Msgf("Amcache File Not Implemented: %v", JSONFileName)
		}
	} else if artifactName == "Windows.Sysinternals.Autoruns" {
		artifact_structs.Process_Windows_Sysinternals_Autoruns(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Registry.UserAssist" {
		artifact_structs.Process_Windows_Registry_UserAssist(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Registry.Sysinternals.Eulacheck" {
		artifact_structs.Process_Windows_Registry_Sysinternals_Eulacheck(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Registry.RDP" {
		if JSONFileName == "Servers.json" {
			artifact_structs.Process_Windows_Registry_RDP_Servers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Mru.json" {
			artifact_structs.Process_Windows_Registry_RDP_Mru(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Windows.Registry.AppCompatCache" {
		artifact_structs.Process_Windows_Registry_AppCompatCache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Network.NetstatEnriched" {
		artifact_structs.Process_Windows_Network_NetstatEnriched(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.KapeFiles.Targets" {
		if JSONFileName == "All File Metadata.json" {
			artifact_structs.Process_Windows_KapeFiles_Targets_AllFileMetadata(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Windows.Forensics.Timeline" {
		artifact_structs.Process_Windows_Forensics_Timeline(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.SRUM" {
		if JSONFileName == "Application Resource Usage.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_ApplicationResourceUsage(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Execution Stats.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_ExecutionStats(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Network Usage.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_NetworkUsage(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Network Connections.json" {
			artifact_structs.Process_Windows_Forensics_SRUM_NetworkConnections(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Windows.Forensics.Shellbags" {
		artifact_structs.Process_Windows_Forensics_Shellbags(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.RecycleBin" {
		artifact_structs.Process_Windows_Forensics_RecycleBin(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.RDPCache" {
		artifact_structs.Process_Windows_Forensics_RDPCache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.Lnk" {
		artifact_structs.Process_Windows_Forensics_Lnk(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.CertUtil" {
		artifact_structs.Process_Windows_Forensics_CertUtil(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.Bam" {
		artifact_structs.Process_Windows_Forensics_Bam(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.EventLogs.AlternateLogon" {
		artifact_structs.Process_Windows_EventLogs_AlternateLogon(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Office.MRU" {
		artifact_structs.Process_Exchange_Windows_Office_MRU(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.EventLogs.RDPClientActivity" {
		artifact_structs.Process_Exchange_Windows_EventLogs_RDPClientActivity(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.EventLogs.LogonSessions" {
		artifact_structs.Process_Exchange_Windows_EventLogs_LogonSessions(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.EventLogs.Bitsadmin" {
		artifact_structs.Process_Exchange_Windows_EventLogs_Bitsadmin(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Custom.Windows.Eventlog.Evtx" {
		artifact_structs.Process_Custom_Windows_Eventlog_Evtx(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if strings.HasPrefix(artifactName, "Custom.Windows.Mft") {
		if arguments["mftlight"].(bool) || arguments["mftfull"].(bool) {
			artifact_structs.Process_Custom_Windows_MFT("Custom.Windows.MFT", clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if strings.HasPrefix(artifactName, "Generic.Client.Info") {
		if JSONFileName == "Users.json" {
			artifact_structs.Process_Generic_Client_Info_Users(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "BasicInformation.json" {
			artifact_structs.Process_Generic_Client_Info_BasicInformation(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "WindowsInfo.json" {
			artifact_structs.Process_Generic_Client_Info_WindowsInfo(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Windows.Applications.Chrome.History" {
		artifact_structs.Process_Windows_Applications_Chrome_History(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Applications.Edge.History" {
		artifact_structs.Process_Windows_Applications_Edge_History(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Applications.Firefox.Downloads" {
		artifact_structs.Process_Windows_Applications_Firefox_Downloads(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Applications.Firefox.History" {
		artifact_structs.Process_Windows_Applications_Firefox_History(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Applications.NirsoftBrowserViewer" {
		artifact_structs.Process_Windows_Applications_NirsoftBrowserViewer(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.EventLogs.PowershellScriptblock" {
		artifact_structs.Process_Windows_EventLogs_PowerShellScriptblock(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.EventLogs.RDPAuth" {
		artifact_structs.Process_Windows_EventLogs_RDPAuth(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if strings.HasPrefix(artifactName, "Windows.Forensics.SAM") {
		if JSONFileName == "CreateTimes.json" {
			artifact_structs.Process_Windows_Forensics_SAM_CreateTimes(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Parsed.json" {
			artifact_structs.Process_Windows_Forensics_SAM_Parsed(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Generic.Forensic.SQLiteHunter" {
		if JSONFileName == "Chromium Browser Bookmarks.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Bookmarks(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser Favicons.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Favicons(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser History_Downloads.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Downloads(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser History_Keywords.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Keywords(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser History_Visits.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_History_Visits(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser Shortcuts.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Chromium_Browser_Shortcuts(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser Extensions.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Extensions(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser Network_Predictor.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Network_Predictor(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Chromium Browser Top Sites.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Chromium_Browser_Top_Sites(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Firefox Cookies.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Cookies(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Firefox Favicons.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Favicons(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Firefox Form History.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Form_History(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Firefox Places.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Places(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Firefox Places_Downloads.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Places_Downloads(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Firefox Places_History.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Places_History(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "IE or Edge WebCacheV01_All Data.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "IE or Edge WebCacheV01_Highlights.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_Highlights(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Windows Search Service_SystemIndex_Gthr.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Windows Search Service_SystemIndex_PropertyStore.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Windows Search Service_SystemIndex_GthrPth.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_GthrPth(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if strings.HasPrefix(artifactName, "Windows.Sys.Drivers") {
		if JSONFileName == "SignedDrivers.json" {
			artifact_structs.Process_Windows_Sys_Drivers_SignedDrivers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "RunningDrivers.json" {
			artifact_structs.Process_Windows_Sys_Drivers_RunningDrivers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Windows.System.Powershell.ModuleAnalysisCache" {
		artifact_structs.Process_Windows_System_Powershell_ModuleAnalysisCache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if strings.HasPrefix(artifactName, "Windows.Analysis.EvidenceOfExecution") {
		if JSONFileName == "Amcache.json" {
			artifact_structs.Process_Windows_Analysis_EvidenceOfExecution_Amcache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "UserAssist.json" {
			artifact_structs.Process_Windows_Registry_UserAssist(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Windows.Analysis.EvidenceOfDownload" {
		artifact_structs.Process_Windows_Analysis_EvidenceOfDownload(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.EventLogs.Evtx" {
		artifact_structs.Process_Windows_EventLogs_Evtx(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.EventLogs.Chainsaw" {
		artifact_structs.Process_Exchange_Windows_EventLogs_Chainsaw(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if strings.HasPrefix(artifactName, "Exchange.Windows.Memory.InjectedThreadEx") {
		if JSONFileName != "RawResults.json" {
			artifact_structs.Process_Exchange_Windows_Memory_InjectedThreadEx(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "RawResults.json" {
			artifact_structs.Process_Exchange_Windows_Memory_InjectedThreadEx_RawResults(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if strings.HasPrefix(artifactName, "Exchange.Windows.EventLogs.Hayabusa") {
		if JSONFileName == "Results.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_Hayabusa(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "Exchange.Windows.Forensics.Trawler" {
		artifact_structs.Process_Exchange_Windows_Forensics_Trawler(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.PersistenceSniper" {
		artifact_structs.Process_Exchange_Windows_Forensics_PersistenceSniper(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.EventLogs.CondensedAccountUsage" {
		artifact_structs.Process_Exchange_Windows_EventLogs_CondensedAccountUsage(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if strings.HasPrefix(artifactName, "Exchange.Windows.EventLogs.EvtxHussar") {
		if JSONFileName == "Accounts_UsersRelatedOperations.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Accounts_UserRelatedOperations(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Antivirus_WindowsDefender.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Antivirus_WindowsDefender(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "BootupRestartShutdown.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_BootupRestartShutdown(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Logons.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Logons(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Powershell_Events.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_PowerShell_Events(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Powershell_ScriptblocksSummary.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Powershell_ScriptblocksSummary(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "RDP.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_RDP(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "Services.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_Services(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "SMB_ClientDestinations.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ClientDestinations(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "SMB_ServerAccessAudit.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerAccessAudit(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "SMB_ServerModifications.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_SMB_ServerModifications(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "WindowsFirewall.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_WindowsFirewall(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		} else if JSONFileName == "WinRM.json" {
			artifact_structs.Process_Exchange_Windows_EventLogs_EvtxHussar_WinRM(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
		}
	} else if artifactName == "DetectRaptor.Windows.Detection.MFT" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_MFT(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Generic.Detection.WebshellYara" {
		artifact_structs.Process_DetectRaptor_Generic_Detection_WebshellYara(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.Amcache" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_Amcache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.Applications" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_Applications(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.BinaryRename" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_BinaryRename(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.Webhistory" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_Webhistory(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.Evtx" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_Evtx(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.Powershell.ISEAutoSave" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.Powershell.PSReadline" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_PowerShell_PSReadline(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.ZoneIdentifier" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_ZoneIdentifier(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.HijackLibsEnv" && JSONFileName == "Suspicious Dll path.json" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.HijackLibsMFT" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_HijackLibsMFT(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.LolDriversVulnerable" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_LolDriversVulnerable(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "DetectRaptor.Windows.Detection.NamedPipes" {
		artifact_structs.Process_DetectRaptor_Windows_Detection_NamedPipes(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Detection.PipeHunter" {
		artifact_structs.Process_Exchange_Windows_Detection_PipeHunter(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Memory.ProcessInfo" {
		artifact_structs.Process_Windows_Memory_ProcessInfo(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.UEFI" {
		artifact_structs.Process_Exchange_Windows_Forensics_UEFI(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.Jumplists_JLECmd" {
		artifact_structs.Process_Exchange_Windows_Forensics_Jumplists_JLECmd(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.ThumbCache" {
		artifact_structs.Process_Exchange_Windows_Forensics_ThumbCache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.UEFI.BootApplication" {
		artifact_structs.Process_Exchange_Windows_Forensics_UEFI_BootApplication(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Custom.Windows.Nirsoft.LastActivityView" {
		artifact_structs.Process_Exchange_Custom_Windows_Nirsoft_LastActivityView(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.Clipboard" {
		artifact_structs.Process_Exchange_Windows_Forensics_Clipboard(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.FileZilla" && JSONFileName == "FileZilla.json" {
		artifact_structs.Process_Exchange_Windows_Forensics_FileZilla_FileZilla(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Forensics.FileZilla" && JSONFileName == "RecentServers.json" {
		artifact_structs.Process_Exchange_Windows_Forensics_FileZilla_RecentServers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Registry.WDigest" {
		artifact_structs.Process_Windows_Registry_WDigest(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.System.PrinterDriver" && JSONFileName == "BinaryCheck.json" {
		artifact_structs.Process_Exchange_Windows_System_PrinterDriver_BinaryCheck(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.System.PrinterDriver" && strings.HasPrefix(JSONFileName, "F.") {
		artifact_structs.Process_Exchange_Windows_System_PrinterDriver(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Detection.Malfind" {
		artifact_structs.Process_Exchange_Windows_Detection_Malfind(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Applications.OfficeMacros" {
		artifact_structs.Process_Windows_Applications_OfficeMacros(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Detection.Mutants" && JSONFileName == "Handles.json" {
		artifact_structs.Process_Windows_Detection_Mutants_Handles(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Detection.Mutants" && JSONFileName == "ObjectTree.json" {
		artifact_structs.Process_Windows_Detection_Mutants_ObjectTree(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Detection.BinaryHunter" {
		artifact_structs.Process_Windows_Detection_BinaryHunter(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Detection.Impersonation" {
		artifact_structs.Process_Windows_Detection_Impersonation(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Detection.PrefetchHunter" {
		artifact_structs.Process_Exchange_Windows_Detection_PrefetchHunter(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Detection.ForwardedImports" {
		artifact_structs.Process_Windows_Detection_ForwardedImports(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Detection.Amcache" {
		artifact_structs.Process_Windows_Detection_Amcache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.Detection.Yara.Zip" {
		artifact_structs.Process_Generic_Detection_Yara_Zip(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.DLLs" {
		artifact_structs.Process_Windows_System_DLLs(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.DNSCache" {
		artifact_structs.Process_Windows_System_DNSCache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.HostsFile" {
		artifact_structs.Process_Windows_System_HostsFile(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.LocalAdmins" {
		artifact_structs.Process_Windows_System_LocalAdmins(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.Powershell.PSReadline" {
		artifact_structs.Process_Windows_System_Powershell_PSReadline(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.Pslist" {
		artifact_structs.Process_Windows_System_Pslist(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.TaskScheduler" && JSONFileName == "Analysis.json" {
		artifact_structs.Process_Windows_System_TaskScheduler(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.StartupItems" {
		artifact_structs.Process_Windows_Sys_StartupItems(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.Interfaces" {
		artifact_structs.Process_Windows_Sys_Interfaces(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.FirewallRules" {
		artifact_structs.Process_Windows_Sys_FirewallRules(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.AllUsers" {
		artifact_structs.Process_Windows_Sys_AllUsers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Registry.PuttyHostKeys" {
		artifact_structs.Process_Windows_Registry_PuttyHostKeys(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Persistence.PermanentWMIEvents" {
		artifact_structs.Process_Windows_Persistence_PermanentWMIEvents(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Network.ArpCache" {
		artifact_structs.Process_Windows_Network_ArpCache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.EventLogs.PowershellModule" {
		artifact_structs.Process_Windows_EventLogs_PowershellModule(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.EventLogs.Modifications" && JSONFileName == "Channels.json" {
		artifact_structs.Process_Windows_EventLogs_Modifications_Channels(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.EventLogs.Modifications" && JSONFileName == "Providers.json" {
		artifact_structs.Process_Windows_EventLogs_Modifications_Providers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Applications.Chrome.Extensions" {
		artifact_structs.Process_Windows_Applications_Chrome_Extensions(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Applications.ChocolateyPackages" {
		artifact_structs.Process_Windows_Applications_ChocolateyPackages(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Network.ExternalIpAddress" {
		artifact_structs.Process_Network_ExternalIpAddress(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.Network.InterfaceAddresses" {
		artifact_structs.Process_Generic_Network_InterfaceAddresses(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.System.WMIProviders" {
		artifact_structs.Process_Exchange_Windows_System_WMIProviders(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Applications.OfficeServerCache" {
		artifact_structs.Process_Exchange_Windows_Applications_OfficeServerCache(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Applications.LECmd" {
		artifact_structs.Process_Exchange_Windows_Applications_LECmd(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.HashRunKeys" {
		artifact_structs.Process_Exchange_HashRunKeys(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.System.Pstree" {
		artifact_structs.Process_Generic_System_Pstree(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Registry.NetshHelperDLLs" {
		artifact_structs.Process_Exchange_Windows_Registry_NetshHelperDLLs(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Registry.Domain" {
		artifact_structs.Process_Exchange_Windows_Registry_Domain(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Registry.COMAutoApprovalList" {
		artifact_structs.Process_Exchange_Windows_Registry_COMAutoApprovalList(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Registry.BackupRestore" {
		artifact_structs.Process_Exchange_Windows_Registry_BackupRestore(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.Client.DiskSpace" {
		artifact_structs.Process_Generic_Client_DiskSpace(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Registry.CapabilityAccessManager" {
		artifact_structs.Process_Exchange_Windows_Registry_CapabilityAccessManager(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.System.Powershell.ISEAutoSave" && strings.HasPrefix(JSONFileName, "F.") {
		artifact_structs.Process_Exchange_Windows_System_Powershell_ISEAutoSave(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.System.Powershell.ISEAutoSave" && JSONFileName == "UserConfig.json" {
		artifact_structs.Process_Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Sys.LoggedInUsers" {
		artifact_structs.Process_Exchange_Windows_Sys_LoggedInUsers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Registry.ScheduledTasks" {
		artifact_structs.Process_Exchange_Windows_Registry_ScheduledTasks(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.System.WindowsErrorReporting" && JSONFileName == "AppCrashReport.json" {
		artifact_structs.Process_Exchange_Windows_System_WindowsErrorReporting(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.NTFS.Timestomp" {
		artifact_structs.Process_Exchange_Windows_NTFS_Timestomp(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.System.BinaryVersion" {
		artifact_structs.Process_Windows_Detection_BinaryVersion(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.NTFS.MFT" && (arguments["mftlight"].(bool) || arguments["mftfull"].(bool)) {
		artifact_structs.Process_Windows_NTFS_MFT(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.Usn" {
		artifact_structs.Process_Windows_Forensics_Usn(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Timeline.MFT" && (arguments["mftlight"].(bool) || arguments["mftfull"].(bool)) {
		artifact_structs.Process_Windows_Timeline_MFT(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Carving.USN" {
		artifact_structs.Process_Windows_Carving_USN(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.Users" {
		artifact_structs.Process_Windows_Sys_Users(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.Handles" {
		artifact_structs.Process_Windows_System_Handles(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.Shares" {
		artifact_structs.Process_Windows_System_Shares(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.DiskInfo" {
		artifact_structs.Process_Windows_Sys_DiskInfo(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.AuditPolicy" {
		artifact_structs.Process_Windows_System_AuditPolicy(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Network.ListeningPorts" {
		artifact_structs.Process_Windows_Network_ListeningPorts(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.PhysicalMemoryRanges" {
		artifact_structs.Process_Windows_Sys_PhysicalMemoryRanges(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.Programs" {
		artifact_structs.Process_Windows_Sys_Programs(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Network.Netstat" {
		artifact_structs.Process_Windows_Network_Netstat(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.Forensic.Timeline" {
		artifact_structs.Process_Generic_Forensic_Timeline(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Forensics.PartitionTable" {
		artifact_structs.Process_Windows_Forensics_PartitionTable(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.System.ProcessSiblings" {
		artifact_structs.Process_Generic_System_ProcessSiblings(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.RootCAStore" {
		artifact_structs.Process_Windows_System_RootCAStore(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Sys.CertificateAuthorities" {
		artifact_structs.Process_Windows_Sys_CertificateAuthorities(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.Applications.Chrome.SessionStorage" {
		artifact_structs.Process_Generic_Applications_Chrome_SessionStorage(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Registry.NTUser" {
		artifact_structs.Process_Windows_Registry_NTUser(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.CatFiles" {
		artifact_structs.Process_Windows_System_CatFiles(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Applications.DefenderDHParser" {
		artifact_structs.Process_Exchange_Windows_Applications_DefenderDHParser(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Registry.RecentDocs" {
		artifact_structs.Process_Windows_Registry_RecentDocs(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.Client.DiskUsage" {
		artifact_structs.Process_Generic_Client_DiskUsage(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Generic.Applications.Office.Keywords" {
		artifact_structs.Process_Generic_Applications_Office_Keywords(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.System.Signers" {
		artifact_structs.Process_Windows_System_Signers(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Windows.Detection.EnvironmentVariables" {
		artifact_structs.Process_Windows_Detection_EnvironmentVariables(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else if artifactName == "Exchange.Windows.Timeline.Prefetch.Improved" {
		artifact_structs.Process_Exchange_Windows_Timeline_Prefetch_Improved(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger)
	} else {
		artifact_structs.Process_Generic_Artifact(artifactName, clientIdentifier, records, recordOutputChannel, arguments, logger, genericHeaders)
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

			/*			if artifactName != "Test.Generic.Parser" {
						continue
					}*/

			_, implemented := vars.ImplementedArtifacts[artifactName]
			if !implemented {
				logger.Info().Msgf("No Parsing Implementation, Trying Generic Dump: %v", artifactName)
			}
			if (artifactName == "Windows.Timeline.MFT" || artifactName == "Windows.NTFS.MFT" || strings.HasPrefix(artifactName, "Custom.Windows.MFT")) && !arguments["mftlight"].(bool) && !arguments["mftfull"].(bool) {
				logger.Info().Msgf("No MFT Flag Enabled, Skipped: %v", artifactName)
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
				artifactHeaders, headerError := GetAppropriateHeaders(filepath.Base(outputFile), JsonFile)
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
					//logger.Info().Msgf("Setting Up Channel: %v", outputDir+outputFile)
					go helpers.ListenOnWriteChannel(artifactRecordChannel, writer, logger, outputF, arguments["writebuffer"].(int), &writeWG)
					// We only want to close a channel once we are done processing all files in their entirety
					go helpers.CloseChannelWhenDone(artifactRecordChannel, &artifactWaiter)
				}
				go ProcessArtifactFile(artifactName, arguments, &artifactWaiter, artifactRecordChannel, logger, clientIdentifier, JsonFile, artifactHeaders)
			}
		}
	}
	writeWG.Wait()

}

func GetAppropriateHeaders(artifact string, inputFile string) ([]string, error) {
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
	} else if artifact == "DetectRaptor.Windows.Detection.MFT.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_MFT.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_MFT{})...)
	} else if artifact == "DetectRaptor.Generic.Detection.WebshellYara.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Generic_Detection_WebshellYara.GetHeaders(artifact_structs.DetectRaptor_Generic_Detection_WebshellYara{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.Amcache.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_Amcache.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_Amcache{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.Applications.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_Applications.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_Applications{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.BinaryRename.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_BinaryRename.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_BinaryRename{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.Webhistory.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_Webhistory.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_Webhistory{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.Evtx.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_Evtx.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_Evtx{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.Powershell.ISEAutoSave.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_PowerShell_ISEAutoSave{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.Powershell.PSReadline.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_PowerShell_PSReadline.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_PowerShell_PSReadline{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.ZoneIdentifier.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_ZoneIdentifier.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_ZoneIdentifier{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.HijackLibsEnv.Suspicious_Dll_path.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_HijackLibsEnv_SuspiciousDLLPath{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.HijackLibsMFT.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_HijackLibsMFT.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_HijackLibsMFT{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.LolDriversVulnerable.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_LolDriversVulnerable.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_LolDriversVulnerable{})...)
	} else if artifact == "DetectRaptor.Windows.Detection.NamedPipes.csv" {
		headers = append(headers, artifact_structs.DetectRaptor_Windows_Detection_NamedPipes.GetHeaders(artifact_structs.DetectRaptor_Windows_Detection_NamedPipes{})...)
	} else if artifact == "Exchange.Windows.Detection.PipeHunter.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Detection_PipeHunter.GetHeaders(artifact_structs.Exchange_Windows_Detection_PipeHunter{})...)
	} else if artifact == "Windows.Memory.ProcessInfo.csv" {
		headers = append(headers, artifact_structs.Windows_Memory_ProcessInfo.GetHeaders(artifact_structs.Windows_Memory_ProcessInfo{})...)
	} else if artifact == "Exchange.Windows.Forensics.UEFI.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_UEFI.GetHeaders(artifact_structs.Exchange_Windows_Forensics_UEFI{})...)
	} else if artifact == "Exchange.Windows.Forensics.Jumplists_JLECmd.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_Jumplists_JLECmd.GetHeaders(artifact_structs.Exchange_Windows_Forensics_Jumplists_JLECmd{})...)
	} else if artifact == "Exchange.Windows.Forensics.ThumbCache.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_ThumbCache.GetHeaders(artifact_structs.Exchange_Windows_Forensics_ThumbCache{})...)
	} else if artifact == "Exchange.Windows.Forensics.UEFI.BootApplication.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_UEFI_BootApplication.GetHeaders(artifact_structs.Exchange_Windows_Forensics_UEFI_BootApplication{})...)
	} else if artifact == "Exchange.Custom.Windows.Nirsoft.LastActivityView.Upload.csv" {
		headers = append(headers, artifact_structs.Exchange_Custom_Windows_Nirsoft_LastActivityView.GetHeaders(artifact_structs.Exchange_Custom_Windows_Nirsoft_LastActivityView{})...)
	} else if artifact == "Exchange.Windows.Forensics.Clipboard.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_Clipboard.GetHeaders(artifact_structs.Exchange_Windows_Forensics_Clipboard{})...)
	} else if artifact == "Exchange.Windows.Forensics.FileZilla.FileZilla.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_FileZilla_FileZilla.GetHeaders(artifact_structs.Exchange_Windows_Forensics_FileZilla_FileZilla{})...)
	} else if artifact == "Exchange.Windows.Forensics.FileZilla.RecentServers.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Forensics_FileZilla_RecentServers.GetHeaders(artifact_structs.Exchange_Windows_Forensics_FileZilla_RecentServers{})...)
	} else if artifact == "Windows.Registry.WDigest.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_WDigest.GetHeaders(artifact_structs.Windows_Registry_WDigest{})...)
	} else if artifact == "Exchange.Windows.System.PrinterDriver.BinaryCheck.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_System_PrinterDriver_BinaryCheck.GetHeaders(artifact_structs.Exchange_Windows_System_PrinterDriver_BinaryCheck{})...)
	} else if artifact == "Exchange.Windows.System.PrinterDriver.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_System_PrinterDriver.GetHeaders(artifact_structs.Exchange_Windows_System_PrinterDriver{})...)
	} else if artifact == "Exchange.Windows.Detection.Malfind.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Detection_Malfind.GetHeaders(artifact_structs.Exchange_Windows_Detection_Malfind{})...)
	} else if artifact == "Windows.Applications.OfficeMacros.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_OfficeMacros.GetHeaders(artifact_structs.Windows_Applications_OfficeMacros{})...)
	} else if artifact == "Windows.Detection.Mutants.Handles.csv" {
		headers = append(headers, artifact_structs.Windows_Detection_Mutants_Handles.GetHeaders(artifact_structs.Windows_Detection_Mutants_Handles{})...)
	} else if artifact == "Windows.Detection.Mutants.ObjectTree.csv" {
		headers = append(headers, artifact_structs.Windows_Detection_Mutants_ObjectTree.GetHeaders(artifact_structs.Windows_Detection_Mutants_ObjectTree{})...)
	} else if artifact == "Windows.Detection.BinaryHunter.csv" {
		headers = append(headers, artifact_structs.Windows_Detection_BinaryHunter.GetHeaders(artifact_structs.Windows_Detection_BinaryHunter{})...)
	} else if artifact == "Windows.Detection.Impersonation.csv" {
		headers = append(headers, artifact_structs.Windows_Detection_Impersonation.GetHeaders(artifact_structs.Windows_Detection_Impersonation{})...)
	} else if artifact == "Exchange.Windows.Detection.PrefetchHunter.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Detection_PrefetchHunter.GetHeaders(artifact_structs.Exchange_Windows_Detection_PrefetchHunter{})...)
	} else if artifact == "Windows.Detection.ForwardedImports.csv" {
		headers = append(headers, artifact_structs.Windows_Detection_ForwardedImports.GetHeaders(artifact_structs.Windows_Detection_ForwardedImports{})...)
	} else if artifact == "Windows.Detection.Amcache.csv" {
		headers = append(headers, artifact_structs.Windows_Detection_Amcache.GetHeaders(artifact_structs.Windows_Detection_Amcache{})...)
	} else if artifact == "Generic.Detection.Yara.Zip.csv" {
		headers = append(headers, artifact_structs.Generic_Detection_Yara_Zip.GetHeaders(artifact_structs.Generic_Detection_Yara_Zip{})...)
	} else if artifact == "Windows.System.DLLs.csv" {
		headers = append(headers, artifact_structs.Windows_System_DLLs.GetHeaders(artifact_structs.Windows_System_DLLs{})...)
	} else if artifact == "Windows.System.DNSCache.csv" {
		headers = append(headers, artifact_structs.Windows_System_DNSCache.GetHeaders(artifact_structs.Windows_System_DNSCache{})...)
	} else if artifact == "Windows.System.HostsFile.csv" {
		headers = append(headers, artifact_structs.Windows_System_HostsFile.GetHeaders(artifact_structs.Windows_System_HostsFile{})...)
	} else if artifact == "Windows.System.LocalAdmins.csv" {
		headers = append(headers, artifact_structs.Windows_System_LocalAdmins.GetHeaders(artifact_structs.Windows_System_LocalAdmins{})...)
	} else if artifact == "Windows.System.Powershell.PSReadline.csv" {
		headers = append(headers, artifact_structs.Windows_System_Powershell_PSReadline.GetHeaders(artifact_structs.Windows_System_Powershell_PSReadline{})...)
	} else if artifact == "Windows.System.Pslist.csv" {
		headers = append(headers, artifact_structs.Windows_System_Pslist.GetHeaders(artifact_structs.Windows_System_Pslist{})...)
	} else if artifact == "Windows.System.TaskScheduler.Analysis.csv" {
		headers = append(headers, artifact_structs.Windows_System_TaskScheduler.GetHeaders(artifact_structs.Windows_System_TaskScheduler{})...)
	} else if artifact == "Windows.Sys.StartupItems.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_StartupItems.GetHeaders(artifact_structs.Windows_Sys_StartupItems{})...)
	} else if artifact == "Windows.Sys.Interfaces.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_Interfaces.GetHeaders(artifact_structs.Windows_Sys_Interfaces{})...)
	} else if artifact == "Windows.Sys.FirewallRules.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_FirewallRules.GetHeaders(artifact_structs.Windows_Sys_FirewallRules{})...)
	} else if artifact == "Windows.Sys.AllUsers.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_AllUsers.GetHeaders(artifact_structs.Windows_Sys_AllUsers{})...)
	} else if artifact == "Windows.Registry.PuttyHostKeys.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_PuttyHostKeys.GetHeaders(artifact_structs.Windows_Registry_PuttyHostKeys{})...)
	} else if artifact == "Windows.Persistence.PermanentWMIEvents.csv" {
		headers = append(headers, artifact_structs.Windows_Persistence_PermanentWMIEvents.GetHeaders(artifact_structs.Windows_Persistence_PermanentWMIEvents{})...)
	} else if artifact == "Windows.Network.ArpCache.csv" {
		headers = append(headers, artifact_structs.Windows_Network_ArpCache.GetHeaders(artifact_structs.Windows_Network_ArpCache{})...)
	} else if artifact == "Windows.EventLogs.PowershellModule.csv" {
		headers = append(headers, artifact_structs.Windows_EventLogs_PowershellModule.GetHeaders(artifact_structs.Windows_EventLogs_PowershellModule{})...)
	} else if artifact == "Windows.EventLogs.Modifications.Channels.csv" {
		headers = append(headers, artifact_structs.Windows_EventLogs_Modifications_Channels.GetHeaders(artifact_structs.Windows_EventLogs_Modifications_Channels{})...)
	} else if artifact == "Windows.EventLogs.Modifications.Providers.csv" {
		headers = append(headers, artifact_structs.Windows_EventLogs_Modifications_Providers.GetHeaders(artifact_structs.Windows_EventLogs_Modifications_Providers{})...)
	} else if artifact == "Windows.Applications.Chrome.Extensions.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_Chrome_Extensions.GetHeaders(artifact_structs.Windows_Applications_Chrome_Extensions{})...)
	} else if artifact == "Windows.Applications.ChocolateyPackages.csv" {
		headers = append(headers, artifact_structs.Windows_Applications_ChocolateyPackages.GetHeaders(artifact_structs.Windows_Applications_ChocolateyPackages{})...)
	} else if artifact == "Network.ExternalIpAddress.csv" {
		headers = append(headers, artifact_structs.Network_ExternalIpAddress.GetHeaders(artifact_structs.Network_ExternalIpAddress{})...)
	} else if artifact == "Generic.Network.InterfaceAddresses.csv" {
		headers = append(headers, artifact_structs.Generic_Network_InterfaceAddresses.GetHeaders(artifact_structs.Generic_Network_InterfaceAddresses{})...)
	} else if artifact == "Exchange.Windows.System.WMIProviders.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_System_WMIProviders.GetHeaders(artifact_structs.Exchange_Windows_System_WMIProviders{})...)
	} else if artifact == "Exchange.Windows.Applications.OfficeServerCache.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Applications_OfficeServerCache.GetHeaders(artifact_structs.Exchange_Windows_Applications_OfficeServerCache{})...)
	} else if artifact == "Exchange.Windows.Applications.LECmd.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Applications_LECmd.GetHeaders(artifact_structs.Exchange_Windows_Applications_LECmd{})...)
	} else if artifact == "Exchange.HashRunKeys.csv" {
		headers = append(headers, artifact_structs.Exchange_HashRunKeys.GetHeaders(artifact_structs.Exchange_HashRunKeys{})...)
	} else if artifact == "Generic.System.Pstree.csv" {
		headers = append(headers, artifact_structs.Generic_System_Pstree.GetHeaders(artifact_structs.Generic_System_Pstree{})...)
	} else if artifact == "Exchange.Windows.Registry.NetshHelperDLLs.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Registry_NetshHelperDLLs.GetHeaders(artifact_structs.Exchange_Windows_Registry_NetshHelperDLLs{})...)
	} else if artifact == "Exchange.Windows.Registry.Domain.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Registry_Domain.GetHeaders(artifact_structs.Exchange_Windows_Registry_Domain{})...)
	} else if artifact == "Exchange.Windows.Registry.COMAutoApprovalList.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Registry_COMAutoApprovalList.GetHeaders(artifact_structs.Exchange_Windows_Registry_COMAutoApprovalList{})...)
	} else if artifact == "Exchange.Windows.Registry.BackupRestore.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Registry_BackupRestore.GetHeaders(artifact_structs.Exchange_Windows_Registry_BackupRestore{})...)
	} else if artifact == "Generic.Client.DiskSpace.csv" {
		headers = append(headers, artifact_structs.Generic_Client_DiskSpace.GetHeaders(artifact_structs.Generic_Client_DiskSpace{})...)
	} else if artifact == "Exchange.Windows.Registry.CapabilityAccessManager.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Registry_CapabilityAccessManager.GetHeaders(artifact_structs.Exchange_Windows_Registry_CapabilityAccessManager{})...)
	} else if artifact == "Exchange.Windows.System.Powershell.ISEAutoSave.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_System_Powershell_ISEAutoSave.GetHeaders(artifact_structs.Exchange_Windows_System_Powershell_ISEAutoSave{})...)
	} else if artifact == "Exchange.Windows.System.Powershell.ISEAutoSave.UserConfig.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig.GetHeaders(artifact_structs.Exchange_Windows_System_Powershell_ISEAutoSave_UserConfig{})...)
	} else if artifact == "Exchange.Windows.Sys.LoggedInUsers.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Sys_LoggedInUsers.GetHeaders(artifact_structs.Exchange_Windows_Sys_LoggedInUsers{})...)
	} else if artifact == "Exchange.Windows.Registry.ScheduledTasks.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Registry_ScheduledTasks.GetHeaders(artifact_structs.Exchange_Windows_Registry_ScheduledTasks{})...)
	} else if artifact == "Exchange.Windows.System.WindowsErrorReporting.AppCrashReport.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_System_WindowsErrorReporting.GetHeaders(artifact_structs.Exchange_Windows_System_WindowsErrorReporting{})...)
	} else if artifact == "Exchange.Windows.NTFS.Timestomp.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_NTFS_Timestomp.GetHeaders(artifact_structs.Exchange_Windows_NTFS_Timestomp{})...)
	} else if artifact == "Exchange.Windows.System.BinaryVersion.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_System_BinaryVersion.GetHeaders(artifact_structs.Exchange_Windows_System_BinaryVersion{})...)
	} else if artifact == "Windows.NTFS.MFT.csv" {
		headers = append(headers, artifact_structs.Windows_NTFS_MFT.GetHeaders(artifact_structs.Windows_NTFS_MFT{})...)
	} else if artifact == "Windows.Forensics.Usn.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_Usn.GetHeaders(artifact_structs.Windows_Forensics_Usn{})...)
	} else if artifact == "Windows.Timeline.MFT.csv" {
		headers = append(headers, artifact_structs.Windows_Timeline_MFT.GetHeaders(artifact_structs.Windows_Timeline_MFT{})...)
	} else if artifact == "Windows.Carving.USN.csv" {
		headers = append(headers, artifact_structs.Windows_Carving_USN.GetHeaders(artifact_structs.Windows_Carving_USN{})...)
	} else if artifact == "Windows.Sys.Users.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_Users.GetHeaders(artifact_structs.Windows_Sys_Users{})...)
	} else if artifact == "Windows.System.Handles.csv" {
		headers = append(headers, artifact_structs.Windows_System_Handles.GetHeaders(artifact_structs.Windows_System_Handles{})...)
	} else if artifact == "Windows.System.Shares.csv" {
		headers = append(headers, artifact_structs.Windows_System_Shares.GetHeaders(artifact_structs.Windows_System_Shares{})...)
	} else if artifact == "Windows.Sys.DiskInfo.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_DiskInfo.GetHeaders(artifact_structs.Windows_Sys_DiskInfo{})...)
	} else if artifact == "Windows.System.AuditPolicy.csv" {
		headers = append(headers, artifact_structs.Windows_System_AuditPolicy.GetHeaders(artifact_structs.Windows_System_AuditPolicy{})...)
	} else if artifact == "Windows.Network.ListeningPorts.csv" {
		headers = append(headers, artifact_structs.Windows_Network_ListeningPorts.GetHeaders(artifact_structs.Windows_Network_ListeningPorts{})...)
	} else if artifact == "Windows.Sys.PhysicalMemoryRanges.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_PhysicalMemoryRanges.GetHeaders(artifact_structs.Windows_Sys_PhysicalMemoryRanges{})...)
	} else if artifact == "Windows.Sys.Programs.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_Programs.GetHeaders(artifact_structs.Windows_Sys_Programs{})...)
	} else if artifact == "Windows.Network.Netstat.csv" {
		headers = append(headers, artifact_structs.Windows_Network_Netstat.GetHeaders(artifact_structs.Windows_Network_Netstat{})...)
	} else if artifact == "Generic.Forensic.Timeline.csv" {
		headers = append(headers, artifact_structs.Generic_Forensic_Timeline.GetHeaders(artifact_structs.Generic_Forensic_Timeline{})...)
	} else if artifact == "Windows.Forensics.PartitionTable.csv" {
		headers = append(headers, artifact_structs.Windows_Forensics_PartitionTable.GetHeaders(artifact_structs.Windows_Forensics_PartitionTable{})...)
	} else if artifact == "Generic.System.ProcessSiblings.csv" {
		headers = append(headers, artifact_structs.Generic_System_ProcessSiblings.GetHeaders(artifact_structs.Generic_System_ProcessSiblings{})...)
	} else if artifact == "Windows.System.RootCAStore.csv" {
		headers = append(headers, artifact_structs.Windows_System_RootCAStore.GetHeaders(artifact_structs.Windows_System_RootCAStore{})...)
	} else if artifact == "Windows.Sys.CertificateAuthorities.csv" {
		headers = append(headers, artifact_structs.Windows_Sys_CertificateAuthorities.GetHeaders(artifact_structs.Windows_Sys_CertificateAuthorities{})...)
	} else if artifact == "Generic.Applications.Chrome.SessionStorage.csv" {
		headers = append(headers, artifact_structs.Generic_Applications_Chrome_SessionStorage.GetHeaders(artifact_structs.Generic_Applications_Chrome_SessionStorage{})...)
	} else if artifact == "Windows.Registry.NTUser.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_NTUser.GetHeaders(artifact_structs.Windows_Registry_NTUser{})...)
	} else if artifact == "Windows.System.CatFiles.csv" {
		headers = append(headers, artifact_structs.Windows_System_CatFiles.GetHeaders(artifact_structs.Windows_System_CatFiles{})...)
	} else if artifact == "Exchange.Windows.Applications.DefenderDHParser.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Applications_DefenderDHParser.GetHeaders(artifact_structs.Exchange_Windows_Applications_DefenderDHParser{})...)
	} else if artifact == "Windows.Registry.RecentDocs.csv" {
		headers = append(headers, artifact_structs.Windows_Registry_RecentDocs.GetHeaders(artifact_structs.Windows_Registry_RecentDocs{})...)
	} else if artifact == "Generic.Client.DiskUsage.csv" {
		headers = append(headers, artifact_structs.Generic_Client_DiskUsage.GetHeaders(artifact_structs.Generic_Client_DiskUsage{})...)
	} else if artifact == "Generic.Applications.Office.Keywords.csv" {
		headers = append(headers, artifact_structs.Generic_Applications_Office_Keywords.GetHeaders(artifact_structs.Generic_Applications_Office_Keywords{})...)
	} else if artifact == "Windows.System.Signers.csv" {
		headers = append(headers, artifact_structs.Windows_System_Signers.GetHeaders(artifact_structs.Windows_System_Signers{})...)
	} else if artifact == "Windows.Detection.EnvironmentVariables.csv" {
		headers = append(headers, artifact_structs.Windows_Detection_EnvironmentVariables.GetHeaders(artifact_structs.Windows_Detection_EnvironmentVariables{})...)
	} else if artifact == "Exchange.Windows.Timeline.Prefetch.Improved.csv" {
		headers = append(headers, artifact_structs.Exchange_Windows_Timeline_Prefetch_Improved.GetHeaders(artifact_structs.Exchange_Windows_Timeline_Prefetch_Improved{})...)
	} else {
		// Generic Header Retrieval
		//fmt.Println(inputFile)
		tmpHeaders, headerErr := artifact_structs.Get_Generic_Headers(false, inputFile)
		if headerErr != nil {
			return headers, headerErr
		}
		headers = append(headers, tmpHeaders...)
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
