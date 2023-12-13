package repository

import (
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
	"github.com/ogabrielrodrigues/gopportunities/internal/domain/entity"
)

func (or *openingRepository) Update(id string, entity entity.Opening) (entity.Opening, *rest.RestErr) {
	found, err := or.Show(id)
	if err != nil {
		logger.Err(err.Error(), err)
		return nil, rest.NewInternalServerErr(err.Error())
	}

	if entity.GetRole() != "" {
		found.SetRole(entity.GetRole())
	}

	if entity.GetDescription() != "" {
		found.SetDescription(entity.GetDescription())
	}

	if entity.GetCompany() != "" {
		found.SetCompany(entity.GetCompany())
	}

	if entity.GetLocation() != "" {
		found.SetLocation(entity.GetLocation())
	}

	if entity.GetRemote() != found.GetRemote() {
		found.SetRemote(entity.GetRemote())
	}

	if entity.GetLink() != "" {
		found.SetLink(entity.GetLink())
	}

	if entity.GetSalary() != 0 {
		found.SetSalary(entity.GetSalary())
	}

	conn, db_err := or.database.Connect()
	if db_err != nil {
		logger.Err(db_err.Error(), db_err)
		return nil, rest.NewInternalServerErr(db_err.Error())
	}
	defer conn.Close()

	_, db_err = conn.Exec(
		`UPDATE openings SET role=$1, description=$2, company=$3, location=$4, remote=$5, link=$6, salary=$7 WHERE id=$8`,
		found.GetRole(),
		found.GetDescription(),
		found.GetCompany(),
		found.GetLocation(),
		found.GetRemote(),
		found.GetLink(),
		found.GetSalary(),
		id,
	)

	if db_err != nil {
		logger.Err(db_err.Error(), db_err)
		return nil, rest.NewInternalServerErr(db_err.Error())
	}

	return found, nil
}
