package measurements

import (
	"encoding/json"
	"fmt"
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

func DecodeMeasurements(body []byte) (RawMeasurements, error) {
	var rawMeasurements RawMeasurements

	err := json.Unmarshal(body, &rawMeasurements)

	if err != nil {
		return RawMeasurements{}, fmt.Errorf("Error parsing the JSON result: %v", err)
	}

	return rawMeasurements, nil
}
