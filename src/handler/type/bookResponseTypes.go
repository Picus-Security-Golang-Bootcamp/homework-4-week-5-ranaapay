package _type

import "PicusHomework4/src/storage/storageType"

type ResponseBookType struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	AuthorId    uint   `json:"authorId"`
	Price       string `json:"price"`
	PublishTime string `json:"publishTime"`
	ISBN        string `json:"isbn"`
	PageNumber  int    `json:"pageNumber"`
	StockCode   string `json:"stockCode"`
	StockNumber int    `json:"stockNumber"`
}

func NewResponseBookType(b *storageType.Book) *ResponseBookType {
	return &ResponseBookType{
		Id:          b.ID,
		Name:        b.Name,
		AuthorId:    b.AuthorId,
		Price:       b.Price,
		PublishTime: b.PublishTime,
		ISBN:        b.ISBN,
		PageNumber:  b.PageNumber,
		StockCode:   b.StockCode,
		StockNumber: b.StockNumber,
	}
}
