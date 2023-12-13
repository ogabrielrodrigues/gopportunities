package repository

import (
	"github.com/ogabrielrodrigues/gopportunities/config/logger"
	"github.com/ogabrielrodrigues/gopportunities/config/rest"
)

func (or *openingRepository) Delete(id string) *rest.RestErr {
	conn, db_err := or.database.Connect()
	if db_err != nil {
		logger.Err(db_err.Error(), db_err)
		return rest.NewInternalServerErr(db_err.Error())
	}
	defer conn.Close()

	tx, err := conn.Exec(
		`DELETE FROM openings WHERE id=$1`,
		id,
	)

	if err != nil {
		logger.Err(err.Error(), err)
		return rest.NewInternalServerErr(err.Error())
	}

	if lines, err := tx.RowsAffected(); lines != 1 {
		logger.Err(err.Error(), err)
		return rest.NewInternalServerErr(err.Error())
	}

	return nil
}
