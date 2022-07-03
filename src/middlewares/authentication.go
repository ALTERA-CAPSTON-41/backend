package middlewares

import (
	"clinic-api/src/configs"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyAuthentication() echo.MiddlewareFunc {
	cfg, _ := configs.LoadServerConfig(".")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return utils.CreateEchoResponse(c, http.StatusUnauthorized, nil)
		},
		SigningKey:  []byte(cfg.JWTsecret),
		ContextKey:  "token",
		Claims:      jwt.MapClaims{},
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	})
}

func GrantDoctor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := utils.GetJwtTokenFromRequest(c)
		claims, _ := utils.ExtractClaims(authToken)

		if claims.Role != types.DOCTOR {
			return utils.CreateEchoResponse(c, http.StatusForbidden, nil)
		}

		return next(c)
	}
}

func GrantAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := utils.GetJwtTokenFromRequest(c)
		claims, _ := utils.ExtractClaims(authToken)

		if claims.Role != types.ADMIN {
			return utils.CreateEchoResponse(c, http.StatusForbidden, nil)
		}

		return next(c)
	}
}
