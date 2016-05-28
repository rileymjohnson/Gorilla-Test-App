package main

import (
	//"fmt"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("adjfhskjdfhaksjdfhskjdfhskdjhf"))

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", LoginHandler).Methods("GET")
	r.HandleFunc("/", LoginHandlerPost).Methods("POST")
	r.HandleFunc("/index", IndexHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", r)
	http.ListenAndServe(":5000", r)
}

var loginTemplate = pongo2.Must(pongo2.FromFile("templates/login.html"))
var indexTemplate = pongo2.Must(pongo2.FromFile("templates/index.html"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "login")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if _, ok := session.Values["login"]; ok {
		http.Redirect(w, r, "/index", 301)
	} else {
		loginTemplate.ExecuteWriter(pongo2.Context{
			"incorrect": false,
		}, w)
	}
}

func LoginHandlerPost(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	if name == "riley" && password == "letmein" {
		session, err := store.Get(r, "login")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		session.Values["login"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/index", 301)
	} else {
		loginTemplate.ExecuteWriter(pongo2.Context{
			"incorrect": true,
		}, w)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if _, ok := session.Values["login"]; ok {
		indexTemplate.ExecuteWriter(pongo2.Context{
			"message": "you're in",
		}, w)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
