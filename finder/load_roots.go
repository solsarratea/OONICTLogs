package finder

import (
	"fmt"
	"io"
	"net/http"

	"github.com/solsarratea/OONICTLogs/common"
	"github.com/solsarratea/OONICTLogs/finder/roots"
)

func QueryRootCertificates(CTLogURI string) ([]byte, error) {
	apiEndpoint := CTLogURI + "/ct/v1/get-roots"

	response, err := http.Get(apiEndpoint)
	if err != nil {
		return []byte{}, fmt.Errorf("error fetching data from API: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("error: API returned status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil

}

// LoadRoots implements the querying and parsing of the CTLogs roots.
// Returns an array of *x509.Certificates.
func LoadRoots(config common.Configuration) (roots.Roots, error) {

	body, err := QueryRootCertificates(config.CTLog.URI)

	if err != nil {
		return roots.Roots{}, fmt.Errorf("error fetching the root nodes from %s", config.CTLog.URI)

	}

	return roots.ParseRootCertificates(body)

}
