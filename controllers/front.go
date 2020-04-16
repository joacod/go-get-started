package controllers

import "net/http"

// RegisterControllers : Register the controllers
func RegisterControllers() {
	userController := newUserController()

	http.Handle("/users", *userController)
	http.Handle("/users/", *userController)
}
