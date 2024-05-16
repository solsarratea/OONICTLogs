package finder

import (
	"fmt"
	"io"
	"net/http"

	"../common"
	"./measurements"
	"./utils"
)

func QueryMeasurements(config common.Configuration) ([]byte, error) {
	//since: 2023-12-20T00%3A00%3A00
	//until: 2024-07-20T00%3A00%3A00

	apiEndpoint := fmt.Sprintf("https://api.ooni.io/api/v1/measurements?test_name=web_connectivity&since=%s&until=%s&failure=false&order_by=measurement_start_time&order=asc&limit=100", config.OONIMeasurements.Since, config.OONIMeasurements.Until)

	fmt.Printf(apiEndpoint)

	response, err := http.Get(apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from API: %v", err)
	}
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

func GetRawMeasurements(config common.Configuration) {
	body, err := QueryMeasurements(config)
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
}
