package main

import (
	"github.com/jmoiron/sqlx"
)

// Author model
type Author struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty" db:"first_name"`
	LastName  string `json:"lastName,omitempty" db:"last_name"`
	PenName   string `json:"penName,omitempty" db:"pen_name"`
}

// GetAuthor return author based on id
func (b *Author) GetAuthor(db *sqlx.DB) error {
	author := Author{}
	err := db.Get(&author, "SELECT * FROM authors WHERE id=$1", b.ID)

	return err
}

// UpdateAuthor update author based on id
func (b *Author) UpdateAuthor(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE authors set first_name=$1, last_name=$2, pen_name=$3 WHERE id=$4", b.FirstName, b.LastName, b.PenName, b.ID)

	return err
}

// DeleteAuthor remove author based on id
func (b *Author) DeleteAuthor(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM authors WHERE id=$1", b.ID)

	return err
}

// CreateAuthor new author
func (b *Author) CreateAuthor(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO authors (first_name, last_name, pen_name) VALUES ($1, $2, $3)", b.FirstName, b.LastName, b.PenName)
	return err
}

// GetAuthors return all authors
func GetAuthors(db *sqlx.DB, start, count int) ([]Author, error) {
	authors := []Author{}
	err := db.Select(&authors, "SELECT id, first_name, last_name, pen_name FROM authors")

	if err != nil {
		return nil, err
	}

	return authors, nil
}
