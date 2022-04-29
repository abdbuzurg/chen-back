package middleware

import (
	"chen/utils/response"
	"chen/utils/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			message := "Authorization header is not provided"
			response.FormatResponse(c, http.StatusUnauthorized, message, false)
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			message := "Invalid authorization header format"
			response.FormatResponse(c, http.StatusForbidden, message, false)
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != "bearer" {
			message := "Unsupported authorization type"
			response.FormatResponse(c, http.StatusForbidden, message, false)
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			response.FormatResponse(c, http.StatusForbidden, err.Error, false)
			return
		}

		c.Set("payload", payload)
	}
}
