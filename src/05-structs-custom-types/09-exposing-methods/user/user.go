package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate cannot be empty")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// The New function is a constructor function that creates a new User instance.
// We change the name of the function to New to follow Go conventions. user.New is a common pattern for constructor functions in Go.
// It takes first name, last name, and birthdate as parameters and returns a pointer to the User instance.
// It also returns an error if any of the parameters are empty.
// Fields in User struct are private, so they cannot be accessed directly from outside the package.
// They can only be accessed through methods defined on the User struct.
// This encapsulation helps to maintain the integrity of the data and allows for better control over how the data is accessed and modified.
