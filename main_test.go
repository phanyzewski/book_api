package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

var a App

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a = App{}
	a.Initialize(os.Getenv("TEST_DATABASE_URL"))

	// ensureTableExists()

	code := m.Run()
	ClearTable()
	os.Exit(code)
}

// func ensureTableExists() {
// 	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
// 		log.Fatal(err)
// 	}
// }

// const tableCreationQuery = `CREATE TABLE IF NOT EXISTS books
// (
//   id SERIAL PRIMARY KEY,
//   title TEXT,
//   published_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
//   rating NUMERIC,
//   book_available TEXT
// )`

func ClearTable() {
	a.DB.Exec("DELETE FROM books")
	a.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	ClearTable()

	req, _ := http.NewRequest("GET", "/books", nil)
	response := ExecuteRequest(req)

	CheckResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentBook(t *testing.T) {
	ClearTable()

	req, _ := http.NewRequest("GET", "/book/11", nil)
	response := ExecuteRequest(req)

	CheckResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["error"] != "Book not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Book not found'. Got '%s'", m["error"])
	}
}

func TestCreateBook(t *testing.T) {
	ClearTable()

	payload := []byte(`{"title":"The Hobbit", "published_date":"1937-09-21"}`)

	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
	response := ExecuteRequest(req)

	CheckResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "The Hobbit" {
		t.Errorf("Expected book title to be 'The Hobbit'. Got '%v'", m["title"])
	}
}

func TestGetBook(t *testing.T) {
	ClearTable()
	AddBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := ExecuteRequest(req)

	CheckResponseCode(t, http.StatusOK, response.Code)
}

func AddBooks(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO books (title) VALUES ($1)", "Book "+strconv.Itoa(i))
	}
}

func TestUpdateBook(t *testing.T) {
	ClearTable()
	AddBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := ExecuteRequest(req)
	var originalBook map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalBook)

	payload := []byte(`{"title":"test book - updated title"}`)

	req, _ = http.NewRequest("PUT", "/book/1", bytes.NewBuffer(payload))
	response = ExecuteRequest(req)

	CheckResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] == originalBook["title"] {
		t.Errorf("Expected the title to change from '%v' to '%v'. Got '%v'", originalBook["title"], m["title"], m["title"])
	}
}

func TestDeleteBook(t *testing.T) {
	ClearTable()
	AddBooks(1)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	response := ExecuteRequest(req)
	CheckResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	response = ExecuteRequest(req)

	CheckResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/book/1", nil)
	response = ExecuteRequest(req)
	CheckResponseCode(t, http.StatusNotFound, response.Code)

}
