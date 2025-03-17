package handlers

import (
	"net/http"
	"sawittree/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DroneHandler struct {
	Service *services.DroneService
}

func (h *DroneHandler) CalculateDronePath(c echo.Context) error {
	estateID := c.Param("id")
	distance, err := h.Service.CalculateDronePath(estateID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "estate not found"})
	}

	return c.JSON(http.StatusOK, map[string]int{"distance": distance})
}

func (h *DroneHandler) CalculateDronePathWithLimit(c echo.Context) error {
	estateID := c.Param("id")
	maxDistanceStr := c.QueryParam("max_distance")

	maxDistance := -1
	if maxDistanceStr != "" { // Jika parameter diberikan, validasi
		var err error
		maxDistance, err = strconv.Atoi(maxDistanceStr)
		if err != nil || maxDistance <= 0 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid max_distance"})
		}
	}

	distance, restX, restY, err := h.Service.CalculateDronePathWithLimit(estateID, maxDistance)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "estate not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"distance": distance,
		"rest": map[string]int{
			"x": restX,
			"y": restY,
		},
	})
}
