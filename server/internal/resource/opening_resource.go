package resource

import (
	"github.com/ogabrielrodrigues/gopportunities/config/database/postgres"
	handler "github.com/ogabrielrodrigues/gopportunities/internal/handler/opening"
	repository "github.com/ogabrielrodrigues/gopportunities/internal/repository/opening"
	service "github.com/ogabrielrodrigues/gopportunities/internal/service/opening"
)

func OpeningResource() handler.OpeningHandler {
	database := postgres.GetDB()
	repository := repository.NewOpeningRepository(database)
	service := service.NewOpeningService(repository)

	return handler.NewOpeningHandler(service)
}
