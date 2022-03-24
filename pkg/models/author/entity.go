package author

import (
	"github.com/horzu/golang/picus-security-bootcamp/homework-3-week-4-horzu/pkg/models/book"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name string
	Books []book.Book
}