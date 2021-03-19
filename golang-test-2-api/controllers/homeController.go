package controllers

import (
	"golang-test-2-api/utils"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, struct {
		Message string `json:"message"`
	}{
		Message: "Go RESTful API",
	})
}
