package handlers

import (
	"net/http"
	"sawittree/services"

	"github.com/labstack/echo/v4"
)

type TreeHandler struct {
	Service *services.TreeService
}

func (h *TreeHandler) AddTree(c echo.Context) error {
	var req struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Height int `json:"height"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	estateID := c.Param("id")
	tree, err := h.Service.AddTree(estateID, req.X, req.Y, req.Height)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": tree.ID.String()})
}
