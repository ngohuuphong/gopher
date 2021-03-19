package controllers

import (
	"encoding/json"
	"golang-test-2-api/auth"
	"golang-test-2-api/models"
	"golang-test-2-api/utils"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	userAuthenticate, err := auth.SignIn(user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	utils.ToJson(w, userAuthenticate)
}
