package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/config/validation"
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
	"github.com/ogabrielrodrigues/gopportunities/internal/handler/dto/request"
	"github.com/ogabrielrodrigues/gopportunities/internal/view"
)

func (oh *openingHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := request.OpeningCreateRequest{}

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

	opening, err := oh.service.Create(opening)
	if err != nil {
		logger.Err("error on opening creation", err)
		rest_err := rest.NewBadRequestErr("error on opening creation", nil)
		rest.JSON(w, rest_err.Code, rest_err)
	}

	w.Header().Set("Location", fmt.Sprintf("/opening/%s", opening.GetID()))
	rest.JSON(w, http.StatusCreated, view.OpeningToView(opening))
}
