package submitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"../common"
	"../finder/utils"
)

type SubmissionResponse struct {
	Timestamp string `json:"timestamp"`
	Signature string `json: "signature"`
}

func PostChain(URI string, data []byte) ([]byte, error) {
	apiEndpoint := URI + "ct/v1/add-chain"
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "lets-encrypt-ct-log-example-1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func ProcessSubmissions(config common.Configuration) {
	dir := config.PathCert
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())

			measurementUID := file.Name()

			fmt.Printf("Submitting entry for: %s\n", measurementUID)

			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", filePath, err)
				continue
			}

			// Make HTTP POST request
			body, err := PostChain(config.CTLog.URI, data)
			if err != nil {
				log.Printf("Failed to submit data for %s: %v", measurementUID, err)
				continue
			}
			// Remove the file after reading
			err = os.Remove(filePath)
			if err != nil {
				log.Printf("Failed to remove file %s: %v", filePath, err)
				continue
			}

			var subresp SubmissionResponse

			err = json.Unmarshal(body, &subresp)
			utils.WriteStringToFile(string(body), "logs.txt")

			if err != nil {
				fmt.Printf("Error parsing response: %v", err)
				return
			}

			fmt.Printf("Success entry for: ", measurementUID)
			time.Sleep(2500 * time.Second)
		}
	}

}

func Start() {
	fmt.Println("Starting Submitter")

	config, err := common.ReadConfigurationFile()

	if err != nil {
		fmt.Println("OOPs, something went wrong... Invalid configuration file!")
		return
	}

	for {
		fmt.Println("Submitter running...")
		ProcessSubmissions(config)
		time.Sleep(500 * time.Second)
	}
}
