package controllers

import (
	"net/http"
	"playground/models"
	"regexp"
)

// UserController : Controller for users
type userController struct {
	userIDPattern *regexp.Regexp
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (uc userController) ServeHTTP(respWritter http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/users" {
		uc.getAll(respWritter, req)
	} else {
		respWritter.Write([]byte("Live long and proper."))
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}
