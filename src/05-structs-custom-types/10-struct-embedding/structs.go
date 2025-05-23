package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin := user.NewAdmin("test@example.com", "password123")
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// The methods OutputUserDetails and ClearUserName belong to the User struct.
// But because of Go's struct embedding, the Admin struct "embeds" User,
// so these methods can be called directly on admin like this:
//     admin.OutputUserDetails()
//     admin.ClearUserName()
//
// This aligns with Go's "composition over inheritance" philosophy.
// Go automatically resolves the Admin → User relationship,
// making User's methods directly accessible on Admin.
//
// If User was included as a named field instead of embedded, like this:
//     type Admin struct {
//         email    string
//         password string
//         user     User  // not embedded
//     }
// then you would have to call the methods like this:
//     admin.user.OutputUserDetails()
//     admin.user.ClearUserName()
