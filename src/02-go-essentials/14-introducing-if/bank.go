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
	}

	fmt.Println("Your choice:", choice)
}

/*
Explanation of `if` Statement in Go:

- The `if` statement in Go is used to execute a block of code based on whether a condition evaluates to true.
- The syntax is:
    if condition {
        code to execute if condition is true
    }

- Curly braces `{}` are mandatory in Go, even if the body contains only one line.
- There is no need for parentheses `()` around the condition.

In this example:
    if choice == 1 {
        fmt.Println("Your current balance is:", accountBalance)
    }

- The condition `choice == 1` checks if the user selected option 1.
- If true, it executes the `fmt.Println(...)` statement and shows the account balance.

Note:
- Unlike some languages (like JavaScript or C), Go enforces cleaner syntax rules, such as no parentheses and mandatory braces.
- This ensures code readability and reduces potential logic bugs.
*/
