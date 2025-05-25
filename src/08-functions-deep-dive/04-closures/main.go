package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// createTransformer returns a function that multiplies by the given factor.
	// double and triple are closures that remember the `factor` value they were created with.
	double := createTransformer(2) // remembers factor = 2
	triple := createTransformer(3) // remembers factor = 3

	// Pass an anonymous function directly as the transform argument.
	// This function doubles the number.
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double) // uses the closure that multiplies by 2
	tripled := transformNumbers(&numbers, triple) // uses the closure that multiplies by 3

	fmt.Println(transformed) // [2 4 6]
	fmt.Println(doubled)     // [2 4 6]
	fmt.Println(tripled)     // [3 6 9]
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

// createTransformer returns a function (closure) that multiplies any number by the captured `factor`
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		// This function *closes over* the `factor` variable — it retains access to it even after createTransformer ends
		return number * factor
	}
}
