package finder

import (
	"fmt"
	"os"
	"time"

	"./measurements"
)

func IsFileEmpty(filename string) (bool, error) {
	// Get file information
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, create it with the specified name
			_, err := os.Create(filename)
			if err != nil {
				return false, fmt.Errorf("error creating file: %v", err)
			}
			return true, nil
		}

		return false, err
	}

	// Check if file size is zero
	if info.Size() == 0 {
		return true, nil // File is empty
	}

	return false, nil // File is not empty
}

func Start() {
	fmt.Println("Starting Chainfinder")

	for {

		isEmpty, _ := IsFileEmpty("raw_measurements.txt")

		if isEmpty {
			fmt.Println("File is empty")
			fmt.Println("Querying raw measurements...")
			body, err := measurements.QueryMeasurements()

			if err != nil {
				fmt.Printf("Failed to query measurements: %v\n", err)
			}

			fmt.Println("Parsing JSON...")
			rawMeasurements, err := measurements.DecodeMeasurements(body)

			measurements.WriteToFile(rawMeasurements)

		} else {
			fmt.Println("Raw measurements should be proccessed...")

		}

		time.Sleep(5 * time.Second)
	}
}
