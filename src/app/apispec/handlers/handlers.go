package apispec_handlers

import "github.com/labstack/echo/v4"

type Handler struct{}

func (h *Handler) GetAPISpec(c echo.Context) error {
	return c.File("src/views/api-spec.html")
}

func NewHandler() *Handler {
	return &Handler{}
}