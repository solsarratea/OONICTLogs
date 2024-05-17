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

func UpdateConfigFile(config common.Configuration, updateSince string) {
	updatedConfig := config
	updatedConfig.OONIMeasurements.Since = updateSince
	utils.WriteStructToJSONFile(updatedConfig, "config.json")

}

func Start() {

	fmt.Println("Starting Chainfinder...")
	config, err := common.ReadConfigurationFile()
	updateSince := config.OONIMeasurements.Since

	if err != nil {
		fmt.Println("OOPs, something went wrong... Invalid configuration file!")
		return
	}

	roots, err := GetRootNodes(config)

	for {

		isEmpty, _ := utils.IsFileEmpty(config.PathMeasurements)

		if isEmpty {
			UpdateConfigFile(config, updateSince)
			GetRawMeasurements(config)

		} else {
			updateSince, err = ProcessMeasurement(config, roots)
			if err != nil {
				fmt.Printf("%v", err)
			}
			Flush(config)
		}
	}

	time.Sleep(20 * time.Second)
}
