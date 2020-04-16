package controllers

import (
	"net/http"
	"regexp"
)

// UserController : Controller for users
type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(respWritter http.ResponseWriter, req *http.Request) {
	respWritter.Write([]byte("Hello from the User Controller!"))
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
