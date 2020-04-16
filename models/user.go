package models

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

// GetUsers : Return all the users
func GetUsers() []*User {
	return users
}

// AddUser : Add the given user to collection users
func AddUser(user User) (User, error) {
	user.ID = nextID
	nextID++
	users = append(users, &user)

	return user, nil
}
