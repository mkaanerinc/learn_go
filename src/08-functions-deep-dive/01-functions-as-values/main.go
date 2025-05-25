package main

import "fmt"

// Define a function type that takes an int and returns an int.
// This allows us to pass any function with this signature as an argument.
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
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

// This function matches the transformFn type.
func double(number int) int {
	return number * 2
}

// This also matches the transformFn type.
func triple(number int) int {
	return number * 3
}
