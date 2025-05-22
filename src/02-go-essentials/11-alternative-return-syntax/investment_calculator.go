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

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {

	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
	// return
}

/*
Explanation of Named Return Values and Assignment inside the Function:

- In this function signature:
  func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64)

  `fv` and `frv` are named return variables of type float64.

- Named return variables mean:
  - These variables are declared implicitly at the start of the function.
  - You can assign values to them directly using `=` inside the function body.
  - When you use a bare `return` (without variables), these named variables are returned automatically.

- Why no `:=` inside the function?
  - `:=` is used for **declaring and initializing** a new variable.
  - Here, since `fv` and `frv` are already declared in the function signature as named return values, you only assign to them using `=`.
  - Using `:=` here would try to declare new variables with the same names, causing a compilation error.

- Summary:
  - Named return values simplify the return statement and can improve readability.
  - Assign to them with `=`, and use a plain `return` to return their current values.
*/
