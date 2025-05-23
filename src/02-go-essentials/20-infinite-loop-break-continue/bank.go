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

		if choice == 1 {
			fmt.Println("Your current balance is:", accountBalance)

		} else if choice == 2 {
			fmt.Print("Enter amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("Deposit amount cannot be negative.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Deposit successful! New balance is:", accountBalance)

		} else if choice == 3 {
			fmt.Print("Enter amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 || withdrawAmount > accountBalance {
				fmt.Println("Invalid withdrawal amount.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Withdrawal successful! New balance is:", accountBalance)

		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for using Go Bank!")
}

/*
Explanation of Infinite `for`, `continue`, `break` and Conditional For Loops in Go:

- `for { ... }`:
    - This is an infinite loop in Go.
    - It keeps the program running in a loop until explicitly broken.
    - Equivalent to `while(true)` in other languages.

- `continue`:
    - Used to skip the rest of the current loop iteration and move to the next one.
    - In this program, it's used to re-show the menu if the user enters an invalid deposit or withdrawal amount.

- `break`:
    - Used to exit the loop.
    - Here, it ends the program when the user selects option 4 (Exit).

- Why this structure?
    - It keeps the program interactive, continuously prompting the user until they choose to exit.
    - Prevents program termination on invalid input and instead allows retrying.

- Note:
    - Unlike `return`, which exits the whole function, `break` only stops the loop.

Conditional For Loops

Besides the for variations introduced in the last lectures, there also is another common variation
(which will also be shown again later in the course):

	for someCondition {
	  do something ...
	}

someCondition is an expression that yields a boolean value or a variable that contains a boolean value (i.e., true or false).

In that case, the loop will continue to execute the code inside the loop body until the condition / variable yields false.
*/
