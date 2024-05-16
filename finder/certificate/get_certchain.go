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
	TestKeys TestKeys `json:"test_keys"`
	URL      string   `json:"input"`
}

func GetCertificateChain(body []byte) ([][]string, error) {
	var response Measurement

	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("Error parsing response: %v", err)
	}

	var handshakesCChain [][]string
	for _, handshake := range response.TestKeys.TLSHandshakes {
		var cchain []string

		for _, cert := range handshake.PeerCertificates {
			pemCert := AppendHeadersFooters(cert.Data)

			_, err := ParsePEMString(pemCert)

			if err != nil {
				return nil, fmt.Errorf("Error parsing certificate: %v", err)
			}

			cchain = append(cchain, pemCert)
		}
		handshakesCChain = append(handshakesCChain, cchain)
	}

	return handshakesCChain, nil
}

func GetURL(body []byte) (string, error) {
	//FIXME: Unnecesary unmarshall
	var response Measurement

	err := json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("Error parsing response: %v", err)
	}
	return response.URL, nil
}
