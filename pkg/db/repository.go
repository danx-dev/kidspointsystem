package db

import (
	"database/sql"
)

func GetAll(query string) *sql.Rows {
	db, _ := sql.Open("sqlite3", "./pointsystem.db")
	row, _ := db.Query(query)
	defer row.Close()
	return row
}
