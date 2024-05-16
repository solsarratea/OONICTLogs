package finder

import (
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"./certificate"
	"./roots"
	"./utils"
)

type ValidSubmission struct {
	URL              string
	Root             *x509.Certificate
	CertificateChain []string
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

func ProcessMeasurement(config Configuration, collection roots.Roots) error {
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
					url, _ := certificate.GetURL(body)
					cert, _ := certificate.ParsePEMString(c)
					hasRoot := roots.FindParent(cert, collection)

					if hasRoot != nil {
						fmt.Printf("\nWEHAVEONE ＼(＾O＾)／	 \n")

						submission := ValidSubmission{
							Root:             hasRoot,
							CertificateChain: subchain,
							URL:              url,
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
