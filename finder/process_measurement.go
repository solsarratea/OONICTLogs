package finder

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"./certchain"
	"./queries"
	"./utils"
)

func ProcessMeasurements(config Configuration) {
	//TODO: Add propper logs of success/lost measurements
	apiEndpoint, _ := utils.ReadLineFromFile(config.PathMeasurements)

	re := regexp.MustCompile(`measurement_uid=([^&]+)`)
	measurement_uid := re.FindStringSubmatch(apiEndpoint)[1]

	fmt.Println("Processing measure with id: " + measurement_uid)
	body, err := queries.QuerySingleMeasurement(apiEndpoint)

	if err != nil {
		fmt.Printf("Failed to query measurements: %v\n", err)

	} else {

		cchain, _ := certchain.GetCertificateChain(body)

		for i, subchain := range cchain {
			content := strings.Join(subchain, "\n")
			path := config.PathCert + measurement_uid + "-chain-" + strconv.Itoa(i)
			utils.WriteStringToFile(content, path)
		}
	}

}
