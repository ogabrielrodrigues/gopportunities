package service

import (
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
)

func (os *openingService) Create(entity entity.Opening) (entity.Opening, *rest.RestErr) {
	id, err := os.repository.Create(entity)
	if err != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}

	entity.SetID(id)

	return entity, nil
}
