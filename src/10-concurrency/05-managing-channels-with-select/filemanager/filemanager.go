package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath) // Attempt to open the file containing raw prices

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

// WriteJSON creates or overwrites a file at the given path and writes the provided data as JSON.
// Returns an error if file creation or encoding fails.
func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath) // Attempt to create (or truncate) the file at the specified path

	if err != nil {
		return errors.New("failed to create file")
	}

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file) // Create a new JSON encoder bound to the file
	err = encoder.Encode(data)       // Encode the data and write it to the file

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
