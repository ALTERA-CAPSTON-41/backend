package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type base struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func CreateEchoResponse(c echo.Context, httpCode int, data interface{}) error {
	response := base{}
	response.Meta.Status = httpCode
	response.Meta.Message = http.StatusText(httpCode)
	response.Data = data
	return c.JSON(httpCode, response)
}
