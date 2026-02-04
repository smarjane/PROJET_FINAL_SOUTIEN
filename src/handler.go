package Projet

import (
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

	http.Redirect(w, r, "/", http.StatusFound)

}

type Cookie struct {
	Name  string
	Value string
}

func SetInfoHandler(w http.ResponseWriter, r *http.Request) {
	pseudo := r.FormValue("pseudo")
	nom := r.FormValue("nom")

	id := InsertValue(nom, pseudo)

	cookie := &http.Cookie{
		Name:  "user",
		Value: strconv.Itoa(id),
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusFound)
}
