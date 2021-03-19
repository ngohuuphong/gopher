package routes

import (
	"net/http"

	"golang-test-4-fullstack-mysql-jwt/api/controllers"
)

var loginRoutes = []Route{
	Route{
		URI:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
