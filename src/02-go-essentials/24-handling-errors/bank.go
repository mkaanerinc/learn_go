package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
		// panic("Can't continue sorry.")
		// return
	}

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
Explanation of Errors, Error Handling, and `nil` in Go:

1. Errors in Go:
   - Go uses the built-in `error` type to represent errors.
   - An error is simply a value that implements the `Error()` method returning a string.
   - Functions that might fail typically return `(value, error)`.

2. Error Handling:
   - After calling a function that returns an error, you check if the error is not `nil`.
   - If it is not `nil`, it means an error occurred, and you handle it (log, return, recover, etc.).
   - Example:
     ```go
     balance, err := getBalanceFromFile()
     if err != nil {
         fmt.Println("Error:", err)
          handle error or provide fallback
     }
     ```

4. `nil` in Go:
   - `nil` represents the absence of a value (e.g., no error).
   - When a function returns an error, `nil` means no error occurred.
   - Always check for `err != nil` before proceeding.

Summary in your program:
- `getBalanceFromFile()` attempts to read and parse the balance file.
- If reading or parsing fails, it returns a default balance (1000) and an error describing what went wrong.
- The main function checks for this error and prints it if present.
- The program continues to work with a default balance if reading the file failed.
*/
