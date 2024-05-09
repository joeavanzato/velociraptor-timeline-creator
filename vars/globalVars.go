package vars

import (
	"time"
)

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

var BaseHeaders = []string{"Timestamp", "Computer", "Artifact", "Event Type", "Event Description", "Source User", "Source Host", "Destination User", "Destination Host", "Source File", "Metadata"}

var LightMFTExtensionsOfInterest = []string{".ps1", ".exe", ".dll", ".hta", ".js", ".vba", ".cpl", ".wsf", ".vbs", ".bat",
	".psm1", ".py", ".psd1", ".cmd", ".scr", ".lnk", ".jar", ".pdf", ".rtf", ".doc", ".xls", ".docx", ".xlsx", ".csv",
	".jpeg", ".png", ".zip", ".gz", ".7z", ".com", ".ocx", ".ps1xml", ".ps2", ".msh", ".msh1", ".msh2", ".mshxml",
	".msh1xml", ".msh2xml", ".jse", ".vb", ".vbe", ".inf", ".reg", ".pif", ".scf", ".msc", ".msi", ".pol", ".hlp",
	".chm", ".ws", ".wsf", ".wsc", ".wsh", ".rar", ".z", ".bz2", ".cab", ".tar", ".ace", ".msp", ".mst", ".msu",
	".ppkg", ".bak", ".tmp", ".ost", ".pst", ".pkg", ".iso", ".img", ".vhd", ".vhdx", ".application", ".lock", ".lck",
	".sln", ".cs", ".csproj", ".rex", ".config", ".resources", ".pdb", ".manifest", ".wbk", ".xlt", ".xlm", ".xla",
	".pot", ".pps", ".ade", ".adp", ".xlam", ".xll", ".xlw", ".ppam"}

var ImplementedArtifacts = map[string]string{
	"Windows.System.Services":                      "Service was Created",
	"Windows.Timeline.Prefetch":                    "Prefetch Execution",
	"Windows.Timeline.Registry.RunMRU":             "Program Execution from RunMRU",
	"Windows.System.Amcache":                       "Amcache Execution",
	"Windows.Sysinternals.Autoruns":                "Autoruns Entry",
	"Windows.Registry.UserAssist":                  "UserAssist Entry",
	"Windows.Registry.Sysinternals.Eulacheck":      "Sysinternals EULA Accepted",
	"Windows.Registry.RDP":                         "Registry RDP Cache Modified",
	"Windows.Registry.AppCompatCache":              "AppCompatCache Entry Modified",
	"Windows.Network.NetstatEnriched":              "Process Started with Network Connection",
	"Windows.KapeFiles.Targets":                    "KAPE Metadata Entry",
	"Windows.Forensics.Timeline":                   "Windows Forensic Timeline Entry",
	"Windows.Forensics.SRUM":                       "SRUM Entry",
	"Windows.Forensics.Shellbags":                  "Shellbags Entry",
	"Windows.Forensics.RecycleBin":                 "File Deleted to RecycleBin",
	"Windows.Forensics.RDPCache":                   "RDP Cache Entry",
	"Windows.Forensics.Lnk":                        "LNK File Modified",
	"Windows.Forensics.CertUtil":                   "Cert Downloaded",
	"Windows.Forensics.Bam":                        "BAM Entry",
	"Windows.Forensics.AlternateLogon":             "Alternate Logon",
	"Exchange.Windows.Office.MRU":                  "Office MRU Entry",
	"Exchange.Windows.EventLogs.RDPClientActivity": "RDP Client Activity",
	"Exchange.Windows.EventLogs.LogonSessions":     "Logon Session Started",
	"Exchange.Windows.EventLogs.Bitsadmin":         "BITS Entry",
	"Custom.Windows.Eventlog.Evtx":                 "EVTX Entry",
	"Custom.Windows.Mft":                           "MFT Entry",
	"Custom.Windows.Mft.C":                         "MFT Entry",
	"Custom.Windows.Mft.D":                         "MFT Entry",
	"Generic.Client.Info":                          "ClientInfo",
	"Windows.Applications.Chrome.History":          "URL Visit (Chrome)",
	"Windows.Applications.Edge.History":            "URL Visit (Edge)",
	"Windows.Applications.Firefox.Downloads":       "Download Started (Firefox)",
	"Windows.Applications.Firefox.History":         "URL Visit (Firefox)",
	"Windows.Applications.NirsoftBrowserViewer":    "URL Visit",
	"Windows.EventLogs.PowershellScriptblock":      "PowerShell Script Block Executed",
	"Windows.EventLogs.RDPAuth":                    "RDP-Related Authentication",
	"Windows.Forensics.SAM":                        "SAMDATA",
	"Generic.Forensic.SQLiteHunter":                "SQLiteHunter",
}
