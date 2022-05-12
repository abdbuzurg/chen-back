package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"chen/pkg/dto"
	"chen/pkg/service"
)

type AuthentificationController interface {
	Register(c *gin.Context) (int, interface{}, bool, error)
	Login(c *gin.Context) (int, interface{}, bool, error)
}

type authentificationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) AuthentificationController {
	return authentificationController{
		authenticationService: service,
	}
}

func (ac authentificationController) Register(c *gin.Context) (int, interface{}, bool, error) {
	var dataForRegistration dto.AuthenticationRegister
	if err := c.ShouldBindJSON(&dataForRegistration); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid body")
	}

	err := ac.authenticationService.AuthRegister(dataForRegistration)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not register: %v", err)
	}

	return http.StatusOK, "You have been registered", true, nil
}

func (ac authentificationController) Login(c *gin.Context) (int, interface{}, bool, error) {
	dataForLogin := dto.AuthenticationLogin{}
	if err := c.ShouldBindJSON(&dataForLogin); err != nil {
		return http.StatusBadRequest, nil, false, errors.New("invalid Data Format")
	}

	token, err := ac.authenticationService.AuthLogin(dataForLogin)
	if err != nil {
		return http.StatusInternalServerError, nil, false, fmt.Errorf("could not log in: %v", err)
	}

	return http.StatusOK, gin.H{"token": token}, true, nil
}
