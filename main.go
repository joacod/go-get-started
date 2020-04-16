package main

import (
	"fmt"
	"net/http"
	"playground/controllers"
	"playground/models"
)

func main() {
	helloGo()
	printUser()

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func helloGo() {
	fmt.Println("Hello Go!")
}

func printUser() {
	user := models.User{
		ID:        2,
		FirstName: "Mr",
		LastName:  "Spock",
	}

	fmt.Println(user)
}
