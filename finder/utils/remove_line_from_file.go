package utils

import (
	"bufio"
	"fmt"
	"os"
)

func RemoveLineFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Skip the first line
	var remainingLines []string
	skipFirstLine := true
	for scanner.Scan() {
		if skipFirstLine {
			skipFirstLine = false
			continue
		}
		remainingLines = append(remainingLines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read the file: %v", err)
	}

	// Write the remaining lines back to the file
	file, err = os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range remainingLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %v", err)
	}

	return nil
}
