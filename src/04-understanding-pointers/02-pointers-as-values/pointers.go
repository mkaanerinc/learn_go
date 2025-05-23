package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age // pointer to the variable

	fmt.Println("Age:", agePointer)  // prints the address of the variable
	fmt.Println("Age:", *agePointer) // prints the value of the variable

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}

/*
A Pointer's Null Value

All values in Go have a so-called "Null Value" - i.e., the value that's set as a default if
no value is assigned to a variable.

For example, the null value of an int variable is 0. Of a float64, it would be 0.0.
Of a string, it's "".

For a pointer, it's nil - a special value built-into Go.

nil represents the absence of an address value - i.e., a pointer pointing at no address /
no value in memory.

*/
