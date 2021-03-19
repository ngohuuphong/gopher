package validations

import (
	"errors"
	"golang-test-2-api/models"
)

var (
	ErrEmptyFields  = errors.New("One or more empty fields")
	ErrInvalidEmail = errors.New("Invalid email")
)

func ValidateNewUser(user models.User) (models.User, error) {
	if IsEmpty(user.Nickname) || IsEmpty(user.Email) || IsEmpty(user.Password) {
		return models.User{}, ErrEmptyFields
	}
	if !IsEmail(user.Email) {
		return models.User{}, ErrInvalidEmail
	}
	return user, nil
}
