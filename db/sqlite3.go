package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	driver string
	dsn    string
}

func (sqlDB Sqlite) InitDB() (db *sql.DB, err error) {
	db, err = sql.Open(sqlDB.driver, sqlDB.dsn)
	return
}
