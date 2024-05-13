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

type Exchange_Windows_System_PrinterDriver_BinaryCheck struct {
	Binary      string   `json:"Binary"`
	DriverNames []string `json:"DriverNames"`
	Hash        struct {
		MD5    string `json:"MD5"`
		SHA1   string `json:"SHA1"`
		SHA256 string `json:"SHA256"`
	} `json:"Hash"`
	PE struct {
		FileHeader struct {
			Machine          string    `json:"Machine"`
			TimeDateStamp    time.Time `json:"TimeDateStamp"`
			TimeDateStampRaw int64     `json:"TimeDateStampRaw"`
			Characteristics  int       `json:"Characteristics"`
			ImageBase        int64     `json:"ImageBase"`
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
			DelayImportsDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int64     `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Delay_Imports_Directory"`
			ExceptionDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Exception_Directory"`
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
			LoadConfigDirectory struct {
				Timestamp    time.Time `json:"Timestamp"`
				TimestampRaw int       `json:"TimestampRaw"`
				Size         int       `json:"Size"`
				FileAddress  int       `json:"FileAddress"`
				SectionName  string    `json:"SectionName"`
			} `json:"Load_Config_Directory"`
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
			VMA        int64  `json:"VMA"`
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
	} `json:"PE"`
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
		ExtraInfo     struct {
			Catalog string `json:"Catalog"`
		} `json:"_ExtraInfo"`
	} `json:"Authenticode"`
}

