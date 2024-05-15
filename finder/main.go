package finder

import (
	"fmt"
	"time"

	"./measurements"
	"./ooniapi"
	"./utils"
)

func Start() {
	fmt.Println("Starting Chainfinder")

	for {

		isEmpty, _ := utils.IsFileEmpty("raw_measurements.txt")

		if isEmpty {
			fmt.Println("File is empty")
			fmt.Println("Querying raw measurements...")

			body, err := ooniapi.QueryMeasurements()
			if err != nil {
				fmt.Printf("Failed to query measurements: %v\n", err)
			} else {

				fmt.Println("Parsing JSON...")
				rawMeasurements, err := measurements.DecodeMeasurements(body)

				if err != nil {
					fmt.Printf("Failed to decode measurements %s", err)
				} else {
					utils.WriteToFile(rawMeasurements)
				}

			}

		} else {
			fmt.Println("Raw measurements should be proccessed...")

		}

		time.Sleep(5 * time.Second)
	}
}
