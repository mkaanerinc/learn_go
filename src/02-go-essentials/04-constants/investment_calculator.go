package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

/*
Constants in Go:

- The `const` keyword is used to declare constant values that cannot be changed after compilation.
- Example: const inflationRate = 2.5
  - `inflationRate` is a constant with a fixed value of 2.5.
  - It must be assigned at the time of declaration.
  - The type is inferred unless explicitly stated (you can optionally write: const inflationRate float64 = 2.5).

Why use constants?
	- Constants improve code clarity by signaling that a value is fixed and should not change.
	- They help prevent accidental modifications to values that are meant to stay constant throughout the program.
	- Useful for fixed configuration values like tax rates, inflation rates, or mathematical constants (e.g., Pi).

Other notes:
	- Constants in Go can be typed or untyped.
	- You cannot use := with constants.
	- Constants are evaluated at compile time and do not take up memory at runtime like variables.
*/
