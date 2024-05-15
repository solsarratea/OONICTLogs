package measurements

//TODO: Move to a common package for handling files
import (
	"fmt"
	"os"
)

func WriteToFile(rawMeasurements RawMeasurements) error {
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
