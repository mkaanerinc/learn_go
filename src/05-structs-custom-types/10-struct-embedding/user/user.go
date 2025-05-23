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

// Admin struct embeds the User struct.
// ✅ Embedding means Admin "inherits" all fields and methods of User.
// - This allows Admin to reuse User’s fields and methods without explicit delegation.
// - It's Go's way of achieving a form of inheritance (composition over inheritance).
type Admin struct {
	email    string
	password string
	User     // Embedded struct — provides all fields and methods of User.
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

// NewAdmin creates a default Admin user.
// ⚠️ Hardcoded User info for simplicity — useful for testing/demo.
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}
