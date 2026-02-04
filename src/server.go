package Projet

import (
	"fmt"
	"net/http"
)

func Server() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/inscription", Inscription)
	http.HandleFunc("/login", Connexion)

	fmt.Println("Server lancé sur : localhost 8080")
	http.ListenAndServe(":8080", nil)

}
