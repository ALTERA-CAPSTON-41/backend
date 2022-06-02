package utils

import (
	"clinic-api/src/configs"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserRole string

const (
	DOCTOR UserRole = "DOCTOR"
	NURSE  UserRole = "NURSE"
	ADMIN  UserRole = "ADMIN"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	Role UserRole
}

func GenerateJwt(userId string, role UserRole) (token string, err error) {
	claims := JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        userId,
		},
		Role: role,
	}

	config, _ := configs.LoadServerConfig(".")
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = rawToken.SignedString([]byte(config.JWTsecret))
	return
}

func SetJwtCookie(ec echo.Context, token string) {
	authCookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	ec.SetCookie(&authCookie)
}

func ExtractClaims(tokenStr string) (JwtCustomClaims, error) {
	config, _ := configs.LoadServerConfig(".")
	hmacSecretString := config.JWTsecret

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecretString), nil
	})

	claims := token.Claims.(JwtCustomClaims)
	return claims, err
}