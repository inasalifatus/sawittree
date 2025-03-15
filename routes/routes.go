package routes

import (
	"sawittree/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, estateHandler *handlers.EstateHandler, treeHandler *handlers.TreeHandler, statsHandler *handlers.StatsHandler, droneHandler *handlers.DroneHandler) {
	e.POST("/estate", estateHandler.CreateEstate)
	e.POST("/estate/:id/tree", treeHandler.AddTree)
	e.GET("/estate/:id/stats", statsHandler.GetEstateStats)
	e.GET("/estate/:id/drone-plan", droneHandler.CalculateDronePath)
	e.GET("/estate/:id/drone-plan", droneHandler.CalculateDronePathWithLimit)
}
