package service

import (
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/entity"
)

func (os *openingService) List() ([]entity.Opening, *rest.RestErr) {
	return os.repository.List()
}
