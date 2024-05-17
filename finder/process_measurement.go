package finder

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"../common"
	"./certificate"
	"./roots"
	"./utils"
)

func QuerySingleMeasurement(apiEndpoint string) ([]byte, error) {
	response, err := http.Get(apiEndpoint)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: API returned status code %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

// ProcessMeasurement implements the querying, finding and storing of valid certificate chains.
//
// Receives config and loaded certificate roots from CTLogs
// Gets single measurement, extract valid certificate chains from tls_handshakes.
// Makes a linear search to find a if resolves to a valid root.
// Writes  entry into specified directory (in config.json) with valid format according to https://crt.sh/gen-add-chain
// Returns the starting measurement time.

func ProcessMeasurement(config common.Configuration, collection roots.Roots) (string, error) {

	apiEndpoint, _ := utils.ReadLineFromFile(config.PathMeasurements)

	re := regexp.MustCompile(`measurement_uid=([^&]+)`)
	measurement_uid := re.FindStringSubmatch(apiEndpoint)[1]

	fmt.Println("Processing measure with id: " + measurement_uid)
	body, err := QuerySingleMeasurement(apiEndpoint)
	updatedSince := config.OONIMeasurements.Since

	if err != nil {
		return updatedSince, fmt.Errorf("Failed to query measurements: %v\n", err)

	} else {
		measurement, _ := certificate.DecodeMeasurement(body)

		cchain, _ := certificate.GetCertificateChain(measurement)
		if len(cchain) > 0 {
			//FIXME: Check TLS handshakes are resolving to the same certificate chain
			subchain := cchain[0]

			if len(subchain) > 0 {
				for _, c := range subchain {
					ch := certificate.AppendHeadersFooters(c)
					cert, _ := certificate.ParsePEMString(ch)
					hasRoot := roots.FindParent(cert, collection)

					if hasRoot != nil {
						fmt.Printf("\nWEHAVEONE ＼(＾O＾)／	 \n")
						base64Cert := base64.StdEncoding.EncodeToString(hasRoot.Raw)

						final := append(subchain, base64Cert)

						submission := map[string]interface{}{
							"chain": final,
						}

						path := config.PathCert + measurement_uid
						err := utils.WriteStructToJSONFile(submission, path)

						if err != nil {
							return updatedSince, fmt.Errorf("Failed to write file with valid measurement: %s\n", err)
						}

						return updatedSince, nil
					}
				}
			}
		}

		return updatedSince, fmt.Errorf("Did not found valid root node")
	}
}
