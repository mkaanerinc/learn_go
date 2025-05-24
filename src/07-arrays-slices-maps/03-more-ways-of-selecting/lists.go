package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[:3] // Slicing the array to get a sub-array
	// This creates a new slice that includes elements from index 0 to 2 (exclusive of 3).
	fmt.Println(featuredPrices)

	featuredPrices2 := prices[1:]            // This creates a new slice that includes elements from index 1 to the end of the array.
	highlightedPrices := featuredPrices2[:1] // This creates a new slice that includes only the first element of featuredPrices2.
	fmt.Println(highlightedPrices)
}

/*
📌 Arrays vs Slices in Go:

- `Array`: Fixed size, defined as `[N]Type`.
- `Slice`: Dynamic view over an array, defined as `[]Type`.

✅ You can create a slice from an array:
  - `arr[:n]` → from start to index n (exclusive)
  - `arr[m:]` → from index m to end
  - `arr[m:n]` → from index m to n (exclusive)

💡 Important:
- Slices are references to the underlying array.
- Changes to slice elements affect the original array.
*/
