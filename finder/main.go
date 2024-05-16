package finder

import (
	"fmt"
	"time"

	"./measurements"
	"./queries"
	"./utils"
)

func GetRawMeasurements() {
	body, err := queries.QueryMeasurements()
	if err != nil {
		fmt.Printf("Failed to query measurements: %v\n", err)

	} else {

		rawMeasurements, err := measurements.DecodeMeasurements(body)

		if err != nil {
			fmt.Printf("Failed to decode raw measurements %v\n", err)
		} else {
			utils.WriteMeasurementsToFile(rawMeasurements)
		}
	}
}

func Flush(config Configuration) {
	utils.RemoveLineFromFile(config.PathMeasurements)
}

func Start() {

	fmt.Println("Starting Chainfinder...")

	config, err := ReadConfigurationFile()

	if err != nil {
		return
	}

	roots, err := LoadRootNodes(config)

	for {

		isEmpty, _ := utils.IsFileEmpty(config.PathMeasurements)

		if isEmpty {
			GetRawMeasurements() //TOOD: Update config

		} else {
			fmt.Println("Processing measurements...")
			ProcessMeasurements(config) //TODO: update config
			Flush(config)
		}
	}

	time.Sleep(10 * time.Second)
}
