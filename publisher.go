package main

import (
	"github.com/jmoiron/sqlx"
)

type Publisher struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (p *Publisher) GetPublisher(db *sqlx.DB) error {
	publisher := Publisher{}
	err := db.Get(&publisher, "SELECT * FROM publishers WHERE id=$1", p.ID)

	return err
}

func (p *Publisher) UpdatePublisher(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE publishers set name=$1 WHERE id=$2", p.Name, p.ID)

	return err
}

func (p *Publisher) DeletePublisher(db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM publishers WHERE id=$1", p.ID)

	return err
}

func (p *Publisher) CreatePublisher(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO publishers (name) VALUES ($1)", p.Name)
	return err
}

func GetPublishers(db *sqlx.DB, start, count int) ([]Publisher, error) {
	publishers := []Publisher{}
	err := db.Select(&publishers, "SELECT id, name FROM publishers")

	if err != nil {
		return nil, err
	}

	return publishers, nil
}
