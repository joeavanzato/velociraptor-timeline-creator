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

type Exchange_Windows_Forensics_UEFI struct {
	Partition struct {
		ImagePath       string `json:"ImagePath"`
		PartitionOffset int    `json:"PartitionOffset"`
		PartitionSize   string `json:"PartitionSize"`
		PartitionName   string `json:"PartitionName"`
	} `json:"Partition"`
	OSPath       string    `json:"OSPath"`
	Size         int       `json:"Size"`
	Mtime        time.Time `json:"Mtime"`
	Atime        time.Time `json:"Atime"`
	Ctime        time.Time `json:"Ctime"`
	Btime        time.Time `json:"Btime"`
	FirstCluster int       `json:"FirstCluster"`
	Attr         string    `json:"Attr"`
	IsDeleted    any       `json:"IsDeleted"`
	ShortName    string    `json:"ShortName"`
	Hash         struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hash"`
	Magic  string `json:"Magic"`
	PEInfo struct {
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
			DebugDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int64     `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Debug_Directory"`
			ExceptionDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Exception_Directory"`
			ExportDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Export_Directory"`
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
			CompanyName      string `json:"CompanyName"`
			FileDescription  string `json:"FileDescription"`
			FileVersion      string `json:"FileVersion"`
			InternalName     string `json:"InternalName"`
			LegalCopyright   string `json:"LegalCopyright"`
			OriginalFilename string `json:"OriginalFilename"`
			ProductName      string `json:"ProductName"`
			ProductVersion   string `json:"ProductVersion"`
		} `json:"VersionInformation"`
		Imports      any    `json:"Imports"`
		Exports      any    `json:"Exports"`
		Forwards     any    `json:"Forwards"`
		ImpHash      string `json:"ImpHash"`
		Authenticode struct {
			Signer struct {
				IssuerName              string `json:"IssuerName"`
				SerialNumber            string `json:"SerialNumber"`
				DigestAlgorithm         string `json:"DigestAlgorithm"`
				AuthenticatedAttributes struct {
					ContentType       string `json:"ContentType"`
					MessageDigest     string `json:"MessageDigest"`
					MessageDigestHex  string `json:"MessageDigestHex"`
					Oid13614131110328 string `json:"Oid: 1.3.6.1.4.1.311.10.3.28"`
					ProgramName       string `json:"ProgramName"`
					MoreInfo          string `json:"MoreInfo"`
				} `json:"AuthenticatedAttributes"`
				UnauthenticatedAttributes struct {
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
					ExtendedKeyUsage struct {
						Critical bool     `json:"Critical"`
						KeyUsage []string `json:"KeyUsage"`
					} `json:"Extended Key Usage"`
					SubjectKeyID struct {
						Critical bool   `json:"Critical"`
						Value    string `json:"Value"`
					} `json:"SubjectKeyId"`
					SubjectAlternativeName struct {
						Critical bool `json:"Critical"`
						DNS      any  `json:"DNS"`
						Email    any  `json:"Email"`
						IP       any  `json:"IP"`
					} `json:"SubjectAlternativeName"`
					AuthorityKeyIdentifier struct {
						Critical bool   `json:"Critical"`
						KeyID    string `json:"KeyId"`
					} `json:"AuthorityKeyIdentifier"`
					CRLDistributionPoints struct {
						Critical bool     `json:"Critical"`
						URI      []string `json:"URI"`
					} `json:"CRLDistributionPoints"`
					BasicConstraints struct {
						Critical   bool `json:"Critical"`
						IsCA       bool `json:"IsCA"`
						MaxPathLen int  `json:"MaxPathLen"`
					} `json:"BasicConstraints"`
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
	} `json:"PEInfo"`
	Authenticode struct {
		Filename      string `json:"Filename"`
		ProgramName   string `json:"ProgramName"`
		PublisherLink string `json:"PublisherLink"`
		MoreInfoLink  string `json:"MoreInfoLink"`
		SerialNumber  string `json:"SerialNumber"`
		IssuerName    string `json:"IssuerName"`
		SubjectName   string `json:"SubjectName"`
		Timestamp     string `json:"Timestamp"`
		Trusted       string `json:"Trusted"`
		ExtraInfo     any    `json:"_ExtraInfo"`
	} `json:"Authenticode"`
}

func (s Exchange_Windows_Forensics_UEFI) StringArray() []string {
	certificates := make([]string, 0)
	for _, v := range s.PEInfo.Authenticode.Certificates {
		certificates = append(certificates, fmt.Sprintf("| SerialNumber: %v, SignatureAlgorithm: %v, Subject: %v, Issuer: %v, NotBefore: %v, NotAfter: %v, PublicKey: %v, ExtendedKeyUsage_Critical: %v, ExtendedKeyUsage_KeyUsage: %v, SubjetKeyID_Critical: %v, SubjectKeyId_Value: %v, SubjectAlternativeName_Critical: %v, SubjectAlternativeName_DNS: %v, SubjectAlternativeName_Email: %v, SubjectAlternativeName_IP: %v, AuthorityKeyIdentifier_Critical: %v, AuthorityKeyIdentifier_KeyID: %v, CRLDistributionPoints_Critical: %v, CRLDistributionPoints_URI: %v, BasicConstraints_Critical: %v, BasicConstraints_IsCA: %v, BasicConstraints_MaxPathLen: %v", v.SerialNumber, v.SignatureAlgorithm, v.Subject, v.Issuer, v.NotBefore, v.NotAfter, v.PublicKey, v.Extensions.ExtendedKeyUsage.Critical, v.Extensions.ExtendedKeyUsage.KeyUsage, v.Extensions.SubjectKeyID.Critical, v.Extensions.SubjectKeyID.Value, v.Extensions.SubjectAlternativeName.Critical, v.Extensions.SubjectAlternativeName.DNS, v.Extensions.SubjectAlternativeName.Email, v.Extensions.SubjectAlternativeName.IP, v.Extensions.AuthorityKeyIdentifier.Critical, v.Extensions.AuthorityKeyIdentifier.KeyID, v.Extensions.CRLDistributionPoints.Critical, v.Extensions.CRLDistributionPoints.URI, v.Extensions.BasicConstraints.Critical, v.Extensions.BasicConstraints.IsCA, v.Extensions.BasicConstraints.MaxPathLen))
	}
	sections := make([]string, 0)
	for _, v := range s.PEInfo.Sections {
		sections = append(sections, fmt.Sprintf("| Perm: %v, Name: %v, FileOffset: %v, VMA: %v, RVA: %v, Size: %v", v.Perm, v.Name, v.FileOffset, v.VMA, v.RVA, v.Size))
	}
	resources := make([]string, 0)
	for _, v := range s.PEInfo.Sections {
		resources = append(resources, fmt.Sprintf("| Perm: %v, Name: %v, FileOffset: %v, VMA: %v, RVA: %v, Size: %v", v.Perm, v.Name, v.FileOffset, v.VMA, v.RVA, v.Size))
	}

	return []string{

		s.Partition.ImagePath, strconv.Itoa(s.Partition.PartitionOffset), s.Partition.PartitionSize, s.Partition.PartitionName,
		s.OSPath, strconv.Itoa(s.Size), s.Mtime.String(), s.Atime.String(), s.Ctime.String(), s.Btime.String(),
		strconv.Itoa(s.FirstCluster), s.Attr, fmt.Sprint(s.IsDeleted), s.ShortName, s.Hash.MD5, s.Hash.SHA1, s.Hash.SHA256,
		s.Magic,
		s.PEInfo.FileHeader.Machine, s.PEInfo.FileHeader.TimeDateStamp.String(), strconv.FormatInt(s.PEInfo.FileHeader.TimeDateStampRaw, 10),
		strconv.Itoa(s.PEInfo.FileHeader.Characteristics), strconv.Itoa(s.PEInfo.FileHeader.ImageBase),
		s.PEInfo.GUIDAge, s.PEInfo.PDB,
		s.PEInfo.Directories.BaseRelocationDirectory.Timestamp.String(),
		strconv.Itoa(s.PEInfo.Directories.BaseRelocationDirectory.TimestampRaw), strconv.Itoa(s.PEInfo.Directories.BaseRelocationDirectory.Size),
		strconv.Itoa(s.PEInfo.Directories.BaseRelocationDirectory.FileAddress), s.PEInfo.Directories.BaseRelocationDirectory.SectionName,
		s.PEInfo.Directories.DebugDirectory.Timestamp.String(),
		strconv.FormatInt(s.PEInfo.Directories.DebugDirectory.TimestampRaw, 10), strconv.Itoa(s.PEInfo.Directories.DebugDirectory.Size),
		strconv.Itoa(s.PEInfo.Directories.DebugDirectory.FileAddress), s.PEInfo.Directories.DebugDirectory.SectionName,
		s.PEInfo.Directories.ExceptionDirectory.Timestamp.String(),
		strconv.Itoa(s.PEInfo.Directories.ExceptionDirectory.TimestampRaw), strconv.Itoa(s.PEInfo.Directories.ExceptionDirectory.Size),
		strconv.Itoa(s.PEInfo.Directories.ExceptionDirectory.FileAddress), s.PEInfo.Directories.ExceptionDirectory.SectionName,
		s.PEInfo.Directories.ExportDirectory.Timestamp.String(),
		strconv.Itoa(s.PEInfo.Directories.ExportDirectory.TimestampRaw), strconv.Itoa(s.PEInfo.Directories.ExportDirectory.Size),
		strconv.Itoa(s.PEInfo.Directories.ExportDirectory.FileAddress), s.PEInfo.Directories.ExportDirectory.SectionName,
		s.PEInfo.Directories.ResourceDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PEInfo.Directories.ResourceDirectory.TimestampRaw)), strconv.Itoa(s.PEInfo.Directories.ResourceDirectory.Size),
		strconv.Itoa(s.PEInfo.Directories.ResourceDirectory.FileAddress), s.PEInfo.Directories.ResourceDirectory.SectionName,
		s.PEInfo.Directories.SecurityDirectory.Timestamp.String(),
		strconv.Itoa(s.PEInfo.Directories.SecurityDirectory.TimestampRaw), strconv.Itoa(s.PEInfo.Directories.SecurityDirectory.Size),
		strconv.Itoa(s.PEInfo.Directories.SecurityDirectory.FileAddress), s.PEInfo.Directories.SecurityDirectory.SectionName,

		fmt.Sprint(sections),
		fmt.Sprint(resources),

		s.PEInfo.VersionInformation.CompanyName, s.PEInfo.VersionInformation.FileDescription, s.PEInfo.VersionInformation.FileVersion,
		s.PEInfo.VersionInformation.InternalName, s.PEInfo.VersionInformation.LegalCopyright, s.PEInfo.VersionInformation.OriginalFilename,
		s.PEInfo.VersionInformation.ProductName, s.PEInfo.VersionInformation.ProductVersion,

		fmt.Sprint(s.PEInfo.Imports), fmt.Sprint(s.PEInfo.Exports), fmt.Sprint(s.PEInfo.Forwards), fmt.Sprint((s.PEInfo.ImpHash)),

		s.PEInfo.Authenticode.Signer.IssuerName, s.PEInfo.Authenticode.Signer.SerialNumber, s.PEInfo.Authenticode.Signer.DigestAlgorithm,
		s.PEInfo.Authenticode.Signer.AuthenticatedAttributes.ContentType, s.PEInfo.Authenticode.Signer.AuthenticatedAttributes.MessageDigest, s.PEInfo.Authenticode.Signer.AuthenticatedAttributes.MessageDigestHex,
		s.PEInfo.Authenticode.Signer.AuthenticatedAttributes.ProgramName, s.PEInfo.Authenticode.Signer.AuthenticatedAttributes.MoreInfo,
		fmt.Sprint(s.PEInfo.Authenticode.Signer.UnauthenticatedAttributes), s.PEInfo.Authenticode.Signer.Subject,
		fmt.Sprint(certificates), s.PEInfo.Authenticode.HashType, s.PEInfo.Authenticode.ExpectedHash, s.PEInfo.Authenticode.ExpectedHashHex,
		s.PEInfo.AuthenticodeHash.MD5, s.PEInfo.AuthenticodeHash.SHA1, s.PEInfo.AuthenticodeHash.SHA256, strconv.FormatBool(s.PEInfo.AuthenticodeHash.HashMatches),
		s.Authenticode.Filename, s.Authenticode.ProgramName, s.Authenticode.PublisherLink, s.Authenticode.MoreInfoLink, s.Authenticode.SerialNumber,
		s.Authenticode.IssuerName, s.Authenticode.SubjectName, s.Authenticode.Timestamp, s.Authenticode.Trusted, fmt.Sprint(s.Authenticode.ExtraInfo),
	}
}

