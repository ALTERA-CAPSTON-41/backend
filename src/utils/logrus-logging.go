package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func CreateLog(c echo.Context, err string) {
	logrus.WithFields(logrus.Fields{
		"method": c.Request().Method,
		"path":   c.Path(),
	}).Error(err)
}
