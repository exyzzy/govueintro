//Auto generated with MetaApi https://github.com/exyzzy/metaapi
package data

import (
	"database/sql"
	"time"
)


//Create Table
func CreateTableTodos(db *sql.DB) (err error) {
	_, err = db.Exec("DROP TABLE IF EXISTS todos")
	if err != nil {
		return
	}
	_, err = db.Exec(`create table todos ( id integer generated always as identity primary key , updated_at timestamptz , done boolean , title text ) ; `)
	return
}

//Struct
type Todo struct {
	Id int32`xml:"Id" json:"id"`
	UpdatedAt time.Time`xml:"UpdatedAt" json:"updatedat"`
	Done bool`xml:"Done" json:"done"`
	Title string`xml:"Title" json:"title"`

}

//Create
func (todo *Todo) CreateTodo(db *sql.DB) (result Todo, err error) {
	stmt, err := db.Prepare("INSERT INTO todos ( updated_at, done, title) VALUES ($1,$2,$3) RETURNING id, updated_at, done, title")
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow( todo.UpdatedAt, todo.Done, todo.Title).Scan( &result.Id, &result.UpdatedAt, &result.Done, &result.Title)
	return
}

//Retrieve
func (todo *Todo) RetrieveTodo(db *sql.DB) (result Todo, err error) {
	result = Todo{}
	err = db.QueryRow("SELECT id, updated_at, done, title FROM todos WHERE (id = $1)", todo.Id).Scan( &result.Id, &result.UpdatedAt, &result.Done, &result.Title)
	return
}

//RetrieveAll
func (todo *Todo) RetrieveAllTodos(db *sql.DB) (todos []Todo, err error) {
	rows, err := db.Query("SELECT id, updated_at, done, title FROM todos ORDER BY id DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		result := Todo{}
		if err = rows.Scan( &result.Id, &result.UpdatedAt, &result.Done, &result.Title); err != nil {
			return
		}
		todos = append(todos, result)
	}
	rows.Close()
	return
}

//Update
func (todo *Todo) UpdateTodo(db *sql.DB) (result Todo, err error) {
	stmt, err := db.Prepare("UPDATE todos SET updated_at = $2, done = $3, title = $4 WHERE (id = $1) RETURNING id, updated_at, done, title")
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow( todo.Id, todo.UpdatedAt, todo.Done, todo.Title).Scan( &result.Id, &result.UpdatedAt, &result.Done, &result.Title)
	return
}

//Delete
func (todo *Todo) DeleteTodo(db *sql.DB) (err error) {
	stmt, err := db.Prepare("DELETE FROM todos WHERE (id = $1)")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Id)
	return
}

//DeleteAll
func DeleteAllTodos(db *sql.DB) (err error) {
	stmt, err := db.Prepare("DELETE FROM todos")
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	return
}

