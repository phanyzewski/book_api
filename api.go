package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID            string        `json:"id,omitempty"`
	Title         string        `json:"title,omitempty"`
	PublishedDate string        `json:"publishedDate,omitempty"`
	Rating        Rating        `json:"rating,omitempty"`
	BookAvailable BookAvailable `"json:'bookAvailable,omitempty"`
	Publisher     *Publisher    `json:"publisher,omitempty"`
	Author        *Author       `json:"author,omitempty"`
}

type Author struct {
	ID        string  `json:"id,omitempty"`
	FirstName string  `json:"firstName,omitempty"`
	LastName  string  `json:"lastName,omitempty"`
	Books     *[]Book `json:"books,omitempty"`
}

type Publisher struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"title,omitempty"`
	Books *[]Book `json:"books,omitempty"`
}

type BookAvailable bool

const (
	CheckedOut BookAvailable = false
	CheckedIn  BookAvailable = true
)

type Rating int

const (
	OneStar Rating = iota + 1
	TwoStars
	ThreeStars
)

var books []Book

// Display all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// Display a single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// Delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}

func main() {
	books = append(books, Book{ID: "1", Title: "The Lord of the Rings", PublishedDate: "07-29-1954", Rating: ThreeStars, BookAvailable: CheckedIn, Author: &Author{ID: "1", LastName: "Tolkien", FirstName: "John"}})
	router := mux.NewRouter()

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9001", router))
}
