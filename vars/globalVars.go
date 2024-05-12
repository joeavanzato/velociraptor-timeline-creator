package vars

import (
	"time"
)

// Tool Logging Output
var LogFile = "velo_timeline_creator.log"

type ShallowRecord struct {
	Timestamp        time.Time
	Computer         string
	Artifact         string
	EventType        string
	EventDescription string
	SourceUser       string
	SourceHost       string
	DestinationUser  string
	DestinationHost  string
	SourceFile       string
	MetaData         string
}

func (sr ShallowRecord) StringArray() []string {
	return []string{sr.Timestamp.String(), sr.Computer, sr.Artifact, sr.EventType, sr.EventDescription, sr.SourceUser, sr.SourceHost, sr.DestinationUser, sr.DestinationHost, sr.SourceFile, sr.MetaData}
}

// If we are parsing full artifacts to individual CSV, this will store a reference to the output channel - if one does not exist yet, we create a new channel and store it here
var ArtifactToChannelMap = make(map[string]chan []string)

// If we are doing a basic super timeline - these are the columns we will use (in order) for the output CSV
var BaseHeaders = []string{"Timestamp", "Computer", "Artifact", "Event Type", "Event Description", "Source User", "Source Host", "Destination User", "Destination Host", "Source File", "Metadata"}

// If we are doing -mftlight, these are the extensions of interest we will include in the super-timeline
var LightMFTExtensionsOfInterest = []string{".ps1", ".exe", ".dll", ".hta", ".js", ".vba", ".cpl", ".wsf", ".vbs", ".bat",
	".psm1", ".py", ".psd1", ".cmd", ".scr", ".lnk", ".jar", ".pdf", ".rtf", ".doc", ".xls", ".docx", ".xlsx", ".csv",
	".jpeg", ".png", ".zip", ".gz", ".7z", ".com", ".ocx", ".ps1xml", ".ps2", ".msh", ".msh1", ".msh2", ".mshxml",
	".msh1xml", ".msh2xml", ".jse", ".vb", ".vbe", ".inf", ".reg", ".pif", ".scf", ".msc", ".msi", ".pol", ".hlp",
	".chm", ".ws", ".wsf", ".wsc", ".wsh", ".rar", ".zip", ".bz2", ".cab", ".tar", ".ace", ".msp", ".mst", ".msu",
	".ppkg", ".bak", ".tmp", ".ost", ".pst", ".pkg", ".iso", ".img", ".vhd", ".vhdx", ".application", ".lock", ".lck",
	".sln", ".cs", ".csproj", ".rex", ".config", ".resources", ".pdb", ".manifest", ".wbk", ".xlt", ".xlm", ".xla",
	".pot", ".pps", ".ade", ".adp", ".xlam", ".xll", ".xlw", ".ppam"}

