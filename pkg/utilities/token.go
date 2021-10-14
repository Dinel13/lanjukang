package utilities

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	jwt.StandardClaims
	Id   int `json:"id"`
	Role int `json:"role"`
}

var APPLICATION_NAME = "LankjukangBe"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func CreateToken(userId int, role int, secretKey string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Id:   userId,
		Role: role,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseToken(tokenString string, secretKey string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, errors.New("signing method invalid")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("token invalid")
	}

	// TODO: check if token is expired

	id := claims["id"].(float64)
	return int(id), nil
}
