package measurements

import (
	"fmt"
	"io"
	"net/http"
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
