package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = 1000

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
Explanation of Writing to Files in Go:

- Go provides the `os` and `io/ioutil` (now deprecated in favor of `os` and `io`) packages for file operations.
- In this program, we use `os.WriteFile` to write the updated balance to a file called `balance.txt`.

- Function used:
    os.WriteFile(filename string, data []byte, perm fs.FileMode) error

    - `filename`: Name of the file to write to.
    - `data`: The content to write (must be a slice of bytes).
    - `perm`: File permissions (e.g., 0644 means: owner read/write, group and others read-only).

- In this program:
    - After every successful deposit or withdrawal, the new balance is written to the file.
    - The balance is converted to a string using `fmt.Sprint()`.
    - Then it's converted to a byte slice with `[]byte(...)` and written to `balance.txt`.

- Why do this?
    - This ensures that the balance is saved between sessions.
    - If the app is restarted, the balance can be restored from the file (if reading is implemented).

- Example:
    balanceText := fmt.Sprint(balance)
    os.WriteFile("balance.txt", []byte(balanceText), 0644)

- Note:
    - `os.WriteFile` will **overwrite** the file if it already exists.
    - To **append** instead, a different approach is required (e.g., opening the file with `os.OpenFile` and using flags).
*/
