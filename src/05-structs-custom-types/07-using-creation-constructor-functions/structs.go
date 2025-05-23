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

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := newUser(userFirstName, userLastName, userBirthdate)

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

// newUser is a constructor function for the user struct.
//
// 🧱 What is a Constructor Function?
// - Go does not have built-in constructors like other OOP languages.
// - Instead, it's common to use a regular function that returns a pointer to a newly created struct.
// - These functions are named like `newTypeName` (e.g., `newUser`).
//
// ✅ Why use a constructor function?
// - It centralizes object creation logic.
// - You can pre-fill default values (like `time.Now()`).
// - It enforces encapsulation and validation logic if needed in the future.
//
// Example:
//     appUser := newUser("Alice", "Smith", "01/01/1990")
//     Returns a pointer to a user struct with prefilled fields.
