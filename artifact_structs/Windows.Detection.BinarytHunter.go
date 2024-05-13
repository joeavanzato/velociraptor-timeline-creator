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

type Windows_Detection_BinaryHunter struct {
	File struct {
		OSPath     string `json:"OSPath"`
		Name       string `json:"Name"`
		Size       int    `json:"Size"`
		Timestamps struct {
			Mtime time.Time `json:"Mtime"`
			Atime time.Time `json:"Atime"`
			Ctime time.Time `json:"Ctime"`
			Btime time.Time `json:"Btime"`
		} `json:"Timestamps"`
	} `json:"File"`
	Authenticode struct {
		Filename      string `json:"Filename"`
		ProgramName   string `json:"ProgramName"`
		PublisherLink string `json:"PublisherLink"`
		MoreInfoLink  string `json:"MoreInfoLink"`
		SerialNumber  string `json:"SerialNumber"`
		IssuerName    string `json:"IssuerName"`
		SubjectName   string `json:"SubjectName"`
		Timestamp     any    `json:"Timestamp"`
		Trusted       string `json:"Trusted"`
		ExtraInfo     any    `json:"_ExtraInfo"`
	} `json:"Authenticode"`
	PE struct {
		FileHeader struct {
			Machine          string    `json:"Machine"`
			TimeDateStamp    time.Time `json:"TimeDateStamp"`
			TimeDateStampRaw int64     `json:"TimeDateStampRaw"`
			Characteristics  int       `json:"Characteristics"`
			ImageBase        int       `json:"ImageBase"`
		} `json:"FileHeader"`
		GUIDAge     string `json:"GUIDAge"`
		PDB         string `json:"PDB"`
		Directories struct {
			BaseRelocationDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Base_Relocation_Directory"`
			ExportDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int64     `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Export_Directory"`
			IATDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"IAT_Directory"`
			ImportDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Import_Directory"`
			ResourceDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Resource_Directory"`
			SecurityDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Security_Directory"`
			TLSDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"TLS_Directory"`
		} `json:"Directories"`
		Sections []struct {
			Perm       string `json:"Perm"`
			Name       string `json:"Name"`
			FileOffset int    `json:"FileOffset"`
			VMA        int    `json:"VMA"`
			RVA        int    `json:"RVA"`
			Size       int    `json:"Size"`
		} `json:"Sections"`
		Resources []struct {
			Type       string `json:"Type"`
			TypeID     int    `json:"TypeId"`
			FileOffset int    `json:"FileOffset"`
			DataSize   int    `json:"DataSize"`
			CodePage   int    `json:"CodePage"`
		} `json:"Resources"`
		VersionInformation struct {
			FileDescription  string `json:"FileDescription"`
			OriginalFilename string `json:"OriginalFilename"`
			CompanyName      string `json:"CompanyName"`
			LegalCopyright   string `json:"LegalCopyright"`
			FileVersion      string `json:"FileVersion"`
			ProductName      string `json:"ProductName"`
			ProductVersion   string `json:"ProductVersion"`
		} `json:"VersionInformation"`
		Imports      []string      `json:"Imports"`
		Exports      []string      `json:"Exports"`
		Forwards     []interface{} `json:"Forwards"`
		ImpHash      string        `json:"ImpHash"`
		Authenticode struct {
			Signer struct {
				IssuerName              string `json:"IssuerName"`
				SerialNumber            string `json:"SerialNumber"`
				DigestAlgorithm         string `json:"DigestAlgorithm"`
				AuthenticatedAttributes struct {
					ProgramName      string `json:"ProgramName"`
					MoreInfo         string `json:"MoreInfo"`
					ContentType      string `json:"ContentType"`
					MessageDigest    string `json:"MessageDigest"`
					MessageDigestHex string `json:"MessageDigestHex"`
				} `json:"AuthenticatedAttributes"`
				UnauthenticatedAttributes struct {
					CounterSignature struct {
						IssuerName              string `json:"IssuerName"`
						SerialNumber            string `json:"SerialNumber"`
						DigestAlgorithm         string `json:"DigestAlgorithm"`
						AuthenticatedAttributes struct {
							ContentType      string    `json:"ContentType"`
							SigningTime      time.Time `json:"SigningTime"`
							MessageDigest    string    `json:"MessageDigest"`
							MessageDigestHex string    `json:"MessageDigestHex"`
						} `json:"AuthenticatedAttributes"`
						UnauthenticatedAttributes struct {
						} `json:"UnauthenticatedAttributes"`
					} `json:"CounterSignature"`
				} `json:"UnauthenticatedAttributes"`
				Subject string `json:"Subject"`
			} `json:"Signer"`
			Certificates []struct {
				SerialNumber       string    `json:"SerialNumber"`
				SignatureAlgorithm string    `json:"SignatureAlgorithm"`
				Subject            string    `json:"Subject"`
				Issuer             string    `json:"Issuer"`
				NotBefore          time.Time `json:"NotBefore"`
				NotAfter           time.Time `json:"NotAfter"`
				PublicKey          string    `json:"PublicKey"`
				Extensions         struct {
					KeyUsage struct {
						Critical bool     `json:"Critical"`
						KeyUsage []string `json:"KeyUsage"`
					} `json:"KeyUsage"`
					BasicConstraints struct {
						Critical   bool `json:"Critical"`
						IsCA       bool `json:"IsCA"`
						MaxPathLen int  `json:"MaxPathLen"`
					} `json:"BasicConstraints"`
					ExtendedKeyUsage struct {
						Critical bool     `json:"Critical"`
						KeyUsage []string `json:"KeyUsage"`
					} `json:"Extended Key Usage"`
					CertificatePolicies struct {
						Critical bool     `json:"Critical"`
						Policy   []string `json:"Policy"`
					} `json:"CertificatePolicies"`
					AuthorityKeyIdentifier struct {
						Critical bool   `json:"Critical"`
						KeyID    string `json:"KeyId"`
					} `json:"AuthorityKeyIdentifier"`
					SubjectKeyID struct {
						Critical bool   `json:"Critical"`
						Value    string `json:"Value"`
					} `json:"SubjectKeyId"`
					CRLDistributionPoints struct {
						Critical bool     `json:"Critical"`
						URI      []string `json:"URI"`
					} `json:"CRLDistributionPoints"`
				} `json:"Extensions"`
			} `json:"Certificates"`
			HashType        string `json:"HashType"`
			ExpectedHash    string `json:"ExpectedHash"`
			ExpectedHashHex string `json:"ExpectedHashHex"`
		} `json:"Authenticode"`
		AuthenticodeHash struct {
			MD5         string `json:"MD5"`
			SHA1        string `json:"SHA1"`
			SHA256      string `json:"SHA256"`
			HashMatches bool   `json:"HashMatches"`
		} `json:"AuthenticodeHash"`
	} `json:"PE"`
	Hash struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hash"`
}

func (s Windows_Detection_BinaryHunter) StringArray() []string {
	certificates := make([]string, 0)
	for _, v := range s.PE.Authenticode.Certificates {
		certificates = append(certificates, fmt.Sprintf("| SerialNumber: %v, SignatureAlgorithm: %v, Subject: %v, Issuer: %v, NotBefore: %v, NotAfter: %v, PublicKey: %v, ExtendedKeyUsage_Critical: %v, ExtendedKeyUsage_KeyUsage: %v, SubjetKeyID_Critical: %v, SubjectKeyId_Value: %v, SubjectAlternativeName_Critical: %v, SubjectAlternativeName_DNS: %v, SubjectAlternativeName_Email: %v, SubjectAlternativeName_IP: %v, AuthorityKeyIdentifier_Critical: %v, AuthorityKeyIdentifier_KeyID: %v, CRLDistributionPoints_Critical: %v, CRLDistributionPoints_URI: %v, BasicConstraints_Critical: %v, BasicConstraints_IsCA: %v, BasicConstraints_MaxPathLen: %v", v.SerialNumber, v.SignatureAlgorithm, v.Subject, v.Issuer, v.NotBefore, v.NotAfter, v.PublicKey, v.Extensions.ExtendedKeyUsage.Critical, v.Extensions.ExtendedKeyUsage.KeyUsage, v.Extensions.SubjectKeyID.Critical, v.Extensions.SubjectKeyID.Value, v.Extensions.AuthorityKeyIdentifier.Critical, v.Extensions.AuthorityKeyIdentifier.KeyID, v.Extensions.CRLDistributionPoints.Critical, v.Extensions.CRLDistributionPoints.URI, v.Extensions.BasicConstraints.Critical, v.Extensions.BasicConstraints.IsCA, v.Extensions.BasicConstraints.MaxPathLen))
	}
	sections := make([]string, 0)
	for _, v := range s.PE.Sections {
		sections = append(sections, fmt.Sprintf("| Perm: %v, Name: %v, FileOffset: %v, VMA: %v, RVA: %v, Size: %v", v.Perm, v.Name, v.FileOffset, v.VMA, v.RVA, v.Size))
	}
	resources := make([]string, 0)
	for _, v := range s.PE.Resources {
		resources = append(resources, fmt.Sprintf("| Type: %v, TypeID: %v, FileOffset: %v, DataSize: %v, CodePage: %v", v.Type, v.TypeID, v.FileOffset, v.DataSize, v.CodePage))
	}

	return []string{
		// File Struct
		s.File.OSPath, strconv.Itoa(s.File.Size), s.File.Timestamps.Mtime.String(), s.File.Timestamps.Atime.String(), s.File.Timestamps.Ctime.String(), s.File.Timestamps.Btime.String(),

		s.Hash.MD5, s.Hash.SHA1, s.Hash.SHA256,

		// PE File Headers
		s.PE.FileHeader.Machine, s.PE.FileHeader.TimeDateStamp.String(), strconv.FormatInt(s.PE.FileHeader.TimeDateStampRaw, 10),
		strconv.Itoa(s.PE.FileHeader.Characteristics), strconv.Itoa(s.PE.FileHeader.ImageBase),
		s.PE.GUIDAge, s.PE.PDB,

		// Possible Directories
		s.PE.Directories.BaseRelocationDirectory.Timestamp.String(),
		strconv.Itoa(s.PE.Directories.BaseRelocationDirectory.TimestampRaw), strconv.Itoa(s.PE.Directories.BaseRelocationDirectory.Size),
		strconv.Itoa(s.PE.Directories.BaseRelocationDirectory.FileAddress), s.PE.Directories.BaseRelocationDirectory.SectionName,
		s.PE.Directories.ExportDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.ExportDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.ExportDirectory.Size),
		strconv.Itoa(s.PE.Directories.ExportDirectory.FileAddress), s.PE.Directories.ExportDirectory.SectionName,
		s.PE.Directories.IATDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.IATDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.IATDirectory.Size),
		strconv.Itoa(s.PE.Directories.IATDirectory.FileAddress), s.PE.Directories.IATDirectory.SectionName,
		s.PE.Directories.ImportDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.ImportDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.ImportDirectory.Size),
		strconv.Itoa(s.PE.Directories.ImportDirectory.FileAddress), s.PE.Directories.ImportDirectory.SectionName,
		s.PE.Directories.ResourceDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.ResourceDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.ResourceDirectory.Size),
		strconv.Itoa(s.PE.Directories.ResourceDirectory.FileAddress), s.PE.Directories.ResourceDirectory.SectionName,
		s.PE.Directories.SecurityDirectory.Timestamp.String(),
		strconv.Itoa(s.PE.Directories.SecurityDirectory.TimestampRaw), strconv.Itoa(s.PE.Directories.SecurityDirectory.Size),
		strconv.Itoa(s.PE.Directories.SecurityDirectory.FileAddress), s.PE.Directories.SecurityDirectory.SectionName,
		s.PE.Directories.TLSDirectory.Timestamp.String(),
		strconv.Itoa(s.PE.Directories.TLSDirectory.TimestampRaw), strconv.Itoa(s.PE.Directories.TLSDirectory.Size),
		strconv.Itoa(s.PE.Directories.TLSDirectory.FileAddress), s.PE.Directories.TLSDirectory.SectionName,

		fmt.Sprint(sections),
		fmt.Sprint(resources),

		s.PE.VersionInformation.CompanyName, s.PE.VersionInformation.FileDescription, s.PE.VersionInformation.FileVersion,
		s.PE.VersionInformation.LegalCopyright, s.PE.VersionInformation.OriginalFilename,
		s.PE.VersionInformation.ProductName, s.PE.VersionInformation.ProductVersion,

		fmt.Sprint(s.PE.Imports), fmt.Sprint(s.PE.Exports), fmt.Sprint(s.PE.Forwards), fmt.Sprint((s.PE.ImpHash)),

		s.PE.Authenticode.Signer.IssuerName, s.PE.Authenticode.Signer.SerialNumber, s.PE.Authenticode.Signer.DigestAlgorithm,
		s.PE.Authenticode.Signer.AuthenticatedAttributes.ContentType, s.PE.Authenticode.Signer.AuthenticatedAttributes.MessageDigest, s.PE.Authenticode.Signer.AuthenticatedAttributes.MessageDigestHex,
		s.PE.Authenticode.Signer.AuthenticatedAttributes.ProgramName, s.PE.Authenticode.Signer.AuthenticatedAttributes.MoreInfo,
		fmt.Sprint(s.PE.Authenticode.Signer.UnauthenticatedAttributes), s.PE.Authenticode.Signer.Subject,
		fmt.Sprint(certificates), s.PE.Authenticode.HashType, s.PE.Authenticode.ExpectedHash, s.PE.Authenticode.ExpectedHashHex,
		s.PE.AuthenticodeHash.MD5, s.PE.AuthenticodeHash.SHA1, s.PE.AuthenticodeHash.SHA256, strconv.FormatBool(s.PE.AuthenticodeHash.HashMatches),

		s.Authenticode.Filename, s.Authenticode.ProgramName, s.Authenticode.PublisherLink, s.Authenticode.MoreInfoLink, s.Authenticode.SerialNumber,
		s.Authenticode.IssuerName, s.Authenticode.SubjectName, fmt.Sprint(s.Authenticode.Timestamp), s.Authenticode.Trusted, fmt.Sprint(s.Authenticode.ExtraInfo),
	}
}

