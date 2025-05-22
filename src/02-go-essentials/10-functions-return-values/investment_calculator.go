package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {

	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}

/*
Parameterized Function with Return Values and Global Constant Usage in Go:

- The function `calculateFutureValues` demonstrates how to define and use a function that:
  - Takes parameters (inputs)
  - Returns multiple values (outputs)

Signature:
	func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64)

Explanation:
	- `investmentAmount`, `expectedReturnRate`, and `years` are input parameters of type `float64`.
	- The function returns two `float64` values:
	1. `fv` - future value of the investment
	2. `frv` - inflation-adjusted future value

Multiple Return Values:
	- Go allows returning multiple values from a function, which is one of its standout features.
	- Useful for returning both result and additional context (like an error or secondary result).

Example Call:
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

Global Constant:
	- The constant `inflationRate` is defined outside the main function (at the package level).
	- This makes it a global constant, accessible by all functions within the same package.
	- We made it global because `calculateFutureValues` needs to access it, and it is not passed as a parameter.
*/
