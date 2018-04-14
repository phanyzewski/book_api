package main

import (
	"database/sql"
	"errors"
)

type book struct {
	ID            string `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	PublishedDate string `json:"publishedDate,omitempty"`
	// Rating        Rating        `json:"rating,omitempty"`
	// BookAvailable BookAvailable `"json:'bookAvailable,omitempty"`
	// Publisher     *Publisher    `json:"publisher,omitempty"`
	// Author        *Author       `json:"author,omitempty"`
}

func (p *book) getBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *book) updateBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *book) deleteBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *book) createBook(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getBooks(db *sql.DB, start, count int) ([]book, error) {
	return nil, errors.New("Not implemented")
}
