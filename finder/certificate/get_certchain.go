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

			//TODO 1: Add a validation function
			// should include - valid chain according RF6962
			// maybe check root certificate not to be signed by CA

			cchain = append(cchain, pemCert)
		}
		handshakesCChain = append(handshakesCChain, cchain)
		//TODO 2: Add a way of chosing the proper/right chain, seems that for all handshakes is the same
	}

	return handshakesCChain, nil
}

//DEBUGGING FUNCTIONS FOR READING CERTIFICATES:

// Print certificate details (optional)
//fmt.Printf("Certificate Details:\n")
//fmt.Printf("  Subject: %s\n", parsedCert.Subject)
//fmt.Printf("  Issuer: %s\n", parsedCert.Issuer)
