package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)

	// If we have a slice (numbers), we cannot pass it directly to a variadic function.
	// Therefore, we use the ... operator to unpack the slice → numbers... becomes 1, 10, 15 when passed.
	anotherSum := sumup(numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// sumup is a variadic function that accepts any number of int arguments.
// "numbers ...int" means it accepts zero or more int values as input.
func sumup(numbers ...int) int {
	sum := 0

	// The variadic parameter "numbers" behaves like a slice.
	for _, val := range numbers {
		sum += val
	}

	return sum
}
