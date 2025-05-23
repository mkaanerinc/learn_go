package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
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
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

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
Explanation of Using Third-Party Packages in Go:

1. What is a Third-Party Package?
   - These are packages not included in the Go standard library.
   - They are created and published by other developers or organizations.

2. go.mod File:
   - This file declares your module path and Go version, and tracks dependencies.
   - Example:
     ```
     module example.com/bank

     go 1.23.2

     require github.com/Pallinder/go-randomdata v1.2.0 // indirect
     ```

3. Importing Third-Party Packages:
   - Use the full import path from the module proxy or repository:
     ```go
     import "github.com/Pallinder/go-randomdata"
     ```

4. Using the Package:
   - Once imported, use its exported functions just like standard packages:
     ```go
     randomdata.PhoneNumber()
     ```

5. Example Use:
   - In this program, `randomdata.PhoneNumber()` generates a random phone number string.
   - It's used to simulate customer service contact info:
     ```go
     fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
     ```

6. Installing Third-Party Dependencies:
   - Run `go get github.com/Pallinder/go-randomdata` to fetch the dependency.
   - Go modules handle versioning and caching automatically.

7. Benefits:
   - Avoid reinventing the wheel.
   - Use well-tested libraries for common tasks like random data generation, HTTP requests, etc.

✅ Summary:
- `go-randomdata` is a third-party library used to generate random contact details.
- It's included via `go.mod`, imported in `main.go`, and enhances the UX with dynamic info.

💡 Tip:
Always review the reliability and maintenance status of third-party packages before using them in production.
*/
