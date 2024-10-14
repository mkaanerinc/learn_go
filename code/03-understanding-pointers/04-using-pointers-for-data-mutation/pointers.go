package main

import "fmt"

func main() {
	age := 32 // regular variable

	//var agePointer *int
	agePointer := &age // Pointer variable. Holds address of the age variable

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
