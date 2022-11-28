package db

import "database/sql"

type DB interface {
	InitDB() (*sql.DB, error)
}
