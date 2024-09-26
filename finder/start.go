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

func Start(logChannel chan<- string) {
	msg := "Starting Chainfinder..."
	logChannel <- msg

	config, err := common.ReadConfigurationFile()

	// updateSince set the lower boundary for raw measurement queries
	updateSince := config.OONIMeasurements.Since

	if err != nil {
		msg := fmt.Sprintf("OOPs, something went wrong... Invalid configuration file: %v ", err)
		logChannel <- msg
		return
	}

	roots, err := LoadRoots(config)
	if err != nil {
		logChannel <- fmt.Sprintf("Error loading roots %v", err)
	}

	for {

		isEmpty, _ := utils.IsFileEmpty(config.PathMeasurements)

		if isEmpty {
			UpdateConfigFile(config, updateSince)
			GetRawMeasurements(config)

		} else {
			// updateSince is updated with the measurement's start time
			updateSince, err = ProcessMeasurement(config, roots, logChannel)

			if err != nil {
				msg := fmt.Sprintf("%v", err)
				logChannel <- msg
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
