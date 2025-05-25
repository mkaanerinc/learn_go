package main

import "fmt"

func main() {
	fact := factorial(3)

	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

/*

factorial(3)
= 3 * factorial(2)
= 3 * (2 * factorial(1))
= 3 * (2 * (1 * factorial(0)))
= 3 * (2 * (1 * 1))
= 3 * (2 * 1)
= 3 * 2
= 6

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

*/
