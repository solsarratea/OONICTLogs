package measurements

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type MeasurementMeta struct {
	Anomaly              bool   `json:"anomaly"`
	Confirmed            bool   `json:"confirmed"`
	Failure              bool   `json:"failure"`
	Input                string `json:"input"`
	MeasurementStartTime string `json:"measurement_start_time"`
	UID                  string `json:"measurement_uid"`
	URL                  string `json:"measurement_url"`
	ProbeASN             string `json:"probe_asn"`
	ProbeCC              string `json:"probe_cc"`
	ReportID             string `json:"report_id"`
	TestName             string `json:"test_name"`
}

type RawMeasurements struct {
	Results []MeasurementMeta `json:"results"`
}

func QueryMeasurements() ([]byte, error) {
	//TODO: Add date and other parameters from config file
	apiEndpoint := "https://api.ooni.io/api/v1/measurements?test_name=web_connectivity&failure=false"

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

func DecodeMeasurements(body []byte) (RawMeasurements, error) {
	var rawMeasurements RawMeasurements

	err := json.Unmarshal(body, &rawMeasurements)

	if err != nil {
		return RawMeasurements{}, fmt.Errorf("Error parsing the JSON result: %v", err)
	}

	return rawMeasurements, nil
}

func WriteToFile(rawMeasurements RawMeasurements) error {
	file, err := os.Create("raw_measurements.txt")
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

	fmt.Println("URLs have been written to measurements.txt")
	return nil
}
