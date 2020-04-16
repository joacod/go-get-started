package main

import (
	"fmt"
	"playground/models"
)

func main() {
	fmt.Println("Hello Go!")

	user := models.User{
		ID:        2,
		FirstName: "Mr",
		LastName:  "Spock",
	}

	fmt.Println(user)
}
