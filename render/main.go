package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

func main() {
	ren := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now."))
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		ren.Data(w, http.StatusOK, []byte("Some binary data here"))
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		ren.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		ren.HTML(w, http.StatusOK, "example", nil)
	})

	http.ListenAndServe(":8080", mux)
}