func (s Exchange_Windows_Forensics_UEFI) GetHeaders() []string {
	return []string{"ImagePath", "PartitionOffset", "PartitionSize", "PartitionName",
		"OSPath", "FileSize", "Mtime", "Atime", "Ctime", "Btime",
		"FirstCluster", "Attr", "IsDeleted", "ShortName", "MD5", "SHA1", "SHA256", "Magic",
		"FileHeader_Machine",
		"FileHeader_TimeDateStamp", "FileHeader_TimeDateStampRaw", "FileHeader_Characteristics", "FileHeader_ImageBase",
		"GUIDAge", "PDB",
		"BaseRelocationDir_Timestamp", "BaseRelocationDir_TimestampRaw", "BaseRelocationDir_Size", "BaseRelocationDir_FileAddress",
		"BaseRelocationDir_SectionName",
		"DebugDirectory_Timestamp", "DebugDirectory_TimestampRaw", "DebugDirectory_Size", "DebugDirectory_FileAddress",
		"DebugDirectory_SectionName",
		"ExceptionDirectory_Timestamp", "ExceptionDirectory_TimestampRaw", "ExceptionDirectory_Size", "ExceptionDirectory_FileAddress",
		"ExceptionDirectory_SectionName",
		"ExportDirectory_Timestamp", "ExportDirectory_TimestampRaw", "ExportDirectory_Size", "ExportDirectory_FileAddress",
		"ExportDirectory_SectionName",
		"ResourceDirectory_Timestamp", "ResourceDirectory_TimestampRaw", "ResourceDirectory_Size", "ResourceDirectory_FileAddress",
		"ResourceDirectory_SectionName",
		"SecurityDirectory_Timestamp", "SecurityDirectory_TimestampRaw", "SecurityDirectory_Size", "SecurityDirectory_FileAddress",
		"SecurityDirectory_SectionName",
		"Sections", "Resources",
		"Version_CompanyName", "Version_FileDescription", "Version_FileVersion", "Version_InternalName", "Version_LegalCopyright",
		"Version_OriginalFileName", "Version_ProductName", "Version_ProductVersion",
		"Imports", "Exports", "Forwards", "ImpHash",
		"Authenticode_AA_IssuerName", "Authenticode_AA_SerialNumber", "Authenticode_AA_DigestAlgorithm",
		"Authenticode_AA_ContentType", "Authenticode_AA_MessageDigest", "Authenticode_AA_MessageDigestHex", "Authenticode_AA_ProgramName", "Authenticode_MoreInfo",
		"Authenticode_UnauthenticatedAttributes", "Authenticode_Subject", "Authenticode_Certificates", "Authenticde_HashType", "Authenticode_ExpectedHash",
		"Authenticode_ExpectedHashHex", "Authenticode_MD5", "Authenticode_SHA1", "Authenticode_SHA256", "Authenticode_HashMatches",
		"Authenticode_Filename", "Authenticode_ProgramName", "Authenticode_PublisherLink", "Authenticode_MoreInfoLink", "Authenticode_IssuerName",
		"Authenticode_SubjectName", "Authenticode_Timestamp", "Authenticode_Trusted", "Authenticode_ExtraInfo"}
}

func Process_Exchange_Windows_Forensics_UEFI(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_Forensics_UEFI{}
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
			EventType:        "UEFI Entry Modified",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.OSPath,
			MetaData:         fmt.Sprintf("Imphash: %v, Magic: %v, Ctime: %v, PDB: %v, Imports: %v", tmp.PEInfo.ImpHash, tmp.Magic, tmp.Ctime, &tmp.PEInfo.PDB, tmp.PEInfo.Imports),
		}
		outputChannel <- tmp2.StringArray()
	}
}
