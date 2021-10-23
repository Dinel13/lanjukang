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
var LOGIN_EXPIRATION_DURATION = time.Duration(10) * time.Hour
var RESET_TOKEN_EXPIRATION_DURATION = time.Duration(10) * time.Minute
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

// CReateToken to create token for auth handler
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

// CReateResePasswordToken to create token for reset password
func CreateResePasswordToken(userId int, secretKey string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(RESET_TOKEN_EXPIRATION_DURATION).Unix(),
		},
		Id: userId,
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

// parse token use JWT library
func parseTokenJwt(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, errors.New("signing method invalid")
		}

		return []byte(secretKey), nil
	})

	return token, err
}

// ParseToken to parse token in auth handler
func ParseToken(tokenString string, secretKey string) (int, int, error) {
	token, err := parseTokenJwt(tokenString, secretKey)

	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, 0, errors.New("token invalid")
	}

	// look the containts of claims
	id := int(claims["id"].(float64))
	role := int(claims["role"].(float64))
	expires_at := int(claims["exp"].(float64))

	// convert expires_at to time.Time
	expires_at_time := time.Unix(int64(expires_at), 0)

	// cek if token expired
	if time.Now().Unix() > expires_at_time.Unix() {
		return 0, 0, errors.New("token expired")
	}

	return int(id), int(role), nil
}

// ParseResePasswordToken to parse token for reset password
func ParseResetPasswordToken(tokenString string, secretKey string) (int, error) {
	token, err := parseTokenJwt(tokenString, secretKey)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("token invalid")
	}

	// look the containts of claims
	id := int(claims["id"].(float64))
	expires_at := int(claims["exp"].(float64))

	// convert expires_at to time.Time
	expires_at_time := time.Unix(int64(expires_at), 0)

	// cek if token expired
	if time.Now().Unix() > expires_at_time.Unix() {
		return 0, errors.New("token expired")
	}

	return int(id), nil
}
