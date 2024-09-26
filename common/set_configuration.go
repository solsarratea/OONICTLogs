package common

// Implements the interface for parsing the `config.json` file at root directory

import (
	"encoding/json"
	"os"
)

type CTLog struct {
	PublicKey string `json:"PublicKey"`
	LogID     string `json:"LogID"`
	URI       string `json:"URI"`
	Start     string `json:"Start"`
	End       string `json:"End"` //parse it with RFC3339
}

type OONIMeasurements struct {
	Since string `json:"Since"`
	Until string `json:"Until"`
	Limit string `json:"Limit"`
}

type Configuration struct {
	PathMeasurements string           `json:"PathMeasurements"`
	PathCert         string           `json:"PathCert"`
	CTLog            CTLog            `json:"CTLog"`
	OONIMeasurements OONIMeasurements `json:"OONImeasurements"`
}

func ReadConfigurationFile() (Configuration, error) {

	file, err := os.Open("config.json")
	if err != nil {
		return Configuration{}, err
	}

	defer file.Close()

	var config Configuration
	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		return Configuration{}, err
	}

	return config, err
}
