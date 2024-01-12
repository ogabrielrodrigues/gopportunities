package handler

import (
	"net/http"

	service "github.com/ogabrielrodrigues/gopportunities/internal/service/opening"
)

type openingHandler struct {
	service service.OpeningService
}

type OpeningHandler interface {
	Create(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func NewOpeningHandler(service service.OpeningService) OpeningHandler {
	return &openingHandler{service}
}
