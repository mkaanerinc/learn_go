package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
}

// Why are maps used in Go?
// Maps are used for quickly looking up, adding, updating, or deleting key-value pairs.
// They provide fast access with an average O(1) time complexity, making them efficient for performance.
// For example, they are ideal for mapping user IDs to names in a database, storing configuration settings, or creating caches.
// Maps are unordered and are a reference type, so changes made in functions affect the original map.
// Note: Care must be taken with concurrent access, as maps are not thread-safe by default.
