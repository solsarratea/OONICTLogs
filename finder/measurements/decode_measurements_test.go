package measurements

import (
	"io/ioutil"
	"testing"
)

func TestDecodeMeasurementsEmptyInput(t *testing.T) {
	jsonData := []byte(`{"result":[]}`)

	rawMeasurements, _ := DecodeMeasurements(jsonData)

	if len(rawMeasurements.Results) != 0 {
		t.Errorf("Expected rawMeasurements.Results length to be 0, got %d", len(rawMeasurements.Results))
	}
}

func TestDecodeMeasurementsValidInput(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("sample_raw_measurements.json")

	rawMeasurements, _ := DecodeMeasurements(jsonData)

	for _, result := range rawMeasurements.Results {
		if result.URL == "" {
			t.Errorf("Unexpected URL decoding for: %v", result.UID)

		}
	}
}

func TestDecodeMeasurementsInvalidInput(t *testing.T) {
	jsonData := []byte(`{"result":["InvalidField": "2"]}`)

	_, err := DecodeMeasurements(jsonData)

	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
