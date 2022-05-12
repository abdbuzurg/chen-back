package globalService

import "github.com/gin-gonic/gin"

type ResponseFormatterFunc func(func(c *gin.Context) (int, interface{}, bool, error)) gin.HandlerFunc

type response struct {
	Data          interface{} `json:"data"`
	Error         string      `json:"error"`
	StatusSuccess bool        `json:"status_success"`
}

type ResponseFormatterService interface {
	SetResponseFormatter(formatterFunc ResponseFormatterFunc)
	GetResponseFormatter() ResponseFormatterFunc
}

type responseFormatterService struct {
	responseFormatterFunc ResponseFormatterFunc
}

func NewResponseFormatterService() GlobalService {
	return &responseFormatterService{
		responseFormatterFunc: DefaultResponseFormatterFunc,
	}
}

func (rs *responseFormatterService) SetResponseFormatter(formatterFunc ResponseFormatterFunc) {
	rs.responseFormatterFunc = formatterFunc
}

func (rs *responseFormatterService) GetResponseFormatter() ResponseFormatterFunc {
	return rs.responseFormatterFunc
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
