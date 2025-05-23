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

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
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
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

// Explanation: Mutation Methods in Go (Pointer Receiver)
//
// - A method with a *pointer receiver* (e.g., `func (u *user) clearUserName()`) allows the method to mutate
//   the original struct it was called on.
//
// - Why pointer receiver (`*user`)?
//   1. Accesses the actual memory address of the struct, not a copy.
//   2. Changes made inside the method affect the original struct data.
//
// - In contrast, a value receiver (e.g., `func (u user) clearUserName()`) would operate on a copy,
//   and any changes would not be reflected outside the method.
//
// - In this example:
//     func (u *user) clearUserName() {
//         u.firstName = ""
//         u.lastName = ""
//     }
//   This method clears the user's name fields by mutating the struct via the pointer.
//
// - Usage in main:
//     appUser.clearUserName()
//   This actually clears `appUser`'s first and last name thanks to the pointer receiver.
//
// 🔁 Summary:
// - Use pointer receivers when you need to:
//   • Modify the struct's fields.
//   • Avoid unnecessary copies of large structs.
// - Use value receivers when you don't need to modify the struct or for small immutable data types.
