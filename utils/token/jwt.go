package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Payload struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

const (
	secretKey = "a;ljg;lsakdfgj;lskdfjg;kdfjgl;kd"
	tokenTTL  = 12 * time.Hour
)

func GenerateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Payload{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token signature")
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, errors.New("invalid token format")
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("invalid token payload")
	}

	return payload, nil
}
