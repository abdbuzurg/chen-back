package middleware

import (
	"chen/pkg/service"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authMiddleware := func(c *gin.Context) (int, error) {

			authHeader := c.GetHeader("Authorization")
			if len(authHeader) == 0 {
				return http.StatusUnauthorized, errors.New("authorization header is not provided")
			}

			fields := strings.Fields(authHeader)
			if len(fields) < 2 {
				return http.StatusUnauthorized, errors.New("invalid authorization header format")
			}

			authType := strings.ToLower(fields[0])
			if authType != "bearer" {
				return http.StatusUnauthorized, errors.New("unsupported authorization type")
			}

			accessToken := fields[1]
			jwtService := service.NewJWTService()
			payload, err := jwtService.VerifyToken(accessToken)
			if err != nil {
				return http.StatusUnauthorized, fmt.Errorf("invalid token: %v", err)
			}

			c.Set("payload", payload)

			return http.StatusOK, nil
		}

		status, err := authMiddleware(c)
		if err != nil {
			c.JSON(status, gin.H{
				"data":    nil,
				"error":   err.Error(),
				"success": false,
			})
			c.Abort()
		}

		c.Next()
	}
}
