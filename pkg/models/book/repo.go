package book

import (
	"encoding/json"
	"errors"
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

func (b *BookRepository) GetBookByID(id uint) (*Book, error) {
	var book Book
	result := b.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &book, nil
}

func (b *BookRepository) FindBookByName(title string) []Book {
	var book []Book
	b.db.Where("title ILIKE ? ", "%"+title+"%").Find(&book)

	return book
}

func (b *BookRepository) Create(book Book) error {
	result := b.db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) Update(book Book) error {
	result := b.db.Save(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) Delete(book Book) error {
	result := b.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) DeleteById(id int) error {
	result := b.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) GetAllBooksWithAuthorInformation() ([]Book, error) {
	var book []Book
	result := b.db.Preload("Authors").Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (b *BookRepository) FindAll() []Book {
	var books []Book
	b.db.Find(&books)

	return books
}