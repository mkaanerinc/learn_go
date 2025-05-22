package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your current balance is:", accountBalance)

	} else if choice == 2 {
		fmt.Print("Enter amount to deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount < 0 {
			fmt.Println("Deposit amount cannot be negative.")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Deposit successful! New balance is:", accountBalance)

	} else if choice == 3 {
		fmt.Print("Enter amount to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 || withdrawAmount > accountBalance {
			fmt.Println("Invalid withdrawal amount.")
			return
		}

		accountBalance -= withdrawAmount
		fmt.Println("Withdrawal successful! New balance is:", accountBalance)

	} else {
		fmt.Println("Goodbye!")
	}
}

/*
Explanation of Nested `if` Statements and `return` Usage in Go:

- A nested `if` is an `if` statement placed inside another `if`, `else if`, or `else` block.
- It enables more granular control when additional conditions must be checked within a broader decision.

- Example structure:
    if condition1 {
        if condition2 {
            do something
        }
    }

- In this example:
    - When the user selects option 2 (Deposit):
        - A nested `if` checks whether the deposit amount is negative.
        - If it is, an error message is printed and `return` is used to terminate the `main` function early.
    - When the user selects option 3 (Withdraw):
        - A nested `if` ensures that the withdrawal amount is:
            - Greater than 0, and
            - Not more than the current account balance.
        - If invalid, an error is printed and `return` stops the program to prevent invalid transactions.

- The `return` statement:
    - Immediately exits the current function.
    - Useful for avoiding deeper nesting or continuing execution when an error or invalid input is detected.

- Summary:
    - `Nested if` improves input validation within specific scenarios.
    - `return` helps safely and cleanly exit logic when conditions are not met.
*/
