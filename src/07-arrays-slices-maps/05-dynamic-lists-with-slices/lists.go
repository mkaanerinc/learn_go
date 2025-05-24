package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)

	fmt.Println(prices)
}

// Explanation of dynamic arrays (slices) in Go:
// In Go, dynamic arrays are referred to as "slices." Slices are a flexible and dynamic data structure built on top of fixed-length arrays.
// A slice consists of three main components:
// 1. **Pointer**: Points to the starting address of the underlying array.
// 2. **Length (len)**: The number of elements currently in the slice.
// 3. **Capacity (cap)**: The total number of elements in the underlying array, starting from the slice's beginning.

// Slices can grow dynamically using the append() function. If the underlying array's capacity is insufficient,
// append() creates a new array and copies the existing elements to it, allowing the slice to expand.

// In the code above:
// - `prices := []float64{10.99, 8.99}`: A slice is declared and initialized with two elements.
// - `prices[0:1]`: Prints a subslice from index 0 to 1 (just 10.99).
// - `prices[1] = 9.99`: Updates the element at index 1 of the slice.
// - `prices = append(prices, 5.99)`: Appends a new element to the slice. If the capacity is insufficient, Go automatically creates a new underlying array.

// Advantages of slices:
// - Dynamic sizing: Can grow easily with append.
// - Flexibility: Supports slicing operations to create subsets.
// - Performance: Efficient memory usage via the underlying array.

// Important note: Slices are reference types. When a slice is passed to a function or assigned to another slice,
// changes to the underlying array affect all references.
