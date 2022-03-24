package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Page      int    `json:"page"`
	Stock     int    `json:"stock"`
	Price     string `json:"price"`
	StockCode string `json:"stockCode"`
	ISBN      string `json:"ISBN"`
	AuthorID  uint   `json:"author"`
}

func (Book) TableName() string {
	return "Book"
}

func (b *Book) toString() string {
	return fmt.Sprintf("ID : %d, Title : %s, Page : %d, Stock : %d, Price : %s, StockCode : %s, ISBN : %s, AuthorID : %d, CreatedAt : %s", 
	b.ID, b.Title, b.Page, b.Stock, b.Price, b.StockCode, b.ISBN, b.AuthorID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (b *Book) beforeDelete(tx *gorm.DB) (err error){
	fmt.Printf("Book (%s) deleting...", b.Title)
	return nil
}