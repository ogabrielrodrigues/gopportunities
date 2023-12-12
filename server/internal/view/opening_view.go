package view

import (
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
	"github.com/ogabrielrodrigues/gopportunities/internal/handler/dto/response"
)

func OpeningToView(opening entity.Opening) *response.OpeningResponse {
	return &response.OpeningResponse{
		ID:          opening.GetID(),
		Role:        opening.GetRole(),
		Description: opening.GetDescription(),
		Company:     opening.GetCompany(),
		Location:    opening.GetLocation(),
		Remote:      opening.GetRemote(),
		Link:        opening.GetLink(),
		Salary:      opening.GetSalary(),
	}
}
