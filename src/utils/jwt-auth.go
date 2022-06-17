package utils

import (
	"clinic-api/src/configs"
	"clinic-api/src/types"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	Name string             `json:"name"`
	NIP  string             `json:"nip"`
	Role types.UserRoleEnum `json:"role"`
}

func GenerateJwt(userId string, name string, NIP string, role types.UserRoleEnum) (token string, err error) {
	claims := JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        userId,
		},
		Name: name,
		NIP:  NIP,
		Role: role,
	}

	config, _ := configs.LoadServerConfig(".")
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = rawToken.SignedString([]byte(config.JWTsecret))
	return
}

func SetJwtCookie(c echo.Context, token string) {
	authCookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(&authCookie)
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