// Artifact Directories that have some type of parsing implemented - artifacts are skipped if not present in thsi list
var ImplementedArtifacts = map[string]string{
	"Windows.System.Services":                               "Service was Created",
	"Windows.Timeline.Prefetch":                             "Prefetch Execution",
	"Windows.Timeline.Registry.RunMRU":                      "Program Execution from RunMRU",
	"Windows.System.Amcache":                                "Amcache Execution",
	"Windows.Sysinternals.Autoruns":                         "Autoruns Entry",
	"Windows.Registry.UserAssist":                           "UserAssist Entry",
	"Windows.Registry.Sysinternals.Eulacheck":               "Sysinternals EULA Accepted",
	"Windows.Registry.RDP":                                  "Registry RDP Cache Modified",
	"Windows.Registry.AppCompatCache":                       "AppCompatCache Entry Modified",
	"Windows.Network.NetstatEnriched":                       "Process Started with Network Connection",
	"Windows.KapeFiles.Targets":                             "KAPE Metadata Entry",
	"Windows.Forensics.Timeline":                            "Windows Forensic Timeline Entry",
	"Windows.Forensics.SRUM":                                "SRUM Entry",
	"Windows.Forensics.Shellbags":                           "Shellbags Entry",
	"Windows.Forensics.RecycleBin":                          "File Deleted to RecycleBin",
	"Windows.Forensics.RDPCache":                            "RDP Cache Entry",
	"Windows.Forensics.Lnk":                                 "LNK File Modified",
	"Windows.Forensics.CertUtil":                            "Cert Downloaded",
	"Windows.Forensics.Bam":                                 "BAM Entry",
	"Windows.EventLogs.AlternateLogon":                      "Alternate Logon",
	"Exchange.Windows.Office.MRU":                           "Office MRU Entry",
	"Exchange.Windows.EventLogs.RDPClientActivity":          "RDP Client Activity",
	"Exchange.Windows.EventLogs.LogonSessions":              "Logon Session Started",
	"Exchange.Windows.EventLogs.Bitsadmin":                  "BITS Entry",
	"Custom.Windows.Eventlog.Evtx":                          "EVTX Entry",
	"Custom.Windows.Mft":                                    "MFT Entry",
	"Custom.Windows.Mft.C":                                  "MFT Entry",
	"Custom.Windows.Mft.D":                                  "MFT Entry",
	"Generic.Client.Info":                                   "ClientInfo",
	"Windows.Applications.Chrome.History":                   "URL Visit (Chrome)",
	"Windows.Applications.Edge.History":                     "URL Visit (Edge)",
	"Windows.Applications.Firefox.Downloads":                "Download Started (Firefox)",
	"Windows.Applications.Firefox.History":                  "URL Visit (Firefox)",
	"Windows.Applications.NirsoftBrowserViewer":             "URL Visit",
	"Windows.EventLogs.PowershellScriptblock":               "PowerShell Script Block Executed",
	"Windows.EventLogs.RDPAuth":                             "RDP-Related Authentication",
	"Windows.Forensics.SAM":                                 "SAMDATA",
	"Generic.Forensic.SQLiteHunter":                         "SQLiteHunter",
	"Windows.Sys.Drivers":                                   "Signed Driver Date",
	"Windows.System.Powershell.ModuleAnalysisCache":         "PowerShell Module Loaded",
	"Windows.Analysis.EvidenceOfExecution":                  "Execution Evidence",
	"Windows.Analysis.EvidenceOfDownload":                   "Download Evidence",
	"Windows.EventLogs.Evtx":                                "EVTX Entry",
	"Exchange.Windows.EventLogs.Chainsaw":                   "Chainsaw Detection",
	"Exchange.Windows.Memory.InjectedThreadEx":              "Thread Injection Analysis",
	"Exchange.Windows.EventLogs.Hayabusa":                   "Hayabusa Detection",
	"Exchange.Windows.Forensics.Trawler":                    "Trawler Detection",
	"Exchange.Windows.Forensics.PersistenceSniper":          "PersistenceSniper Detection",
	"Exchange.Windows.EventLogs.CondensedAccountUsage":      "Account Usage",
	"Exchange.Windows.EventLogs.EvtxHussar":                 "Hussar Parser",
	"DetectRaptor.Windows.Detection.MFT":                    "DetectRaptor - MFT",
	"DetectRaptor.Generic.Detection.WebshellYara":           "DetectRaptor - WebshellYara",
	"DetectRaptor.Windows.Detection.Amcache":                "DetectRaptor - Amcache",
	"DetectRaptor.Windows.Detection.Applications":           "DetectRaptor - Applications",
	"DetectRaptor.Windows.Detection.BinaryRename":           "DetectRaptor - BinaryRename",
	"DetectRaptor.Windows.Detection.Webhistory":             "DetectRaptor - Webhistory",
	"DetectRaptor.Windows.Detection.Evtx":                   "DetectRaptor - Evtx",
	"DetectRaptor.Windows.Detection.Powershell.ISEAutoSave": "DetectRaptor - Powershell - ISEAutoSave",
	"DetectRaptor.Windows.Detection.Powershell.PSReadline":  "DetectRaptor - Powershell - PSReadline",
	"DetectRaptor.Windows.Detection.ZoneIdentifier":         "DetectRaptor - ZoneIdentifier",
	"DetectRaptor.Windows.Detection.HijackLibsEnv":          "DetectRaptor - HijackLibsEnv",
	"DetectRaptor.Windows.Detection.HijackLibsMFT":          "DetectRaptor - HijackLibsMFT",
}
