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

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate (in %): ")
	outputText("Expected Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
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

func outputText(text string) {
	fmt.Print(text)
}

/*
Functions in Go:

- Functions are reusable blocks of code designed to perform specific tasks.
- They help break down complex problems into manageable pieces, improve readability, and reduce repetition.

Syntax:
	func functionName(parameters) returnType {
    	function body
	}

In this example:
	- A simple function `outputText` is defined to display a prompt message to the user.

		func outputText(text string) {
			fmt.Print(text)
		}

	- `text string` indicates the function takes one parameter of type `string`.
	- The function doesn't return any value (no return type specified), which is valid in Go.
	- Inside the function, it simply calls `fmt.Print(text)` to output the given text.

Why use a function here?
	- To avoid repeating `fmt.Print(...)` multiple times in `main()`.
	- To make `main()` cleaner and more readable.
	- Easy to extend: if you later want to add logging, styling, or localization to your output, you only need to update the function, not every print call.

Best Practices:
	- Use functions to isolate logic (input, calculations, output).
	- Prefer short, descriptive names that clearly state the function's purpose.
	- Aim for single-responsibility: each function should do one thing.

You can continue modularizing this program by creating separate functions for:
	- Input gathering
	- Future value calculation
	- Output formatting
*/
