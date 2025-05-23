package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	for i := 0; i < 200; i++ {
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
}

/*
Explanation of `for` Loop Usage in Go:

- In Go, `for` is the only looping construct and it can be used like a `while`, `do-while`, or classic `for` loop depending on how it's written.
- Here, the loop is defined as:
    for i := 0; i < 200; i++ {
        logic
    }
*/
