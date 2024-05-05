package repository

import (
	"database/sql"
	"testAgent/entities"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./weather.db")
	entities.CommonError("Failed to open database\n", err)

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS weather (
	  city TEXT PRIMARY KEY,
	  temperature REAL,
	)`)
	entities.CommonError("The database already exists\n", err)
}
