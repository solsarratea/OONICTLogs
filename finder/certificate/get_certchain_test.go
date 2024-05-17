package certificate

import (
	"io/ioutil"
	"testing"
)

func TestGetCChainFromEmpty(t *testing.T) {
	jsonData := []byte(`{"result":[]}`)
	measurement, _ := DecodeMeasurement(jsonData)
	output, _ := GetCertificateChain(measurement)

	if len(output) != 0 {
		t.Errorf("Expected rawMeasurements.Results length to be 0, got %d", len(output))
	}
}

func TestGetCChainFromValidInput(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("sample_single_measurement.json")
	measurement, _ := DecodeMeasurement(jsonData)
	output, err := GetCertificateChain(measurement)

	if err != nil {
		t.Errorf("Expected error %s", err)
	}

	if len(output) == 0 {
		t.Errorf("No results shown up")
	}

	for _, cchain := range output {
		if len(cchain) < 1 {
			t.Errorf("Empty chain")
		}
	}
}
