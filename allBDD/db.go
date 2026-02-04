package test

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "MyGameList.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	CreateDB()
}

func CreateDB() {
	createTableusers := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT
	);
	`
	_, err := db.Exec(createTableusers)
	if err != nil {
		panic(err)
	}

	createTablegames := `
	CREATE TABLE IF NOT EXISTS games(
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT, 
		platform TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`

	_, err = db.Exec(createTablegames)
	if err != nil {
		panic(err)
	}
}
