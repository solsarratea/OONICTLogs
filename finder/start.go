package finder

import (
	"fmt"
	"time"

	"github.com/solsarratea/OONICTLogs/common"
	"github.com/solsarratea/OONICTLogs/finder/utils"
)

//  Start orchastrates the:
//  - Loading in memory root certificates from CTLog specified in config file
//  - Updating on the `raw_measurements.txt`, which acts as a stack of measurements
//  - Calls for processing individual measurement

func Start() {

	fmt.Println("Starting Chainfinder...")
	config, err := common.ReadConfigurationFile()

	// updateSince set the lower boundary for raw measurement queries
	updateSince := config.OONIMeasurements.Since

	if err != nil {
		fmt.Println("OOPs, something went wrong... Invalid configuration file!")
		return
	}

	roots, err := LoadRoots(config)

	for {

		isEmpty, _ := utils.IsFileEmpty(config.PathMeasurements)

		if isEmpty {
			UpdateConfigFile(config, updateSince)
			GetRawMeasurements(config)

		} else {
			// updateSince is updated with the measurement's start time
			updateSince, err = ProcessMeasurement(config, roots)

			if err != nil {
				fmt.Printf("%v", err)
			}

			Flush(config)
		}
	}

	time.Sleep(20 * time.Second)
}

// Empty entry from raw_measurements.txt
func Flush(config common.Configuration) {
	utils.RemoveLineFromFile(config.PathMeasurements)
}

// Dynamic update "config.json" for making sure raw measurements window query is being updated
func UpdateConfigFile(config common.Configuration, updateSince string) {
	updatedConfig := config
	updatedConfig.OONIMeasurements.Since = updateSince
	utils.WriteStructToJSONFile(updatedConfig, "config.json")

}
