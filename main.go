package main

import "net/http"

func main() {
	http.HandleFunc("GET /events", nil)
	http.HandleFunc("POST /events", nil)
	http.HandleFunc("GET /events/{id}", nil)

	http.ListenAndServe(":5000", nil)
}