package author

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model `gorm:"foreignKey:author_id;references:id"`
	Name string `json:"Name"`
}

// func (Author) TableName() string {
// 	return "Author"
// }