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
		accountBalance += depositAmount
		fmt.Println("Deposit successful! New balance is:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter amount to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Withdrawal successful! New balance is:", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}

/*
Explanation of `else` in Go:

- `else` is used to handle the default case when all previous `if` and `else if` conditions are false.
- It does not require a condition — it simply runs when no other branch matches.
- Syntax:
    if condition1 {
        code block 1
    } else if condition2 {
        code block 2
    } else {
        code block 3 (executed if none of the above conditions are true)
    }

- In this example:
    - If the user enters 1, 2, or 3, those specific operations are executed.
    - If the user enters 4 or any other number, the `else` block is triggered, printing "Goodbye!".
    - This acts as a graceful fallback to end the program or handle invalid input.
*/
