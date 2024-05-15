package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLineFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		return scanner.Text(), nil
	} else if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read the file: %v", err)
	}

	return "", fmt.Errorf("File is empty")
}
