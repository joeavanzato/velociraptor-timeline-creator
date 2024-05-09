package main

import (
	"bufio"
	"encoding/csv"
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

func parseArgs(logger zerolog.Logger) (map[string]any, error) {
	velodir := flag.String("velodir", "", "")
	fullparse := flag.Bool("fullparse", false, "")
	maxgoperfile := flag.Int("maxgoperfile", 20, "Maximum number of goroutines to spawn on a per-file basis for concurrent processing of data.")
	batchsize := flag.Int("batchsize", 500, "Maximum number of lines to read at a time for processing within each spawned goroutine per file.")
	outputdir := flag.String("outputdir", "", "")
	writebuffer := flag.Int("writebuffer", 2000, "How many lines to queue at a time for writing to output CSV")
	mftlight := flag.Bool("mftlight", false, "")
	mftfull := flag.Bool("mftfull", false, "")

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
	for _, v := range artifactPaths {
		artifactWaiter.Add(1)
		go ProcessArtifact(v, arguments, &artifactWaiter, recordChannel, logger, clientIdentifier)
	}
	helpers.CloseChannelWhenDone(recordChannel, &artifactWaiter)
	writeWG.Wait()
	return nil
}

func ProcessArtifact(artifactDir string, arguments map[string]any, artifactWaiter *sync.WaitGroup, recordOutputChannel chan []string, logger zerolog.Logger, clientIdentifier string) {

	// Here we will actually read the relevant artifact file (if it exists) and process records through the appropriate artifact helper func
	paths := helpers.GetAllJSONFromDirectory(artifactDir)
	defer artifactWaiter.Done()
	for _, artifactJSON := range paths {
		artifactWaiter.Add(1)
		go ProcessArtifactFile(artifactDir, arguments, artifactWaiter, recordOutputChannel, logger, clientIdentifier, artifactJSON)
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
	} else if artifactName == "Windows.Forensics.AlternateLogon" {
		artifact_structs.Process_Windows_Forensics_AlternateLogon("Windows.Forensics.AlternateLogon", clientIdentifier, records, recordOutputChannel, arguments)
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
		} else if JSONFileName == "Firefox Cookies.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Cookies("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Firefox Form History.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Firefox_Form_History("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "IE or Edge WebCacheV01_All Data.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_IE_Edge_WebCache_All_Data("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Windows Search Service_SystemIndex_Gthr.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_Gthr("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		} else if JSONFileName == "Windows Search Service_SystemIndex_PropertyStore.json" {
			artifact_structs.Process_Generic_Forensic_SQLiteHunter_Windows_Search_Service_SystemIndex_PropertyStore("Generic.Forensic.SQLiteHunter", clientIdentifier, records, recordOutputChannel, arguments)
		}
	} else if strings.HasPrefix(artifactName, "Windows.Sys.Drivers") {
		if JSONFileName == "SignedDrivers.json" {
			artifact_structs.Process_Windows_Sys_Drivers_SignedDrivers("Windows.Sys.Drivers", clientIdentifier, records, recordOutputChannel, arguments)
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
		}
	}
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
	if velo_dir_err != nil {
		logger.Error().Msgf(velo_clients_err.Error())
		return
	}
	clientWaiter := sync.WaitGroup{}
	for _, path := range clientArtifactPaths {
		//
		clientWaiter.Add(1)
		ProcessClientArtifactPath(path, arguments, &clientWaiter, logger)
	}

	clientWaiter.Wait()
	logger.Info().Msgf("Done!")

}
