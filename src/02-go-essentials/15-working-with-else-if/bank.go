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
	}

	fmt.Println("Your choice:", choice)
}

/*
Explanation of `else if` and `+=` in Go:

1. `else if`:
   - Used to handle multiple conditional branches.
   - Follows an `if` block to check another condition if the first one is false.
   - Syntax:
       if condition1 {
           code if condition1 is true
       } else if condition2 {
           code if condition2 is true
       }

   - In this example:
       - If `choice == 1`, the balance is shown.
       - If `choice == 2`, it proceeds to deposit money.

2. `+=` Operator:
   - A shorthand for incrementing a variable by a specific value.
   - Equivalent to: `accountBalance = accountBalance + depositAmount`
   - In this code:
       accountBalance += depositAmount
     increases the current account balance by the deposit amount.

These constructs make the code cleaner and more concise while maintaining readability.
*/
