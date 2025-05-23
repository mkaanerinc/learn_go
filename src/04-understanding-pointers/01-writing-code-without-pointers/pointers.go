package main

import "fmt"

func main() {
	age := 32 // regular variable

	fmt.Println("Age:", age)

	// Since we're passing `age` by value, a copy of `age` is made
	// and passed to the `getAdultYears` function.
	// This means there are now two copies of the value 32:
	// one in `main`, and one inside `getAdultYears`
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

// This function receives a copy of the `age` variable.
// The original `age` in main() is not affected or modified.
func getAdultYears(age int) int {
	return age - 18
}

/*
Why is this important?

- In this example, it's not a big deal because `int` is a small, simple value.
- But if we were passing a large struct or array, copying it could hurt performance.
- In such cases, it's better to pass a pointer (e.g., *MyStruct) to avoid the copy
  and allow the function to modify the original value if needed.
*/
