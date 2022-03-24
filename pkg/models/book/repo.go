package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/gorm"
)

type BookRepository struct{
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (b *BookRepository) GetAllBooks() []Book{
	var books []Book
	b.db.Find(&books)

	return books
}

func (b *BookRepository) Migration() {
	b.db.AutoMigrate(&Book{})
}

func (b *BookRepository) InsertSampleData() {
	jsonFile, err := os.Open("./pkg/mocks/books.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	values, _ := ioutil.ReadAll(jsonFile)
	books := []Book{}
	json.Unmarshal(values, &books)

	for _, book := range books {
		b.db.FirstOrCreate(&book)
	}
}