package author

import "gorm.io/gorm"

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository{
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) Migration() {
	a.db.AutoMigrate(&Author{})
}