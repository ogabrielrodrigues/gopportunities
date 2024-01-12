package handler

import (
	"net/http"

	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/dto/response"
	"github.com/ogabrielrodrigues/gopportunities/internal/view"
)

// List Opening godoc
// @Summary      List openings
// @Description  List all openings
// @Tags         Opening
// @Produce      json
// @Success      200  {object}  []response.OpeningResponse
// @Failure      400  {object}  rest.RestErr
// @Failure      500  {object}  rest.RestErr
// @Router       /opening [get]
func (oh *openingHandler) List(w http.ResponseWriter, r *http.Request) {
	openings, err := oh.service.List()
	if err != nil {
		logger.Err("error on list openings", err)
		rest_err := rest.NewBadRequestErr("error on list openings", nil)
		rest.JSON(w, rest_err.Code, rest_err)
		return
	}

	var openings_view []*response.OpeningResponse

	for _, opening := range openings {
		openings_view = append(openings_view, view.OpeningToView(opening))
	}

	rest.JSON(w, http.StatusOK, openings_view)
}
