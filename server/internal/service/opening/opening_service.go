package service

import (
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/entity"
	repository "github.com/ogabrielrodrigues/gopportunities/internal/repository/opening"
)

type openingService struct {
	repository repository.OpeningRepository
}

type OpeningService interface {
	Create(entity.Opening) (entity.Opening, *rest.RestErr)
	Show(string) (entity.Opening, *rest.RestErr)
	List() ([]entity.Opening, *rest.RestErr)
	Update(string, entity.Opening) (entity.Opening, *rest.RestErr)
	Delete(string) *rest.RestErr
}

func NewOpeningService(repository repository.OpeningRepository) OpeningService {
	return &openingService{repository}
}
