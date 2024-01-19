package repository

import (
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/entity"
)

func (or *openingRepository) Create(entity entity.Opening) (entity.Opening, *rest.RestErr) {
	conn, err := or.database.Connect()
	if err != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(
		`INSERT INTO openings (role, description, company, location, remote, link, salary) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		entity.GetRole(),
		entity.GetDescription(),
		entity.GetCompany(),
		entity.GetLocation(),
		entity.GetRemote(),
		entity.GetLink(),
		entity.GetSalary(),
	)

	if row.Err() != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var id string
	if err = row.Scan(&id); err != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}

	entity.SetID(id)

	return entity, nil
}
