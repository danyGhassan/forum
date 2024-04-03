package models

import (
	"database/sql"
	structs "forum/struct"
)

func GetID(connected string) (int, error) {
	// Connexion à la base de données SQLite
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		return 0, err
	}
	defer db.Close()
	var userID int
	err = db.QueryRow("SELECT id FROM Users WHERE pseudo=?", connected).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func GetCommentsFromDBIfConnected() ([]structs.Comments, error) {
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, content, likes, dislikes, idUsers FROM Comments WHERE idUsers = ?", structs.IdConnected)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []structs.Comments
	for rows.Next() {
		var infos structs.Comments
		if err := rows.Scan(&infos.ID, &infos.Content, &infos.Likes, &infos.Dislikes, &infos.IdUsers); err != nil {
			return nil, err
		}
		comments = append(comments, infos)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func GetPostsFromDBIfConnected() ([]structs.Posts, error) {
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, content, category, likes, dislikes, idUsers FROM posts WHERE idUsers = ?", structs.IdConnected)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var post []structs.Posts
	for rows.Next() {
		var infos structs.Posts
		if err := rows.Scan(&infos.ID, &infos.Content, &infos.Category, &infos.Likes, &infos.Dislikes, &infos.IdUsers); err != nil {
			return nil, err
		}
		post = append(post, infos)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return post, nil
}

func GetUsersFromDBIfConnected() ([]structs.User, error) {
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, pseudo, email FROM Users WHERE pseudo = ?", structs.Connected)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersFromDB() ([]structs.User, error) {
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, pseudo, email FROM Users ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetPostsFromDB() ([]structs.Posts, error) {
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, content, category, likes, dislikes, idUsers FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var post []structs.Posts
	for rows.Next() {
		var infos structs.Posts
		if err := rows.Scan(&infos.ID, &infos.Content, &infos.Category, &infos.Likes, &infos.Dislikes, &infos.IdUsers); err != nil {
			return nil, err
		}
		post = append(post, infos)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return post, nil
}

func GetCommentsFromDB() ([]structs.Comments, error) {
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, content, likes, dislikes, idUsers FROM Comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []structs.Comments
	for rows.Next() {
		var infos structs.Comments
		if err := rows.Scan(&infos.ID, &infos.Content, &infos.Likes, &infos.Dislikes, &infos.IdUsers); err != nil {
			return nil, err
		}
		comments = append(comments, infos)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
