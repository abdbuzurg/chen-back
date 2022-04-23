package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

const (
	signingKey = "a;ljg;lsakdfgj;lskdfjg;kdfjgl;kd"
	tokenTTL   = 12 * time.Hour
)

func GenerateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	return token.SignedString([]byte(signingKey))
}
