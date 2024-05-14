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
Since MFT can be a very 'heavy' artifact, we exclude related json by default and only parse it when -mftlight or -mftfull is enabled at the command-line.

-mftfull parses the entire artifact without any exclusions.
-mftlight only includes files that have 'interesting' extensions in the output - see \vars\globalVars.go for a list of specific extensions that are included.


### My XYZ artifact is not supported?
* I am working on a 'generic' artifact parser to help dump  artifacts to individual CSV but this will not include a presence in the super-timeline.
* If you have an artifact you want to include, open an Issue with the name of the artifact as well as a single event from the resultant JSON (usually from a location like $VELODATASTORE$\clients\C.*\artifacts\ARTIFACTHERE\SomeFile.JSON)
  * Ideally, you provide me the entire JSON output (sanitized where needed) so I can include it in some parsing tests
* 