package book

import "gorm.io/gorm"

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