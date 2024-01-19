package database

import "database/sql"

type DB interface {
	Connect() (*sql.DB, error)
}
