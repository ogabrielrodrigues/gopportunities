package resource

import (
	"github.com/ogabrielrodrigues/gopportunities/config/database/postgres"
	repository "github.com/ogabrielrodrigues/gopportunities/internal/domain/repository/opening_repository"
	service "github.com/ogabrielrodrigues/gopportunities/internal/domain/service/opening_service"
	handler "github.com/ogabrielrodrigues/gopportunities/internal/handler/opening_handler"
)

func OpeningResource() handler.OpeningHandler {
	database := postgres.GetDB()
	repository := repository.NewOpeningRepository(database)
	service := service.NewOpeningService(repository)

	return handler.NewOpeningHandler(service)
}
