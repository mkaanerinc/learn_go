/*
Ask for revenue, expenses, tax rate
Calculate earnings before tax (EBT) and earnings after tax (profit)
Calculate ratio (EBT / profit)
Output EBT, profit, and ratio
*/

/*
Goals
	1. Validate user input
		=> Shown error message & exit if invalid input is provided
			- No negative values
			- No zero values
	2. Store calculated results into file "results.txt"
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, errorForRevenue := getUserInput("Enter revenue: ")

	if errorForRevenue != nil {
		fmt.Println(errorForRevenue)
		return
		// panic(errorForRevenue)
	}

	expenses, errorForExpenses := getUserInput("Enter expenses: ")

	if errorForExpenses != nil {
		fmt.Println(errorForExpenses)
		return
		// panic(errorForExpenses)
	}

	taxRate, errorForTaxRate := getUserInput("Enter tax rate (in %): ")

	if errorForTaxRate != nil {
		fmt.Println(errorForTaxRate)
		return
		// panic(errorForTaxRate)
	}

	EBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeResultsToFile(EBT, profit, ratio)

	fmt.Printf("EBT: %.1f\n", EBT)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio (EBT / Profit): %.3f\n", ratio)
}

func getUserInput(prompt string) (input float64, err error) {
	fmt.Print(prompt)
	fmt.Scan(&input)

	if input <= 0 {
		err = errors.New("invalid input: must be greater than zero")
		return 0, err
	}
	return input, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit

	return EBT, profit, ratio
}

func writeResultsToFile(EBT, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio (EBT / Profit): %.3f\n", EBT, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
