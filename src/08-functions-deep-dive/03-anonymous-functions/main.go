package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// Pass an anonymous function directly as the transform argument.
	// This function doubles the number.
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)
}

// transformNumbers takes a slice of numbers and applies the transform function to each element.
// The transform function is passed as an argument (can be named or anonymous).
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
