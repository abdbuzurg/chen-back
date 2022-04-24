package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chen/model"
	"chen/pkg/service"
	"chen/utils/response"
)

func Register(c *gin.Context) {
	var dataForRegistration model.RegisterData
	if err := c.ShouldBind(&dataForRegistration); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Data Format", false)
		return
	}

	err := service.AuthRegiseter(dataForRegistration)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, "Server Cannot handle your requests", false)
		return
	}

	response.FormatResponse(c, http.StatusOK, "You have been registered", true)
}

func Login(c *gin.Context) {
	var dataForLogin model.LoginData
	if err := c.ShouldBind(&dataForLogin); err != nil {
		response.FormatResponse(c, http.StatusBadRequest, "Invalid Data Format", false)
		return
	}

	message, err := service.AuthLogin(dataForLogin)
	if err != nil {
		response.FormatResponse(c, http.StatusInternalServerError, message, false)
		return
	}

	response.FormatResponse(c, http.StatusOK, gin.H{
		"token": message,
	}, true)
}
