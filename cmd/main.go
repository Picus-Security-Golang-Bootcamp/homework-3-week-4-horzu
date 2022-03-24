package main

import (
	"log"

	postgres "github.com/horzu/golang/picus-security-bootcamp/homework-3-week-4-horzu/pkg/db"
	"github.com/horzu/golang/picus-security-bootcamp/homework-3-week-4-horzu/pkg/models/author"
	"github.com/horzu/golang/picus-security-bootcamp/homework-3-week-4-horzu/pkg/models/book"
	"github.com/joho/godotenv"
)

func main() {
	// Set Environment Variables
	err := godotenv.Load()
	if err!=nil{
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := postgres.NewPsqlDB()
	if err!=nil{
		log.Fatalf("Postgres cannot init: %s", err)
	}
	log.Printf("Connected to Postgres Database.")

	// Repositories
	bookRepo := book.NewBookRepository(db)
	bookRepo.Migration()

	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
}