package service

import (
	"chen/pkg/dto"
	"chen/pkg/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	AuthRegister(registrationData dto.AuthenticationRegister) error
	AuthLogin(loginData dto.AuthenticationLogin) (string, error)
}

type authenticationService struct {
	authenticationRepository repository.AuthenticationRepository
	jwtService               JWTService
}

func NewAuthenticationService(repo repository.AuthenticationRepository, jwtService JWTService) AuthenticationService {
	return authenticationService{
		authenticationRepository: repo,
		jwtService:               jwtService,
	}
}

func (as authenticationService) AuthRegister(registrationData dto.AuthenticationRegister) error {
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

func (as authenticationService) AuthLogin(loginData dto.AuthenticationLogin) (string, error) {
	user, err := as.authenticationRepository.UserFindByUsername(loginData.Username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := as.jwtService.GenerateToken(user.ID, user.RoleID)
	if err != nil {
		return "", errors.New("could not sign in")
	}

	return token, nil
}
