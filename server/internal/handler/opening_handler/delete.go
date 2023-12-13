package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/config/validation"
)

func (oh *openingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := validation.Validate.Var(id, "uuid4"); err != nil {
		logger.Err("param id not provided", err)
		rest_err := validation.ValidateVar("id", err)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	if err := oh.service.Delete(id); err != nil {
		logger.Err("error deleting opening", err)
		rest_err := rest.NewBadRequestErr("error deleting opening", nil)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
