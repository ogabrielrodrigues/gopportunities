package service

import (
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
)

func (os *openingService) Show(id string) (entity.Opening, *rest.RestErr) {
	return os.repository.Show(id)
}
