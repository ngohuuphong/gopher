package auto

import (
	"golang-test-4-fullstack-mysql-jwt/api/models"
)

var users = []models.User{
	models.User{Nickname: "supperman", Email: "ngohuuphong@gmail.com", Password: "123456"},
}

var posts = []models.Post{
	models.Post{
		Title:   "Welcome",
		Content: "Hello World",
	},
}
