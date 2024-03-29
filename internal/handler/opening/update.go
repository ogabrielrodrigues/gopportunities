package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/config/validation"
	"github.com/ogabrielrodrigues/gopportunities/internal/dto/request"
	"github.com/ogabrielrodrigues/gopportunities/internal/entity"
	"github.com/ogabrielrodrigues/gopportunities/internal/view"
)

// Update Opening godoc
// @Summary      Update opening
// @Description  Receive user request body to update a opening
// @Tags         Opening
// @Accept       json
// @Produce      json
// @Param        id path string true "Request Param"
// @Param        request body request.CreateOpeningRequest true "Request Body"
// @Success      200  {object}  response.OpeningResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /opening/{id} [put]
func (oh *openingHandler) Update(w http.ResponseWriter, r *http.Request) {
	body := request.UpdateOpeningRequest{}
	id := chi.URLParam(r, "id")

	if err := validation.Validate.Var(id, "uuid4"); err != nil {
		logger.Err("param id not provided", err)
		rest_err := validation.ValidateVar("id", err)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Err("error decoding request body", err)
		rest_err := rest.NewInternalServerErr("error decoding request body")
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	if err := validation.Validate.Struct(body); err != nil {
		logger.Err("error on request validation", err)
		rest_err := validation.ValidateErr(err)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	opening := entity.NewOpening(
		body.Role,
		body.Description,
		body.Company,
		body.Location,
		body.Link,
		body.Remote,
		body.Salary,
	)

	opening, err := oh.service.Update(id, opening)
	if err != nil {
		logger.Err("error on opening update", err)
		rest_err := rest.NewBadRequestErr("error on opening update", nil)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	rest.JSON(w, http.StatusOK, view.OpeningToView(opening))
}
