package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

//AppendHeadersFooters implements the interface of creating a PEM cert out of raw certificate string
func AppendHeadersFooters(cert string) string {
	return fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----", cert)
}

//ParsePEMString implements the interface of decoding a PEM certificate into x.509
func ParsePEMString(pemCert string) (*x509.Certificate, error) {

	block, _ := pem.Decode([]byte(pemCert))
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("Failed to decode PEM block containing certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse x.509: %w", err)
	}

	return cert, nil
}
