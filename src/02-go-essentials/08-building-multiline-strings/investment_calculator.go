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

	//formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	//formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	// outputs information

	//fmt.Println("Future Value:", futureValue)
	//fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)

	//fmt.Printf("Future Value: %v\n", futureValue)
	//fmt.Printf("Future Value (adjusted for Inflation): %v\n", futureRealValue)

	//fmt.Printf("Future Value: %v\nFuture Value (adjusted for Inflation):%v", futureValue, futureRealValue)

	//fmt.Printf("Future Value: %f\nFuture Value (adjusted for Inflation):%f", futureValue, futureRealValue)

	//fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for Inflation):%.1f", futureValue, futureRealValue)

	//fmt.Print(formattedFV, formattedFRV)

	fmt.Printf(`Future Value: %.1f
	Future Value (adjusted for Inflation):%.1f`, futureValue, futureRealValue)
}

/*
Multi-line Output Formatting in Go:

- In this version, a raw string literal (enclosed in backticks ``) is used inside `fmt.Printf` to format multi-line output.
- Raw string literals preserve whitespace and newlines as written, making them ideal for formatted block output.

Example:
fmt.Printf(`Future Value: %.1f
	Future Value (adjusted for Inflation):%.1f`, futureValue, futureRealValue)

Advantages:
- Improves readability of the output.
- Keeps format structure inside the string literal.
- Especially useful when generating structured text (e.g., reports, logs, templates).

Other Commented Options:
- Multiple alternatives are shown using `fmt.Println`, `fmt.Printf` with various format verbs (%v, %f, %.1f), and `fmt.Sprintf`.
- These alternatives demonstrate how different formatting methods work and can be used based on the context (simple output vs precise control vs storing formatted strings).

Note:
- Using `%.1f` ensures that floating-point numbers are printed with one digit after the decimal.
- Aligning output properly is important for user-facing CLI programs.
*/
