package main

import "fmt"

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
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go Bank!")
			return
		}
	}
}

/*
Explanation of `switch` Statement in Go:

- `switch` is used for cleaner, more readable branching when you have multiple specific values to check.
- It's a great alternative to multiple `if-else if` statements.

- Syntax:
    switch variable {
    case value1:
        code for value1
    case value2:
        code for value2
    default:
        code if no case matches
    }
*/
