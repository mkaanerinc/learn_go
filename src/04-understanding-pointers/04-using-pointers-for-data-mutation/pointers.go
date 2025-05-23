package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age // pointer to the variable

	fmt.Println("Age Address:", agePointer) // prints the address of the variable
	fmt.Println("Age:", *agePointer)        // prints the value of the variable

	editAgeToAdultYears(agePointer) // passing the pointer to the function
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
