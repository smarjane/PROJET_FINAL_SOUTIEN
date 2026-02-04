package Projet

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "MyGameList.db")
	if err != nil {
		panic(err)
	}
}

func CreateDB() {
	InitDB()
	createTableusers := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		pseudo TEXT
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
	defer db.Close()
}

func InsertValue(user string, pseudo string) int {
	InitDB()
	insertQuery := `INSERT INTO users(username, pseudo) VALUES(?, ?)`
	res, err := db.Exec(insertQuery, user, pseudo)

	if err != nil {
		panic(err)
	}

	id, _ := res.LastInsertId()
	fmt.Println(id)

	defer db.Close()
	return int(id)

}
