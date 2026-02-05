package Projet

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func LoadGames() ([]Game, error) {
	resp, err := http.Get("https://www.freetogame.com/api/games")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data []Game
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

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

func SetInscription(w http.ResponseWriter, r *http.Request) {
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

	pseudo := r.FormValue("pseudo")

	id, username, err := GetUserByPseudo(pseudo) //gestion d'erreur si pa de connextion par le pseudo
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound) //renvoie de /login si user pas connecter
		return
	}

	cookie := &http.Cookie{ //si ok on creer le cookie
		Name:  "user",
		Value: strconv.Itoa(id),
	}
	http.SetCookie(w, cookie)

	fmt.Println("connexion réussie", username)

	http.Redirect(w, r, "/", http.StatusFound)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "user",
		Value:  "",
		MaxAge: -1, // supprime cookie
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusFound)
}

func PlayerHandler(w http.ResponseWriter, r *http.Request) {

	games, err := LoadGames()

	tmpl, err := template.ParseFiles("pages/collection.html", "pages/templates/data.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, games)

}
