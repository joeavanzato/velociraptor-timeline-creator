# Velociraptor Timeline Creator

### What?

A command-line tool for dumping data out of Velociraptor data store files into both super-timelines as well as individual artifact CSVs.

### Why?

It's often useful to use Velociraptor (either online or offline) as a tool for collecting data, than to use said data through more traditional tooling such as Excel, Timeline Explorer, Timesketch, Splunk, Elastic, etc.

While Velociraptor does support certain integrations, I found it easier to work with the data by mass-extracting it from individual JSON files into CSV - hence, VTC.

### Use-Cases
* Building a summarized super-timeline from any supported artifacts present in Velociraptor (per-client)
  * ```velo-timeline-creator.exe -velodir "C:\velodatastore"```
  * ```velo-timeline-creator.exe -velodir "C:\velodatastore" -mftlight```
  * ```velo-timeline-creator.exe -velodir "C:\velodatastore" -mftfull```
* Dumping out all supported artifacts across all clients to individual CSV files
  * ```velo-timeline-creator.exe -velodir "C:\velodatastore" -artifactdump```
  * ```velo-timeline-creator.exe -velodir "C:\velodatastore" -artifactdump -mftlight```
  * ```velo-timeline-creator.exe -velodir "C:\velodatastore" -artifactdump -mftfull```

### MFT 
Since MFT can be a very 'heavy' artifact, we exclude related artifacts by default and only parse these when -mftlight or -mftfull is enabled at the command-line.

-mftfull parses the entire artifact without any exclusions.
-mftlight only includes files that have 'interesting' extensions in the output

https://github.com/joeavanzato/velociraptor-timeline-creator/blob/e913633718b3eda690090dd79c4fcd4416b67b85/vars/globalVars.go#L35


### My XYZ artifact is not supported?
* I am working on a 'generic' artifact parser to help dump  artifacts to individual CSV but this will not include a presence in the super-timeline.
* If you have an artifact you want to include, open an Issue with the name of the artifact as well as a single event from the resultant JSON (usually from a location like C:\VELODATASTORE\clients\C.*\artifacts\ARTIFACTHERE\SomeFile.JSON)
  * Ideally, you provide me the entire JSON output (sanitized where needed) so I can include it in some parsing tests

Build Link: https://github.com/joeavanzato/velociraptor-timeline-creator/releases/download/pre-release/velo-timeline-creator.zip

### Example Outputs

<h4 align="center">Per-Client Super-Timeline Output Examples (Named after ClientID)</h4>
<p align="center">
<img src="images/example1.png">
</p>
<h4 align="center">Super-Timeline Column Format</h4>
<p align="center">
<img src="images/columnExamples.png">
</p>
<h4 align="center">Cross-Artifact Super-Timeline Record Examples</h4>
<p align="center">
<img src="images/dataExamples.png">
</p>
<h4 align="center">Cross-Client Artifact Dump Examples</h4>
<p align="center">
<img src="images/example2.png">
</p>

