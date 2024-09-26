package submitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/solsarratea/OONICTLogs/common"
	"github.com/solsarratea/OONICTLogs/finder/utils"
)

// SubmissionResponse represents  structure of the response from API call POST chain log
// Specified in RFC 6962 Section 4.1
type SubmissionResponse struct {
	SCTVerion string `json:"sct_verion"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	Id        string `json:"id"`
}

// PostChain implements an API call for posting certificate chain
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ValidSubmission struct {
	Certificates []string `json:"chain"`
}

// ProcessSubmissions implements the reading and posting of valid certificate chain
func ProcessSubmissions(config common.Configuration, logChannel chan<- string) {
	dir := config.PathCert
	files, err := os.ReadDir(dir)

	if err != nil {
		msg := fmt.Sprintf("Failed to read directory: %v", err)
		logChannel <- msg
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())

			measurementUID := file.Name()

			msg := fmt.Sprintf("Submitting entry for: %s\n", measurementUID)
			logChannel <- msg

			data, err := os.ReadFile(filePath)
			if err != nil {
				msg := fmt.Sprintf("Failed to read file %s: %v", filePath, err)
				logChannel <- msg
				return
			}

			//Check if the format is still valid
			var submission ValidSubmission
			if err := json.Unmarshal(data, &submission); err != nil {
				msg := fmt.Sprintf("Failed to unmarshal JSON from file %s: %v", filePath, err)
				logChannel <- msg

				// Remove the file after reading
				err = os.Remove(filePath)
				if err != nil {
					msg := fmt.Sprintf("Failed to remove file %s: %v", filePath, err)
					logChannel <- msg
					continue
				}

				return
			}

			// Make HTTP POST request
			body, err := PostChain(config.CTLog.URI, data)
			if err != nil {
				msg := fmt.Sprintf("Failed to submit data for %s: %v", measurementUID, err)
				logChannel <- msg
				return
			}
			// Remove the file after reading
			err = os.Remove(filePath)
			if err != nil {
				msg := fmt.Sprintf("Failed to remove file %s: %v", filePath, err)
				logChannel <- msg
				continue
			}

			var subresp SubmissionResponse

			err = json.Unmarshal(body, &subresp)
			surl := "submission-response-" + measurementUID + ".txt"
			utils.WriteStringToFile(string(body), surl)

			msg = fmt.Sprintf("Response has been written at:  %s \n", surl)
			logChannel <- msg

			if err != nil {
				msg := fmt.Sprintf("Error parsing response: %v", err)
				logChannel <- msg
				return
			}

			msg = fmt.Sprintf("Success entry for: %s", measurementUID)
			logChannel <- msg
			time.Sleep(500 * time.Second)
		}
	}

}

// Start orchastrates the submission of valid entries shared directory
func Start(logChannel chan<- string) {
	logChannel <- "Starting Submitter"

	config, err := common.ReadConfigurationFile()

	if err != nil {
		logChannel <- "OOPs, something went wrong... Invalid configuration file!"
		return
	}

	for {
		ProcessSubmissions(config, logChannel)
		time.Sleep(500 * time.Second)
	}
}
