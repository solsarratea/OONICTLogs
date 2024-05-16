package finder

import (
	"fmt"
	"time"

	"./utils"
)

func Flush(config Configuration) {
	utils.RemoveLineFromFile(config.PathMeasurements)
}

func Start() {

	fmt.Println("Starting Chainfinder...")

	config, err := ReadConfigurationFile()

	if err != nil {
		return
	}

	roots, err := GetRootNodes(config)

	for {

		isEmpty, _ := utils.IsFileEmpty(config.PathMeasurements)

		if isEmpty {
			GetRawMeasurements(config)

		} else {
			fmt.Println("Processing measurements...")
			ProcessMeasurement(config, roots) //TODO: Update config
			Flush(config)
		}
	}

	time.Sleep(10 * time.Second)
}
