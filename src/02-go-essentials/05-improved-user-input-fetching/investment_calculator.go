package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

/*
User Input in Go with fmt.Scan:

- The `fmt.Scan` function is used to read input from the user via the standard input (keyboard).
- It stores the entered value in the **variable whose memory address is passed** to it.
  Example:
    fmt.Scan(&investmentAmount)

- The `&` symbol is the **address-of operator**. It passes the memory address of the variable so that `fmt.Scan` can modify its value directly.
  Without `&`, Go would try to copy the value, which doesn't allow writing input into it.

Key Points:
	- Each call to `fmt.Scan` reads input until whitespace or newline.
	- Variables must be declared before scanning input into them.

In this program:
	- The user is asked to input the investment amount, expected return rate, and investment duration (in years).
	- These inputs are then used to calculate the future value and the real future value (adjusted for inflation).

fmt.Scan() Limitations

	- The fmt.Scan() function is a great function for easily fetching & using user input through the command line.
	- But this function also has an important limitation: You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.
	- For the moment, we only need single words / digits as input, so that's no problem.
	- Later, when we work on projects where more complex input values are required, you'll therefore learn about an alternative to fmt.Scan().
*/
