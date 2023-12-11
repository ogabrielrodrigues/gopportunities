package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/gopportunities/config"
	"github.com/ogabrielrodrigues/gopportunities/internal/routes"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	routes.RegisterMiddlewares(router)

	http.ListenAndServe(config.GetServerConfig().Port, router)
}