func (s Exchange_Windows_System_PrinterDriver_BinaryCheck) StringArray() []string {
	certificates := make([]string, 0)
	for _, v := range s.PE.Authenticode.Certificates {
		certificates = append(certificates, fmt.Sprintf("| SerialNumber: %v, SignatureAlgorithm: %v, Subject: %v, Issuer: %v, NotBefore: %v, NotAfter: %v, PublicKey: %v, ExtendedKeyUsage_Critical: %v, ExtendedKeyUsage_KeyUsage: %v, SubjetKeyID_Critical: %v, SubjectKeyId_Value: %v, SubjectAlternativeName_Critical: %v, SubjectAlternativeName_DNS: %v, SubjectAlternativeName_Email: %v, SubjectAlternativeName_IP: %v, AuthorityKeyIdentifier_Critical: %v, AuthorityKeyIdentifier_KeyID: %v, CRLDistributionPoints_Critical: %v, CRLDistributionPoints_URI: %v, BasicConstraints_Critical: %v, BasicConstraints_IsCA: %v, BasicConstraints_MaxPathLen: %v", v.SerialNumber, v.SignatureAlgorithm, v.Subject, v.Issuer, v.NotBefore, v.NotAfter, v.PublicKey, v.Extensions.ExtendedKeyUsage.Critical, v.Extensions.ExtendedKeyUsage.KeyUsage, v.Extensions.SubjectKeyID.Critical, v.Extensions.SubjectKeyID.Value, v.Extensions.SubjectAlternativeName.Critical, v.Extensions.SubjectAlternativeName.DNS, v.Extensions.SubjectAlternativeName.Email, v.Extensions.SubjectAlternativeName.IP, v.Extensions.AuthorityKeyIdentifier.Critical, v.Extensions.AuthorityKeyIdentifier.KeyID, v.Extensions.CRLDistributionPoints.Critical, v.Extensions.CRLDistributionPoints.URI, v.Extensions.BasicConstraints.Critical, v.Extensions.BasicConstraints.IsCA, v.Extensions.BasicConstraints.MaxPathLen))
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

		s.Binary, fmt.Sprint(s.DriverNames), s.Hash.MD5, s.Hash.SHA1, s.Hash.SHA256,
		s.PE.FileHeader.Machine, s.PE.FileHeader.TimeDateStamp.String(), strconv.FormatInt(s.PE.FileHeader.TimeDateStampRaw, 10),
		strconv.Itoa(s.PE.FileHeader.Characteristics), strconv.Itoa(int(s.PE.FileHeader.ImageBase)),
		s.PE.GUIDAge, s.PE.PDB,
		s.PE.Directories.BaseRelocationDirectory.Timestamp.String(),
		strconv.Itoa(s.PE.Directories.BaseRelocationDirectory.TimestampRaw), strconv.Itoa(s.PE.Directories.BaseRelocationDirectory.Size),
		strconv.Itoa(s.PE.Directories.BaseRelocationDirectory.FileAddress), s.PE.Directories.BaseRelocationDirectory.SectionName,
		s.PE.Directories.DebugDirectory.Timestamp.String(),
		strconv.FormatInt(s.PE.Directories.DebugDirectory.TimestampRaw, 10), strconv.Itoa(s.PE.Directories.DebugDirectory.Size),
		strconv.Itoa(s.PE.Directories.DebugDirectory.FileAddress), s.PE.Directories.DebugDirectory.SectionName,
		s.PE.Directories.DelayImportsDirectory.Timestamp.String(),
		strconv.FormatInt(s.PE.Directories.DelayImportsDirectory.TimestampRaw, 10), strconv.Itoa(s.PE.Directories.DelayImportsDirectory.Size),
		strconv.Itoa(s.PE.Directories.DelayImportsDirectory.FileAddress), s.PE.Directories.DelayImportsDirectory.SectionName,
		s.PE.Directories.ExceptionDirectory.Timestamp.String(),
		strconv.Itoa(s.PE.Directories.ExceptionDirectory.TimestampRaw), strconv.Itoa(s.PE.Directories.ExceptionDirectory.Size),
		strconv.Itoa(s.PE.Directories.ExceptionDirectory.FileAddress), s.PE.Directories.ExceptionDirectory.SectionName,
		s.PE.Directories.ExportDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.ExportDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.ExportDirectory.Size),
		strconv.Itoa(s.PE.Directories.ExportDirectory.FileAddress), s.PE.Directories.ExportDirectory.SectionName,
		s.PE.Directories.IATDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.IATDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.IATDirectory.Size),
		strconv.Itoa(s.PE.Directories.IATDirectory.FileAddress), s.PE.Directories.IATDirectory.SectionName,
		s.PE.Directories.ImportDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.ImportDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.ImportDirectory.Size),
		strconv.Itoa(s.PE.Directories.ImportDirectory.FileAddress), s.PE.Directories.ImportDirectory.SectionName,
		s.PE.Directories.LoadConfigDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.LoadConfigDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.LoadConfigDirectory.Size),
		strconv.Itoa(s.PE.Directories.LoadConfigDirectory.FileAddress), s.PE.Directories.LoadConfigDirectory.SectionName,
		s.PE.Directories.ResourceDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.ResourceDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.ResourceDirectory.Size),
		strconv.Itoa(s.PE.Directories.ResourceDirectory.FileAddress), s.PE.Directories.ResourceDirectory.SectionName,
		s.PE.Directories.TLSDirectory.Timestamp.String(),
		strconv.Itoa(int(s.PE.Directories.TLSDirectory.TimestampRaw)), strconv.Itoa(s.PE.Directories.TLSDirectory.Size),
		strconv.Itoa(s.PE.Directories.TLSDirectory.FileAddress), s.PE.Directories.TLSDirectory.SectionName,
		s.PE.Directories.SecurityDirectory.Timestamp.String(),
		strconv.Itoa(s.PE.Directories.SecurityDirectory.TimestampRaw), strconv.Itoa(s.PE.Directories.SecurityDirectory.Size),
		strconv.Itoa(s.PE.Directories.SecurityDirectory.FileAddress), s.PE.Directories.SecurityDirectory.SectionName,

		fmt.Sprint(sections),
		fmt.Sprint(resources),

		s.PE.VersionInformation.CompanyName, s.PE.VersionInformation.FileDescription, s.PE.VersionInformation.FileVersion,
		s.PE.VersionInformation.InternalName, s.PE.VersionInformation.LegalCopyright, s.PE.VersionInformation.OriginalFilename,
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

