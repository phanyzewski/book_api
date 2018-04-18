package main

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Book model
type Book struct {
	ID            int        `json:"id,omitempty"`
	Title         string     `json:"title,omitempty"`
	PublishedDate time.Time  `json:"publishedDate,omitempty" db:"published_date"`
	Rating        Rating     `json:"rating,omitempty"`
	Status        Status     `json:"bookAvailable,omitempty"`
	Publisher     *Publisher `json:"publisher,omitempty" db:"publisher_id"`
	Author        *Author    `json:"author,omitempty" db:"author_id"`
}

// Status checked in or checked out
type Status int

// valid statuses
const (
	CheckedOut Status = iota
	CheckedIn
)

// Rating one to three stars
type Rating int

// valid ratings
const (
	OneStar Rating = iota + 1
	TwoStars
	ThreeStars
)

// String interface returns status in english
func (s Status) String() string {
	names := [...]string{
		"CheckedOut",
		"CheckedIn",
	}

	if s != CheckedOut && s != CheckedIn {
		return "Unknown"
	}

	return names[s]
}

// GetBook returns a book
func (b *Book) GetBook(db *sqlx.DB) error {
	book := Book{}
	err := db.Get(&book, "SELECT * FROM books WHERE id=$1", b.ID)

	return err
}

// UpdateBook updates a book
func (b *Book) UpdateBook(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE books set title=$1 WHERE id=$2", b.Title, b.ID)

	return err
}

// DeleteBook removes a book
func (b *Book) DeleteBook(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM books WHERE id=$1", b.ID)

	return err
}

// CreateBook inserts a new record
func (b *Book) CreateBook(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO books (title, published_date) VALUES ($1, $2)", b.Title, b.PublishedDate)
	return err
}

// GetBooks returns all books
func GetBooks(db *sqlx.DB, start, count int) ([]Book, error) {
	books := []Book{}
	err := db.Select(&books, "SELECT id, title FROM books")

	if err != nil {
		return nil, err
	}

	return books, nil
}
