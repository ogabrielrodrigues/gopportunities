package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/gopportunities/internal/resource"
)

func RegisterRoutes(router *chi.Mux) {
	router.Mount("/opening", openingRoutes())
}

func openingRoutes() http.Handler {
	or := chi.NewRouter()
	oh := resource.OpeningResource()

	or.Post("/", oh.Create)
	or.Get("/{id}", oh.Show)
	or.Get("/", oh.List)
	or.Put("/{id}", oh.Update)
	or.Delete("/{id}", oh.Delete)

	return or
}
