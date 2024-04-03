package controllers

import (
	"database/sql"
	"fmt"
	hash "forum/hash"
	models "forum/model"
	structs "forum/struct"
	"html/template"
	"net/http"
)

func PublicationHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/publication" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if structs.Connect {
		db, err := sql.Open("sqlite3", "./site.db")
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de la connexion à la base de données : %v", err), http.StatusInternalServerError)
			return
		}
		defer db.Close()
		contenu := r.FormValue("content")
		category := r.FormValue("category")
		_, err = db.Exec("INSERT INTO Posts (content, category,idUsers) VALUES (?, ?,?)", contenu, category, structs.IdConnected)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de la connexion à la base de données : %v", err), http.StatusInternalServerError)
			return
		}
		_, err = db.Exec("INSERT INTO comments (content,idUsers) VALUES (?, ?)", "", structs.IdConnected)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de la connexion à la base de données : %v", err), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func HomePageHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la connexion à la base de données : %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	if structs.Connect {
		err = db.QueryRow("SELECT id FROM Users WHERE pseudo = ?", structs.Connected).Scan(&structs.IdConnected)
		if err != nil {
			if err == sql.ErrNoRows {
				// Gérer le cas où aucune ligne n'a été trouvée
			} else {
				http.Error(w, fmt.Sprintf("Erreur lors de la recherche dans la base de données : %v", err), http.StatusInternalServerError)
				return
			}
		}
	}
	users, err := models.GetUsersFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts, err := models.GetPostsFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comments, err := models.GetCommentsFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := structs.PageData{
		Users:    users,
		Posts:    posts,
		Comments: comments,
	}
	tmpl, err := template.ParseFiles("./view/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginAccountHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		password = hash.Hash(password)
		if username == "" || password == "" {
			http.Redirect(w, r, "/", http.StatusFound)
		}
		if models.EmailExist(username, password) {
			http.Redirect(w, r, "/home", http.StatusFound)
			structs.Connected = username
			structs.Connect = true
			return
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

	} else {
		tmpl, err := template.ParseFiles("./view/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func LoginPageHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	tmpl, err := template.ParseFiles("./view/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateAccountHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/creation" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la connexion à la base de données : %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	email := r.FormValue("email")
	username := r.FormValue("pseudo")
	password := r.FormValue("password")
	password = hash.Hash(password)
	uuid := models.GenerateUUID()
	_, err = db.Exec("INSERT INTO Users (email, pseudo, password, uuid) VALUES (?, ?, ?, ?)", email, username, password, uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'insertion des données dans la base de données : %v", err), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func DeconnectionHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/deco" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	structs.Connect = false
	http.Redirect(w, r, "/", http.StatusFound)
}

func ProfilAccountHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profil" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if structs.Connect {
		users, err := models.GetUsersFromDBIfConnected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts, err := models.GetPostsFromDBIfConnected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		comments, err := models.GetCommentsFromDBIfConnected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := structs.PageData{
			Users:    users,
			Posts:    posts,
			Comments: comments,
		}
		tmpl, err := template.ParseFiles("./view/profil.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comments" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if structs.Connect {
		commentID := r.FormValue("commentID")
		comments := r.FormValue("commentaire")

		db, err := sql.Open("sqlite3", "./site.db")
		if err != nil {
			http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Exécuter une requête SELECT pour récupérer les informations sur le post
		var comment structs.Comments
		err = db.QueryRow("SELECT id, content FROM comments WHERE id = ?", commentID).Scan(&comment.ID, &comment.Content)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des informations sur le commentaire", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("UPDATE comments SET content = ? WHERE id = ?", comments, commentID)
		if err != nil {
			http.Error(w, "Erreur lors de la mise à jour du commentaire dans la base de données", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LikePostHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likePost" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	if structs.Connect {
		// Ouvrir la connexion à la base de données
		postID := r.FormValue("postID")
		db, err := sql.Open("sqlite3", "./site.db")
		if err != nil {
			http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
			return
		}
		defer db.Close()
		var post structs.Posts
		err = db.QueryRow("SELECT id, content, category, likes FROM posts WHERE id = ?", postID).Scan(&post.ID, &post.Category, &post.Content, &post.Likes)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des informations sur le post", http.StatusInternalServerError)
			return
		}

		// Exécuter une requête UPDATE pour incrémenter le nombre de likes du post
		_, err = db.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", postID)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DislikePostHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dislikePost" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if structs.Connect {
		postID := r.FormValue("postID")
		db, err := sql.Open("sqlite3", "./site.db")
		if err != nil {
			http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var post structs.Posts
		err = db.QueryRow("SELECT id, category, content, dislikes FROM posts WHERE id = ?", postID).Scan(&post.ID, &post.Category, &post.Content, &post.Dislikes)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des informations sur le post", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("UPDATE posts SET dislikes = dislikes + 1 WHERE id = ?", postID)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LikeComHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likecom" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if structs.Connect {
		commentID := r.FormValue("commentID")
		db, err := sql.Open("sqlite3", "./site.db")
		if err != nil {
			http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var comment structs.Comments
		err = db.QueryRow("SELECT id, content, likes FROM comments WHERE id = ?", commentID).Scan(&comment.ID, &comment.Content, &comment.Likes)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des informations sur le commentaire", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("UPDATE comments SET likes = likes + 1 WHERE id = ?", commentID)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DislikeComHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dislikecom" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	if structs.Connect {
		commentID := r.FormValue("commentID")
		db, err := sql.Open("sqlite3", "./site.db")
		if err != nil {
			http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var comment structs.Comments
		err = db.QueryRow("SELECT id, content, dislikes FROM comments WHERE id = ?", commentID).Scan(&comment.ID, &comment.Content, &comment.Dislikes)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des informations sur le commentaire", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("UPDATE comments SET dislikes = dislikes + 1 WHERE id = ?", commentID)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ChangeUsernameHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/changePseudo" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	newPseudo := r.FormValue("nouveauPseudo")
	structs.Connected = newPseudo
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE Users SET pseudo = ? WHERE id = ?", newPseudo, structs.IdConnected)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func ChangeEmailHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/changeEmail" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	newEmail := r.FormValue("nouveauMail")
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE Users SET email = ? WHERE id = ?", newEmail, structs.IdConnected)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func ChangeMDPHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/changeMDP" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	newMDP := r.FormValue("nouveauMDP")
	newMDP = hash.Hash(newMDP)
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE Users SET password = ? WHERE id = ?", newMDP, structs.IdConnected)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func ChangePostHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/changePost" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	newPost := r.FormValue("post")
	postID := r.FormValue("postID")
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE posts SET content = ? WHERE id = ?", newPost, postID)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func ChangeComHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/changeCom" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	newCom := r.FormValue("commentaire")
	commentID := r.FormValue("commentaireID")
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE comments SET content = ? WHERE id = ?", newCom, commentID)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func DeleteComHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/deleteCom" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	commentID := r.FormValue("commentaireID")
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM comments WHERE id = ?", commentID)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func DeletePostHd(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/deletePost" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	postID := r.FormValue("postID")
	db, err := sql.Open("sqlite3", "./site.db")
	if err != nil {
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM posts WHERE id = ?", postID)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}

func Error404(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
