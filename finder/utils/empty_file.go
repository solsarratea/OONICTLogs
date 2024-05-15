package utils

import (
	"fmt"
	"os"
)

func IsFileEmpty(filename string) (bool, error) {
	// Get file information
	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, create it with the specified name
			_, err := os.Create(filename)
			if err != nil {
				return false, fmt.Errorf("error creating file: %v", err)
			}
			return true, nil
		}

		return false, err
	}

	// Check if file size is zero
	if info.Size() == 0 {
		return true, nil // File is empty
	}

	return false, nil // File is not empty
}
