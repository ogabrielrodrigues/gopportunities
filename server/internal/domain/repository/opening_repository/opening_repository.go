package repository

import (
	"github.com/ogabrielrodrigues/gopportunities/config/database"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
)

type openingRepository struct {
	database database.DB
}

type OpeningRepository interface {
	Create(entity.Opening) (string, *rest.RestErr)
	Show(string) (entity.Opening, *rest.RestErr)
}

func NewOpeningRepository(database database.DB) OpeningRepository {
	return &openingRepository{database}
}
