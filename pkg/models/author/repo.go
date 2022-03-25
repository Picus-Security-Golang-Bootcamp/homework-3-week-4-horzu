package author

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) Migration() {
	a.db.AutoMigrate(&Author{})
}

func (a *AuthorRepository) InsertSampleData() {
	jsonFile, err := os.Open("./pkg/mocks/authors.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	values, _ := ioutil.ReadAll(jsonFile)
	authors := []Author{}
	json.Unmarshal(values, &authors)

	for _, author := range authors {
		a.db.FirstOrCreate(&author)
	}
}

func (a *AuthorRepository) GetAuthorByID(id uint) (*Author, error) {
	var author Author
	result := a.db.First(&author, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &author, nil
}

func (a *AuthorRepository) FindAuthorByName(name string) []Author {
	var author []Author
	a.db.Where("name ILIKE ? ", "%"+name+"%").Find(&author)

	return author
}

func (a *AuthorRepository) Create(author Author) error {
	result := a.db.Create(author)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *AuthorRepository) Update(author Author) error {
	result := a.db.Save(author)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *AuthorRepository) Delete(author Author) error {
	result := a.db.Delete(author)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *AuthorRepository) DeleteById(id int) error {
	result := a.db.Delete(&Author{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *AuthorRepository) GetAuthorsWithBookInformation()  {
	result, err := a.db.Table("authors").Select("*").Joins("left join books on authors.id = books.author_id").Rows()
	if err!=nil{
		return 
	}
	for result.Next() {
		fmt.Println(*result)
	}
	
}