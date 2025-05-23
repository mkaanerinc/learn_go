package main

import (
	"fmt"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	outputUserDetails(firstName, lastName, birthdate)
}

func outputUserDetails(firstName, lastName, birthdate string) {
	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

/*

What is a struct in Go?

A struct (short for "structure") in Go is a user-defined composite data type that allows you to combine fields (variables)
of different types under a single name. Think of it like a blueprint for creating objects that hold related pieces of information.

For example, if you wanted to represent a person, you might need their name (a string), their age (an integer), and their height (a floating-point number).
Instead of managing these as separate, unrelated variables, a struct lets you bundle them all into one "person" object. This makes your code more organized,
readable, and manageable, especially as your programs become more complex.

In essence, structs enable you to define custom data types that accurately model real-world entities, significantly improving code organization and clarity.

*/
