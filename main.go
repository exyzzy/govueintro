package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func AddAllMainRoutes(r *mux.Router) {

	// defined in route_main.go
	r.HandleFunc("/", HomeHandler).Methods("GET")

	// defined in route_todo.go
	r.HandleFunc("/api/todo", TodoPostHandler).Methods("POST")
	r.HandleFunc("/api/todo/{id}", TodoGetHandler).Methods("GET")
	r.HandleFunc("/api/todo/{id}", TodoPutHandler).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", TodoDeleteHandler).Methods("DELETE")
	r.HandleFunc("/api/todos", TodosGetHandler).Methods("GET")
	r.HandleFunc("/api/clean", ApiCleanHandler).Methods("DELETE")

	r.HandleFunc("/todo/todos", TodoTodosHandler).Methods("GET")

	//root
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(config.Static))))

}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	AddAllMainRoutes(router)

	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		IdleTimeout:    time.Duration(config.IdleTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	logger.Println(">>>>>>>", config.App, "("+config.SiteName+")", "started at", addr, "<<<<<<<")

	log.Fatal(server.ListenAndServe())
}
