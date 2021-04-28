package controllers

import (
	"net/http"
	"plural_sight/code/models"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller!"))
}
func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(),2)
}
func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeResponseAsJSON(u, w)
	}
}
func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err 
}
func newUserController () *userController {
	return &userController {
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
