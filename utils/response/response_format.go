package response

import (
	"github.com/gin-gonic/gin"
)

func FormatResponse(c *gin.Context, statusCode int, data interface{}, success bool) {

	dataFormat := data
	if !success {
		dataFormat = gin.H{
			"error": data,
		}
	}

	c.JSON(statusCode, gin.H{
		"success": success,
		"data":    dataFormat,
	})
}