func (s Windows_Detection_BinaryHunter) GetHeaders() []string {
	return []string{
		"OSPath", "Name", "FileSize", "Mtime", "Atime", "Ctime", "Btime",
		"MD5", "SHA1", "SHA256",
		"FileHeader_Machine",
		"FileHeader_TimeDateStamp", "FileHeader_TimeDateStampRaw", "FileHeader_Characteristics", "FileHeader_ImageBase",
		"GUIDAge", "PDB",
		"BaseRelocationDir_Timestamp", "BaseRelocationDir_TimestampRaw", "BaseRelocationDir_Size", "BaseRelocationDir_FileAddress",
		"BaseRelocationDir_SectionName",
		"ExportDirectory_Timestamp", "ExportDirectory_TimestampRaw", "ExportDirectory_Size", "ExportDirectory_FileAddress",
		"ExportDirectory_SectionName",
		"IATDirectory_Timestamp", "IATDirectory_TimestampRaw", "IATDirectory_Size", "IATDirectory_FileAddress",
		"IATDirectory_SectionName",
		"ImportDirectory_Timestamp", "ImportDirectory_TimestampRaw", "ImportDirectory_Size", "ImportDirectory_FileAddress",
		"ImportDirectory_SectionName",
		"ResourceDirectory_Timestamp", "ResourceDirectory_TimestampRaw", "ResourceDirectory_Size", "ResourceDirectory_FileAddress",
		"ResourceDirectory_SectionName",
		"SecurityDirectory_Timestamp", "SecurityDirectory_TimestampRaw", "SecurityDirectory_Size", "SecurityDirectory_FileAddress",
		"SecurityDirectory_SectionName",
		"TLSDirectory_Timestamp", "TLSDirectory_TimestampRaw", "TLSDirectory_Size", "TLSDirectory_FileAddress",
		"TLSDirectory_SectionName",
		"Sections", "Resources",
		"Version_CompanyName", "Version_FileDescription", "Version_FileVersion", "Version_LegalCopyright",
		"Version_OriginalFileName", "Version_ProductName", "Version_ProductVersion",
		"Imports", "Exports", "Forwards", "ImpHash",
		"Authenticode_AA_IssuerName", "Authenticode_AA_SerialNumber", "Authenticode_AA_DigestAlgorithm",
		"Authenticode_AA_ContentType", "Authenticode_AA_MessageDigest", "Authenticode_AA_MessageDigestHex", "Authenticode_AA_ProgramName", "Authenticode_MoreInfo",
		"Authenticode_UnauthenticatedAttributes", "Authenticode_Subject", "Authenticode_Certificates", "Authenticde_HashType", "Authenticode_ExpectedHash",
		"Authenticode_ExpectedHashHex", "Authenticode_MD5", "Authenticode_SHA1", "Authenticode_SHA256", "Authenticode_HashMatches",
		"Authenticode_Filename", "Authenticode_ProgramName", "Authenticode_PublisherLink", "Authenticode_MoreInfoLink", "Authenticode_IssuerName",
		"Authenticode_SubjectName", "Authenticode_Timestamp", "Authenticode_Trusted", "Authenticode_ExtraInfo"}
}

func Process_Windows_Detection_BinaryHunter(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Detection_BinaryHunter{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.File.Timestamps.Mtime.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.File.Timestamps.Mtime,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        vars.ImplementedArtifacts[artifactName],
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.File.OSPath,
			MetaData:         fmt.Sprintf("Imphash: %v, Ctime: %v, PDB: %v, Imports: %v", tmp.PE.ImpHash, tmp.File.Timestamps.Ctime, tmp.PE.PDB, tmp.PE.Imports),
		}
		outputChannel <- tmp2.StringArray()
	}
}
