package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	controllers "forum/controller"
	models "forum/model"
)

func main() {
	models.Sqlite()
	mux.NewRouter()

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/home", controllers.HomePageHd)
	http.HandleFunc("/login", controllers.LoginAccountHd)
	http.HandleFunc("/", controllers.LoginPageHd)
	http.HandleFunc("/creation", controllers.CreateAccountHd)
	http.HandleFunc("/profil", controllers.ProfilAccountHd)
	http.HandleFunc("/publication", controllers.PublicationHd)
	http.HandleFunc("/comments", controllers.AddCommentHandler)
	http.HandleFunc("/deco", controllers.DeconnectionHd)
	http.HandleFunc("/likePost", controllers.LikePostHd)
	http.HandleFunc("/dislikePost", controllers.DislikePostHd)
	http.HandleFunc("/likecom", controllers.LikeComHd)
	http.HandleFunc("/dislikecom", controllers.DislikeComHd)
	http.HandleFunc("/changePseudo", controllers.ChangeUsernameHd)
	http.HandleFunc("/changeEmail", controllers.ChangeEmailHd)
	http.HandleFunc("/changeMDP", controllers.ChangeMDPHd)
	http.HandleFunc("/changePost", controllers.ChangePostHd)
	http.HandleFunc("/changeCom", controllers.ChangeComHd)
	http.HandleFunc("/deletePost", controllers.DeletePostHd)
	http.HandleFunc("/deleteCom", controllers.DeleteComHd)
	http.HandleFunc("/404", controllers.Error404)
	port := 8080
	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
