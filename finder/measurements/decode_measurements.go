package measurements

import (
	"encoding/json"
	"fmt"
)

func DecodeMeasurements(body []byte) (RawMeasurements, error) {
	var rawMeasurements RawMeasurements

	err := json.Unmarshal(body, &rawMeasurements)

	if err != nil {
		return RawMeasurements{}, fmt.Errorf("Error parsing the JSON result: %v", err)
	}

	return rawMeasurements, nil
}
