package roots

import (
	"io/ioutil"
	"testing"
)

func TestParseRootCertificatesFromEmpty(t *testing.T) {
	jsonData := []byte(`{"result":[]}`)

	output, _ := ParseRootCertificates(jsonData)

	if len(output) != 0 {
		t.Errorf("Expeted length result to be 0, got %d", len(output))
	}
}

func TestParseRootCertificatesFromValidInput(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("sample_roots.json")

	_, err := ParseRootCertificates(jsonData)

	if err != nil {
		t.Errorf("Invalid parsing of json file error %s", err)
	}

}
