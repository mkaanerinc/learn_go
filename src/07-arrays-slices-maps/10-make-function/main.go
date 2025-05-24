package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Kaan")
	userNames = append(userNames, "Mustafa")

	fmt.Println(userNames)
}

// The make function is used in Go to initialize slices, maps, and channels.
// It allocates memory and sets up the data structure with initial values.
// Syntax: make(type, length, [capacity])
// - type: The type of the data structure (e.g., []string for a string slice).
// - length: The initial number of elements in the slice or channel.
// - capacity: (Optional) The maximum size of the underlying array for a slice.
// Example: make([]string, 2, 5) creates a string slice.
// - Length 2: The slice has 2 elements (default to empty strings: ["", ""]).
// - Capacity 5: The slice can grow up to 5 elements without reallocation.
