package artifact_structs

import (
	"encoding/json"
	"fmt"
	"github.com/joeavanzato/velo-timeline-creator/helpers"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type Windows_Sys_CertificateAuthorities struct {
	Store   string `json:"Store"`
	IsCA    bool   `json:"IsCA"`
	Subject struct {
		Country            []string `json:"Country"`
		Organization       []string `json:"Organization"`
		OrganizationalUnit []string `json:"OrganizationalUnit"`
		Locality           []string `json:"Locality"`
		Province           []string `json:"Province"`
		StreetAddress      any      `json:"StreetAddress"`
		PostalCode         any      `json:"PostalCode"`
		SerialNumber       string   `json:"SerialNumber"`
		CommonName         string   `json:"CommonName"`
		Names              []struct {
			Type  []int  `json:"Type"`
			Value string `json:"Value"`
		} `json:"Names"`
		ExtraNames any `json:"ExtraNames"`
	} `json:"Subject"`
	SubjectKeyID   string `json:"SubjectKeyId"`
	AuthorityKeyID string `json:"AuthorityKeyId"`
	Issuer         struct {
		Country            []string `json:"Country"`
		Organization       []string `json:"Organization"`
		OrganizationalUnit []string `json:"OrganizationalUnit"`
		Locality           []string `json:"Locality"`
		Province           []string `json:"Province"`
		StreetAddress      any      `json:"StreetAddress"`
		PostalCode         any      `json:"PostalCode"`
		SerialNumber       string   `json:"SerialNumber"`
		CommonName         string   `json:"CommonName"`
		Names              []struct {
			Type  []int  `json:"Type"`
			Value string `json:"Value"`
		} `json:"Names"`
		ExtraNames any `json:"ExtraNames"`
	} `json:"Issuer"`
	KeyUsageString     string    `json:"KeyUsageString"`
	IsSelfSigned       bool      `json:"IsSelfSigned"`
	SHA1               string    `json:"SHA1"`
	SignatureAlgorithm int       `json:"SignatureAlgorithm"`
	PublicKeyAlgorithm int       `json:"PublicKeyAlgorithm"`
	KeyStrength        int       `json:"KeyStrength"`
	NotBefore          time.Time `json:"NotBefore"`
	NotAfter           time.Time `json:"NotAfter"`
	HexSerialNumber    string    `json:"HexSerialNumber"`
}

func (s Windows_Sys_CertificateAuthorities) StringArray() []string {
	return []string{s.Store, fmt.Sprint(s.IsCA), fmt.Sprint(s.Subject.Country), fmt.Sprint(s.Subject.Organization), fmt.Sprint(s.Subject.OrganizationalUnit),
		fmt.Sprint(s.Subject.Locality), fmt.Sprint(s.Subject.Province), fmt.Sprint(s.Subject.StreetAddress), fmt.Sprint(s.Subject.PostalCode), s.Subject.SerialNumber, s.Subject.CommonName,
		fmt.Sprint(s.Subject.Names), fmt.Sprint(s.Subject.ExtraNames), s.SubjectKeyID, s.AuthorityKeyID, fmt.Sprint(s.Issuer.Country), fmt.Sprint(s.Issuer.Organization), fmt.Sprint(s.Issuer.OrganizationalUnit),
		fmt.Sprint(s.Issuer.Locality), fmt.Sprint(s.Issuer.Province), fmt.Sprint(s.Issuer.StreetAddress), fmt.Sprint(s.Issuer.PostalCode), s.Issuer.SerialNumber, s.Issuer.CommonName,
		fmt.Sprint(s.Issuer.Names), fmt.Sprint(s.Issuer.ExtraNames), s.KeyUsageString, fmt.Sprint(s.IsSelfSigned), s.SHA1, strconv.Itoa(s.SignatureAlgorithm), strconv.Itoa(s.PublicKeyAlgorithm), strconv.Itoa(s.KeyStrength), s.NotBefore.String(), s.NotAfter.String(),
		s.HexSerialNumber}
}

func (s Windows_Sys_CertificateAuthorities) GetHeaders() []string {
	return []string{"Store", "IsCA", "Subject_Country", "Subject_Organization", "Subject_OrganizationUnit", "Subject_Locality", "Subject_Province",
		"Subject_StreetAddress", "Subject_PostalCode", "Subject_SerialNumber", "Subject_CommonName", "Subject_Names", "Subject_ExtraNames",
		"SubjectKeyID", "AuthorityKeyID", "Issuer_Country", "Issuer_Organization", "Issuer_OrganizationUnit", "Issuer_Locality", "Issuer_Province", "Issuer_StreetAddress",
		"Issuer_PostalCode", "Issuer_SerialNumber", "Issuer_CommonName", "Issuer_Names", "Issuer_ExtraNames", "KeyUsageString", "IsSelfSigned", "SHA1", "SignatureAlgorithm", "PublicKeyAlgorithm",
		"KeyStrength", "NotBefore", "NotAfter", "HexSerialNumber"}
}

func Process_Windows_Sys_CertificateAuthorities(artifactName string, clientIdentifier string, inputLines []string, outputChannel chan<- []string, arguments map[string]any, logger zerolog.Logger) {
	// Receives lines from a file, unmarshalls to appropriate struct and sends the newly constructed array of ShallowRecords string to the output channel
	for _, line := range inputLines {
		tmp := Windows_Sys_CertificateAuthorities{}
		err := json.Unmarshal([]byte(line), &tmp)
		if err != nil {
			logger.Error().Msgf(err.Error())
			continue
		}
		if arguments["artifactdump"].(bool) {
			helpers.BuildAndSendArtifactRecord("", clientIdentifier, "", tmp.StringArray(), outputChannel)
			continue
		}
	}
}
