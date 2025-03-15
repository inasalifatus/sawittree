package handlers

import (
	"net/http"
	"sawittree/services"

	"github.com/labstack/echo/v4"
)

type StatsHandler struct {
	Service *services.StatsService
}

func (h *StatsHandler) GetEstateStats(c echo.Context) error {
	estateID := c.Param("id")
	stats, err := h.Service.GetStats(estateID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "estate not found"})
	}

	return c.JSON(http.StatusOK, stats)
}
