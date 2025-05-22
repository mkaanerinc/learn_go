/*
Ask for revenue, expenses, tax rate
Calculate earnings before tax (EBT) and earnings after tax (profit)
Calculate ratio (EBT / profit)
Output EBT, profit, and ratio
*/

package main

import "fmt"

func main() {
	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate (in %): ")

	EBT, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", EBT)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio (EBT / Profit): %.3f\n", ratio)
}

func getUserInput(prompt string) (input float64) {
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func calculateFinancials(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit

	return EBT, profit, ratio
}
