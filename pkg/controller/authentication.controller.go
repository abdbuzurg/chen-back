package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chen/pkg/dto"
	"chen/pkg/service"
	"chen/utils/response"
)

type AuthentificationController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authentificationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) AuthentificationController {
	return authentificationController{
		authenticationService: service,
	}
}

func (ac authentificationController) Register(c *gin.Context) {
	var dataForRegistration dto.AuthenticationRegisterDTO
	if err := c.ShouldBindJSON(&dataForRegistration); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Data Format", false)
		return
	}

	err := ac.authenticationService.AuthRegister(dataForRegistration)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Server Cannot handle your requests", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "You have been registered", true)
}

func (ac authentificationController) Login(c *gin.Context) {
	dataForLogin := dto.AuthenticationLoginDTO{}
	if err := c.ShouldBindJSON(&dataForLogin); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Data Format", false)
		return
	}

	token, err := ac.authenticationService.AuthLogin(dataForLogin)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, err.Error(), false)
		return
	}

	response.FormatResponse(c, http.StatusOK, gin.H{
		"token": token,
	}, true)
}
