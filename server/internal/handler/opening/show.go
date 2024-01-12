package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/config/validation"
	"github.com/ogabrielrodrigues/gopportunities/internal/view"
)

// Show Opening godoc
// @Summary      Show opening
// @Description  Receive user param to delete opening
// @Tags         Opening
// @Produce      json
// @Param        id path string true "Request Param"
// @Success      200  {object}  response.OpeningResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /opening/{id} [get]
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
