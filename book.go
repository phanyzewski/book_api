package main

import (
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID            int    `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	PublishedDate string `json:"publishedDate,omitempty"`
	// Rating        Rating        `json:"rating,omitempty"`
	// BookAvailable BookAvailable `"json:'bookAvailable,omitempty"`
	// Publisher     *Publisher    `json:"publisher,omitempty"`
	Author *Author `json:"author,omitempty"`
}

func (b *Book) GetBook(db *sqlx.DB) error {
	book := Book{}
	err := db.Get(&book, "SELECT title FROM books WHERE id=$1", b.ID)

	return err
}

func (b *Book) UpdateBook(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE books set title=$1 WHERE id=$2", b.Title, b.ID)

	return err
}

func (b *Book) DeleteBook(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM books WHERE id=$1", b.ID)

	return err
}

func (b *Book) CreateBook(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO books (title, published_date) VALUES ($1, $2)", b.Title, b.PublishedDate)
	return err
}

func GetBooks(db *sqlx.DB, start, count int) ([]Book, error) {
	books := []Book{}
	err := db.Select(&books, "SELECT id, title FROM books")

	if err != nil {
		return nil, err
	}

	return books, nil
}