### What artifacts are included in the super-timeline (if they have a timestamp field)?
If one of the below artifacts does not have time-element, then it will **not** appear in the super-timeline but it will be dumped when using -artifactdump.
* DetectRaptor.Generic.Detection.WebshellYara 
* DetectRaptor.Windows.Detection.Amcache
* DetectRaptor.Windows.Detection.Applications
* DetectRaptor.Windows.Detection.BinaryRename
* DetectRaptor.Windows.Detection.Evtx
* DetectRaptor.Windows.Detection.HijackLibsEnv
* DetectRaptor.Windows.Detection.HijackLibsMFT
* DetectRaptor.Windows.Detection.LolDriversVulnerable
* DetectRaptor.Windows.Detection.MFT
* DetectRaptor.Windows.Detection.NamedPipes
* DetectRaptor.Windows.Detection.Powershell.ISEAutoSave
* DetectRaptor.Windows.Detection.Powershell.PSReadline
* DetectRaptor.Windows.Detection.Webhistory
* DetectRaptor.Windows.Detection.ZoneIdentifier
* Exchange.Custom.Windows.Nirsoft.LastActivityView
* Exchange.HashRunKeys
* Exchange.Windows.Applications.DefenderDHParser
* Exchange.Windows.Applications.LECmd
* Exchange.Windows.Applications.OfficeServerCache
* Exchange.Windows.Detection.Malfind
* Exchange.Windows.Detection.PipeHunter
* Exchange.Windows.Detection.PrefetchHunter
* Exchange.Windows.EventLogs.Bitsadmin
* Exchange.Windows.EventLogs.Chainsaw
* Exchange.Windows.EventLogs.CondensedAccountUsage
* Exchange.Windows.EventLogs.EvtxHussar
* Exchange.Windows.EventLogs.Hayabusa
* Exchange.Windows.EventLogs.LogonSessions
* Exchange.Windows.EventLogs.RDPClientActivity
* Exchange.Windows.Forensics.Clipboard
* Exchange.Windows.Forensics.FileZilla
* Exchange.Windows.Forensics.Jumplists_JLECmd
* Exchange.Windows.Forensics.PersistenceSniper
* Exchange.Windows.Forensics.ThumbCache
* Exchange.Windows.Forensics.Trawler
* Exchange.Windows.Forensics.UEFI
* Exchange.Windows.Forensics.UEFI.BootApplication
* Exchange.Windows.Memory.InjectedThreadEx
* Exchange.Windows.NTFS.Timestomp
* Exchange.Windows.Office.MRU
* Exchange.Windows.Registry.BackupRestore
* Exchange.Windows.Registry.CapabilityAccessManager
* Exchange.Windows.Registry.COMAutoApprovalList
* Exchange.Windows.Registry.Domain
* Exchange.Windows.Registry.NetshHelperDLLs
* Exchange.Windows.Registry.ScheduledTasks
* Exchange.Windows.Sys.LoggedInUsers
* Exchange.Windows.System.BinaryVersion
* Exchange.Windows.System.Powershell.ISEAutoSave
* Exchange.Windows.System.PrinterDriver
* Exchange.Windows.System.WindowsErrorReporting
* Exchange.Windows.System.WMIProviders
* Generic.Applications.Chrome.SessionStorage
* Generic.Applications.Office.Keywords
* Generic.Client.DiskSpace
* Generic.Client.DiskUsage
* Generic.Client.Info
* Generic.Detection.Yara.Zip
* Generic.Forensic.SQLiteHunter
* Generic.Forensic.Timeline
* Generic.Network.InterfaceAddresses
* Generic.System.ProcessSiblings
* Generic.System.Pstree
* Network.ExternalIpAddress
* Windows.Analysis.EvidenceOfDownload
* Windows.Analysis.EvidenceOfExecution
* Windows.Applications.ChocolateyPackages
* Windows.Applications.Chrome.Extensions
* Windows.Applications.Chrome.History
* Windows.Applications.Edge.History
* Windows.Applications.Firefox.Downloads
* Windows.Applications.Firefox.History
* Windows.Applications.NirsoftBrowserViewer
* Windows.Applications.OfficeMacros
* Windows.Carving.USN
* Windows.Detection.Amcache
* Windows.Detection.BinaryHunter
* Windows.Detection.ForwardedImports
* Windows.Detection.Impersonation
* Windows.Detection.Mutants
* Windows.EventLogs.AlternateLogon
* Windows.EventLogs.Evtx
* Windows.EventLogs.Modifications
* Windows.EventLogs.PowershellModule
* Windows.EventLogs.PowershellScriptblock
* Windows.EventLogs.RDPAuth
* Windows.Forensics.Bam
* Windows.Forensics.CertUtil
* Windows.Forensics.Lnk
* Windows.Forensics.PartitionTable
* Windows.Forensics.RDPCache
* Windows.Forensics.RecycleBin
* Windows.Forensics.SAM
* Windows.Forensics.Shellbags
* Windows.Forensics.SRUM
* Windows.Forensics.Timeline
* Windows.Forensics.Usn
* Windows.KapeFiles.Targets
* Windows.Memory.ProcessInfo
* Windows.Network.ArpCache
* Windows.Network.ListeningPorts
* Windows.Network.Netstat
* Windows.Network.NetstatEnriched
* Windows.NTFS.MFT
* Windows.Persistence.PermanentWMIEvents
* Windows.Registry.AppCompatCache
* Windows.Registry.NTUser
* Windows.Registry.PuttyHostKeys
* Windows.Registry.RDP
* Windows.Registry.RecentDocs
* Windows.Registry.Sysinternals.Eulacheck
* Windows.Registry.UserAssist
* Windows.Registry.WDigest
* Windows.Sys.AllUsers
* Windows.Sys.CertificateAuthorities
* Windows.Sys.DiskInfo
* Windows.Sys.Drivers
* Windows.Sys.FirewallRules
* Windows.Sys.Interfaces
* Windows.Sys.PhysicalMemoryRanges
* Windows.Sys.Programs
* Windows.Sys.StartupItems
* Windows.Sys.Users
* Windows.Sysinternals.Autoruns
* Windows.System.Amcache
* Windows.System.AuditPolicy
* Windows.System.CatFiles
* Windows.System.DLLs
* Windows.System.DNSCache
* Windows.System.Handles
* Windows.System.HostsFile
* Windows.System.LocalAdmins
* Windows.System.Powershell.ModuleAnalysisCache
* Windows.System.Powershell.PSReadline
* Windows.System.Pslist
* Windows.System.RootCAStore
* Windows.System.Services
* Windows.System.Shares
* Windows.System.Signers
* Windows.System.TaskScheduler
* Windows.Timeline.MFT
* Windows.Timeline.Prefetch
* Windows.Timeline.Registry.RunMRU


