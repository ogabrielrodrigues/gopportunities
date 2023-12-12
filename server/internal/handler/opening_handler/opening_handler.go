package handler

import (
	"net/http"

	service "github.com/ogabrielrodrigues/gopportunities/internal/domain/service/opening_service"
)

type OpeningHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type openingHandler struct {
	service service.OpeningService
}

func NewOpeningHandler(service service.OpeningService) OpeningHandler {
	return &openingHandler{service}
}
