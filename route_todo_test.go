//https://elithrar.github.io/article/testing-http-handlers-go/
//https://github.com/kelvins/GoApiTutorial
//https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/exyzzy/govueintro/data"
	"github.com/gorilla/mux"
)

//test data
var testTodo = [2]data.Todo{{0, time.Now().UTC(), false, "MyTestTodo Foo"}, {0, time.Now().UTC(), true, "MyTestTodo Bar"}}

type Stack []int32

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// https://github.com/golang/go/wiki/SliceTricks

func (s *Stack) Push(x int32) int32 {
	*s = append(*s, x)
	return x
}

func (s *Stack) Pop() int32 {
	if s.IsEmpty() {
		return -1
	}
	var x int32
	x, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return x
}

func (s *Stack) Peek() int32 {
	if s.IsEmpty() {
		return -1
	}
	return (*s)[len(*s)-1]
}

// stack for accumulated test records, for reference and to cleanup later
var testTodoId Stack

//https://godoc.org/github.com/gorilla/mux#SetURLVars
//must set URLVars with Gorilla

func TestTodoPostHandler(t *testing.T) {

	// todoJson := `{"title": "` + testTodoTitle[0] + `"}`
	fmt.Println("==TestTodoPostHander")

	js, err := json.Marshal(testTodo[0])

	reader := strings.NewReader(string(js))

	req, err := http.NewRequest("POST", "/api/todo", reader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TodoPostHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("cannot ioutil.ReadAll(request.Body): " + err.Error())
	}

	var response data.Todo

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		t.Errorf("cannot Unmarshall body: " + err.Error())
	}

	var expected data.Todo

	//use actual Id and SubmittedAt
	expected.Id = response.Id
	expected.UpdatedAt = response.UpdatedAt
	expected.Title = testTodo[0].Title
	expected.Done = testTodo[0].Done

	testTodoId.Push(response.Id)

	expectJson, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("cannot json.Marshal(expected): " + err.Error())
	}

	fmt.Println("  Response: ", string(body))
	fmt.Println("  Expected: ", string(expectJson))

	if strings.TrimSpace(string(body)) != strings.TrimSpace(string(expectJson)) {

		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), string(expectJson))
	}

}

func TestTodoGetHandler(t *testing.T) {

	fmt.Println("==TestTodoGetHander")

	req, err := http.NewRequest("GET", "/api/todo/", nil)
	if err != nil {
		t.Fatal(err)
	}

	//note id is not parsed, must use mux.SetURLVars()
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(testTodoId.Peek()))})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TodoGetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("cannot ioutil.ReadAll(request.Body): " + err.Error())
	}

	var response data.Todo

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		t.Errorf("cannot Unmarshall body: " + err.Error())
	}

	var expected data.Todo

	expected.Id = response.Id
	expected.UpdatedAt = response.UpdatedAt
	expected.Title = testTodo[0].Title
	expected.Done = testTodo[0].Done

	expectJson, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("cannot json.Marshal(expected): " + err.Error())
	}

	fmt.Println("  Response: ", string(body))
	fmt.Println("  Expected: ", string(expectJson))

	if strings.TrimSpace(string(body)) != strings.TrimSpace(string(expectJson)) {

		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), string(expectJson))
	}

}

func TestTodoPutHandler(t *testing.T) {

	// todoJson := `{"title": "` + testTodoTitle[1] + `","done":` + strconv.FormatBool(testTodoDone[1]) + `}`

	fmt.Println("==TestToDoPutHander")

	js, err := json.Marshal(testTodo[1])

	reader := strings.NewReader(string(js))

	req, err := http.NewRequest("PUT", "/api/todo", reader)
	if err != nil {
		t.Fatal(err)
	}

	//note id is not parsed, must use mux.SetURLVars()
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(testTodoId.Peek()))})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TodoPutHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("cannot ioutil.ReadAll(request.Body): " + err.Error())
	}

	var response data.Todo

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		t.Errorf("cannot Unmarshall body: " + err.Error())
	}

	var expected data.Todo

	expected.Id = response.Id
	expected.UpdatedAt = response.UpdatedAt
	expected.Title = testTodo[1].Title
	expected.Done = testTodo[1].Done

	expectJson, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("cannot json.Marshal(expected): " + err.Error())
	}

	fmt.Println("  Response: ", string(body))
	fmt.Println("  Expected: ", string(expectJson))

	if strings.TrimSpace(string(body)) != strings.TrimSpace(string(expectJson)) {

		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), string(expectJson))
	}

}

func TestTodoDeleteHandler(t *testing.T) {

	fmt.Println("==TestTodoDeleteHander")
	fmt.Println("  Deleting: ", testTodoId.Peek())

	req, err := http.NewRequest("DELETE", "/api/todo", nil)
	if err != nil {
		t.Fatal(err)
	}

	//note id is not parsed, must use mux.SetURLVars()
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(testTodoId.Peek()))})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TodoDeleteHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//ensure it's gone
	req, err = http.NewRequest("GET", "/api/todo/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(testTodoId.Pop()))})

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(TodoGetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

}

func TestTodosGetHandler(t *testing.T) {

	fmt.Println("==TestTodosGetHander")

	t.Run("  Add Todo1", TestTodoPostHandler) //add back in so more than 1
	t.Run("  Add Todo2", TestTodoPostHandler)

	req, err := http.NewRequest("GET", "/api/todos/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TodosGetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("cannot ioutil.ReadAll(request.Body): " + err.Error())
	}

	var response []data.Todo

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		t.Errorf("cannot Unmarshall body: " + err.Error())
	}

	//simple test
	if len(response) < 2 {
		t.Errorf("expected at least 2 records")
	}

	//cleanup
	for !testTodoId.IsEmpty() {
		t.Run("  Cleanup", TestTodoDeleteHandler) //add [0] back in so more than 1
	}

}
