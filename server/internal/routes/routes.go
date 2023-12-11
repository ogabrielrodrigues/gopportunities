package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux) {
	router.Mount("/", openingRoutes())
}

func openingRoutes() http.Handler {
	or := chi.NewRouter()
	return or
}
