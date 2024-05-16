package certificate

import (
	"io/ioutil"
	"testing"
)

func TestGetCChainFromEmpty(t *testing.T) {
	jsonData := []byte(`{"result":[]}`)

	output, _ := GetCertificateChain(jsonData)

	if len(output) != 0 {
		t.Errorf("Expected rawMeasurements.Results length to be 0, got %d", len(output))
	}
}

func TestGetCChainFromValidInput(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("sample_single_measurement.json")

	output, err := GetCertificateChain(jsonData)

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
