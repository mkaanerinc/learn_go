package main

import "fmt"

// Define a function type that takes an int and returns an int.
// This allows us to pass any function with this signature as an argument.
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled) // [2 4 6 8]
	fmt.Println(tripled) // [3 6 9 12]

	// Get dynamic transformation functions based on slice content
	transformerFn1 := getTransformerFunction(&numbers)     // returns double (since numbers[0] == 1)
	transformerFn2 := getTransformerFunction(&moreNumbers) // returns triple (since moreNumbers[0] == 5)

	// Apply the returned transformation functions
	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	transformedMoreNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)     // [2 4 6 8]
	fmt.Println(transformedMoreNumbers) // [15 3 6]
}

// transformNumbers takes a pointer to a slice of ints and a transform function.
// It applies the transform function to each element of the slice and returns a new slice.
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

// getTransformerFunction returns a function based on a condition.
// This demonstrates returning a function value.
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// This function matches the transformFn type.
func double(number int) int {
	return number * 2
}

// This also matches the transformFn type.
func triple(number int) int {
	return number * 3
}
