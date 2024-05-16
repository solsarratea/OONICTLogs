package finder

import (
	"fmt"
	"time"

	"../common"
	"./utils"
)

func Flush(config common.Configuration) {
	utils.RemoveLineFromFile(config.PathMeasurements)
}

func Start() {

	fmt.Println("Starting Chainfinder...")
	config, err := common.ReadConfigurationFile()

	if err != nil {
		fmt.Println("OOPs, something went wrong... Invalid configuration file!")
		return
	}

	roots, err := GetRootNodes(config)

	for {

		isEmpty, _ := utils.IsFileEmpty(config.PathMeasurements)

		if isEmpty {
			GetRawMeasurements(config) //TODO: Update config

		} else {
			fmt.Println("Processing measurements...")
			err = ProcessMeasurement(config, roots)
			if err != nil {
				fmt.Printf("%v", err)
			}
			Flush(config)
		}
	}

	time.Sleep(20 * time.Second)
}
