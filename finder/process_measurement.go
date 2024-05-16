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

type ValidSubmission struct {
	Chain []string
}

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

func ProcessMeasurement(config common.Configuration, collection roots.Roots) error {
	apiEndpoint, _ := utils.ReadLineFromFile(config.PathMeasurements)

	re := regexp.MustCompile(`measurement_uid=([^&]+)`)
	measurement_uid := re.FindStringSubmatch(apiEndpoint)[1]

	fmt.Println("Processing measure with id: " + measurement_uid)
	body, err := QuerySingleMeasurement(apiEndpoint)

	if err != nil {
		return fmt.Errorf("Failed to query measurements: %v\n", err)

	} else {
		cchain, _ := certificate.GetCertificateChain(body)
		if len(cchain) > 0 {
			//FIXME: Check TLS handshakes are resolving to the same certificate chane
			subchain := cchain[0]

			if len(subchain) > 0 {
				for _, c := range subchain {
					ch := certificate.AppendHeadersFooters(c)
					cert, _ := certificate.ParsePEMString(ch)
					hasRoot := roots.FindParent(cert, collection)

					if hasRoot != nil {
						fmt.Printf("\nWEHAVEONE ＼(＾O＾)／	 \n")
						base64Cert := base64.StdEncoding.EncodeToString(hasRoot.Raw)
						//	fmt.Printf(base64Cert)
						final := append(subchain, base64Cert)
						//	fmt.Printf(subchain[0])

						submission := map[string]interface{}{
							"chain": final,
						}

						path := config.PathCert + measurement_uid
						err := utils.WriteStructToJSONFile(submission, path)

						if err != nil {
							return fmt.Errorf("Failed to write file with valid measurement: %s\n", err)
						}

						return nil
					}
				}
			}
		}
		return fmt.Errorf("Did not found valid root node")

		/* Obsolete code for writing and analysing cert chains
		for i, subchain := range cchain {
			content := strings.Join(subchain, "\n")

		}*/
	}
}
