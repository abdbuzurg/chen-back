package server

import (
	globalService "chen/pkg/service/global"

	"github.com/gin-gonic/gin"
)

type ResponseFormatterFunc = globalService.ResponseFormatterFunc

type response struct {
	Data          interface{} `json:"data"`
	Error         string      `json:"error"`
	StatusSuccess bool        `json:"status_success"`
}

func DefaultResponseFormatterFunc(controller func(c *gin.Context) (int, interface{}, bool, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		statusCode, data, statusSuccess, err := controller(c)

		response := response{
			Data:          data,
			StatusSuccess: statusSuccess,
		}

		if !statusSuccess {
			response.Error = err.Error()
		}

		c.JSON(statusCode, response)
	}
}
