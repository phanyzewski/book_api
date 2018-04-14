package main

type Book struct {
	ID            string `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	PublishedDate string `json:"publishedDate,omitempty"`
	// Rating        Rating        `json:"rating,omitempty"`
	// BookAvailable BookAvailable `"json:'bookAvailable,omitempty"`
	// Publisher     *Publisher    `json:"publisher,omitempty"`
	// Author        *Author       `json:"author,omitempty"`
}

func (b *Book) getBook(db *sqlx.DB) error {
	book := Book{}
	err := db.Get(&book, "SELECT title, FROM books WHERE id=$1", b.ID)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *Book) updateBook(db *sqlx.DB) error {
	_, err := db.MustExec("UPDATE books set title=$1 WHERE id=$2", b.Title, b.ID)

	return err
}

func (b *Book) deleteBook(db *sqlx.DB) error {
	_, err := db.MustExec("DELETE FROM books WHERE id=$1", b.Title, b.ID)

	return err
}

func (b *Book) createBook(db *sqlx.DB) error {
	_, err := db.MustExec("INSERT INTO books(title) VALUES($1) RETURNING id ", b.Title)

	return err
}

func getBooks(db *sqlx.DB, start, count int) ([]Book, error) {
	books := []Book{}
	err := db.Select(&books, "SELECT id, title FROM books")

	if err != nil {
		return nil, err
	}

	return books, nil
}
