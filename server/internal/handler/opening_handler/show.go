package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/config/validation"
	"github.com/ogabrielrodrigues/gopportunities/internal/view"
)

func (oh *openingHandler) Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := validation.Validate.Var(id, "uuid4"); err != nil {
		logger.Err("param id not provided", err)
		rest_err := validation.ValidateVar("id", err)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	opening, err := oh.service.Show(id)
	if err != nil {
		logger.Err("error on show opening", err)
		rest_err := rest.NewBadRequestErr("error on show opening", nil)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	rest.JSON(w, http.StatusOK, view.OpeningToView(opening))
}
