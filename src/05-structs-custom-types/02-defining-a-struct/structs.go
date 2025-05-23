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

	outputUserDetails(appUser)
}

func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

/*
Explanation of `struct` in Go:

- A `struct` (short for "structure") is a composite data type in Go that groups together related values under one type.
- Each value inside a struct is called a **field**, and it has a name and a type.

In this program:
- The `user` struct is defined to represent a user with:
  - `firstName` and `lastName` as strings,
  - `birthdate` as a string (can be time.Time for more accuracy),
  - `createdAt` as a time.Time field that captures when the user was created.

Purpose:
- Grouping related user information in one type improves organization and readability.
- Instead of managing multiple variables (`firstName`, `lastName`, `birthdate`), we store all related data in a single `user` object.

Why Use Structs?
1. **Logical grouping**: Combines multiple fields into one object.
2. **Code clarity**: More meaningful and readable function signatures.
3. **Easier expansion**: New fields can be added without changing a lot of code.
4. **Support for methods**: You can attach functions to structs to work with their data.

Summary:
- Structs allow you to model real-world entities (like a user) in a clean and structured way.
- `appUser := user{...}` shows how to initialize and use the struct.
*/
