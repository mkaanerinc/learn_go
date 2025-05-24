package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(productNames)
	fmt.Println(prices[2])

	// Create a slice starting from index 1 to the end
	featuredPrices := prices[1:] // → [9.99 45.99 20]

	// Modify the first element of the slice
	featuredPrices[0] = 199.99              // this modifies prices[1] as well (because slice points to the same array)
	highlightedPrices := featuredPrices[:1] // slice with only the first element of featuredPrices

	fmt.Println(highlightedPrices) // → [199.99]
	fmt.Println(prices)            // → [10.99 199.99 45.99 20] → underlying array modified!

	// Print length and capacity of slices
	fmt.Println(len(featuredPrices), cap(featuredPrices))       // len=3 (1:), cap=3 (from index 1 to end of array)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // len=1, cap=3 (still based on same underlying array)

	// Extend highlightedPrices to include more elements within its capacity
	// In Go, a slice has a length (number of elements it currently references) and a capacity (number of elements
	// in the underlying array it can potentially reference, starting from its starting index).
	// Here, highlightedPrices was created as featuredPrices[:1], so it starts at index 1 of the underlying array
	// (prices) and currently has length 1 ([199.99]). However, its capacity is 3 because the underlying array
	// (prices) has 3 elements from index 1 to the end ([199.99, 45.99, 20]).
	// By re-slicing with highlightedPrices[:3], we extend the slice to include all elements up to its capacity
	// (from index 1 to index 3 of the underlying array), increasing its length from 1 to 3.
	// This does not create a new array; it just adjusts the slice's length to include more of the existing underlying
	// array's elements, so highlightedPrices now references [199.99, 45.99, 20].
	highlightedPrices = highlightedPrices[:3]

	// After re-slicing, len(highlightedPrices) is now 3 because it includes 3 elements.
	// The capacity remains 3 because it is still limited by the underlying array's size from index 1 to the end.// Now includes [199.99, 45.99, 20]
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // len=3, cap=3

	// Print the extended slice, which now includes all elements within its capacity.
	fmt.Println(highlightedPrices) // [199.99 45.99 20]
}
