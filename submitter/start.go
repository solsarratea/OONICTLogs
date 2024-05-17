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

// SubmissionResponse represents  structure of the response from API call POST chain log
// Specified in RFC 6962 Section 4.1
type SubmissionResponse struct {
	SCTVerion string `json:"sct_verion"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	Id        string `json:"id"`
}

//PostChain implements an API call for posting certificate chain
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

type ValidSubmission struct {
	Certificates []string `json:"chain"`
}

// ProcessSubmissions implements the reading and posting of valid certificate chain
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
				return
			}

			//Check if the format is still valid
			var submission ValidSubmission
			if err := json.Unmarshal(data, &submission); err != nil {
				log.Printf("Failed to unmarshal JSON from file %s: %v", filePath, err)

				// Remove the file after reading
				err = os.Remove(filePath)
				if err != nil {
					log.Printf("Failed to remove file %s: %v", filePath, err)
					continue
				}

				return
			}

			// Make HTTP POST request
			body, err := PostChain(config.CTLog.URI, data)
			if err != nil {
				log.Printf("Failed to submit data for %s: %v", measurementUID, err)
				return
			}
			// Remove the file after reading
			err = os.Remove(filePath)
			if err != nil {
				log.Printf("Failed to remove file %s: %v", filePath, err)
				continue
			}

			var subresp SubmissionResponse

			err = json.Unmarshal(body, &subresp)
			surl := "submission-response-" + measurementUID + ".txt"
			utils.WriteStringToFile(string(body), surl)

			fmt.Printf("Response has been written at:  %s \n", surl)

			if err != nil {
				fmt.Printf("Error parsing response: %v", err)
				return
			}

			fmt.Printf("Success entry for: ", measurementUID)
			time.Sleep(500 * time.Second)
		}
	}

}

// Start orchastrates the submission of valid entries shared directory
func Start() {
	fmt.Println("Starting Submitter")

	config, err := common.ReadConfigurationFile()

	if err != nil {
		fmt.Println("OOPs, something went wrong... Invalid configuration file!")
		return
	}

	for {
		ProcessSubmissions(config)
		time.Sleep(500 * time.Second)
	}
}
