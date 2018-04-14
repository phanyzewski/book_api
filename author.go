package main

import (
	"github.com/jmoiron/sqlx"
)

type Author struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty" db:"first_name"`
	LastName  string `json:"lastName,omitempty" db:"last_name"`
	PenName   string `json:"penName,omitempty" db:"pen_name"`
}

func (a *Author) GetAuthor(db *sqlx.DB) error {
	author := Author{}
	err := db.Get(&author, "SELECT first_name, last_name, pen_name FROM authors WHERE id=$1", a.ID)
	if err != nil {
		return err
	}

	return err
}

func (a *Author) UpdateAuthor(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE authors set first_name=$1, last_name=$2, pen_name=$3 WHERE id=$4", a.FirstName, a.LastName, a.PenName, a.ID)

	return err
}

func (a *Author) DeleteAuthor(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM authors WHERE id=$1", a.ID)

	return err
}

func (a *Author) CreateAuthor(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO authors (first_name, last_name, pen_name) VALUES ($1, $2, $3)", a.FirstName, a.LastName, a.PenName)

	return err
}

func GetAuthors(db *sqlx.DB, start, count int) ([]Author, error) {
	authors := []Author{}
	err := db.Select(&authors, "SELECT id, first_name, last_name, pen_name FROM authors")

	if err != nil {
		return nil, err
	}

	return authors, nil
}
