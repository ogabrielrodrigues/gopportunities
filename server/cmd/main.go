package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/gopportunities/config"
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/internal/routes"
)

// @title Gopportunities
// @version 1.0
// @description Gopportunities. Find a new job opportunity
// @BasePath /api/v1
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	routes.RegisterRoutes(router)

	logger.Info(fmt.Sprintf("server running on %s", config.GetServerConfig().Port))
	if err := http.ListenAndServe(config.GetServerConfig().Port, router); err != nil {
		logger.Err("error initializing server", err)
	}
}
