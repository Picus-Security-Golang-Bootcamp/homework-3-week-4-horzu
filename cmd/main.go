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
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatalf("Postgres cannot init: %s", err)
	}
	log.Printf("Connected to Postgres Database.")

	// Initialize Repositories
	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migration()
	bookRepo := book.NewBookRepository(db)
	bookRepo.Migration()
	// bookRepo.InsertSampleData()
	// authorRepo.InsertSampleData()

	// Initialize methods
	// fmt.Println(authorRepo.GetAuthorByID(1))
	// fmt.Println(authorRepo.FindAuthorByName("es"))
	// fmt.Println(authorRepo.GetAuthorsWithBookInformation())
	// fmt.Println(authorRepo.GetDeletedAuthorsWithBookInformation())
	// fmt.Println(authorRepo.GetAuthorWithBookInformationByID(3))
	// authorRepo.DeleteById(2)
	// authorRepo.GetAuthorByID(1)
	// fmt.Println(authorRepo.FindAll())
	// fmt.Println(authorRepo.GetAuthorsCount())


	// bookRepo.InsertSampleData()
	// fmt.Println(bookRepo.GetBookByID(1))
	// fmt.Println(bookRepo.FindBookByName("decoder"))
	// result, _ := bookRepo.GetAllBooksWithAuthorInformation()
	// for _, v := range result {
	// 	fmt.Printf("Book: %s, Author: %s\n",v.Title, v.Authors[0].Name)
	// }
	// bookRepo.DeleteById(2)
	// bookRepo.GetBookByID(1)
	// fmt.Println(bookRepo.FindAll())
	// fmt.Println(bookRepo.GetBooksCount())
	// fmt.Println(bookRepo.GetBooksWithAuthorInformation())
	// fmt.Println(bookRepo.GetDeletedBooksWithAuthorInformation())
	// fmt.Println(bookRepo.GetBookWithAuthorInformationByID(3))
	// fmt.Println(bookRepo.GetBooksByPagesLessThenWithAuthorInformation(500))
	// bookRepo.GetStockCodeByTitle("the")
	fmt.Println(bookRepo.BuyBookByID(3,50))

}
