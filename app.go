package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func (a *App) Initialize(dataSourceName string) {
	var err error
	fmt.Printf("data source: %s \n", dataSourceName)

	a.DB, err = sqlx.Connect("postgres", "postgres://dev@localhost/book_development?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}