func (s Exchange_Windows_System_PrinterDriver_BinaryCheck) GetHeaders() []string {
	return []string{
		"Binary", "DriverNames", "MD5", "SHA1", "SHA256",
		"FileHeader_Machine",
		"FileHeader_TimeDateStamp", "FileHeader_TimeDateStampRaw", "FileHeader_Characteristics", "FileHeader_ImageBase",
		"GUIDAge", "PDB",
		"BaseRelocationDir_Timestamp", "BaseRelocationDir_TimestampRaw", "BaseRelocationDir_Size", "BaseRelocationDir_FileAddress",
		"BaseRelocationDir_SectionName",
		"DebugDirectory_Timestamp", "DebugDirectory_TimestampRaw", "DebugDirectory_Size", "DebugDirectory_FileAddress",
		"DebugDirectory_SectionName",
		"DelayImportsDirectory_Timestamp", "DelayImportsDirectory_TimestampRaw", "DelayImportsDirectory_Size", "DelayImportsDirectory_FileAddress",
		"DelayImportsDirectory_SectionName",
		"ExceptionDirectory_Timestamp", "ExceptionDirectory_TimestampRaw", "ExceptionDirectory_Size", "ExceptionDirectory_FileAddress",
		"ExceptionDirectory_SectionName",
		"ExportDirectory_Timestamp", "ExportDirectory_TimestampRaw", "ExportDirectory_Size", "ExportDirectory_FileAddress",
		"ExportDirectory_SectionName",
		"IATDirectory_Timestamp", "IATDirectory_TimestampRaw", "IATDirectory_Size", "IATDirectory_FileAddress",
		"IATDirectory_SectionName",
		"ImportDirectory_Timestamp", "ImportDirectory_TimestampRaw", "ImportDirectory_Size", "ImportDirectory_FileAddress",
		"ImportDirectory_SectionName",
		"LoadConfigDirectory_Timestamp", "LoadConfigDirectory_TimestampRaw", "LoadConfigDirectory_Size", "LoadConfigDirectory_FileAddress",
		"LoadConfigDirectory_SectionName",
		"ResourceDirectory_Timestamp", "ResourceDirectory_TimestampRaw", "ResourceDirectory_Size", "ResourceDirectory_FileAddress",
		"ResourceDirectory_SectionName",
		"TLSDirectory_Timestamp", "TLSDirectory_TimestampRaw", "TLSDirectory_Size", "TLSDirectory_FileAddress",
		"TLSDirectory_SectionName",
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

func Process_Exchange_Windows_System_PrinterDriver_BinaryCheck(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_System_PrinterDriver_BinaryCheck{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord(tmp.PE.FileHeader.TimeDateStamp.String(), clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		tmp2 := vars.ShallowRecord{
			Timestamp:        tmp.PE.FileHeader.TimeDateStamp,
			Computer:         clientIdentifier,
			Artifact:         artifactName,
			EventType:        "Print Driver",
			EventDescription: "",
			SourceUser:       "",
			SourceHost:       "",
			DestinationUser:  "",
			DestinationHost:  "",
			SourceFile:       tmp.Binary,
			MetaData:         fmt.Sprintf("MD5: %v", tmp.Hash.MD5),
		}
		outputChannel <- tmp2.StringArray()
	}
}

type Exchange_Windows_System_PrinterDriver struct {
	Name              string `json:"Name"`
	SupportedPlatform string `json:"SupportedPlatform"`
	Version           int    `json:"Version"`
	DriverPath        string `json:"DriverPath"`
	ConfigFile        string `json:"ConfigFile"`
	DataFile          string `json:"DataFile"`
}

func (s Exchange_Windows_System_PrinterDriver) StringArray() []string {
	return helpers.GetStructValuesAsStringSlice(s)
}

func (s Exchange_Windows_System_PrinterDriver) GetHeaders() []string {
	return helpers.GetStructHeadersAsStringSlice(s)
}

func Process_Exchange_Windows_System_PrinterDriver(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Exchange_Windows_System_PrinterDriver{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}

		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
		continue
	}
}
