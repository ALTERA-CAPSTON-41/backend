package middlewares

import (
	"clinic-api/src/configs"
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
