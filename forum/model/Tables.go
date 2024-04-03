package models

import (
	"database/sql"
	"fmt"
	"log"
)

func Sqlite() {

	db, err := sql.Open("sqlite3", "site.db")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
			uuid TEXT,
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            pseudo TEXT,
            email TEXT,
			password TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            content TEXT,
            category TEXT,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0,
			idUsers INTEGER,
        	FOREIGN KEY (idUsers) REFERENCES Users(id)
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            content TEXT,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0,
			idUsers INTEGER,
			idPosts INTEGER,
			FOREIGN KEY (idUsers) REFERENCES Users(id),
			FOREIGN KEY (idPosts) REFERENCES Posts(id)
        )
    `)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables crées avec succès")

}
