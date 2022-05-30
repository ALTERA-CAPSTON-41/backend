package utils

import (
	"github.com/labstack/echo/v4"
)

type base struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func CreateEchoResponse(ectx echo.Context, httpCode int, reason string, data interface{}) error {
	response := base{}
	response.Meta.Status = httpCode
	response.Meta.Message = reason
	response.Data = data
	return ectx.JSON(httpCode, response)
}
