package middlewares

import (
	"clinic-api/src/configs"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyAuthentication() echo.MiddlewareFunc {
	cfg, _ := configs.LoadServerConfig(".")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return utils.CreateEchoResponse(c, http.StatusUnauthorized, err.Error())
		},
		SigningKey:  []byte(cfg.JWTsecret),
		ContextKey:  "token",
		Claims:      utils.JwtCustomClaims{},
		TokenLookup: "cookie:token",
	})
}

func GrantDoctor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("token")
		claims, _ := utils.ExtractClaims(cookie.Value)

		if claims.Role != types.DOCTOR {
			return utils.CreateEchoResponse(c, http.StatusForbidden, nil)
		}

		return next(c)
	}
}
