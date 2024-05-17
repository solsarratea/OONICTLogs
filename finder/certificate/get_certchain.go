package certificate

import (
	"encoding/json"
	"fmt"
)

type Certificate struct {
	Data string `json:"data"`
}

type TLSHandshake struct {
	PeerCertificates []Certificate `json:"peer_certificates"`
}

type TestKeys struct {
	TLSHandshakes []TLSHandshake `json:"tls_handshakes"`
}

type Measurement struct {
	TestKeys  TestKeys `json:"test_keys"`
	URL       string   `json:"input"`
	StartTime string   `json:"test_start_time"`
}

func DecodeMeasurement(body []byte) (Measurement, error) {
	var response Measurement

	err := json.Unmarshal(body, &response)
	if err != nil {
		return Measurement{}, fmt.Errorf("Error parsing response: %v", err)
	}

	return response, nil
}

func GetCertificateChain(response Measurement) ([][]string, error) {

	var handshakesCChain [][]string
	for _, handshake := range response.TestKeys.TLSHandshakes {
		var cchain []string

		for _, cert := range handshake.PeerCertificates {

			cchain = append(cchain, cert.Data)
		}
		handshakesCChain = append(handshakesCChain, cchain)
	}

	return handshakesCChain, nil
}

//DEBUGGING FUNCTIONS FOR READING CERTIFICATES:

// Print certificate details (optional)
//fmt.Printf("Certificate Details:\n")
//fmt.Printf("  Subject: %s\n", parsedCert.Subject)
//fmt.Printf("  Issuer: %s\n", parsedCert.Issuer)
