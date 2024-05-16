package ooniapi

import (
	"fmt"
	"io"
	"net/http"
)

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
