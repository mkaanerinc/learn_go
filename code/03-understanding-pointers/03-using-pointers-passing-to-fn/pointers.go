package main

import "fmt"

func main() {
	age := 32 // regular variable

	//var agePointer *int
	agePointer := &age // Pointer variable. Holds address of the age variable

	fmt.Println("Age: ", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
