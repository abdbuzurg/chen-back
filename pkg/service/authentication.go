package service

import (
	"chen/model"
	"chen/pkg/repository"
	"chen/utils/token"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	AuthRegister(registrationData model.RegisterData) error
	AuthLogin(loginData model.LoginData) (string, error)
}

type authenticationService struct {
	authenticationRepository repository.AuthenticationRepository
}

func NewAuthenticationService(repo repository.AuthenticationRepository) AuthenticationService {
	return authenticationService{
		authenticationRepository: repo,
	}
}

func (as authenticationService) AuthRegister(registrationData model.RegisterData) error {
	password, err := bcrypt.GenerateFromPassword([]byte(registrationData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	registrationData.Password = string(password)

	err = as.authenticationRepository.Create(registrationData)
	if err != nil {
		return err
	}

	return nil
}

// CHANGE RETURN TO VALID JWT TOKEN
func (as authenticationService) AuthLogin(loginData model.LoginData) (string, error) {
	user, err := as.authenticationRepository.UserFindByUsername(loginData.Username)
	if err != nil {
		return "User does not exist", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "Incorrect password", err
	}

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		return "Error logging in", err
	}

	return token, nil
}
