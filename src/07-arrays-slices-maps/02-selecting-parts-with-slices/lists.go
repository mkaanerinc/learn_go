package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[1:3] // Slicing the array to get a sub-array
	// This creates a new slice that includes elements from index 1 to 2 (exclusive of 3).
	fmt.Println(featuredPrices)
	// Slicing allows you to create a new slice from an existing array or slice.
}
