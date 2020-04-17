package models

import (
	"fmt"

	"github.com/pkg/errors"
)

// User : User model struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// CreateSomeUsers : Create some users
func CreateSomeUsers() {
	user1 := User{
		FirstName: "Mr",
		LastName:  "Spock",
	}
	user2 := User{
		FirstName: "James",
		LastName:  "Kirk",
	}
	user3 := User{
		FirstName: "Leonard",
		LastName:  "McCoy",
	}
	AddUser(user1)
	AddUser(user2)
	AddUser(user3)
}

// GetUsers : Return all the users
func GetUsers() []*User {
	CreateSomeUsers()
	return users
}

// AddUser : Add the given user to collection users
func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("The user already have an ID")
	}

	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}

// GetUserByID : Get a User by a given ID
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser : Update a user
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID : Remove a user
func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
