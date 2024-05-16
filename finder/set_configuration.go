package finder

import (
	"encoding/json"
	"fmt"
	"os"
)

type CTLog struct {
	PublicKey string `json:"publicKey"`
	LogID     string `json:"logID"`
	URI       string `json:"URI"`
	Start     string `json:"start"`
	End       string `json:"end"` //parse it with RFC3339
}

type Configuration struct {
	PathMeasurements string `json:"path-measurements"`
	PathCert         string `json:"path-cert"`
	CTLog            CTLog  `json:"ctlog"`
}

func ReadConfigurationFile() (Configuration, error) {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Failed opening `config.json`:", err)
		return Configuration{}, err
	}

	defer file.Close()

	var config Configuration
	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		fmt.Println("Failed parsing `config.json`:", err)
		return Configuration{}, err
	}

	return config, err
}
