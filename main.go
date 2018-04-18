package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting Book API...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := App{}
	fmt.Printf("env: DATABASE_URL%v\n", os.Getenv("DATABASE_URL"))
	a.Initialize(os.Getenv("DATABASE_URL"))
	a.Run(":8080")
}
