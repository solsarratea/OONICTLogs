package finder

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/solsarratea/OONICTLogs/common"
	"github.com/solsarratea/OONICTLogs/finder/measurements"
)

func QueryMeasurements(config common.Configuration) ([]byte, error) {

	//  Querying window for raw measurements depend on config
	apiEndpoint := "https://api.ooni.io/api/v1/measurements?test_name=web_connectivity&since=" + url.QueryEscape(config.OONIMeasurements.Since) + "&until=" + url.QueryEscape(config.OONIMeasurements.Until) + "&failure=false&order_by=measurement_start_time&order=asc&limit=100"

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

// GetRawMeasurements extract URLs for single measurement reports and write it into a text file.
func GetRawMeasurements(config common.Configuration) {
	body, err := QueryMeasurements(config)
	if err != nil {
		fmt.Printf("Failed to query measurements: %v\n", err)

	} else {

		rawMeasurements, err := measurements.DecodeMeasurements(body)

		if err != nil {
			fmt.Printf("Failed to decode raw measurements %v\n", err)
		} else {
			writeMeasurementsToFile(config.PathMeasurements, rawMeasurements)
		}
	}
}

func writeMeasurementsToFile(url string, rawMeasurements measurements.RawMeasurements) error {
	file, err := os.Create(url)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	defer file.Close()

	for _, result := range rawMeasurements.Results {
		_, err := file.WriteString(result.URL + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}

	fmt.Println("URLs have been written to %s \n", url)
	return nil
}
