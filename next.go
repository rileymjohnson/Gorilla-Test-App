package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("adjfhskjdfhaksjdfhskjdfhskdjhf"))

func login(w http.ResponseWriter, r *http.Request, name string) bool {
	session, err := store.Get(r, "validation")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return false
	}
	session.Values[name] = true
	session.Save(r, w)
	return true
}

func logout(w http.ResponseWriter, r *http.Request, name string) bool {
	session, err := store.Get(r, "validation")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return false
	}
	session.Values[name] = false
	session.Save(r, w)
	return true
}

func isIn(w http.ResponseWriter, r *http.Request, name string) bool {
	session, err := store.Get(r, "validation")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return false
	}
	if val, ok := session.Values[name]; ok {
		if val == true {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)

	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/test", TestHandler)

	http.ListenAndServe(":5000", r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	login(w, r, "test")
	fmt.Fprintf(w, "login handler")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logout(w, r, "test")
	fmt.Fprintf(w, "logout handler")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	if isIn(w, r, "test") == true {
		fmt.Fprintf(w, "this is a test")
	} else {
		fmt.Fprintf(w, "not logged in")
	}
}
