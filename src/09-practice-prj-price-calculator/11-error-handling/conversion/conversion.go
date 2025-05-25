package conversion

import (
	"errors"
	"strconv"
)

// StringsToFloats converts a slice of strings to a slice of float64 values.
// Returns an error if any of the strings fail to convert.
func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64) // Convert each line to a float64

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		// Append successfully parsed float to the result slice
		floats = append(floats, floatVal)
	}

	// Return the fully parsed slice with no error
	return floats, nil
}
