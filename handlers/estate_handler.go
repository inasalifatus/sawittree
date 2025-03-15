package handlers

import (
	"net/http"
	"sawittree/services"

	"github.com/labstack/echo/v4"
)

type EstateHandler struct {
	Service *services.EstateService
}

func (h *EstateHandler) CreateEstate(c echo.Context) error {
	var req struct {
		Width  int `json:"width"`
		Length int `json:"length"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	estate, err := h.Service.CreateEstate(req.Width, req.Length)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": estate.ID.String()})
}
