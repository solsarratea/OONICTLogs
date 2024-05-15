package ooniapi

import (
	"fmt"
	"io"
	"net/http"
)

func QuerySingleMeasurement(string apiEndpoint) ([]byte, error) {
	response, err := http.Get(apiEndpoint)
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
