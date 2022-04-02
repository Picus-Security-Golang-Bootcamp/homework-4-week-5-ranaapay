package helper

import (
	"PicusHomework4/src/storage/storageType"
)

func SetBookUpdateOptions(book storageType.Book) map[string]interface{} {
	updateOptions := map[string]interface{}{
		"name":         book.GetBookName(),
		"author_id":    book.GetAuthorId(),
		"price":        book.GetPrice(),
		"publish_time": book.GetPublishTime(),
		"isbn":         book.GetISBN(),
		"page_number":  book.GetPageNumber(),
		"stock_code":   book.GetStockCode(),
		"stock_number": book.GetStockNumber(),
	}
	return updateOptions
}

func SetAuthorUpdateOptions(author storageType.Author) map[string]interface{} {
	updateOptions := map[string]interface{}{
		"name":      author.Name,
		"last_name": author.LastName,
		"email":     author.Email,
		"gender":    author.Gender,
	}
	return updateOptions
}
