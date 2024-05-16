package finder

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"./certchain"
	"./measurements"
	"./queries"
	"./utils"
)

type CTLog struct {
	PublicKey string `json:"publicKey"`
	LogID     string `json:"logID"`
	URI       string `json:"URI"`
	Start     string `json:"start"`
	End       string `json:"end"` //parse it with RFC3339
}

type Configuration struct {
	PathMeasurements string `json:"path-measurements"`
	PathCert         string `json:"path-cert"`
	CTLog            CTLog  `json:"ctlog"`
}

func ReadConfigurationFile() (Configuration, error) {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Failed opening `config.json`:", err)
		return Configuration{}, err
	}

	defer file.Close()

	var config Configuration
	err = json.NewDecoder(file).Decode(&config)

	if err != nil {
		fmt.Println("Failed parsing `config.json`:", err)
		return Configuration{}, err
	}

	return config, err
}

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

func ProcessMeasurements(config Configuration) {
	//TODO: Add propper logs of success/lost measurements
	apiEndpoint, _ := utils.ReadLineFromFile(config.PathMeasurements)

	re := regexp.MustCompile(`measurement_uid=([^&]+)`)
	measurement_uid := re.FindStringSubmatch(apiEndpoint)[1]

	fmt.Println("Processing measure with id: " + measurement_uid)
	body, err := queries.QuerySingleMeasurement(apiEndpoint)

	if err != nil {
		fmt.Printf("Failed to query measurements: %v\n", err)

	} else {

		cchain, _ := certchain.GetCertificateChain(body)

		for i, subchain := range cchain {
			content := strings.Join(subchain, "\n")
			path := config.PathCert + measurement_uid + "-chain-" + strconv.Itoa(i)
			utils.WriteStringToFile(content, path)
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
