package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"../measurements"
)

func WriteMeasurementsToFile(rawMeasurements measurements.RawMeasurements) error {
	file, err := os.Create("raw_measurements.txt")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}

	defer file.Close()

	for _, result := range rawMeasurements.Results {
		_, err := file.WriteString(result.URL + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}

	fmt.Println("URLs have been written to measurements.txt")
	return nil
}

func WriteStringToFile(content, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

func WriteStructToJSONFile(data interface{}, filename string) error {
	// Convert the struct to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
