package roots

import (
	"crypto/x509"
	"encoding/json"
	"fmt"

	"../certificate"
)

type RawRoots struct {
	Certificates []string `json:"certificates"`
}

type Roots []*x509.Certificate

func ParseRootCertificates(body []byte) (Roots, error) {
	var rroots RawRoots

	err := json.Unmarshal(body, &rroots)
	if err != nil {
		return Roots{}, fmt.Errorf("Error parsing the JSON result")
	}

	parsedRoots := Roots{}
	for _, cert := range rroots.Certificates {
		pemCert := certificate.AppendHeadersFooters(cert)

		parsedCert, _ := certificate.ParsePEMString(pemCert)

		parsedRoots = append(parsedRoots, parsedCert)
	}

	return parsedRoots, nil
}
