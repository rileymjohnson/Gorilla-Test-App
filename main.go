package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{item}", ProductsHandler)
	r.HandleFunc("/arguments", ArgumentsHandler)
	r.HandleFunc("/method", MethodHandlerGet).Methods("get")
	r.HandleFunc("/method", MethodHandlerPost).Methods("post")
	http.ListenAndServe(":5000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func MethodHandlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hi</h1>")
}

func MethodHandlerPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	last := r.FormValue("last")
	fmt.Fprintf(w, "the name is: "+name+last)
}

func ProductsHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	item := vars["item"]
	fmt.Fprintf(w, "the item: "+item)
}

func ArgumentsHandler(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	hostname := args.Get("key")
	fmt.Fprintf(w, hostname)
}
