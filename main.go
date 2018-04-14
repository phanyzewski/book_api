package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// type Author struct {
// 	ID        string  `json:"id,omitempty"`
// 	FirstName string  `json:"firstName,omitempty"`
// 	LastName  string  `json:"lastName,omitempty"`
// 	Books     *[]Book `json:"books,omitempty"`
// }

// type Publisher struct {
// 	ID    string  `json:"id,omitempty"`
// 	Name  string  `json:"title,omitempty"`
// 	Books *[]Book `json:"books,omitempty"`
// }

// type BookAvailable bool

// const (
// 	CheckedOut BookAvailable = false
// 	CheckedIn  BookAvailable = true
// )

// type Rating int

// const (
// 	OneStar Rating = iota + 1
// 	TwoStars
// 	ThreeStars
// )

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := App{}
	a.Initialize(os.Getenv("DATABASE_URL"))
	a.Run(":8080")
}
