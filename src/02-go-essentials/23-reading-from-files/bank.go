package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance := getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is:", accountBalance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("Deposit amount cannot be negative.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Deposit successful! New balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 || withdrawAmount > accountBalance {
				fmt.Println("Invalid withdrawal amount.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Withdrawal successful! New balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go Bank!")
			return
		}
	}
}

/*
Explanation of Reading Files, Using `_`, and `strconv` in Go:

1. Reading from a File:
   - We use `os.ReadFile` to read the contents of a file into memory.
   - It returns two values: the data (as a byte slice) and an error.
   - Example: `data, err := os.ReadFile("balance.txt")`

2. Ignoring the Error with `_`:
   - In Go, if a function returns multiple values, all of them must be used.
   - If you intentionally want to ignore a returned value (e.g., an error), you can use the blank identifier `_`.
   - Example: `data, _ := os.ReadFile("balance.txt")`
   - ⚠️ Note: Ignoring errors is fine for learning/demo purposes, but in production code, always handle them properly.

3. Converting String to Float using `strconv.ParseFloat`:
   - The content read from the file is a string. We need to convert it to a float.
   - `strconv.ParseFloat(string, bitSize)` is used for this.
     - `bitSize` is typically 64 for `float64`.
   - It also returns a float and an error (which we’re ignoring here for simplicity).
   - Example:
     ```go
     balanceText := string(data)
     balance, _ := strconv.ParseFloat(balanceText, 64)
     ```

Summary of flow in `getBalanceFromFile()`:
- Read the file content (`os.ReadFile`)
- Convert byte slice to string (`string(data)`)
- Convert string to float (`strconv.ParseFloat`)
- Return the float as the current balance

This allows the program to persist the balance between runs.
*/
