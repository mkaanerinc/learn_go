package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	// outputs information

	//fmt.Println("Future Value:", futureValue)
	//fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)

	//fmt.Printf("Future Value: %v\n", futureValue)
	//fmt.Printf("Future Value (adjusted for Inflation): %v\n", futureRealValue)

	//fmt.Printf("Future Value: %v\nFuture Value (adjusted for Inflation):%v", futureValue, futureRealValue)

	//fmt.Printf("Future Value: %f\nFuture Value (adjusted for Inflation):%f", futureValue, futureRealValue)

	//fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation):%.1f", futureValue, futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}

/*
Formatted Output in Go:

In Go, the `fmt` package provides multiple ways to print output:

1. fmt.Println:
	- Prints each argument with a space and adds a newline at the end.
	- Example: fmt.Println("Future Value:", futureValue)
	- Simple and readable, but offers limited control over formatting.

2. fmt.Printf:
	- Prints a formatted string directly to the console.
	- Uses format verbs like %v (default), %f (float), %.1f (float with 1 decimal place), etc.
	- Example: fmt.Printf("Future Value: %.1f\n", futureValue)
	- Good for precise output control.

3. fmt.Sprintf:
	- Works like `Printf`, but returns the formatted string instead of printing it.
	- Useful when you want to store or manipulate the formatted string before printing.
	- Example: formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)

Why use formatting?
	- To control how numbers are displayed (e.g., number of decimal places).
	- To improve readability and presentation.
	- To build output strings that can be reused, stored, or printed later.

In this program:
	- Future value and real future value are both formatted with one decimal place using `Sprintf`.
	- The formatted strings are then printed together using `fmt.Print`.

Bonus:
	- You can also format values into strings to write to files, logs, or UI elements using `fmt.Sprintf`.
*/
