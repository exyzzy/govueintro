package main

import (
	"database/sql"
	"encoding/json"

	// "fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/exyzzy/govueintro/data"
	"github.com/gorilla/mux"
)

// Todos page template
// r.HandleFunc("/todo/todos", TodoTodosHandler).Methods("GET")
func TodoTodosHandler(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, nil, "toolbar.public.html", "todos.html", "todos.vue.js")
}

// ToDo CRUD API ===

// r.HandleFunc("/api/todos", TodosGetHandler).Methods("GET")
func TodosGetHandler(writer http.ResponseWriter, request *http.Request) {
	var todo data.Todo

	response, err := todo.RetrieveAllTodos(data.Db)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(writer, http.StatusNotFound, "Todos not found")
		default:
			respondWithError(writer, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(writer, http.StatusOK, response)
}

// r.HandleFunc("/api/todo", TodoPostHandler).Methods("POST")
func TodoPostHandler(writer http.ResponseWriter, request *http.Request) {

	var todo data.Todo
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	todo.UpdatedAt = time.Now().UTC()

	response, err := todo.CreateTodo(data.Db)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, response)
}

// r.HandleFunc("/api/todo/{id}", TodoGetHandler).Methods("GET")
func TodoGetHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	var todo data.Todo
	todo.Id = int32(id)

	response, err := todo.RetrieveTodo(data.Db)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(writer, http.StatusNotFound, "Todo not found")
		default:
			respondWithError(writer, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(writer, http.StatusOK, response)
}

// r.HandleFunc("/api/todo/{id}", TodoPutHandler).Methods("PUT")
func TodoPutHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	var todo data.Todo
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	todo.Id = int32(id)
	todo.UpdatedAt = time.Now().UTC()

	response, err := todo.UpdateTodo(data.Db)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, response)
}

// r.HandleFunc("/api/todo/{id}", TodoDeleteHandler).Methods("DELETE")
func TodoDeleteHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid todo ID")
		return
	}

	var todo data.Todo
	todo.Id = int32(id)

	err = todo.DeleteTodo(data.Db)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}

// r.HandleFunc("/api/clean", ApiCleanHandler).Methods("DELETE")
func ApiCleanHandler(writer http.ResponseWriter, request *http.Request) {

	err := data.CreateTableTodos(data.Db)
	if err != nil {
		respondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(writer, http.StatusOK, map[string]string{"result": "success"})
}
