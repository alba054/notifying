package database

import (
	"alba054/kartjis-notify/shared"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(database_url string) *sql.DB {
	db, err := sql.Open("mysql", database_url)
	shared.ThrowError(err)

	return db
}
