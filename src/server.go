package Projet

import (
	"fmt"
	"net/http"
)

type Game struct {
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Genre            string `json:"genre"`
}

func Server() {

	LoadGames()

	http.HandleFunc("/", Home)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/setInfo", SetInscription)
	http.HandleFunc("/setconnect", SetConnexion)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/player", PlayerHandler)

	fmt.Println("Server lancé sur : localhost 8080")
	http.ListenAndServe(":8080", nil)

}
