package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path) // Attempt to open the file containing raw prices

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file) // Scanner to read file line by line

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Collect each line into the lines slice
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}
