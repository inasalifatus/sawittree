package main

import (
	"sawittree/config"
	"sawittree/handlers"
	"sawittree/repositories"
	"sawittree/routes"
	"sawittree/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.InitDb()

	// Initialize repositories
	estateRepo := &repositories.EstateRepository{DB: config.DB}

	// Initialize services
	estateService := &services.EstateService{Repo: estateRepo}
	treeService := &services.TreeService{Repo: estateRepo}
	statsService := &services.StatsService{Repo: estateRepo}
	droneService := &services.DroneService{Repo: estateRepo}

	// Initialize handlers
	estateHandler := &handlers.EstateHandler{Service: estateService}
	treeHandler := &handlers.TreeHandler{Service: treeService}
	statsHandler := &handlers.StatsHandler{Service: statsService}
	droneHandler := &handlers.DroneHandler{Service: droneService}

	// Register routes
	routes.InitRoutes(e, estateHandler, treeHandler, statsHandler, droneHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
