package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
		// panic("Can't continue sorry.")
		// return
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go Bank!")
			return
		}
	}
}

/*
Explanation of Exporting and Importing in Go:

1. Packages in Go:
   - Go code is organized into **packages**. Each `.go` file belongs to a package.
   - A package is defined using the `package` keyword (e.g., `package fileops`).

2. Exporting Functions:
   - To **export** a function (make it accessible from other packages), its name must **start with an uppercase letter**.
   - In this example:
     - `GetFloatFromFile` and `WriteFloatToFile` are **exported** because they start with uppercase letters.
     - They can be used from other packages (like `main`) by importing `fileops`.

3. Importing Packages:
   - You can **import** a package using its relative module path:
     ```go
     import "your_module_path/fileops"
     ```
   - Then use its exported functions like:
     ```go
     value, err := fileops.GetFloatFromFile("balance.txt")
     ```

4. Reusability:
   - Splitting logic into reusable packages like `fileops` improves **modularity** and **code organization**.
   - File I/O logic is separated from main banking logic, making the code cleaner and easier to maintain.

5. Error Handling:
   - These functions still handle file read and parse errors gracefully using the same `(value, error)` return pattern.

In summary:
- You created a reusable utility package `fileops`.
- Exported functions (`GetFloatFromFile`, `WriteFloatToFile`) can be used in `main` or any other package.
- This approach is idiomatic and encouraged in Go for clean architecture and separation of concerns.
*/
