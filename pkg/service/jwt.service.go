package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID, roleID uint) (string, error)
	VerifyToken(toke string) (*Payload, error)
}

type jWTService struct{}

func NewJWTService() JWTService {
	return &jWTService{}
}

type Payload struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

const (
	secretKey = "a;ljg;lsakdfgj;lskdfjg;kdfjgl;kd"
	tokenTTL  = 12 * time.Hour
)

func (js jWTService) GenerateToken(userID, roleID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Payload{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
		roleID,
	})
	return token.SignedString([]byte(secretKey))
}

func (js jWTService) VerifyToken(token string) (*Payload, error) {
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
