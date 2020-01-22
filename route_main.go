package main

import (
	// "fmt"
	"net/http"
)

// r.HandleFunc("/", HomeHandler).Methods("GET")
func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, nil, "toolbar.public.html", "home.html", "home.vue.js")
}
