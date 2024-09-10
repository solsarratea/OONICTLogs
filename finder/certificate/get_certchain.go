package certificate

import (
	"encoding/json"
	"fmt"
)

// Certificate represents child structure of the response from API call get measurement from OONI
type Certificate struct {
	Data string `json:"data"`
}

// TLSHandshake represents child structure of the response from API call get measurement from OONI
type TLSHandshake struct {
	PeerCertificates []Certificate `json:"peer_certificates"`
}

// TestKeys represents child structure of the response from API call get measurement from OONI
type TestKeys struct {
	TLSHandshakes []TLSHandshake `json:"tls_handshakes"`
}

// Measurement represents child structure of the response from API call get measurement from OONI
type Measurement struct {
	TestKeys  TestKeys `json:"test_keys"`
	URL       string   `json:"input"`
	StartTime string   `json:"test_start_time"`
}

// DecodeMeasurement implements the json.Unmarshal interface for Measurements
func DecodeMeasurement(body []byte) (Measurement, error) {
	var response Measurement

	err := json.Unmarshal(body, &response)
	if err != nil {
		return Measurement{}, fmt.Errorf("Error parsing response: %v", err)
	}

	return response, nil
}

// GetCertificateChain extract the raw certificate data from the tls_handshakes
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
