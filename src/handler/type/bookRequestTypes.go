package _type

import (
	"PicusHomework4/src/pkg/httpErrors"
	"PicusHomework4/src/storage/storageType"
)

type RequestBookType struct {
	Name        string `json:"name"`
	AuthorId    uint   `json:"authorId"`
	Price       string `json:"price"`
	PublishTime string `json:"publishTime"`
	ISBN        string `json:"isbn"`
	PageNumber  int    `json:"pageNumber"`
	StockCode   string `json:"stockCode"`
	StockNumber int    `json:"stockNumber"`
}

func (book *RequestBookType) NewRequestTypeToBook() storageType.Book {
	return storageType.Book{
		Name:        book.Name,
		AuthorId:    book.AuthorId,
		Price:       book.Price,
		PublishTime: book.PublishTime,
		ISBN:        book.ISBN,
		PageNumber:  book.PageNumber,
		StockCode:   book.StockCode,
		StockNumber: book.StockNumber,
	}
}

func (book *RequestBookType) ValidateBookRequest() error {
	if book.Name == "" {
		return httpErrors.NameValidationError
	}
	if book.AuthorId == 0 {
		return httpErrors.AuthorIdValidationError
	}
	if book.Price == "" {
		return httpErrors.PriceValidationError
	}
	if book.PublishTime == "" {
		return httpErrors.PublishTimeValidationError
	}
	if book.ISBN == "" {
		return httpErrors.ISBNValidationError
	}
	if book.PageNumber <= 0 {
		return httpErrors.PageNumberValidationError
	}
	if book.StockCode == "" {
		return httpErrors.StockCodeValidationError
	}
	if book.StockNumber <= 0 {
		return httpErrors.StockNumberValidationError
	}
	return nil
}
