package finder

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"./certchain"
	"./measurements"
	"./ooniapi"
	"./utils"
)

func Start() {

	fmt.Println("Starting Chainfinder")
	RAW_MEASUREMENTS_PATH := "raw_measurements.txt"
	CERTIFICATES_PATH := "_certificates/"

	for {

		isEmpty, _ := utils.IsFileEmpty(RAW_MEASUREMENTS_PATH)

		if isEmpty {

			body, err := ooniapi.QueryMeasurements()
			if err != nil {
				fmt.Printf("Failed to query measurements: %v\n", err)

			} else {

				rawMeasurements, err := measurements.DecodeMeasurements(body)

				if err != nil {
					fmt.Printf("Failed to decode raw measurements %v\n", err)
				} else {
					utils.WriteMeasurementsToFile(rawMeasurements)
				}

			}

		} else {
			//TODO: Add propper logs of success/lost measurements
			apiEndpoint, _ := utils.ReadLineFromFile(RAW_MEASUREMENTS_PATH)

			re := regexp.MustCompile(`measurement_uid=([^&]+)`)
			measurement_uid := re.FindStringSubmatch(apiEndpoint)[1]

			body, err := ooniapi.QuerySingleMeasurement(apiEndpoint)

			if err != nil {
				fmt.Printf("Failed to query measurements: %v\n", err)

			} else {

				cchain, _ := certchain.GetCertificateChain(body)

				for i, subchain := range cchain {
					content := strings.Join(subchain, "\n")
					path := CERTIFICATES_PATH + measurement_uid + "-chain-" + strconv.Itoa(i)
					utils.WriteStringToFile(content, path)
				}
			}

			utils.RemoveLineFromFile(RAW_MEASUREMENTS_PATH)

		}
	}
	time.Sleep(10 * time.Second)
}
