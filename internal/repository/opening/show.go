package repository

import (
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/entity"
)

func (or *openingRepository) Show(id string) (entity.Opening, *rest.RestErr) {
	conn, err := or.database.Connect()
	if err != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM openings WHERE id=$1`, id)

	if row.Err() != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var oid, role, description, company, location, link string
	var remote bool
	var salary uint64

	if err = row.Scan(
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

	return opening, nil
}
