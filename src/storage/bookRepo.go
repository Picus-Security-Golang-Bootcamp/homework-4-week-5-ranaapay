package storage

import (
	"PicusHomework4/src/pkg/httpErrors"
	"PicusHomework4/src/storage/helper"
	"PicusHomework4/src/storage/storageType"
	"errors"
	"gorm.io/gorm"
	"time"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	bookRepo := BookRepository{db: db}
	bookRepo.migrations()
	bookRepo.insertSampleData()
	return bookRepo
}

func (b *BookRepository) FindBookById(id int) (*storageType.Book, error) {
	var book storageType.Book
	result := b.db.First(&book, id).Where("deleted_at = ?", "null")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, httpErrors.NotFound
	}
	return &book, nil
}

func (b *BookRepository) CreateBook(book storageType.Book) (uint, error) {
	result := b.db.Create(&book)
	if result.Error != nil {
		return 0, httpErrors.RecordCreateError
	}
	return book.ID, nil
}

func (b *BookRepository) UpdateBookById(book storageType.Book, options map[string]interface{})  error {
	result := b.db.Model(&book).Updates(options).Where("deleted_at = ?", "null")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return httpErrors.NotFound
	}
	if result.Error != nil {
		return httpErrors.RecordUpdateError
	}
	return nil
}

func (b* BookRepository) DeleteBookById(id uint) error {
	result := b.db.Model(storageType.Book{}).Where("id = ?", id).Updates(storageType.Book{
		Model:       gorm.Model{
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

func (b *BookRepository) migrations() {
	b.db.AutoMigrate(&storageType.Book{})
}

func (b *BookRepository) insertSampleData() error {
	books, err := helper.ReadJSONForBook("book.json")
	if err != nil {
		return err
	}
	b.db.Create(&books)
	return nil
}

//___________________________________________________________________

func (b *BookRepository) FindBookByIdWithAuthor(id int) (*storageType.Book,error) {
	var book storageType.Book
	//result := b.db.Unscoped().Preload("authors").Where("id = ?", uint(id)).First(&book)
	result := b.db.Where("id = ?", uint(id)).Preload("authors", "id = ?", book.AuthorId).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil

}

func (b *BookRepository) GetAllBooks() ([]storageType.Book, error) {
	var books []storageType.Book
	result := b.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (b *BookRepository) FindBooksWithLimitOffset(limit int, offset int) []storageType.Book {
	var books []storageType.Book
	b.db.Limit(limit).Offset(offset).Find(&books)
	return books
}

func (b *BookRepository) FindByName(name string) []storageType.Book {
	var books []storageType.Book
	b.db.Where("Name LIKE ? ", "%"+name+"%").
		Find(&books)
	return books
}

func (b *BookRepository) FindBookByItsStockCode(requestCode string) storageType.Book {
	var book storageType.Book
	b.db.Where("stock_code = ?", requestCode).First(&book)
	return book
}

func (b *BookRepository) BuyBookByItsId(id int, quantity int) {
	var book storageType.Book
	b.db.Model(&book).Where("id = ?", id).
		Where("stock_number > ?", quantity).
		UpdateColumn("stock_number", gorm.Expr("stock_number - ?", quantity))
}

