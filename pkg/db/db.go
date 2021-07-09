package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var kidSchema = `
	CREATE TABLE IF NOT EXISTS kid(
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		points INTEGER, 
		createdAt TEXT
	)
`
var parrentSchema = `
	CREATE TABLE IF NOT EXISTS parrent(
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		points INTEGER,
		createdAt TEXT
	)
`
var assingmentSchema = `
	CREATE TABLE IF NOT EXISTS assignment(
		id INTEGER PRIMARY KEY, 
		description TEXT, 
		points INTEGER,
		createdat TEXT,
		status TEXT, 
		kidid INTEGER
	)
`

func GetDatabase() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "./pointsystem.db")
	if err != nil {
		log.Fatalln("Failed to Connect: $v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(kidSchema)
	db.MustExec(parrentSchema)
	db.MustExec(assingmentSchema)
	return db
}

func ErrorHandling(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
