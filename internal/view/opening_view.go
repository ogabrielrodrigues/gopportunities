package view

import (
	"github.com/ogabrielrodrigues/gopportunities/internal/dto/response"
	"github.com/ogabrielrodrigues/gopportunities/internal/entity"
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
