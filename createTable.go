package bookservice

import (
	"database/sql"
	"log"
)

const database string = "./data/mydb.sqlite"

func CreateTable () bool {
	const create string = `
		CREATE TABLE IF NOT EXISTS books (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		author TEXT NOT NULL,
		description TEXT NOT NULL 
	);`

	db, err := sql.Open("sqlite3", database)

	if err != nil {
		return false
	}

	if _, err := db.Exec(create); err != nil {
		return false
	}
	
	db.Close()
	log.Println("Created the books table.")
	return true

}