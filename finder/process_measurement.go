package finder

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

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

func ProcessMeasurement(config Configuration, collection roots.Roots) {
	//TODO: Add propper logs of success/lost measurements
	apiEndpoint, _ := utils.ReadLineFromFile(config.PathMeasurements)

	re := regexp.MustCompile(`measurement_uid=([^&]+)`)
	measurement_uid := re.FindStringSubmatch(apiEndpoint)[1]

	fmt.Println("Processing measure with id: " + measurement_uid)
	body, err := QuerySingleMeasurement(apiEndpoint)

	if err != nil {
		fmt.Printf("Failed to query measurements: %v\n", err)

	} else {

		cchain, _ := certificate.GetCertificateChain(body)
		if len(cchain) > 0 {
			subchain := cchain[0]

			if len(subchain) > 0 {
				for _, c := range subchain {
					cert, _ := certificate.ParsePEMString(c)
					isRoot := roots.Contained(cert, collection)
					if isRoot {
						fmt.Println("wehavetheone")
					} else {
						fmt.Println("ni")
					}
				}
			}
		}
		//for i, subchain := range cchain {

		//	content := strings.Join(subchain, "\n")
		//	path := config.PathCert + measurement_uid + "-chain-" + strconv.Itoa(i)
		//	utils.WriteStringToFile(content, path)
		//}
	}

}
