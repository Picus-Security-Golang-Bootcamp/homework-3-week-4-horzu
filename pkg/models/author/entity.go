package author

import (
	"github.com/horzu/golang/picus-security-bootcamp/homework-3-week-4-horzu/pkg/models/book"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name string `json:"Name"`
	Books []book.Book `gorm:"foreignKey:AuthorID;references:id"`
}

// func (Author) TableName() string {
// 	return "Author"
// }