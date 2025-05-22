package main

import (
	"fmt"
	"math"
)

func main() {
	// var investmentAmount, years float64 = 1000, 10
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0
	//var years float64 = 10

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Println(futureValue)
}

/*
Variables in Go:

Go supports both explicit and implicit variable declarations.

1. Explicit Declaration:
	Example: var investmentAmount float64 = 1000
	- 'var' is the keyword used to declare a variable.
	- 'investmentAmount' is the variable name.
	- 'float64' is the type (64-bit floating-point number).
	- '= 1000' assigns the initial value.
	This form is useful when you want to clearly state the variable type.

2. Short Declaration (Type Inference):
	Example: expectedReturnRate := 5.5
	- ':=' declares and initializes the variable in one step.
	- Go automatically infers the type from the value.
	- This can only be used inside functions, not at package level.

3. Multiple Variable Declaration:
	Example: investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	- You can declare and assign multiple variables in a single line.
	- All variables must be assigned values of appropriate types.

Notes:
	- Go is statically typed, so every variable has a fixed type determined at compile time.
	- Unused variables cause a compile-time error. Every declared variable must be used.
	- Use explicit declarations when clarity is needed; use short declarations for concise code inside functions.
*/
