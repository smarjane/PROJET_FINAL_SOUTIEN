package Projet

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/inscription.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func Connexion(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/connexion.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func SetInfoHandler(w http.ResponseWriter, r *http.Request) {
	pseudo := r.FormValue("pseudo")
	nom := r.FormValue("nom")

	id := InsertValue(nom, pseudo) //inscription

	cookie := &http.Cookie{
		Name:  "user",
		Value: strconv.Itoa(id),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}

func SetConnexion(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(`SELECT username, pseudo FROM users`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pseudo string
		var username string
		err := rows.Scan(&pseudo, &username)

		if err != nil {
			panic(err)
		}
		fmt.Println(pseudo, username)
	}
}

func Comparaison() {

}
