package main

import "fmt"

func main() {
	result := add(5, 10)
	fmt.Println(result)
}

// add is a generic function that can add two values of the same type.
// It uses type parameters to allow for different types like int, float64, or string.
func add[T int | float64 | string](a, b T) T {
	return a + b
}

/*
Generics in Go:

- Allow functions and types to be written in a way that works with different data types without code duplication.
- The syntax `func add[T int | float64 | string](...)` defines a generic function with a type constraint.
- Constraints (like `int | float64 | string`) define what operations (like +) are valid on T.
- Useful for collections, algorithms, and utilities that are type-agnostic.
*/
