package controllers

import (
	"encoding/json"
	"net/http"

	"golang-test-4-fullstack-mysql-jwt/api/auth"

	"golang-test-4-fullstack-mysql-jwt/api/models"
	"golang-test-4-fullstack-mysql-jwt/api/responses"
)

// Login is the signIn method
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}
