package models

import (
	"database/sql"
	"log"
)

func EmailExist(email, password string) bool {
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var storedPassword string
	err = db.QueryRow("SELECT password FROM Users WHERE pseudo = ?", email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatal(err)
	}
	return password == storedPassword
}
