package controllers

import (
    "net/http"
    "regexp"
)

type userController struct {
    userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from the User Controller!")) //Should be returning a byte array, but let's see...
}

func newUserController() *userController {
    return &userController{
        userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
    }
}
