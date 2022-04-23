package service

import (
	"chen/model"
	"chen/pkg/repository"
	"chen/utils/token"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func AuthRegiseter(registrationData model.RegisterData) error {
	password, err := bcrypt.GenerateFromPassword([]byte(registrationData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	registrationData.Password = string(password)

	err = repository.UserCreate(registrationData)
	if err != nil {
		return err
	}

	return nil
}

// CHANGE RETURN TO VALID JWT TOKEN
func AuthLogin(loginData model.LoginData) (string, error) {
	user, err := repository.UserFindByUsername(loginData.Username)
	if err != nil {
		return "User does not exist", err
	}

	fmt.Println(loginData.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "Incorrect password", err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "Error logging in", err
	}

	return token, nil
}
