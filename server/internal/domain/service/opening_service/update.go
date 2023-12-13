package service

import (
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
)

func (os *openingService) Update(id string, entity entity.Opening) (entity.Opening, *rest.RestErr) {
	return os.repository.Update(id, entity)
}
