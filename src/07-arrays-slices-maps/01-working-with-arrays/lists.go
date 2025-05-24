package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(productNames)
	fmt.Println(prices[2])
}

/*
📌 Arrays in Go:

- Arrays are fixed-size collections of elements of the same type.
- Syntax: `[size]Type` — e.g., `[4]string` means an array of 4 strings.
- Indexing starts from 0.
- Arrays cannot change in size after declaration.

💡 Notes:
- You can assign values at specific indexes later.
- The rest of the elements will have their zero values (`""` for string, `0.0` for float64).
*/
