package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers : Register the controllers
func RegisterControllers() {
	userController := newUserController()

	http.Handle("/users", *userController)
	http.Handle("/users/", *userController)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}