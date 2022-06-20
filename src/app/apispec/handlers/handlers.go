package apispec_handlers

import (
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (h *Handler) GetAPISpec(c echo.Context) error {
	return c.File("public/html/api-spec.html")
}

func (h *Handler) ServeDocsFile(c echo.Context) error {
	return c.File("libs/api-spec/api-spec.yml")
}

func NewHandler() *Handler {
	return &Handler{}
}
