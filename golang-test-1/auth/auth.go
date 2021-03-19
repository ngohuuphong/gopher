package auth

import (
	"errors"
	"fmt"
	"golang-test-1/models"
	"golang-test-1/utils"
	"strings"
)

var (
	ErrEmailNotFound   = errors.New("Email doesn't Exist")
	ErrInvalidPassword = errors.New("Invalid Password")
	ErrEmptyFields     = errors.New("Please fill in")
)

func Singin(email, password string) (models.User, error) {
	err := validateFields(strings.ToLower(email), password)
	if err != nil {
		return models.User{}, err
	}
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, ErrEmailNotFound
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	fmt.Println(user)
	return user, nil
}

func validateFields(email, password string) error {
	if models.IsEmpty(models.Trim(email)) || models.IsEmpty(password) {
		return ErrEmptyFields
	}
	if !models.IsEmail(email) {
		return models.ErrInvalidEmail
	}
	return nil
}
