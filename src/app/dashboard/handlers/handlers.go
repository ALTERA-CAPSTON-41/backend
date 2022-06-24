package dashboard_handlers

import (
	"clinic-api/src/app/dashboard"
	"clinic-api/src/app/dashboard/handlers/response"
	"clinic-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services dashboard.Services
}

// onShowTotal
func (h *Handler) ShowTotalHandler(c echo.Context) error {
	feature := c.Param("feature")
	data, err := h.services.GetTotal(feature)
	if err != nil {
		if err.Error() == "forbidden" {
			return utils.CreateEchoResponse(c, http.StatusForbidden, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.Response{Total: data})
}

func NewHandler(service dashboard.Services) *Handler {
	return &Handler{service}
}
