package service

import "github.com/ogabrielrodrigues/gopportunities/config/rest"

func (os *openingService) Delete(id string) *rest.RestErr {
	return os.repository.Delete(id)
}
