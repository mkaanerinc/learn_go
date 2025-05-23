package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

// Explanation: Methods with Receiver Structs in Go
//
// - In Go, a method is a function that is associated with a specific type (usually a struct).
// - This is done by defining a **receiver** before the method name.
// - Syntax:
//     func (r ReceiverType) methodName() { ... }
//
// - In this code:
//     func (u user) outputUserDetails() { ... }
//   The receiver is of type `user`, and `u` is the instance on which the method operates.
//
// - This allows us to call the method like this:
//     appUser.outputUserDetails()
//
// - Why use receiver methods?
//   1. Encapsulates behavior with data.
//   2. Improves code organization and readability.
//   3. Enables object-oriented-like design in Go.
//
// - Value vs Pointer Receiver:
//   - (u user): receives a copy of the struct.
//   - (u *user): receives a pointer, so it can modify the original data.
//   - Use a pointer receiver when:
//       * You need to mutate the struct fields.
//       * You want to avoid copying large structs.
