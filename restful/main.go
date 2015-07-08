package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)

	// Posts Collection
	posts := r.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(PostsIndexHandler)
	posts.Methods("POST").HandlerFunc(PostsCreateHandler)

	// Posts singular
	post := r.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(PostEditHandler)
	post.Methods("GET").HandlerFunc(PostsShowHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Posts Index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Posts Create")
}

func PostsShowHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintln(rw, "Showing Post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Post Update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Post Delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Post Edit")
}
