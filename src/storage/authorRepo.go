package storage

import (
	"PicusHomework4/src/pkg/httpErrors"
	"PicusHomework4/src/storage/helper"
	"PicusHomework4/src/storage/storageType"
	"errors"
	"gorm.io/gorm"
	"time"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	authorRepo := AuthorRepository{db: db}
	authorRepo.migrations()
	authorRepo.insertSampleData()
	return &authorRepo
}

func (a *AuthorRepository) FindAuthorById(id int) (*storageType.Author, error) {
	var author storageType.Author
	result := a.db.First(&author, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, httpErrors.NotFound
	}
	return &author, nil
}

func (a *AuthorRepository) CreateAuthor(author storageType.Author) (uint, error) {
	result := a.db.Create(&author)
	if result.Error != nil {
		return 0, httpErrors.RecordCreateError
	}
	return author.ID, nil
}

func (a *AuthorRepository) UpdateAuthorById(author storageType.Author, options map[string]interface{}) error {
	result := a.db.Model(&author).Updates(options).Where("deleted_at = ?", "null")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return httpErrors.NotFound
	}
	if result.Error != nil {
		return httpErrors.RecordUpdateError
	}
	return nil
}

func (a *AuthorRepository) DeleteAuthorById(id uint) error {
	result := a.db.Model(storageType.Author{}).Where("id = ?", id).Updates(storageType.Author{
		Model: gorm.Model{
			DeletedAt: gorm.DeletedAt{
				Time:  time.Now(),
				Valid: true,
			},
		},
	})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return httpErrors.NotFound
	}
	if result.Error != nil {
		return httpErrors.RecordDeleteError
	}
	return nil
}

func (a *AuthorRepository) GetAuthorByIdWithBooks(id int) (*storageType.Author, error) {
	var author storageType.Author
	result := a.db.Preload("Books").Where("id = ?", uint(id)).First(&author)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, httpErrors.NotFound
	}
	if result.Error != nil {
		return nil, httpErrors.FindAuthorError
	}
	return &author, nil
}

func (a *AuthorRepository) insertSampleData() error {
	authors, err := helper.ReadCSVForAuthor("author.csv")
	if err != nil {
		return err
	}
	for _, author := range authors {
		a.db.Where(storageType.Author{Name: author.Name, LastName: author.LastName}).
			Create(&author)
	}
	return nil
}

func (a *AuthorRepository) migrations() {
	a.db.AutoMigrate(&storageType.Author{})
}

/////////////////////////////////////////////////////////////

func (a *AuthorRepository) FindByName(name string) []storageType.Author {
	var authors []storageType.Author
	a.db.Where("firstName LIKE ? ", "%"+name+"%").
		Find(&authors)
	return authors
}

func (a *AuthorRepository) GetAllAuthorsWithBookInfo() (storageType.Author, error) {
	var authors []storageType.Author
	result := a.db.Preload("books").Find(&authors)
	if result.Error != nil {
		return storageType.Author{}, result.Error
	}
	return authors[1], nil
}
