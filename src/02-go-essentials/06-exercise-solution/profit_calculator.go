/*
Ask for revenue, expenses, tax rate
Calculate earnings before tax (EBT) and earnings after tax (profit)
Calculate ratio (EBT / profit)
Output EBT, profit, and ratio
*/

package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate (in %): ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit := EBT * (1 - taxRate/100)
	ratio := EBT / profit

	fmt.Println("EBT:", EBT)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio (EBT / Profit):", ratio)
}
