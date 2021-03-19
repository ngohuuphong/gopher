package auth

import (
	"errors"
	"golang-test-2-api/config"
	"golang-test-2-api/models"
	"golang-test-2-api/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidPassword = errors.New("Invalid password")
)

type Auth struct {
	User    models.User `json:"user"`
	Token   string      `json:"token"`
	IsValid bool        `json:"is_valid"`
}

var configs = config.LoadConfigs()

func SignIn(user models.User) (Auth, error) {
	password := user.Password
	user, err := models.GetUserByEmail(user.Email)
	if err != nil {
		return Auth{IsValid: false}, err
	}
	err = utils.IsPassword(user.Password, password)
	if err != nil {
		return Auth{IsValid: false}, ErrInvalidPassword
	}
	token, err := GenerateJWT(user)
	if err != nil {
		return Auth{IsValid: false}, err
	}
	return Auth{user, token, true}, nil
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) //claims := make(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = user.UID                          // inform some unique data of the user who will receive the token
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // time the token will expire
	return token.SignedString(configs.Jwt.SecretKey)
}
