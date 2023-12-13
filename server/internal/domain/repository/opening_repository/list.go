package repository

import (
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
)

func (or *openingRepository) List() ([]entity.Opening, *rest.RestErr) {
	conn, err := or.database.Connect()
	if err != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row, err := conn.Query(`SELECT * FROM openings`)
	if err != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var openings []entity.Opening
	var oid, role, description, company, location, link string
	var remote bool
	var salary uint64

	for row.Next() {
		if err := row.Scan(
			&oid,
			&role,
			&description,
			&company,
			&location,
			&remote,
			&link,
			&salary,
		); err != nil {
			logger.Err(err.Error(), err)
			return nil, rest.NewInternalServerErr(err.Error())
		}

		opening := entity.NewOpening(role, description, company, location, link, remote, salary)
		opening.SetID(oid)

		openings = append(openings, opening)
	}

	return openings, nil
}
