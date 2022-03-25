package main

import (
	"fmt"
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
	
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
	authorRepo.InsertSampleData()
	// fmt.Println(authorRepo.GetAuthorByID(1))
	fmt.Println(authorRepo.FindAuthorByName("es"))

	bookRepo := book.NewBookRepository(db)
	bookRepo.Migration()
	bookRepo.InsertSampleData()
	// fmt.Println(bookRepo.GetBookByID(1))
	fmt.Println(bookRepo.FindBookByName("decoder"))

}