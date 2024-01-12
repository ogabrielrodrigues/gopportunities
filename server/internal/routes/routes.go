package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/ogabrielrodrigues/gopportunities/docs"
	"github.com/ogabrielrodrigues/gopportunities/internal/middleware"
	"github.com/ogabrielrodrigues/gopportunities/internal/resource"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(router *chi.Mux) {
	middleware.RegisterMiddlewares(router)

	router.Mount("/api/v1/swagger", httpSwagger.WrapHandler)
	router.Mount("/api/v1/opening", openingRoutes())
}

func openingRoutes() http.Handler {
	or := chi.NewRouter()
	oh := resource.OpeningResource()

	or.Post("/", oh.Create)
	or.Get("/{id}", oh.Show)
	or.Get("/list", oh.List)
	or.Put("/{id}", oh.Update)
	or.Delete("/{id}", oh.Delete)

	return or
}
